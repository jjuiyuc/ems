#!/bin/bash

die() {
  echo $*
  exit 1
}

version=${1:-latest}

[ -z "${version}" ] && die "Usage: $0 VERSION"

cd $(dirname $0)

gcloud auth activate-service-account hes-build-415@ubiik-hes-dev.iam.gserviceaccount.com --key-file=service-account.json
gcloud config set project ubiik-hes-dev
gcloud auth configure-docker

docker build -t gcr.io/ubiik-hes-dev/der-ems-test-env:${version} .
docker push gcr.io/ubiik-hes-dev/der-ems-test-env:${version}
