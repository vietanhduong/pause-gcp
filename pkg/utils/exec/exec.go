package exec

import (
	"bytes"
	"fmt"
	"github.com/vietanhduong/pause-gcp/pkg/utils/env"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const DefaultTimeout = 10 * time.Minute

var (
	timeout time.Duration
)

type CmdOpts struct {
	Timeout  time.Duration
	Redactor func(text string) string
}

var DefaultCmdOpts = CmdOpts{
	Timeout:  time.Duration(0),
	Redactor: Unredacted,
}

func init() {
	timeout = env.ParseDurationFromEnv("EXEC_TIMEOUT", DefaultTimeout, time.Second, 24*time.Hour)
}

func Redact(items []string) func(text string) string {
	return func(text string) string {
		for _, item := range items {
			text = strings.Replace(text, item, "******", -1)
		}
		return text
	}
}

// RunCommandExt is a convenience function to run/log a command and return/log stderr in an error upon
// failure.
func RunCommandExt(cmd *exec.Cmd, opts CmdOpts) (string, error) {
	redactor := DefaultCmdOpts.Redactor
	if opts.Redactor != nil {
		redactor = opts.Redactor
	}

	// log in a way we can copy-and-paste into a terminal
	args := strings.Join(cmd.Args, " ")

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Start()
	if err != nil {
		return "", err
	}

	done := make(chan error)
	go func() { done <- cmd.Wait() }()

	// Start a timer
	timeout := DefaultCmdOpts.Timeout

	if opts.Timeout != time.Duration(0) {
		timeout = opts.Timeout
	}

	var timoutCh <-chan time.Time
	if timeout != 0 {
		timoutCh = time.NewTimer(timeout).C
	}

	select {
	// noinspection ALL
	case <-timoutCh:
		_ = cmd.Process.Kill()
		output := stdout.String()
		err = newCmdError(redactor(args), fmt.Errorf("timeout after %v", timeout), "")
		return strings.TrimSuffix(output, "\n"), err
	case err := <-done:
		if err != nil {
			output := stdout.String()
			err = newCmdError(redactor(args), errors.New(redactor(err.Error())), strings.TrimSpace(redactor(stderr.String())))
			return strings.TrimSuffix(output, "\n"), err
		}
	}

	output := stdout.String()
	return strings.TrimSuffix(output, "\n"), nil
}

func Run(cmd *exec.Cmd) (string, error) {
	opts := CmdOpts{Timeout: timeout}
	return RunCommandExt(cmd, opts)
}

func Command(name string, args ...string) *exec.Cmd { return exec.Command(name, args...) }

// WaitPIDOpts are options to WaitPID
type WaitPIDOpts struct {
	PollInterval time.Duration
	Timeout      time.Duration
}

// WaitPID waits for a non-child process id to exit
func WaitPID(pid int, opts ...WaitPIDOpts) error {
	if runtime.GOOS != "linux" {
		return errors.Errorf("Platform '%s' unsupported", runtime.GOOS)
	}
	var timeout time.Duration
	pollInterval := time.Second
	if len(opts) > 0 {
		if opts[0].PollInterval != 0 {
			pollInterval = opts[0].PollInterval
		}
		if opts[0].Timeout != 0 {
			timeout = opts[0].Timeout
		}
	}
	path := fmt.Sprintf("/proc/%d", pid)

	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()
	var timoutCh <-chan time.Time
	if timeout != 0 {
		timoutCh = time.NewTimer(timeout).C
	}
	for {
		select {
		case <-ticker.C:
			_, err := os.Stat(path)
			if err != nil {
				if os.IsNotExist(err) {
					return nil
				}
				return errors.WithStack(err)
			}
		case <-timoutCh:
			return ErrWaitPIDTimeout
		}
	}
}
