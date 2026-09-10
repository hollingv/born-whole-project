#!/bin/bash
set -e

if ! command -v cog &> /dev/null; then
    echo "Error: cog is not installed. Run 'make init' to install it."
    exit 1
fi

echo "[ pre-push ] Checking commit messages..."
cog check

echo "[ pre-push ] Running tests..."
make test

echo "[ pre-push ] All checks passed."
