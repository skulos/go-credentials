#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "$0")/.." && pwd)"

"$ROOT_DIR/scripts/install_goreleaser.sh"

VERSION="v1.26.2"
GORELEASER="$ROOT_DIR/bin/goreleaser_${VERSION}/goreleaser"

ARGS=("release" "--clean")

# Parse flags
for arg in "$@"; do
  case $arg in
    --snapshot)
      ARGS=("release" "--snapshot" "--clean")
      ;;
    --skip-publish)
      ARGS+=("--skip=publish")
      ;;
  esac
done

echo "🚀 Running goreleaser ${ARGS[*]}"

"$GORELEASER" "${ARGS[@]}"