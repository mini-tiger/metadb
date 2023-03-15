

手动导入导出数据测试
```shell
mongoexport --db cmdb --collection cc_ObjAsst --out mycollection.json --username cc --password cc --host 172.22.50.25 --port 27117 --authenticationDatabase=cmdb
mongoimport --db cmdb  --file cc_ObjAsst.json -c cc_ObjAsst --username cc --password cc --host 172.22.50.25 --port 27117 --authenticationDatabase=cmdb

```


开发环境导入数据

1. http://172.22.50.191:60004/migrate/v3/migrate/community/0
2. 
```shell
cd src/configcenter/deploy/cmdb-helm-b28test/init-mongodb-data
./check.sh  http://172.22.50.191:60004/find/init_mongodb_data 172.22.50.25:27117

```