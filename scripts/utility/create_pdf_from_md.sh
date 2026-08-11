#!/bin/bash
set -euo pipefail

SOURCE="$1"
DEST="$2"

if [[ "$SOURCE" != *.md ]]; then
    echo "Error: source file must be a .md file (got: $SOURCE)"
    exit 1
fi

if [[ "$DEST" != *.pdf ]]; then
    echo "Error: destination file must be a .pdf file (got: $DEST)"
    exit 1
fi

if [[ ! -f "$SOURCE" ]]; then
    echo "Error: source file not found: $SOURCE"
    exit 1
fi

docker run --rm -v "$(pwd):/data" pandoc/latex "$SOURCE" -o "$DEST"
echo "Generated $DEST"
