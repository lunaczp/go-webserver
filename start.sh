#!/usr/bin/env bash

if [ ! -f start.sh ]; then
    echo "must be run within its container folder" 1>&2
    exit 1
fi

dist/bin/server

