#!/bin/bash

# fail fast on errors
set -e

export DYNO=$1

source ./lib/run_hooks $(pwd)
