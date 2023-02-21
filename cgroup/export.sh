#!/usr/bin/env bash

function test() {
  go build -buildmode=c-archive -o cgroup.a
}

function production() {
  GOOS=linux GOARCH=amd64 go build -buildmode=c-archive -o cgroup.a
}

$1
