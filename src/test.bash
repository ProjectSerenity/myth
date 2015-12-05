#!/usr/bin/env bash
# Copyright 2015 The Myth Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# Automation script for testing individual (or all) myth packages.

set -e

if [ ! -f test.bash ]; then
    echo 'test.bash must be run from $MYTHROOT/src' 1>&2
    exit 1
fi

# Test for Go
hash go 2>/dev/null || { echo >&2 "Building Myth requires Go. See golang.org for more info."; exit 1; }

if [ $# -eq 0 ]; then
    echo "Usage: ./test.bash <package>..."
    exit 1
fi

GOPATH="$(cd .. && pwd)" go test $@
