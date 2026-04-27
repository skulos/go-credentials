#!/usr/bin/env bash
set -euo pipefail

VERSION="v2.15.4"
BIN_DIR="$(pwd)/bin"
INSTALL_DIR="$BIN_DIR/goreleaser_${VERSION}"
BINARY="$INSTALL_DIR/goreleaser"

mkdir -p "$INSTALL_DIR"

if [ -f "$BINARY" ]; then
  echo "✅ goreleaser already installed ($VERSION)"
  exit 0
fi

echo "📦 Installing goreleaser $VERSION"

# Install specific version
go install github.com/goreleaser/goreleaser/v2@${VERSION}

# Find where Go installed it
GO_BIN="$(go env GOBIN)"
if [ -z "$GO_BIN" ]; then
  GO_BIN="$(go env GOPATH)/bin"
fi

# Move binary into your bin folder
mv "${GO_BIN}/goreleaser" "$BINARY"

chmod +x "$BINARY"

echo "✅ Installed goreleaser -> $BINARY"