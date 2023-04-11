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

package local

import (
	"configcenter/src/common"
	"configcenter/src/common/blog"
	"configcenter/src/common/metadata"
	"configcenter/src/source_controller/coreservice/bussiness/transactionBus"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"time"
)

func (c *Mongo) StartTransaction(ctx context.Context, cap *metadata.TxnCapable) error {
	wc := writeconcern.New(writeconcern.WMajority())
	rc := readconcern.Snapshot()
	//var maxCommitTime = cap.Timeout
	//txnOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc).SetMaxCommitTime(&maxCommitTime)
	// 将会话ID转换为bsoncore.Document类型

	sessionOpts := &options.SessionOptions{}
	sessionOpts.SetDefaultWriteConcern(wc).SetDefaultReadConcern(rc).SetDefaultMaxCommitTime(&cap.Timeout)

	session1, err := c.Client().StartSession(sessionOpts)
	if err != nil {
		return err
	}
	sessionContext := mongo.NewSessionContext(ctx, session1)
	if err = sessionContext.StartTransaction(); err != nil {
		return err
	}
	txnNumber, err := c.tm.GenTxnNumber(cap.SessionID, cap.Timeout)
	if err != nil {
		return err
	}
	sessTxnCtx := &metadata.SessionTxnCtx{
		SessionContext: sessionContext,
		Time:           time.Now(),
		TxnNumber:      txnNumber,
		SessionID:      session1.ID().String(),
	}

	//使用topo 生成的sessionID
	transactionBus.TxnMapSess.Set(cap.SessionID, sessTxnCtx)

	//fmt.Println(txnNumber)
	return nil
}

func (c *Mongo) ExportTransactionSess(ctx context.Context, cap *metadata.TxnCapable) (mongo.Session, error) {
	val, exist := transactionBus.TxnMapSess.Get(cap.SessionID)
	if !exist {
		err := errors.New(fmt.Sprintf("Not Found Sessionid:%v ", cap.SessionID))
		//blog.Errorf("commit transaction, but prepare transaction failed, err: %v, rid: %v", err.Error(), rid)
		return nil, err
	}

	if _, ok := val.(*metadata.SessionTxnCtx); !ok {
		err := errors.New(fmt.Sprintf("Sessionid:%v convert type err", cap.SessionID))
		//blog.Errorf("commit transaction, but prepare transaction failed, err: %v, rid: %v", err.Error(), rid)
		return nil, err
	}

	session, _ := val.(*metadata.SessionTxnCtx)

	reloadSession := mongo.SessionFromContext(session.SessionContext)
	return reloadSession, nil
}

// CommitTransaction 提交事务
func (c *Mongo) CommitTransaction(ctx context.Context, cap *metadata.TxnCapable) error {
	rid := ctx.Value(common.ContextRequestIDField)

	//reloadSession, err := c.tm.PrepareTransaction(cap, c.dbc)
	//if err != nil {
	//	blog.Errorf("commit transaction, but prepare transaction failed, err: %v, rid: %v", err, rid)
	//	return err
	//}
	reloadSession, err := c.ExportTransactionSess(ctx, cap)
	if err != nil {
		blog.Errorf("commit transaction, but prepare transaction failed, err: %v, rid: %v", err.Error(), rid)
	}

	// reset the transaction state, so that we can commit the transaction after start the
	// transaction immediately.
	mongo.CmdbPrepareCommitOrAbort(reloadSession)

	// we commit the transaction with the session id
	err = reloadSession.CommitTransaction(ctx)
	if err != nil {
		return fmt.Errorf("commit transaction: %s failed, err: %v, rid: %v", cap.SessionID, err, rid)
	}

	err = c.tm.RemoveSessionKey(cap.SessionID)
	if err != nil {
		// this key has ttl, it's ok if we not delete it, cause this key has a ttl.
		blog.Errorf("commit transaction, but delete txn session: %s key failed, err: %v, rid: %v", cap.SessionID, err, rid)
		// do not return.
	}
	transactionBus.TxnMapSess.Remove(cap.SessionID)

	return nil
}

// AbortTransaction 取消事务
func (c *Mongo) AbortTransaction(ctx context.Context, cap *metadata.TxnCapable) (bool, error) {
	rid := ctx.Value(common.ContextRequestIDField)
	//reloadSession, err := c.tm.PrepareTransaction(cap, c.dbc)
	//if err != nil {
	//	blog.Errorf("abort transaction, but prepare transaction failed, err: %v, rid: %v", err, rid)
	//	return false, err
	//}
	reloadSession, err := c.ExportTransactionSess(ctx, cap)
	if err != nil {
		blog.Errorf("abort transaction, but prepare transaction failed, err: %v, rid: %v", err.Error(), rid)
	}
	// reset the transaction state, so that we can abort the transaction after start the
	// transaction immediately.
	mongo.CmdbPrepareCommitOrAbort(reloadSession)

	// we abort the transaction with the session id
	err = reloadSession.AbortTransaction(ctx)
	if err != nil {
		return false, fmt.Errorf("abort transaction: %s failed, err: %v, rid: %v", cap.SessionID, err, rid)
	}

	err = c.tm.RemoveSessionKey(cap.SessionID)
	if err != nil {
		// this key has ttl, it's ok if we not delete it, cause this key has a ttl.
		blog.Errorf("abort transaction, but delete txn session: %s key failed, err: %v, rid: %v", cap.SessionID, err, rid)
		// do not return.
	}
	transactionBus.TxnMapSess.Remove(cap.SessionID)
	errorType := c.tm.GetTxnError(sessionKey(cap.SessionID))
	switch errorType {
	// retry when the transaction error type is write conflict, which means the transaction conflicts with another one
	case WriteConflictType:
		return true, nil
	}

	return false, nil
}
