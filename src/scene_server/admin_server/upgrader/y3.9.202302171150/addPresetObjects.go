/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package y3_9_202302171150

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"configcenter/src/common"
	"configcenter/src/common/blog"
	"configcenter/src/common/metadata"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

func addPresetObjects(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	err = addClassifications(ctx, db, conf)
	if err != nil {
		return err
	}

	err = addObjDesData(ctx, db, conf)
	if err != nil {
		return err
	}
	err = addPropertyGroupData(ctx, db, conf)
	if err != nil {
		return err
	}
	err = addObjAttDescData(ctx, db, conf)
	if err != nil {
		return err
	}

	//关联
	err = addAsstData(ctx, db, conf)
	if err != nil {
		return err
	}
	err = addUniqueData(ctx, db, conf)
	if err != nil {
		return err
	}

	//err = addObjectBaseData(ctx, db, conf)
	//if err != nil {
	//	return err
	//}

	return nil
}
func addObjectBaseData(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {
	tablename := "cc_ObjectBase"
	blog.Errorf("add data for  %s table ", tablename)
	rows := getAddObjectBaseData(conf.OwnerID)
	for _, row := range rows {
		// topo mainline could be changed,so need to ignore bk_asst_obj_id
		//_, _, err := upgrader.Upsert(ctx, db, tablename, row, "id", []string{}, []string{})
		err := upgrader.InsertData(ctx, db, tablename, row)
		if nil != err {
			blog.Errorf("add data for  %s table error  %s", tablename, err)
			return err
		}
	}

	blog.Errorf("add data for  %s table  ", tablename)
	return nil
}

func addUniqueData(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {
	tablename := common.BKTableNameObjUnique
	blog.Errorf("add data for  %s table ", tablename)
	rows := getAddUniqueData(conf.OwnerID)
	for _, row := range rows {
		// topo mainline could be changed,so need to ignore bk_asst_obj_id
		//_, _, err := upgrader.Upsert(ctx, db, tablename, row, "id", []string{}, []string{})
		err := upgrader.InsertData(ctx, db, tablename, row)
		if nil != err {
			blog.Errorf("add data for  %s table error  %s", tablename, err)
			return err
		}
	}

	blog.Errorf("add data for  %s table  ", tablename)
	return nil
}

func addAsstData(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {
	tablename := common.BKTableNameObjAsst
	blog.Errorf("add data for  %s table ", tablename)
	rows := getAddAsstData(conf.OwnerID)
	for _, row := range rows {
		// topo mainline could be changed,so need to ignore bk_asst_obj_id
		_, _, err := upgrader.Upsert(ctx, db, tablename, row, "id", []string{common.BKObjIDField, common.BKObjAttIDField, common.BKOwnerIDField}, []string{"id", "bk_asst_obj_id"})
		if nil != err {
			blog.Errorf("add data for  %s table error  %s", tablename, err)
			return err
		}
	}

	blog.Errorf("add data for  %s table  ", tablename)
	return nil
}

func addObjAttDescData(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {
	tablename := common.BKTableNameObjAttDes
	blog.Infof("add data for  %s table ", tablename)
	rows := getObjAttDescData(conf.OwnerID)
	for _, row := range rows {
		err := upgrader.InsertData(ctx, db, tablename, row)
		//upgrader.Insert()
		//_, _, err := upgrader.Upsert(ctx, db, tablename, row, "id", []string{common.BKObjIDField, common.BKPropertyIDField}, []string{})
		if nil != err {
			blog.Errorf("add data for  %s table error  %s", tablename, err)
			return err
		}
	}
	selector := map[string]interface{}{
		common.BKObjIDField: map[string]interface{}{
			common.BKDBIN: []string{"bk_switch",
				"bk_router",
				"bk_load_balance",
				"bk_firewall",
			},
		},
		common.BKPropertyIDField: "bk_name",
	}

	db.Table(tablename).Delete(ctx, selector)

	return nil
}

func addObjDesData(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {
	tablename := common.BKTableNameObjDes
	blog.Errorf("add data for  %s table ", tablename)
	rows := getObjectDesData(conf.OwnerID)
	for _, row := range rows {
		if _, _, err := upgrader.Upsert(ctx, db, tablename, row, "id", []string{common.BKObjIDField, common.BKClassificationIDField, common.BKOwnerIDField}, []string{"id"}); err != nil {
			blog.Errorf("add data for  %s table error  %s", tablename, err)
			return err
		}
	}

	return nil
}

func addClassifications(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	tablename := common.BKTableNameObjClassification
	blog.Infof("add %s rows", tablename)
	for _, row := range classificationRows {
		if _, _, err = upgrader.Upsert(ctx, db, tablename, row, "id", []string{common.BKClassificationIDField}, []string{"id"}); err != nil {
			blog.Errorf("add data for  %s table error  %s", tablename, err)
			return err
		}
	}
	return
}

func addPropertyGroupData(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {
	tablename := common.BKTableNamePropertyGroup
	blog.Errorf("add data for  %s table ", tablename)
	rows := getPropertyGroupData(conf.OwnerID)
	for _, row := range rows {
		if _, _, err := upgrader.Upsert(ctx, db, tablename, row, "id", []string{common.BKObjIDField, "bk_group_id"}, []string{"id"}); err != nil {
			blog.Errorf("add data for  %s table error  %s", tablename, err)
			return err
		}
	}
	return nil
}
func getObjectDesData(ownerID string) []*metadata.Object {

	dataRows := []*metadata.Object{
		&metadata.Object{IsCache: false, ObjCls: "lmanager", ObjectID: "region", ObjectName: "地域", ObjIcon: "icon-cc-default", IsPre: false, IsHidden: false, IsPaused: false},
		&metadata.Object{IsCache: false, ObjCls: "lmanager", ObjectID: "zone", ObjectName: "可用区", ObjIcon: "icon-cc-triangle", IsPre: false, IsHidden: false, IsPaused: false},
		&metadata.Object{IsCache: false, ObjCls: "lmanager", ObjectID: "cmdb_host", ObjectName: "主机", ObjIcon: "icon-cc-default", IsPre: false, IsHidden: false, IsPaused: false},
		&metadata.Object{IsCache: false, ObjCls: "lmanager", ObjectID: "cluster", ObjectName: "集群", ObjIcon: "icon-cc-default", IsPre: false, IsHidden: false, IsPaused: false},
		&metadata.Object{IsCache: false, ObjCls: "lmanager", ObjectID: "storageBackend", ObjectName: "存储后端", ObjIcon: "icon-cc-default", IsPre: false, IsHidden: false, IsPaused: false},
		&metadata.Object{IsCache: false, ObjCls: "lmanager", ObjectID: "storageBackendPool", ObjectName: "存储后端池", ObjIcon: "icon-cc-default", IsPre: false, IsHidden: false, IsPaused: false},
		//&metadata.Object{ObjCls: "HVAC", ObjectID: "water_cooled_chiller", ObjectName: "水冷型冷水机组", ObjIcon: "icon-cc-default", IsPre: false, IsHidden: false},
		//&metadata.Object{ObjCls: "HVAC", ObjectID: "refrigeration_pump", ObjectName: "冷冻泵", ObjIcon: "icon-cc-default", IsPre: false, IsHidden: false},
		//&metadata.Object{ObjCls: "power", ObjectID: "low_voltage_switchgear", ObjectName: "低压开关柜", ObjIcon: "icon-cc-default", IsPre: false, IsHidden: false},
		//&metadata.Object{ObjCls: "power", ObjectID: "modular_UPS", ObjectName: "模块化UPS", ObjIcon: "icon-cc-default", IsPre: false, IsHidden: false},
		//&metadata.Object{ObjCls: "spatial_relationship", ObjectID: "idc", ObjectName: "数据中心", ObjIcon: "icon-cc-idc", IsPre: false, IsHidden: false},
		//&metadata.Object{ObjCls: "spatial_relationship", ObjectID: "building", ObjectName: "楼栋", ObjIcon: "icon-cc-company", IsPre: false, IsHidden: false},
		//&metadata.Object{ObjCls: "spatial_relationship", ObjectID: "floor", ObjectName: "楼层", ObjIcon: "icon-cc-firewall", IsPre: false, IsHidden: false},
		//&metadata.Object{ObjCls: "spatial_relationship", ObjectID: "room", ObjectName: "机房", ObjIcon: "icon-cc-engine-room", IsPre: false, IsHidden: false},
		//&metadata.Object{ObjCls: "spatial_relationship", ObjectID: "line", ObjectName: "机列", ObjIcon: "icon-cc-engine-room", IsPre: false, IsHidden: false},
		//&metadata.Object{ObjCls: "spatial_relationship", ObjectID: "cabinet", ObjectName: "机柜", ObjIcon: "icon-cc-cabinet", IsPre: false, IsHidden: false},
		//&metadata.Object{ObjCls: "bk_host_manage", ObjectID: common.BKInnerObjIDHost, ObjectName: "主机", IsPre: true, ObjIcon: "icon-cc-host", Position: `{"bk_host_manage":{"x":-600,"y":-650}}`},
		//&metadata.Object{ObjCls: "bk_biz_topo", ObjectID: common.BKInnerObjIDModule, ObjectName: "模块", IsPre: true, ObjIcon: "icon-cc-module", Position: ``},
		//&metadata.Object{ObjCls: "bk_biz_topo", ObjectID: common.BKInnerObjIDSet, ObjectName: "集群", IsPre: true, ObjIcon: "icon-cc-set", Position: ``},
		//&metadata.Object{ObjCls: "bk_organization", ObjectID: common.BKInnerObjIDApp, ObjectName: "业务", IsPre: true, ObjIcon: "icon-cc-business", Position: `{"bk_organization":{"x":-100,"y":-100}}`},
		//&metadata.Object{ObjCls: "bk_host_manage", ObjectID: common.BKInnerObjIDProc, ObjectName: "进程", IsPre: true, ObjIcon: "icon-cc-process", Position: `{"bk_host_manage":{"x":-450,"y":-650}}`},
		//&metadata.Object{ObjCls: "bk_host_manage", ObjectID: common.BKInnerObjIDPlat, ObjectName: "云区域", IsPre: true, ObjIcon: "icon-cc-subnet", Position: `{"bk_host_manage":{"x":-600,"y":-500}}`},
		//&metadata.Object{ObjCls: "bk_network", ObjectID: common.BKInnerObjIDSwitch, ObjectName: "交换机", ObjIcon: "icon-cc-switch2", Position: `{"bk_network":{"x":-200,"y":-50}}`},
		//&metadata.Object{ObjCls: "bk_network", ObjectID: common.BKInnerObjIDRouter, ObjectName: "路由器", ObjIcon: "icon-cc-router", Position: `{"bk_network":{"x":-350,"y":-50}}`},
		//&metadata.Object{ObjCls: "bk_network", ObjectID: common.BKInnerObjIDBlance, ObjectName: "负载均衡", ObjIcon: "icon-cc-balance", Position: `{"bk_network":{"x":-500,"y":-50}}`},
		//&metadata.Object{ObjCls: "bk_network", ObjectID: common.BKInnerObjIDFirewall, ObjectName: "防火墙", ObjIcon: "icon-cc-firewall", Position: `{"bk_network":{"x":-650,"y":-50}}`},
	}
	t := metadata.Now()
	for _, r := range dataRows {
		r.CreateTime = &t
		r.LastTime = &t
		r.IsPaused = false
		r.Creator = common.CCSystemOperatorUserName
		r.OwnerID = ownerID
		r.Description = ""
		r.Modifier = ""
	}

	return dataRows
}

// Association for purpose of this structure not change by other, copy here
type Association struct {
	ID       int64  `field:"id" json:"id" bson:"id"`
	ObjectID string `field:"bk_obj_id" json:"bk_obj_id" bson:"bk_obj_id"`
	OwnerID  string `field:"bk_supplier_account" json:"bk_supplier_account" bson:"bk_supplier_account"`
	//AsstForward      string `field:"bk_asst_forward" json:"bk_asst_forward" bson:"bk_asst_forward"`
	AsstObjID string `field:"bk_asst_obj_id" json:"bk_asst_obj_id" bson:"bk_asst_obj_id"`
	AsstName  string `field:"bk_asst_name" json:"bk_asst_name" bson:"bk_asst_name"`
	//ObjectAttID      string `field:"bk_object_att_id" json:"bk_object_att_id" bson:"bk_object_att_id"`
	ClassificationID string `field:"bk_classification_id" bson:"-"`
	ObjectIcon       string `field:"bk_obj_icon" bson:"-"`
	ObjectName       string `field:"bk_obj_name" bson:"-"`
	IsPre            string `field:"ispre" bson:"-"`
	ObjAsstId        string `field:"bk_obj_asst_id" bson:"bk_obj_asst_id"`
	AsstId           string `field:"bk_asst_id" bson:"bk_asst_id"`
	Mapping          string `field:"mapping" bson:"mapping"`
}

func getAddAsstData(ownerID string) []Association {
	dataRows := []Association{
		//{ID: 1, OwnerID: ownerID, ObjectID: "jigui", AsstObjID: "jifang", ObjAsstId: "jigui_belong_jifang", AsstId: "belong", Mapping: "n:n"},
		{ID: 10, OwnerID: ownerID, ObjectID: "zone", AsstObjID: "region", ObjAsstId: "zone_connect_region", AsstId: "connect", Mapping: "n:n"},
		{ID: 11, OwnerID: ownerID, ObjectID: "cluster", AsstObjID: "zone", ObjAsstId: "cluster_connect_zone", AsstId: "connect", Mapping: "n:n"},
	}
	return dataRows
}

func getAddObjectBaseData(ownerID string) []map[string]interface{} {
	dataRows := make([]map[string]interface{}, 0)
	ct := new(time.Time)

	dataRows = []map[string]interface{}{
		{
			"bk_obj_id":                          "water_cooled_chiller",
			"refrigerating_capacity":             545,
			"rated_temperature_of_chilled_water": 7,
			"model":                              "DFDDG",
			"SN":                                 "12345789",
			"system_design_pressure":             3,
			"bk_inst_name":                       "1#冷水机组",
			"bk_supplier_account":                "0",
			"manufacturer":                       "约克",
			"rated_power":                        4000,
			"bk_inst_id":                         9,
			"create_time":                        ct,
			"last_time":                          ct,
		},
		{
			"bk_inst_name":        "1AA1变压器进线柜",
			"bk_supplier_account": "0",
			"manufacturer":        "施耐德",
			"rated_power":         1231,
			"bk_inst_id":          10,
			"output_voltage":      235,
			"bk_obj_id":           "low_voltage_switchgear",
			"nominal_voltage":     400,
			"model":               "SADG",
			"create_time":         ct,
			"last_time":           ct,
			"SN":                  "12456",
		},
		{
			"bk_inst_name":        "2#冷冻水泵",
			"bk_supplier_account": "0",
			"manufacturer":        "格兰富",
			"rated_speed":         532,
			"model":               "SAFsf",
			"bk_inst_id":          11,
			"nominal_voltage":     400,
			"SN":                  "123456",
			"bk_obj_id":           "refrigeration_pump",
			"create_time":         ct,
			"last_time":           ct,
		},
		{
			"bk_obj_id":               "building",
			"floor_high":              6,
			"seismic_grade":           "8",
			"bk_inst_id":              12,
			"area":                    3758.34,
			"bk_inst_name":            "1栋",
			"bk_supplier_account":     "0",
			"building_structure_type": "1",
			"name":                    "智航楼",
			"create_time":             ct,
			"last_time":               ct,
		},
		{

			"function":            "",
			"bk_supplier_account": "0",
			"build_time":          "2020-03-17",
			"handle_weight":       500,
			"name":                "1层",
			"prod_time":           "2021-03-09",
			"bk_inst_id":          13,
			"create_time":         ct,
			"last_time":           ct,
			"bk_inst_name":        "1层",
			"bk_obj_id":           "floor",
		},
		{
			"bk_obj_id":           "room",
			"is_cold_channel":     true,
			"name":                "001",
			"bk_inst_id":          14,
			"bk_inst_name":        "001室",
			"customer_name":       "",
			"create_time":         ct,
			"last_time":           ct,
			"bk_supplier_account": "0",
		},
		{
			"bk_obj_id":           "room",
			"is_cold_channel":     true,
			"name":                "002",
			"bk_inst_id":          15,
			"bk_inst_name":        "002室",
			"customer_name":       "",
			"create_time":         ct,
			"last_time":           ct,
			"bk_supplier_account": "0",
		},
	}
	var n int32 = 1
	for i := 16; i <= 20; i++ {
		dataRows = append(dataRows, map[string]interface{}{
			"bk_inst_name":        fmt.Sprintf("%d列", n),
			"bk_obj_id":           "line",
			"bk_supplier_account": "0",
			"manufacturer":        "华为",
			"bk_inst_id":          i,
			"create_time":         ct,
			"last_time":           ct,
		})
		atomic.AddInt32(&n, 1)
	}

	n = 1
	for ii := 21; ii <= 28; ii++ {
		dataRows = append(dataRows, map[string]interface{}{
			"create_time":         ct,
			"last_time":           ct,
			"bk_inst_name":        fmt.Sprintf("%d列", n),
			"bk_obj_id":           "cabinet",
			"bk_supplier_account": "0",
			"depth":               800,
			"hight":               2000,
			"bk_inst_id":          ii,
		})
		atomic.AddInt32(&n, 1)
	}
	//row, err := upgrader.JsonToMap(data)
	//if err == nil {

	//}

	return dataRows
}
func getAddUniqueData(ownerID string) []metadata.ObjectUnique {
	dataRows := []metadata.ObjectUnique{
		{
			ID:        18,
			ObjID:     "region",
			MustCheck: true,
			Keys: []metadata.UniqueKey{
				{
					Kind: metadata.UniqueKeyKindProperty,
					ID:   152,
				},
			},
			Ispre:    false,
			OwnerID:  ownerID,
			LastTime: metadata.Now(),
		},
		{
			ID:        25,
			ObjID:     "region",
			MustCheck: false,
			Keys: []metadata.UniqueKey{
				{
					Kind: metadata.UniqueKeyKindProperty,
					ID:   158,
				},
			},
			Ispre:    false,
			OwnerID:  ownerID,
			LastTime: metadata.Now(),
		},
		{
			ID:        26,
			ObjID:     "region",
			MustCheck: false,
			Keys: []metadata.UniqueKey{
				{
					Kind: metadata.UniqueKeyKindProperty,
					ID:   159,
				},
			},
			Ispre:    false,
			OwnerID:  ownerID,
			LastTime: metadata.Now(),
		},
		{
			ID:        27,
			ObjID:     "zone",
			MustCheck: true,
			Keys: []metadata.UniqueKey{
				{
					Kind: metadata.UniqueKeyKindProperty,
					ID:   190,
				},
			},
			Ispre:    false,
			OwnerID:  ownerID,
			LastTime: metadata.Now(),
		},
		{
			ID:        28,
			ObjID:     "zone",
			MustCheck: false,
			Keys: []metadata.UniqueKey{
				{
					Kind: metadata.UniqueKeyKindProperty,
					ID:   193,
				},
			},
			Ispre:    false,
			OwnerID:  ownerID,
			LastTime: metadata.Now(),
		},
		{
			ID:        30,
			ObjID:     "cmdb_host",
			MustCheck: true,
			Keys: []metadata.UniqueKey{
				{
					Kind: metadata.UniqueKeyKindProperty,
					ID:   300,
				},
			},
			Ispre:    false,
			OwnerID:  ownerID,
			LastTime: metadata.Now(),
		},
		{
			ID:        31,
			ObjID:     "cluster",
			MustCheck: true,
			Keys: []metadata.UniqueKey{
				{
					Kind: metadata.UniqueKeyKindProperty,
					ID:   210,
				},
			},
			Ispre:    false,
			OwnerID:  ownerID,
			LastTime: metadata.Now(),
		},
		{
			ID:        32,
			ObjID:     "storageBackend",
			MustCheck: true,
			Keys: []metadata.UniqueKey{
				{
					Kind: metadata.UniqueKeyKindProperty,
					ID:   270,
				},
			},
			Ispre:    false,
			OwnerID:  ownerID,
			LastTime: metadata.Now(),
		},
		{
			ID:        33,
			ObjID:     "storageBackendPool",
			MustCheck: true,
			Keys: []metadata.UniqueKey{
				{
					Kind: metadata.UniqueKeyKindProperty,
					ID:   250,
				},
			},
			Ispre:    false,
			OwnerID:  ownerID,
			LastTime: metadata.Now(),
		},
	}
	return dataRows
}

func getObjAttDescData(ownerID string) []*Attribute {

	predataRows := SetRow()
	//predataRows = append(predataRows, SetRow()...)
	//predataRows = append(predataRows, ModuleRow()...)
	//predataRows = append(predataRows, HostRow()...)
	//predataRows = append(predataRows, ProcRow()...)
	//predataRows = append(predataRows, PlatRow()...)

	//dataRows := SwitchRow()
	//dataRows = append(dataRows, RouterRow()...)
	//dataRows = append(dataRows, LoadBalanceRow()...)
	//dataRows = append(dataRows, FirewallRow()...)

	t := new(time.Time)
	*t = time.Now()
	for _, r := range predataRows {
		r.OwnerID = ownerID
		if false != r.IsPre {
			r.IsPre = true
		}
		//if false != r.IsEditable {
		//	r.IsEditable = true
		//}
		//if true != r.IsSystem {
		//	r.IsEditable = false
		//}
		// xxx edit
		r.IsEditable = true
		r.IsReadOnly = false
		r.CreateTime = t
		//r.Creator = common.CCSystemOperatorUserName
		r.Creator = common.CCSystemUserName
		r.LastTime = r.CreateTime
		r.Description = ""
		r.IsSystem = false
		r.BizID = 0
	}
	//for _, r := range dataRows {
	//	r.OwnerID = ownerID
	//	if false != r.IsEditable {
	//		r.IsEditable = true
	//	}
	//	r.IsReadOnly = false
	//	r.CreateTime = t
	//	r.Creator = common.CCSystemOperatorUserName
	//	r.LastTime = r.CreateTime
	//	r.Description = ""
	//}

	return append(predataRows)
}

func getPropertyGroupData(ownerID string) []*metadata.Group {
	objectIDs := make(map[string]map[string]string)

	dataRows := []*metadata.Group{
		//netdata
		&metadata.Group{BizID: 0, ObjectID: "region", GroupID: "default", GroupName: "Default", GroupIndex: -1, OwnerID: ownerID, IsDefault: true, IsPre: false, IsCollapse: false, ID: 20},

		&metadata.Group{BizID: 0, ObjectID: "zone", GroupID: "default", GroupName: "Default", GroupIndex: -1, OwnerID: ownerID, IsDefault: true, IsPre: false, IsCollapse: false, ID: 21},

		&metadata.Group{BizID: 0, ObjectID: "storageBackend", GroupID: "default", GroupName: "Default", GroupIndex: -1, OwnerID: ownerID, IsDefault: true, IsPre: false, IsCollapse: false, ID: 22},

		&metadata.Group{BizID: 0, ObjectID: "storageBackendPool", GroupID: "default", GroupName: "Default", GroupIndex: -1, OwnerID: ownerID, IsDefault: true, IsPre: false, IsCollapse: false, ID: 23},
		&metadata.Group{BizID: 0, ObjectID: "cmdb_host", GroupID: "default", GroupName: "Default", GroupIndex: -1, OwnerID: ownerID, IsDefault: true, IsPre: false, IsCollapse: false, ID: 24},
		&metadata.Group{BizID: 0, ObjectID: "cluster", GroupID: "default", GroupName: "Default", GroupIndex: -1, OwnerID: ownerID, IsDefault: true, IsPre: false, IsCollapse: false, ID: 25},
		//&metadata.Group{BizID: 0, ObjectID: "water_cooled_chiller", GroupID: "default", GroupName: "基础信息", GroupIndex: -1, OwnerID: ownerID, IsDefault: true, IsPre: false, IsCollapse: false, ID: 18},
		//&metadata.Group{BizID: 0, ObjectID: "low_voltage_switchgear", GroupID: "default", GroupName: "基础信息", GroupIndex: -1, OwnerID: ownerID, IsDefault: true, IsPre: false, IsCollapse: false, ID: 36},
		//&metadata.Group{BizID: 0, ObjectID: "water_cooled_chiller", GroupID: "1b4717fe-7e71-4fd9-a725-70d454b80803", GroupName: "运维信息", GroupIndex: -1, OwnerID: ownerID, IsDefault: false, IsPre: false, IsCollapse: false, ID: 37},
		//
		//&metadata.Group{BizID: 0, ObjectID: "low_voltage_switchgear", GroupID: "1b99731e-bbb3-4258-b7fc-2cc3be29185b", GroupName: "运维信息", GroupIndex: -1, OwnerID: ownerID, IsDefault: false, IsPre: false, IsCollapse: false, ID: 38},
		//
		//&metadata.Group{BizID: 0, ObjectID: "idc", GroupID: "default", GroupName: "基础信息", GroupIndex: -1, OwnerID: ownerID, IsDefault: true, IsPre: false, IsCollapse: false, ID: 39},
		//&metadata.Group{BizID: 0, ObjectID: "building", GroupID: "default", GroupName: "基础信息", GroupIndex: -1, OwnerID: ownerID, IsDefault: true, IsPre: false, IsCollapse: false, ID: 40},
		//
		//&metadata.Group{BizID: 0, ObjectID: "refrigeration_pump", GroupID: "default", GroupName: "基础信息", GroupIndex: -1, OwnerID: ownerID, IsDefault: true, IsPre: false, IsCollapse: false, ID: 41},
		//&metadata.Group{BizID: 0, ObjectID: "refrigeration_pump", GroupID: "18746577-f5f5-4b8f-972f-82d24a3a5139", GroupName: "运维信息", GroupIndex: -1, OwnerID: ownerID, IsDefault: false, IsPre: false, IsCollapse: false, ID: 42},
		//
		//&metadata.Group{BizID: 0, ObjectID: "floor", GroupID: "default", GroupName: "基础信息", GroupIndex: -1, OwnerID: ownerID, IsDefault: true, IsPre: false, IsCollapse: false, ID: 43},
		//&metadata.Group{BizID: 0, ObjectID: "modular_UPS", GroupID: "default", GroupName: "基础信息", GroupIndex: -1, OwnerID: ownerID, IsDefault: true, IsPre: false, IsCollapse: false, ID: 44},
		//&metadata.Group{BizID: 0, ObjectID: "room", GroupID: "default", GroupName: "基础信息", GroupIndex: -1, OwnerID: ownerID, IsDefault: true, IsPre: false, IsCollapse: false, ID: 45},
		//
		//&metadata.Group{BizID: 0, ObjectID: "modular_UPS", GroupID: "fc2a4990-5750-4cbe-b402-081864185a7c", GroupName: "运维信息", GroupIndex: -1, OwnerID: ownerID, IsDefault: false, IsPre: false, IsCollapse: false, ID: 46},
		//&metadata.Group{BizID: 0, ObjectID: "line", GroupID: "default", GroupName: "基础信息", GroupIndex: -1, OwnerID: ownerID, IsDefault: true, IsPre: false, IsCollapse: false, ID: 47},
		//&metadata.Group{BizID: 0, ObjectID: "cabinet", GroupID: "default", GroupName: "基础信息", GroupIndex: -1, OwnerID: ownerID, IsDefault: true, IsPre: false, IsCollapse: false, ID: 48},
		//
		//

		//app
		//&metadata.Group{ObjectID: common.BKInnerObjIDApp, GroupID: mCommon.BaseInfo, GroupName: mCommon.BaseInfoName, GroupIndex: 1, OwnerID: ownerID, IsDefault: true},
		//&metadata.Group{ObjectID: common.BKInnerObjIDApp, GroupID: mCommon.AppRole, GroupName: mCommon.AppRoleName, GroupIndex: 2, OwnerID: ownerID, IsDefault: true},
		//
		////set
		//&metadata.Group{ObjectID: common.BKInnerObjIDSet, GroupID: mCommon.BaseInfo, GroupName: mCommon.BaseInfoName, GroupIndex: 1, OwnerID: ownerID, IsDefault: true},
		//
		////module
		//&metadata.Group{ObjectID: common.BKInnerObjIDModule, GroupID: mCommon.BaseInfo, GroupName: mCommon.BaseInfoName, GroupIndex: 1, OwnerID: ownerID, IsDefault: true},
		//
		////host
		//&metadata.Group{ObjectID: common.BKInnerObjIDHost, GroupID: mCommon.BaseInfo, GroupName: mCommon.BaseInfoName, GroupIndex: 1, OwnerID: ownerID, IsDefault: true},
		//&metadata.Group{ObjectID: common.BKInnerObjIDHost, GroupID: mCommon.HostAutoFields, GroupName: mCommon.HostAutoFieldsName, GroupIndex: 3, OwnerID: ownerID, IsDefault: true},
		//
		////proc
		//&metadata.Group{ObjectID: common.BKInnerObjIDProc, GroupID: mCommon.BaseInfo, GroupName: mCommon.BaseInfoName, GroupIndex: 1, OwnerID: ownerID, IsDefault: true},
		//&metadata.Group{ObjectID: common.BKInnerObjIDProc, GroupID: mCommon.ProcPort, GroupName: mCommon.ProcPortName, GroupIndex: 2, OwnerID: ownerID, IsDefault: true},
		//&metadata.Group{ObjectID: common.BKInnerObjIDProc, GroupID: mCommon.ProcGsekitBaseInfo, GroupName: mCommon.ProcGsekitBaseInfoName, GroupIndex: 3, OwnerID: ownerID, IsDefault: true},
		//&metadata.Group{ObjectID: common.BKInnerObjIDProc, GroupID: mCommon.ProcGsekitManageInfo, GroupName: mCommon.ProcGsekitManageInfoName, GroupIndex: 4, OwnerID: ownerID, IsDefault: true},
		//
		////plat
		//&metadata.Group{ObjectID: common.BKInnerObjIDPlat, GroupID: mCommon.BaseInfo, GroupName: mCommon.BaseInfoName, GroupIndex: 1, OwnerID: ownerID, IsDefault: true},
		//
		////bk_switch
		//&metadata.Group{ObjectID: common.BKInnerObjIDSwitch, GroupID: mCommon.BaseInfo, GroupName: mCommon.BaseInfoName, GroupIndex: 1, OwnerID: ownerID, IsDefault: true},
		//&metadata.Group{ObjectID: common.BKInnerObjIDRouter, GroupID: mCommon.BaseInfo, GroupName: mCommon.BaseInfoName, GroupIndex: 1, OwnerID: ownerID, IsDefault: true},
		//&metadata.Group{ObjectID: common.BKInnerObjIDBlance, GroupID: mCommon.BaseInfo, GroupName: mCommon.BaseInfoName, GroupIndex: 1, OwnerID: ownerID, IsDefault: true},
		//&metadata.Group{ObjectID: common.BKInnerObjIDFirewall, GroupID: mCommon.BaseInfo, GroupName: mCommon.BaseInfoName, GroupIndex: 1, OwnerID: ownerID, IsDefault: true},
	}
	for objID, kv := range objectIDs {
		index := int64(1)
		for id, name := range kv {
			row := &metadata.Group{ObjectID: objID, GroupID: id, GroupName: name, GroupIndex: index, OwnerID: ownerID, IsDefault: true}
			dataRows = append(dataRows, row)
			index++
		}

	}

	return dataRows

}

var classificationRows = []*metadata.Classification{
	//&metadata.Classification{ClassificationID: "bk_host_manage", ClassificationName: "主机管理", ClassificationType: "inner", ClassificationIcon: "icon-cc-host"},
	//&metadata.Classification{ClassificationID: "bk_biz_topo", ClassificationName: "业务拓扑", ClassificationType: "inner", ClassificationIcon: "icon-cc-business"},
	//&metadata.Classification{ClassificationID: "bk_organization", ClassificationName: "组织架构", ClassificationType: "inner", ClassificationIcon: "icon-cc-organization"},
	//&metadata.Classification{ClassificationID: "bk_network", ClassificationName: "网络", ClassificationType: "inner", ClassificationIcon: "icon-cc-network-equipment"},
	//&metadata.Classification{ClassificationID: "HVAC", ClassificationName: "暖通", ClassificationType: "inner", ClassificationIcon: "icon-cc-business"},
	//&metadata.Classification{ClassificationID: "power", ClassificationName: "电力", ClassificationType: "inner", ClassificationIcon: "icon-cc-business"},
	//&metadata.Classification{ClassificationID: "spatial_relationship", ClassificationName: "空间关系", ClassificationType: "inner", ClassificationIcon: "icon-cc-network-equipment"},
	&metadata.Classification{ClassificationID: "lmanager", ClassificationName: "lmanager", ClassificationType: "inner", ClassificationIcon: "icon-cc-business"},
}
