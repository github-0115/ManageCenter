package detailmodel

import (
	db "ManageCenter/service/db"
	"fmt"
	"time"

	log "github.com/inconshreveable/log15"
	"gopkg.in/mgo.v2/bson"
)

type Detail struct {
	CreatedAt time.Time `bson:"created_at"`
	User      string    `bson:"user"`
	Scene     string    `bson:"scene"`
	Ctype     []string  `bson:"ctype"`
	Result    *Result   `bson:"result"`
}

type Result struct {
	Md5        string      `bson:"md5"`
	Version    string      `bson:"version"`
	Tags       *Tags       `bson:"tags"`
	Porn       *Porn       `bson:"porn"`
	Screenshot *Screenshot `bson:"screenshot"`
	Live       *Live       `bson:"live"`
}

type Tags struct {
	P         float32 `bson:"p"`
	S         float32 `bson:"s"`
	N         float32 `bson:"n"`
	IsScreen  float32 `bson:"isScreen"`
	NotScreen float32 `bson:"notScreen"`
}

type Porn struct {
	Res    string `bson:"res"`
	Review bool   `bson:"review"`
}

type Screenshot struct {
	Res string `bson:"res"`
}

type Live struct {
	StreamId string `bson:"stream_id"`
	FrameId  string `bson:"frame_id"`
}

func GetDetail(username string, startTime time.Time, endTime time.Time, rType string, review bool, sort string, sord int, pageIndex int, pageSize int) ([]*Detail, error) {
	s := db.Stats.GetSession()
	defer s.Close()
	var result []*Detail
	var sortString string
	if sord == -1 {
		if sort == "created_at" {
			sortString = "-created_at"
		} else {
			sortString = "-result.tags." + rType
		}
	} else {
		if sort == "created_at" {
			sortString = "created_at"
		} else {
			sortString = "result.tags." + rType
		}
	}
	err := s.DB(db.Stats.DB).C("detail").Find(bson.M{
		"user": username,
		"created_at": bson.M{
			"$gte": startTime, "$lt": endTime},
		"result.porn.res":    rType,
		"result.porn.review": review,
	}).Sort(sortString).Skip((pageIndex - 1) * pageSize).Limit(pageSize).All(&result)
	if err != nil {
		log.Error(fmt.Sprintf("find detail err ", err))
		return nil, err
	}
	return result, nil
}
func GetDetailCount(username string, startTime time.Time, endTime time.Time, rType string, review bool) (int, error) {
	s := db.Stats.GetSession()
	defer s.Close()
	count, err := s.DB(db.Stats.DB).C("detail").Find(bson.M{
		"user": username,
		"created_at": bson.M{
			"$gte": startTime, "$lt": endTime},
		"result.porn.res":    rType,
		"result.porn.review": review,
	}).Count()

	if err != nil {
		log.Error(fmt.Sprintf("find detail err ", err))
		return 0, err
	}
	return count, nil
}

func GetAllDetail(username string, startTime time.Time, endTime time.Time, rType string, sort string, sord int, pageIndex int, pageSize int) ([]*Detail, error) {
	s := db.Stats.GetSession()
	defer s.Close()
	var result []*Detail
	var sortString string
	if sord == -1 {
		if sort == "created_at" {
			sortString = "-created_at"
		} else {
			sortString = "-result.tags." + rType
		}
	} else {
		if sort == "created_at" {
			sortString = "created_at"
		} else {
			sortString = "result.tags." + rType
		}
	}
	err := s.DB(db.Stats.DB).C("detail").Find(bson.M{
		"user": username,
		"created_at": bson.M{
			"$gte": startTime, "$lt": endTime},
		"result.porn.res": rType,
	}).Sort(sortString).Skip((pageIndex - 1) * pageSize).Limit(pageSize).All(&result)
	if err != nil {
		log.Error(fmt.Sprintf("find detail err ", err))
		return nil, err
	}
	return result, nil
}

func GetAllDetailCount(username string, startTime time.Time, endTime time.Time, rType string) (int, error) {
	s := db.Stats.GetSession()
	defer s.Close()
	count, err := s.DB(db.Stats.DB).C("detail").Find(bson.M{
		"user": username,
		"created_at": bson.M{
			"$gte": startTime, "$lt": endTime},
		"result.porn.res": rType,
	}).Count()

	if err != nil {
		log.Error(fmt.Sprintf("find detail err ", err))
		return 0, err
	}
	return count, nil
}

func GetScreenDetail(username string, startTime time.Time, endTime time.Time, sort string, sord int, pageIndex int, pageSize int) ([]*Detail, error) {
	s := db.Stats.GetSession()
	defer s.Close()
	var result []*Detail
	var sortString string
	if sord == -1 {
		if sort == "created_at" {
			sortString = "-created_at"
		} else {
			sortString = "-result.tags.isScreen"
		}
	} else {
		if sort == "created_at" {
			sortString = "created_at"
		} else {
			sortString = "result.tags.isScreen"
		}
	}
	err := s.DB(db.Stats.DB).C("detail").Find(bson.M{
		"user": username,
		"created_at": bson.M{
			"$gte": startTime, "$lt": endTime},
		"result.screenshot.res": "isScreen",
	}).Sort(sortString).Skip((pageIndex - 1) * pageSize).Limit(pageSize).All(&result)
	if err != nil {
		log.Error(fmt.Sprintf("find detail err ", err))
		return nil, err
	}
	return result, nil
}

func GetScreenDetailCount(username string, startTime time.Time, endTime time.Time) (int, error) {
	s := db.Stats.GetSession()
	defer s.Close()
	count, err := s.DB(db.Stats.DB).C("detail").Find(bson.M{
		"user": username,
		"created_at": bson.M{
			"$gte": startTime, "$lt": endTime},
		"result.screenshot.res": "isScreen",
	}).Count()
	if err != nil {
		log.Error(fmt.Sprintf("find detail err ", err))
		return 0, err
	}
	return count, nil
}
