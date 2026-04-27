#!/usr/bin/env bash
set -euo pipefail

source "$(dirname "$0")/common.sh"

VERSION="v2.11.3"
INSTALL_DIR="$BIN_DIR/golangci-lint_${VERSION}"
BINARY="$INSTALL_DIR/golangci-lint"

ensure_bin

if [ -f "$BINARY" ]; then
  echo "✅ golangci-lint already installed ($VERSION)"
  exit 0
fi

echo "📦 Installing golangci-lint $VERSION"

mkdir -p "$INSTALL_DIR"

curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
  | sh -s -- -b "$INSTALL_DIR" "$VERSION"

echo "✅ Installed golangci-lint -> $BINARY"