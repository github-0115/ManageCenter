package db

import (
	c "ManageCenter/config"
	"fmt"

	log "github.com/inconshreveable/log15"
	"gopkg.in/mgo.v2"
)

var (
	Session   *mgo.Session
	DBUri     string
	DBUserUri string
	Stats     *MongoSession
	Manager   *MongoSession
	User      *MongoSession
	Demo      *MongoSession
)

func init() {
	DBUri = c.DBCfg.ManageCenterMongoTask.DB
	DBUserUri = c.DBCfg.UserCenterMongoTask.DB
	mongoUrl := c.DBCfg.ManageCenterMongoTask.String()

	var err error
	Session, err = mgo.Dial(mongoUrl)
	if err != nil {
		log.Error(fmt.Sprintf("connect Mongo error", err))
		panic(err)
	}

	Stats = new(MongoSession).Init(
		c.DBCfg.StatsMongo,
		"porn_hour",
	)
	Manager = new(MongoSession).Init(
		c.DBCfg.ManageCenterMongoTask,
		"manager",
	)
	User = new(MongoSession).Init(
		c.DBCfg.UserCenterMongoTask,
		"user",
	)
	Demo = new(MongoSession).Init(
		c.DBCfg.DemoMongo,
		"demo",
	)
}
