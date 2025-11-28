#!/usr/bin/env bash

set -e

psql -v ON_ERROR_STOP=1 -U "${POSTGRES_USER}"  <<-EOF
  CREATE DATABASE hiload;
EOF