package service

import (
	"configcenter/src/common"
	"configcenter/src/common/blog"
	"configcenter/src/common/metadata"
	"configcenter/src/common/util"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/scene_server/admin_server/upgrader/y3.9.209912302358"
	"errors"
	"github.com/emicklei/go-restful"
	"net/http"
	"time"
)

/**
 * @Author: Tao Jun
 * @Since: 2023/2/26
 * @Desc: initdata.go
**/
type initData struct {
	ID          interface{} `bson:"_id,omitempty"`
	VersionDate time.Time   `json:"VersionDate"`
}

func (s *Service) FindInitBaseData(req *restful.Request, resp *restful.Response) {
	//rHeader := req.Request.Header
	//rid := util.GetHTTPCCRequestID(rHeader)
	//defErr := s.CCErr.CreateDefaultCCErrorIf(util.GetLanguage(rHeader))

	// 判断初始化metadb basedata finish
	cond := map[string]interface{}{"current_version": "y3.9.209912302358"}

	data, err := s.db.Table(common.BKTableNameSystem).Find(cond).Count(s.ctx)
	if err != nil {
		_ = resp.WriteError(http.StatusBadRequest, err)
		return
	}

	if data == 0 {
		_ = resp.WriteError(http.StatusBadRequest, errors.New("data len 0, init Metadb Base Data not finish"))
		return
	}
	_ = resp.WriteEntity(metadata.NewSuccessResp("success"))

}

// SearchConfigAdmin search the config
func (s *Service) FindCollectionVer(req *restful.Request, resp *restful.Response) {
	rHeader := req.Request.Header
	rid := util.GetHTTPCCRequestID(rHeader)
	defErr := s.CCErr.CreateDefaultCCErrorIf(util.GetLanguage(rHeader))

	cond := map[string]interface{}{
		"_id": req.PathParameter("collection"),
	}

	//var initDataCount &uint64
	initDataCount, err := s.db.Table(common.BKTableNameInitDataVer).Find(cond).Count(s.ctx)

	if err != nil {
		blog.Errorf("FindCollectionVer failed, err: %+v, rid: %s", err, rid)
		result := &metadata.RespError{
			Msg: defErr.Error(common.CCErrCommDBSelectFailed),
		}
		_ = resp.WriteError(http.StatusBadRequest, result)
		return
	}
	//fmt.Println(initDataCount)
	var insertData = initData{
		ID:          cond["_id"],
		VersionDate: time.Now(),
	}
	//map[string]interface{}{"_id": cond["_id"], "version": time.Now()}
	if initDataCount == 0 {
		s.db.Table(common.BKTableNameInitDataVer).Insert(s.ctx, &insertData)
		_ = resp.WriteEntity(metadata.NewSuccessResp("not found verion,mongoimport exec"))
		return
	}

	// 已经有插入记录
	_ = resp.WriteError(http.StatusCreated, errors.New("found version,mongoimport skip"))
	return

}

func (s *Service) UpdateSequenceView(req *restful.Request, resp *restful.Response) {
	rHeader := req.Request.Header
	ownerid := util.GetOwnerID(rHeader)
	user := util.GetUser(rHeader)

	// 将id 列的类型转换为 long
	y3_9_209912302358.UpdateFieldTypeLong(s.ctx, s.db, &upgrader.Config{
		OwnerID: ownerid,
		User:    user,
	})

	err := y3_9_209912302358.UpdateSequence(s.ctx, s.db, &upgrader.Config{
		OwnerID: ownerid,
		User:    user,
	})
	if err != nil {
		_ = resp.WriteError(http.StatusBadRequest, err)
		return
	}
	_ = resp.WriteEntity(metadata.NewSuccessResp("success"))
	return

}
