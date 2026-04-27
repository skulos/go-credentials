#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "$0")/.." && pwd)"

"$ROOT_DIR/scripts/install_golangci_lint.sh"

VERSION="v2.11.3"
LINT_BIN="$ROOT_DIR/bin/golangci-lint_${VERSION}/golangci-lint"

echo "🚀 Linting code"
"$LINT_BIN" run