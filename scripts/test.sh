#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "$0")/.." && pwd)"
REPO_NAME="$(basename "$(git config --get remote.origin.url)" .git)"

source "$ROOT_DIR/scripts/common.sh"

# Ensure bin exists
ensure_bin

# Install dependencies
"$ROOT_DIR/scripts/install_tparse.sh"

TPARSE_VERSION="0.18.0"
TPARSE_PATH="$BIN_DIR/tparse_v${TPARSE_VERSION}/tparse"

echo "🚀 Running tests"

go test ./... -json | "$TPARSE_PATH"
#go test ./... -json | "$TPARSE_PATH" \
#  -trimpath "github.com/ingka-group-digital/${REPO_NAME}/internal/"