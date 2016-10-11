package db

import (
	"fmt"

	c "ManageCenter/config"

	log "github.com/inconshreveable/log15"
	"gopkg.in/mgo.v2"
)

type MongoSession struct {
	MongoUri   string
	DB         string
	Collection string
	Session    *mgo.Session
}

func (t *MongoSession) Init(cfg *c.MongoCfg, collection string) *MongoSession {
	t.MongoUri = cfg.String()
	t.DB = cfg.DB
	t.Collection = collection

	t.InitSession()
	return t
}

func (t *MongoSession) InitSession() {
	var err error
	t.Session, err = mgo.Dial(t.MongoUri)
	if err != nil {
		log.Error(fmt.Sprintf("connect %s  Mongo  error: %s", t.MongoUri, err.Error()))
		panic(err)
	}
}

func (t *MongoSession) GetSession() *mgo.Session {
	return t.Session.Clone()
}

func (t *MongoSession) CreateIndex(index *mgo.Index) *MongoSession {
	s := t.GetSession()
	defer s.Close()

	db := s.DB(t.DB)
	err := db.C(t.Collection).EnsureIndex(*index)
	if err != nil {
		log.Error(fmt.Sprintf("build  %s  Index  error: %s", t.MongoUri, err.Error()))
		panic(err)
	}
	return t
}
