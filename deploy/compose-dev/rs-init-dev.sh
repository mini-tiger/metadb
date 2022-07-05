#!/bin/bash

mongo <<EOF
var config = {
    "_id": "rs0",
    "version": 1,
    "members": [
        {
            "_id": 1,
            "host": "172.22.50.25:27021",
            "priority": 3
        },
        {
            "_id": 2,
            "host": "172.22.50.25:27022",
            "priority": 2
        },
        {
            "_id": 3,
            "host": "172.22.50.25:27023",
            "priority": 1
        }
    ]
};
rs.initiate(config, { force: true });
rs.status();
EOF
