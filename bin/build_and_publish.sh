#!/usr/bin/env bash
set -euxo pipefail

die() { echo "$1" >&2; exit "${2:-1}"; }

hash docker || die "docker not found"


version=$(cat VERSION)
tag=${1:-$version}

docker build --platform linux/amd64 --build-arg version="$version" -t "ghcr.io/gorginz/community-library:$tag" .

docker push "ghcr.io/gorginz/community-library:$tag"

#for local GCR
# docker build --platform linux/amd64 --build-arg version="$version" -t "asia.gcr.io/gorg-364804/community-library:$tag" .

# docker push "asia.gcr.io/gorg-364804/community-library:$tag"