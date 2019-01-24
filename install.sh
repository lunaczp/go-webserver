#!/usr/bin/env bash

if [ ! -f install.sh ]; then
    echo "must be run within its container folder" 1>&2
    exit 1
fi

rm -rf dist/*
mkdir -p dist/bin

gofmt -w src
go build -v -o dist/bin/server src/main.go
cp -r src/app/view dist/view
cp -r config dist/config

echo  "done"
