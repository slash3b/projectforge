#!/bin/bash
# Code generated by Project Forge, see https://projectforge.dev for details.

## Runs goreleaser in test mode

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/../..

[[ -f "$HOME/bin/oauth" ]] && . $HOME/bin/oauth

goreleaser -f ./tools/release/.goreleaser.yml --snapshot --skip-publish --rm-dist
