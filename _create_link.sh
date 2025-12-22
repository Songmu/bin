#!/bin/bash
set -eo pipefail

bindir="$HOME/bin"

cd $(dirname $0)
for exe in *; do
    [[ -x "$exe" && ! "$exe" =~ ^_ ]] || continue
    ln -is "$PWD/$exe" "$bindir/$exe"
done
