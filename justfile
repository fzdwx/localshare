#!/usr/bin/env just --justfile

update:
  go get -u
  go mod tidy -v

build:
    @cd web && npx tailwindcss -i input.css -o output.css
    go build .

css-watch:
    cd web && npx tailwindcss -i input.css -o output.css --watch