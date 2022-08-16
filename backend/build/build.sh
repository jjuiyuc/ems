#!/bin/bash

GIT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
GIT_COMMIT=$(git rev-parse --short HEAD)

echo "branch: $GIT_BRANCH"
echo "commit: $GIT_COMMIT"

for i in $(ls -d cmd/daemon/* | cut -f3 -d'/'); do
    WORKER=${i%%/}
    echo "$WORKER"
    go build -ldflags "-X der-ems/config.gitBranch=$GIT_BRANCH -X der-ems/config.gitCommit=$GIT_COMMIT" -o dist/$WORKER cmd/daemon/$WORKER/$WORKER.go
done