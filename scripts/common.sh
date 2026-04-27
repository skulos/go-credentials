#!/usr/bin/env bash
set -euo pipefail

OS="$(uname | tr '[:upper:]' '[:lower:]')"
REAL_ARCH="$(uname -m)"

# Normalize architecture
case "$REAL_ARCH" in
  x86_64) ARCH="amd64" ;;
  aarch64 | arm64) ARCH="arm64" ;;
  *) ARCH="$REAL_ARCH" ;;
esac

BIN_DIR="$(pwd)/bin"

ensure_bin() {
  if [ ! -d "$BIN_DIR" ]; then
    echo "📁 Creating bin directory"
    mkdir -p "$BIN_DIR"
  fi
}