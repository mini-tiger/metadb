// Copyright (C) MongoDB, Inc. 2017-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package session

import (
	"time"

	"go.mongodb.org/mongo-driver/internal/uuid"
	"go.mongodb.org/mongo-driver/mongo/description"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

// Server is an open session with the server.
type Server struct {
	SessionID bsoncore.Document
	TxnNumber int64
	LastUsed  time.Time
	Dirty     bool
}

// returns whether or not a session has expired given a timeout in minutes
// a session is considered expired if it has less than 1 minute left before becoming stale
func (ss *Server) expired(topoDesc topologyDescription) bool {
	// There is no server monitoring in LB mode, so we do not track session timeout minutes from server hello responses
	// and never consider sessions to be expired.
	if topoDesc.kind == description.LoadBalanced {
		return false
	}

	if topoDesc.timeoutMinutes <= 0 {
		return true
	}
	timeUnused := time.Since(ss.LastUsed).Minutes()
	return timeUnused > float64(topoDesc.timeoutMinutes-1)
}

// update the last used time for this session.
// must be called whenever this server session is used to send a command to the server.
func (ss *Server) updateUseTime() {
	ss.LastUsed = time.Now()
}

func GetUUID() (id uuid.UUID,err error) {
	id, err = uuid.New()
	return
}
// 生成一个新的Session ID
// 从给定的字节数组中生成一个新的Session ID
func GenerateSessionIDFromBytes(b []byte) bsoncore.Document {
	var id []byte
	idx,id := bsoncore.AppendDocumentStart(id)
	//id = bsoncore.AppendString(id, "_id")
	id = bsoncore.AppendBinaryElement(id,"id", UUIDSubtype,b)
	id, _ = bsoncore.AppendDocumentEnd(id, idx)
	return bsoncore.Document(id)
}

func newServerSession() (*Server, error) {
	id, err := uuid.New()
	if err != nil {
		return nil, err
	}

	idx, idDoc := bsoncore.AppendDocumentStart(nil)
	idDoc = bsoncore.AppendBinaryElement(idDoc, "id", UUIDSubtype, id[:])
	idDoc, _ = bsoncore.AppendDocumentEnd(idDoc, idx)

	return &Server{
		SessionID: idDoc,
		LastUsed:  time.Now(),
	}, nil
}


func NewServerSession() (*Server, error) {
	id, err := uuid.New()
	if err != nil {
		return nil, err
	}

	idx, idDoc := bsoncore.AppendDocumentStart(nil)
	idDoc = bsoncore.AppendBinaryElement(idDoc, "id", UUIDSubtype, id[:])
	idDoc, _ = bsoncore.AppendDocumentEnd(idDoc, idx)

	return &Server{
		SessionID: idDoc,
		LastUsed:  time.Now(),
	}, nil
}

//func GetServerSession(id []byte) (*Server, error) {
//
//
//	idx, idDoc := bsoncore.AppendDocumentStart(nil)
//	idDoc = bsoncore.AppendBinaryElement(idDoc, "id", UUIDSubtype, id[:])
//	idDoc, _ = bsoncore.AppendDocumentEnd(idDoc, idx)
//
//	return &Server{
//		SessionID: idDoc,
//		LastUsed:  time.Now(),
//	}, nil
//}

//func GetRealodServerSession(id []byte) (*Server, error) {
//
//
//	idx, idDoc := bsoncore.AppendDocumentElementStart(id,"id")
//	idDoc = bsoncore.AppendBinaryElement(idDoc, "id", UUIDSubtype, id[:])
//	idDoc, _ = bsoncore.AppendDocumentEnd(idDoc, idx)
//
//	return &Server{
//		SessionID: idDoc,
//		LastUsed:  time.Now(),
//	}, nil
//}

// IncrementTxnNumber increments the transaction number.
func (ss *Server) IncrementTxnNumber() {
	ss.TxnNumber++
}

// MarkDirty marks the session as dirty.
func (ss *Server) MarkDirty() {
	ss.Dirty = true
}

// UUIDSubtype is the BSON binary subtype that a UUID should be encoded as
const UUIDSubtype byte = 4
