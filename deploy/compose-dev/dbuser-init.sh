#!/bin/bash

mongo <<EOF
use cmdb
db.createUser({user: "cc",pwd: "cc",roles: [ { role: "readWrite", db: "cmdb" } ]})
EOF
