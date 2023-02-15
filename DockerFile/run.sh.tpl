#!/bin/bash
ip=127.0.0.1
#port=

cd /data/cmdb
#python init.py --database cmdb --redis_ip redis-master --redis_port 6379 --redis_pass cc --mongo_ip mongo-mongodb-headless --mongo_port 27017 --mongo_user cc --mongo_pass cc --blueking_cmdb_url http://${ip}:8090 --listen_port 8090 --user_info admin:admin --auth_enabled false  --full_text_search off --log_level 0
python init.py --database cmdb --mongo_user cc --mongo_pass cc --redis_ip ${redisIp} --redis_port ${redisPort} --redis_pass ${redisPass} --mongo_ip ${mongoIp} --mongo_port ${mongoPort} --blueking_cmdb_url http://${ip}:8090 --listen_port 8090 --user_info admin:admin --auth_enabled false  --full_text_search off --log_level 0
{{if .AdminServer}}
cd /data/cmdb/cmdb_adminserver/configures/

#sed -i 's/opensource/skip-login/g' common.conf
#sed -i 's/opensource/skip-login/g' common.yaml
{{end}}

cd /data/cmdb/{{.AppName}}


# init data
{{if .AdminServer}}
./start.sh &
sleep 2
./init_db.sh
tail -f /dev/null
{{else}}
# start cmdb

./start.sh
{{end}}





