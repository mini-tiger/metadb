package transactionBus

import (
	"configcenter/src/common/blog"
	"configcenter/src/common/mapstr"
	"configcenter/src/common/metadata"
	"fmt"
	"github.com/pkg/errors"
	"time"
)

/**
 * @Author: Tao Jun
 * @Since: 2023/3/8
 * @Desc: session.go
**/

var TxnMapSess mapstr.MapStr = make(map[string]interface{}, 0)

const maxtime float64 = 86400

func PrintTxnMap() {
	go func() {
		for {
			blog.Infof("TxnMapSess %+v", TxnMapSess)
			time.Sleep(1 * time.Hour)
		}
	}()
}

// GC
func CleanTimeOutTxn() {
	clean := func(key string, val interface{}) error {
		if txnctx, ok := val.(*metadata.SessionTxnCtx); ok {
			if time.Now().Sub(txnctx.Time).Seconds()-maxtime > 0 {
				blog.Errorf("GC CleanTimeOut session id : %v", key)
				TxnMapSess.Remove(key)
				return nil
			}
		} else {
			blog.Errorf("GC Can't Convert session id : %v", key)
			TxnMapSess.Remove(key)
			return errors.New(fmt.Sprintf("clean process convert type err,del key %v", key))
		}
		return nil
	}
	go func() {
		for {
			//a := &metadata.SessionTxnCtx{
			//	SessionContext: nil,
			//	SessionID:      "",
			//	Time:           time.Now().AddDate(0, 0, -1),
			//	TxnNumber:      0,
			//}
			//TxnMapSess.Set("abc", a)
			time.Sleep(1 * time.Hour)

			TxnMapSess.ForEach(clean)

		}
	}()
}
