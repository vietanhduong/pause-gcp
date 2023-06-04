# Pause GCP

A tool support you pause/unpause your GCP resources which save your cost. This tool is designed for running on CI.

Basically, Pause GCP rely-on **Gcloud CLI** that means, you must install the `gcloud` command first.

Currently, we support `gke`, `vm` and `cloud sql` resources.

## Usage

### Pause a GKE cluster
```console
$ pause-gcp gke pause --help

Pause a GKE cluster.
This command require '--location' and '--project' flags.

Usage:
  pause-gcp gke pause [CLUSTER_NAME] [flags]

Examples:
# write output from stdout
$ pause-gcp gke pause dev-cluster -l asia-southeast1 -p develop-project > output_state.json

# write output to gcs bucket
$ pause-gcp gke pause dev-cluster -l asia-southeast -p develop-project --output-dir=gs://bucket-name/gke-states

# write output to a directory, pause-gcp will try to create the output dir if it not exists
$ pause-gcp gke pause dev-cluster -p project --output-dir=output_states

# pause cluster with some except pools
$ pause-gcp gke pause dev-cluster -p project --except-pools=critical-pool


Flags:
      --except-pools strings                                      except node pools
  -h, --help                                                      help for pause
  -l, --location string                                           the cluster location (default "asia-southeast1")
      --output-dir gke_<project>_<location>_<cluster_name>.json   the output directory to write the cluster state. If no path is specified, this will skip the write-to-file process. The output state file has named by format gke_<project>_<location>_<cluster_name>.json
  -p, --project string                                            the project where contain the cluster
```

This command will print the previous state of the input cluster after it is paused. This state is used to recover the cluster in the unpause command.

If the `--output-dir` is a GCS bucket (start with `gs://`), this tool will push the state file to the destination directly.

### Unpause a GKE cluster

```console
$ pause-gcp gke unpause --help

Unpause a GKE cluster.
This command requires a GKE state file which is created when you pause the cluster.

Usage:
  pause-gcp gke unpause [STATE_FILE] [flags]

Examples:

# STATE_FILE from local
$ pause-gcp gke unpause ./gke-states/gke_develop_asia-southeast1_dev-cluster.state.json

# STATE_FILE from a gcs bucket
$ pause-gcp gke unpause gs://bucket/path/json_file.json --rm


Flags:
  -h, --help   help for unpause
      --rm     Remove the cluster state after complete
```
The input file must be the previos state of a cluster. The input file can be a GCS url (with `gs://` prefix).

## FAQ

### Why do I need this tool?
This will save your money.

### Can I use `gcloud` CLI instead?
Yes, you can. This tool builds on top of `gcloud` CLI. If you only need to turn off a VM or a cloud SQL instance,
you can use cloud CLI instead of this tool. But if you need to turn off a `GKE` cluster, I recommend you use `pause-gcp` because turning off a cluster is more complicated than a VM and cloud SQL.
