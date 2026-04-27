#!/usr/bin/env bash
set -euo pipefail

source "$(dirname "$0")/common.sh"

TPARSE_VERSION="0.18.0"
TPARSE_DIR="$BIN_DIR/tparse_v${TPARSE_VERSION}"
TPARSE_PATH="$TPARSE_DIR/tparse"

ensure_bin

if [ -f "$TPARSE_PATH" ]; then
  echo "✅ tparse already installed ($TPARSE_VERSION)"
  exit 0
fi

echo "📦 Installing tparse $TPARSE_VERSION"

mkdir -p "$TPARSE_DIR"

URL="https://github.com/mfridman/tparse/releases/download/v${TPARSE_VERSION}/tparse_${OS}_${REAL_ARCH}"

curl -sSL "$URL" -o "$TPARSE_PATH"
chmod +x "$TPARSE_PATH"

echo "✅ Installed tparse -> $TPARSE_PATH"