#!/usr/bin/env just --justfile

build:
    @cd web && npx tailwindcss -i input.css -o output.css
    go build .


dev:
    LOCALSHARE_DEV=true go run .

css-watch:
    cd web && npx tailwindcss -i input.css -o output.css --watch
