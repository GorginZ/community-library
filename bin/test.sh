#!/usr/bin/env bash
set -exo pipefail

die() { echo "$1" >&2; exit "${2:-1}"; }

hash docker || die "docker not found"

# docker compose run --rm go test -v ./...
echo "hello testing, uncomment when functional options pr merged"

