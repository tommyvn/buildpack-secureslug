#!/bin/bash

# fail fast on errors
set -e

export PROC=$1

source ./lib/run_hooks $(pwd)
