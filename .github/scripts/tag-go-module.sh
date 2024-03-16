#!/bin/bash

test "${1}" = "" && echo "Usage: $0 <version>" && exit 1

# needs to be exported so the subshells of the find command can see it
export VERSION=${1}

# Tag root module
git tag "v${VERSION}"
