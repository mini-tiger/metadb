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

package y3_9_202203171605

import (
	"context"

	"configcenter/src/common"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
	"configcenter/src/storage/dal/types"

	"gopkg.in/mgo.v2"
)

func createTable(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	for tableName, indexes := range tables {
		exists, err := db.HasTable(ctx, tableName)
		if err != nil {
			return err
		}
		if !exists {
			if err = db.CreateTable(ctx, tableName); err != nil && !mgo.IsDup(err) {
				return err
			}
		}
		for index := range indexes {
			if err = db.Table(tableName).CreateIndex(ctx, indexes[index]); err != nil && !db.IsDuplicatedError(err) {
				return err
			}
		}
	}
	return nil
}

var tables = map[string][]types.Index{

	common.BKTableNameObjAsst: {
		types.Index{Name: "", Keys: map[string]int32{common.BKObjIDField: 1}, Background: true},
		types.Index{Name: "", Keys: map[string]int32{common.BKAsstObjIDField: 1}, Background: true},
		types.Index{Name: "", Keys: map[string]int32{common.BkSupplierAccount: 1}, Background: true},
	},
	common.BKTableNameObjAttDes: {
		types.Index{Name: "", Keys: map[string]int32{common.BKObjIDField: 1}, Background: true},
		types.Index{Name: "", Keys: map[string]int32{common.BkSupplierAccount: 1}, Background: true},
		types.Index{Name: "", Keys: map[string]int32{common.BKFieldID: 1}, Background: true},
	},
	common.BKTableNameObjClassification: {
		types.Index{Name: "", Keys: map[string]int32{common.BKClassificationIDField: 1}, Background: true},
		types.Index{Name: "", Keys: map[string]int32{common.BKClassificationNameField: 1}, Background: true},
	},
	common.BKTableNameObjDes: {
		types.Index{Name: "", Keys: map[string]int32{common.BKObjIDField: 1}, Background: true},
		types.Index{Name: "", Keys: map[string]int32{common.BKClassificationIDField: 1}, Background: true},
		types.Index{Name: "", Keys: map[string]int32{common.BKObjNameField: 1}, Background: true},
		types.Index{Name: "", Keys: map[string]int32{common.BkSupplierAccount: 1}, Background: true},
	},

	common.BKTableNamePropertyGroup: {
		types.Index{Name: "", Keys: map[string]int32{common.BKObjIDField: 1}, Background: true},
		types.Index{Name: "", Keys: map[string]int32{common.BkSupplierAccount: 1}, Background: true},
		types.Index{Name: "", Keys: map[string]int32{common.BKPropertyGroupIDField: 1}, Background: true},
	},
	common.BKTableNameObjUnique: {
		types.Index{
			Keys:       map[string]int32{common.BKFieldID: int32(1)},
			Unique:     true,
			Background: true,
			Name:       "idx_unique_id",
		},
		types.Index{Name: "", Keys: map[string]int32{common.BKObjIDField: 1}, Background: true},
	},
}
