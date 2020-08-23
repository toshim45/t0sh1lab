#!/bin/bash

if [[ ${PWD##*/} != "competix" ]]; then echo "pls run from competix folder"; exit 1; fi

go run cmd/findsubstring/main.go < cmd/findsubstring/input.txt