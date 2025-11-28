#!/usr/bin/env bash

set -e

psql -v ON_ERROR_STOP=1 -U "${POSTGRES_USER}" -d "${DB_NAME}" <<-EOF
INSERT INTO customer (name, email) VALUES
    ('Bob Cody', 'cody@truth.io'),
    ('O.W Grant', 'one@wish.you'),
    ('Neal Oliver', 'atr@forever.to'),
    ('Lynn Linden', 'love@seeker.to');
EOF
