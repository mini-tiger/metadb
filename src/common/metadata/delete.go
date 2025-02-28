/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.,
 * Copyright (C) 2017,-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the ",License",); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an ",AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package metadata

import "configcenter/src/common/mapstr"

// DeleteOption common delete condition options
type DeleteOption struct {
	Condition mapstr.MapStr `json:"condition"`
}

// DeletedCountResult delete  api http response return result struct
type DeletedOptionResult struct {
	BaseResp `json:",inline"`
	Data     DeletedCount `json:"data"`
}

type DelAndInsertOption struct {
	Condition mapstr.MapStr   `json:"condition"`
	Datas     []mapstr.MapStr `json:"datas"`
}
