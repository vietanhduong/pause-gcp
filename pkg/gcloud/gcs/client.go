package gcs

import (
	"bytes"
	"context"
	"io"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/iterator"
)

var (
	InvalidUrlErr   = errors.New("invalid url")
	FileNotFoundErr = errors.New("file not found")
)

type Client struct{}

type BucketUrl struct {
	BucketName string
	Path       string
}

func NewClient() *Client { return &Client{} }

func (c *Client) Cat(url string) ([]byte, error) {
	conn, err := c.newStorageClient()
	if err != nil {
		return nil, errors.Wrap(err, "cat")
	}

	br, err := parseBucketUrl(url)
	if err != nil {
		return nil, errors.Wrap(err, "cat")
	}

	if br.Path == "" {
		return nil, errors.Wrap(InvalidUrlErr, "cat")
	}

	reader, err := conn.Bucket(br.BucketName).Object(br.Path).NewReader(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "cat")
	}
	defer reader.Close()
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, errors.Wrap(err, "cat")
	}
	return b, nil
}

func (c *Client) Upload(dstUrl string, content []byte) error {
	conn, err := c.newStorageClient()
	if err != nil {
		return errors.Wrap(err, "upload")
	}
	br, err := parseBucketUrl(dstUrl)
	if err != nil {
		return errors.Wrap(err, "upload")
	}
	if br.Path == "" {
		return errors.New("upload: missing file name")
	}
	writer := conn.Bucket(br.BucketName).Object(br.Path).NewWriter(context.Background())
	writer.ChunkSize = 0
	if _, err = io.Copy(writer, bytes.NewBuffer(content)); err != nil {
		return errors.Wrap(err, "upload")
	}
	if err = writer.Close(); err != nil {
		return errors.Wrap(err, "upload")
	}
	return nil
}

func (c *Client) Remove(url string, concurrent bool) error {
	conn, err := c.newStorageClient()
	if err != nil {
		return errors.Wrap(err, "remove")
	}
	br, err := parseBucketUrl(url)
	if err != nil {
		return errors.Wrap(err, "remove")
	}

	it := conn.Bucket(br.BucketName).Objects(context.Background(), &storage.Query{Prefix: br.Path})
	var eg errgroup.Group
	for {
		obj, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return errors.Wrap(err, "remove")
		}
		if concurrent {
			eg.Go(func() error {
				return conn.Bucket(obj.Bucket).Object(obj.Name).Delete(context.Background())
			})
		} else {
			if err = conn.Bucket(obj.Bucket).Object(obj.Name).Delete(context.Background()); err != nil {
				return errors.Wrap(err, "remove")
			}
		}
	}
	if concurrent {
		if err = eg.Wait(); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) newStorageClient() (*storage.Client, error) {
	return storage.NewClient(context.Background())
}

func parseBucketUrl(url string) (*BucketUrl, error) {
	if !strings.HasPrefix(url, "gs://") {
		return nil, InvalidUrlErr
	}
	url = strings.TrimPrefix(url, "gs://")
	var ret BucketUrl
	parts := strings.Split(url, "/")
	ret.BucketName = parts[0]
	if len(parts) > 1 {
		ret.Path = strings.Join(parts[1:], "/")
	}
	return &ret, nil
}
