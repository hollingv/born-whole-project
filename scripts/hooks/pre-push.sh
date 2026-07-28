#!/bin/bash

if ! command -v cog &> /dev/null; then
    echo "Error: cog is not installed. Run 'make init' to install it."
    exit 1
fi

cog check v0.0.4..HEAD
