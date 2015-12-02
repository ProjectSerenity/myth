#!/usr/bin/env bash
# Copyright 2015 The Myth Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# Environment variables that control make.bash:
#
# MYTHROOT_FINAL: The expected final Myth root, baked into binaries.
# The default is the location of the Myth tree during the build.
#
# ARCH: The target architecture for installed packages and tools.
#
# MYTH_LDFLAGS: Additional go tool link arguments to use when
# building the commands.

set -e

if [ ! -f build.bash ]; then
	echo 'build.bash must be run from $MYTHROOT/src' 1>&2
	exit 1
fi

# Test for Windows.
case "$(uname)" in
*MINGW* | *WIN32* | *CYGWIN*)
	echo 'ERROR: Do not use build.bash to build on Windows.'
	echo 'Use build.bat instead.'
	echo
	exit 1
	;;
esac

# Test for Go.
hash go 2>/dev/null || { echo >&2 "Building Myth requires Go. See golang.org for more info."; exit 1; }

# Run the build.
GOPATH="$(cd .. && pwd)" go install cmd/myth
