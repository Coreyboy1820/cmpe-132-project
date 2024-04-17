#!/bin/bash
rm rbac.db
sqlite3 rbac.db <<EOF
.read sql/users.sql
.read sql/test.sql
.quit
EOF