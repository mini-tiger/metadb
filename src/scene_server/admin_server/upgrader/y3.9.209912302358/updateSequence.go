package y3_9_209912302358

import (
	"configcenter/src/common"
	"configcenter/src/common/blog"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/source_controller/coreservice/core/operation"
	"configcenter/src/storage/dal"
	"configcenter/src/storage/dal/mongo/local"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

/**
 * @Author: Tao Jun
 * @Since: 2022/7/13
 * @Desc: updateSequence.go
**/

func GetMaxID(dbc dal.RDB, collection, idField string) (id int64, err error) {

	pipeline := []operation.M{

		{"$group": operation.M{"_id": "_id", "id": operation.M{"$max": fmt.Sprintf("$%s", idField)}}}, //排序后 取第一个 唯一字段

	}
	var result map[string]interface{}
	if err = dbc.Table(collection).AggregateOne(context.Background(), pipeline, &result); err != nil {
		return 0, err
	}
	var ok bool
	if id, ok = result["id"].(int64); !ok {
		return 0, nil
	}
	return id, nil

}

// NextSequence 获取新序列号(非事务)
func UpdateSequenceCol(dbc *mongo.Client, ctx context.Context, sequenceName string, id int64) (uint64, error) {

	// 直接使用新的context，确保不会用到事务,不会因为context含有session而使用分布式事务，防止产生相同的序列号
	ctx = context.Background()

	coll := dbc.Database("cmdb").Collection("cc_idgenerator")

	Update := bson.M{
		//"$inc":         bson.M{"SequenceID": int64(1)},
		"$setOnInsert": bson.M{"create_time": time.Now()},
		"$set":         bson.M{"last_time": time.Now(), "SequenceID": id},
	}
	filter := bson.M{"_id": sequenceName}
	upsert := true
	returnChange := options.After
	opt := &options.FindOneAndUpdateOptions{
		Upsert:         &upsert,
		ReturnDocument: &returnChange,
	}

	doc := local.Idgen{}
	err := coll.FindOneAndUpdate(ctx, filter, Update, opt).Decode(&doc)
	if err != nil {
		return 0, err
	}
	return doc.SequenceID, err
}
func UpdateSequence(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	for _, tableNameMap := range alltables {

		for tableName, idField := range tableNameMap {
			exists, err := db.HasTable(ctx, tableName)
			if err != nil {
				return err
			}
			if !exists {
				break
			}

			id, err := GetMaxID(db, tableName, idField)
			if err != nil {
				blog.Errorf("collction:%s GetMaxID err:%v\n", tableName, err)
				break
			}
			if id == 0 {
				blog.Infof("collction:%s GetMaxID:%d not found document\n", tableName, 0)
				break
			}

			mongoCli := db.Client()
			sid, err := UpdateSequenceCol(mongoCli, ctx, tableName, id)
			if err != nil {
				return err
			}
			blog.Infof("collction:%s seqid:%d\n", tableName, sid)
		}

	}
	return nil
}

func UpdateFieldTypeLong(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	for _, tableNameMap := range alltables {

		for tableName, idField := range tableNameMap {
			//err = db.Table(tableName).Update(ctx, operation.M{},
			//	types.ModeUpdate{Op: "set", Doc: operation.M{idField: operation.M{"$toLong": fmt.Sprintf("$%s", idField)}}})
			_, err = db.Client().Database("cmdb").Collection(tableName).UpdateMany(ctx, operation.M{},
				primitive.A{operation.M{"$set": operation.M{idField: operation.M{"$toLong": fmt.Sprintf("$%s", idField)}}}})

			if err != nil {
				continue
			}
		}
	}
	return nil
}

var alltables = []map[string]string{
	map[string]string{common.BKTableNameBaseInst: "bk_inst_id"},
	map[string]string{common.BKTableNameBaseModule: "bk_module_id"},
	map[string]string{common.BKTableNameObjAsst: "id"},
	map[string]string{common.BKTableNameObjAttDes: "id"},
	map[string]string{common.BKTableNameObjClassification: "id"},
	map[string]string{common.BKTableNameObjDes: "id"},

	map[string]string{common.BKTableNamePropertyGroup: "id"},
	map[string]string{common.BKTableNameObjUnique: "id"},
	map[string]string{common.BKTableNameAuditLog: "id"},
}
