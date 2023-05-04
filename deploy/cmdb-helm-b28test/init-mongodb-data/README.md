

手动导入导出数据测试
```shell
mongoexport --db cmdb --collection cc_ObjAsst --out mycollection.json --username cc --password cc --host 172.22.50.25 --port 27117 --authenticationDatabase=cmdb
mongoimport --db cmdb  --file cc_ObjAsst.json -c cc_ObjAsst --username cc --password cc --host 172.22.50.25 --port 27117 --authenticationDatabase=cmdb

b28-实验环境
mongoexport --db cmdb --collection cc_ObjAsst --out cc_ObjAsst.json --username neolink --password Ne01ink2022! --host 172.60.3.144 --port 27000 --authenticationDatabase=admin
mongoexport --db cmdb --collection cc_ObjAttDes --out cc_ObjAttDes.json --username neolink --password Ne01ink2022! --host 172.60.3.144 --port 27000 --authenticationDatabase=admin
mongoexport --db cmdb --collection cc_ObjClassification --out cc_ObjClassification.json --username neolink --password Ne01ink2022! --host 172.60.3.144 --port 27000 --authenticationDatabase=admin
mongoexport --db cmdb --collection cc_ObjDes --out cc_ObjDes.json --username neolink --password Ne01ink2022! --host 172.60.3.144 --port 27000 --authenticationDatabase=admin
mongoexport --db cmdb --collection cc_ObjectUnique --out cc_ObjectUnique.json --username neolink --password Ne01ink2022! --host 172.60.3.144 --port 27000 --authenticationDatabase=admin
mongoexport --db cmdb --collection cc_PropertyGroup --out cc_PropertyGroup.json --username neolink --password Ne01ink2022! --host 172.60.3.144 --port 27000 --authenticationDatabase=admin

```


开发环境导入数据

1. http://172.22.50.191:60004/migrate/v3/migrate/community/0
2. 
```shell
cd src/configcenter/deploy/cmdb-helm-b28test/init-mongodb-data
./check.sh  http://172.22.50.191:60004/find/init_mongodb_data 172.22.50.25:27117

```