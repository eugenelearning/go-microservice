#!/usr/bin/env bash

set -e

psql -v ON_ERROR_STOP=1 -U "${POSTGRES_USER}" -d "${POSTGRES_DB}" <<-EOF
  CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
  CREATE EXTENSION IF NOT EXISTS pgcrypto;

  CREATE TABLE customer (
    id    uuid PRIMARY KEY DEFAULT uuid_generate_v4(), 
    name  varchar(100) NOT NULL,
    email varchar(250) NOT NULL
  );
EOF