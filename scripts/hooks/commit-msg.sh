#!/bin/bash
set -e

if ! command -v cog &> /dev/null; then
    echo "Error: cog is not installed. Run 'make init' to install it."
    exit 1
fi

cog verify "$(cat "$1")"
