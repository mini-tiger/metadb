package y3_9_202203171605

import (
	"context"
	"errors"
	"fmt"

	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"

	"configcenter/src/common"
	"configcenter/src/common/blog"
	"gopkg.in/mgo.v2"
)

//
func dropObjects(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	for tableName, _ := range tables {
		exists, err := db.HasTable(ctx, tableName)
		if err != nil {
			return err
		}
		if !exists {
			continue
		}
		if err = db.DropTable(ctx, tableName); err != nil && !mgo.IsDup(err) {
			return err
		}

	}
	return nil
}

func delObjects(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	err = delClassifications(ctx, db, conf)
	if err != nil {
		return errors.New(fmt.Sprintf("table:%s, err:%s", common.BKTableNameObjClassification, err))
	}
	err = delPropertyGroupData(ctx, db, conf)
	if err != nil {
		return err
	}
	//err = addObjDesData(ctx, db, conf)
	//if err != nil {
	//	return err
	//}
	//
	//err = addObjAttDescData(ctx, db, conf)
	//if err != nil {
	//	return err
	//}

	//err = addAsstData(ctx, db, conf)
	//if err != nil {
	//	return err
	//}

	return nil
}

func delClassifications(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	tablename := common.BKTableNameObjClassification
	subcond := map[string]interface{}{"$ne": "bk_biz_topo"}
	condition := map[string]interface{}{"bk_classification_id": subcond}
	blog.Infof("del %s rows", condition)
	err = db.Table(tablename).Delete(ctx, condition)

	return err
}

func delPropertyGroupData(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	tablename := common.BKTableNamePropertyGroup
	subcond := map[string]interface{}{"$ne": "biz"}
	condition := map[string]interface{}{"bk_obj_id": subcond}
	blog.Infof("del %s rows", condition)
	err = db.Table(tablename).Delete(ctx, condition)

	return err
}
