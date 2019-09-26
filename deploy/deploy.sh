#! /bin/bash

set -ue

host=travtech-1
zone=europe-north1-a
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

gcloud compute scp \
  --zone $zone \
  $dir/docker-compose.yml \
  $dir/api-manager.sh \
  $host:~/

gcloud compute ssh \
  --zone $zone \
  $host \
  --command 'bash api-manager.sh'
