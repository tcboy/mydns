#!/bin/bash

export GO111MODULE=on

mkdir -p ./output

go build -a -o ./output/mydns
