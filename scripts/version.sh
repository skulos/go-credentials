#!/usr/bin/env bash
set -euo pipefail

# Usage:
# ./scripts/version.sh patch
# ./scripts/version.sh minor
# ./scripts/version.sh major
# OR
# ./scripts/version.sh v1.2.3

VERSION_FILE="cmd/version.go"

# Extract current version
CURRENT_VERSION=$(grep -Eo 'v[0-9]+\.[0-9]+\.[0-9]+' "$VERSION_FILE")

echo "Current version: $CURRENT_VERSION"

bump_version() {
  local version=$1
  local part=$2

  # strip leading v
  version="${version#v}"

  IFS='.' read -r major minor patch <<< "$version"

  case "$part" in
    major)
      major=$((major + 1))
      minor=0
      patch=0
      ;;
    minor)
      minor=$((minor + 1))
      patch=0
      ;;
    patch)
      patch=$((patch + 1))
      ;;
    *)
      echo "Invalid bump type: $part"
      exit 1
      ;;
  esac

  echo "v${major}.${minor}.${patch}"
}

# Determine new version
if [[ "$1" =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
  NEW_VERSION="$1"
else
  NEW_VERSION=$(bump_version "$CURRENT_VERSION" "$1")
fi

echo "Bumping to: $NEW_VERSION"

# Update version in Go file
sed -i.bak "s/${CURRENT_VERSION}/${NEW_VERSION}/" "$VERSION_FILE"
rm "${VERSION_FILE}.bak"

# Commit change
git add "$VERSION_FILE"
git commit -m "chore: bump version to ${NEW_VERSION}"

# Create tag
git tag "$NEW_VERSION"

git push
git push --tags

echo "✅ Version updated and tagged: $NEW_VERSION"
