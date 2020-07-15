#!/usr/bin/env bash

cd internal/gen/
./templates.sh
cd ../../
make build


# /Users/domzmac/go/src/github.com/kryptodirect/database/schema/