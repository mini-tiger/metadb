#!/bin/bash
ip=127.0.0.1
#port=

cd /data/cmdb
python init.py --database cmdb --redis_ip redis-master --redis_port 6379 --redis_pass cc --mongo_ip mongo-mongodb --mongo_port 27017 --mongo_user cc --mongo_pass cc --blueking_cmdb_url http://${ip}:8090 --listen_port 8090 --user_info admin:admin --auth_enabled false  --full_text_search off --log_level 3
{{if .AdminServer}}
cd /data/cmdb/cmdb_adminserver/configures/

#sed -i 's/opensource/skip-login/g' common.conf
#sed -i 's/opensource/skip-login/g' common.yaml
{{end}}

# start cmdb
cd /data/cmdb/{{.AppName}}

./start.sh

# init data
{{if .AdminServer}}
sleep 2
./init_db.sh
{{end}}

tail -f /dev/null
