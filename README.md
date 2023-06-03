# Pause GCP

A tool support you pause/unpause your GCP resources which save your cost. This tool is designed for running on CI.

Basically, Pause GCP rely-on **Gcloud CLI** that means, you must install the `gcloud` command first.

Currently, we support `gke`, `vm` and `cloud sql` resources.

## FAQ

### Why do I need this tool?
This will save your money.

### Can I use `gcloud` CLI instead?
Yes, you can. This tool builds on top of `gcloud` CLI. If you only need to turn off a VM or a cloud SQL instance,
you can use cloud CLI instead of this tool. But if you need to turn off a `GKE` cluster, I recommend you use `pause-gcp b ecause turning off a cluster is more complicated than a VM and cloud SQL.


