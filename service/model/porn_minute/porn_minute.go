package porn_minute

import (
	db "ManageCenter/service/db"
	"errors"
	"fmt"
	"time"

	log "github.com/inconshreveable/log15"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	PornMinuteNotFount = errors.New("result of pornMinute not found")
)

type PornMinute struct {
	User       string
	CreatedAt  time.Time `bson:"created_at"`
	Review     int64     `bson:"review"`
	Total      int64
	P          int64               `bson:"p"`
	S          int64               `bson:"s"`
	N          int64               `bson:"n"`
	PReview    int64               `bson:"p_review"`
	SReview    int64               `bson:"s_review"`
	PConfirm   int64               `bson:"p_confirm"`
	SConfirm   int64               `bson:"s_confirm"`
	Detail     []*PornMinuteDetail `bson:"detail"`
	InkeReview int64               `bson:"inke_review"`
}

type PornMinuteDetail struct {
	Res     string
	Review  bool
	P       float32 `bson:"p,omitempty"`
	S       float32 `bson:"s,omitempty"`
	N       float32 `bson:"n,omitempty"`
	Md5     string  `bson:"md5,omitempty"`
	Version string  `bson:"version,omitempty"`
}

type MinuteDetail struct {
	Detail *PornMinuteDetail `bson:"detail"`
}

func Mintue(username string) ([]*PornMinute, error) {

	s := db.Stats.GetSession()
	defer s.Close()

	var result []*PornMinute
	err := s.DB(db.Stats.DB).C("porn_minute").Find(bson.M{
		"user": username,
	}).All(&result)
	if err != nil {
		log.Error(fmt.Sprintf("find user pornMinute err ", err))
		if err == mgo.ErrNotFound {
			return nil, PornMinuteNotFount
		}

		return nil, err
	}
	return result, nil
}

func GetMintueDetail(username string, startTime time.Time, endTime time.Time, rType string, sort int, pageIndex int, pageSize int) ([]*MinuteDetail, error) {
	s := db.Stats.GetSession()
	defer s.Close()
	var result []*MinuteDetail
	err := s.DB(db.Stats.DB).C("porn_minute").Pipe([]bson.M{
		{
			"$match": bson.M{
				"user": username,
				"created_at": bson.M{
					"$gt": startTime, "$lt": endTime}},
		}, {
			"$project": bson.M{
				"detail": 1},
		}, {
			"$unwind": "$detail",
		}, {
			"$match": bson.M{
				"detail.res": rType, "detail.review": true},
		}, {
			"$sort": bson.M{
				"detail." + rType: sort},
		}, {
			"$skip": (pageIndex - 1) * pageSize,
		}, {
			"$limit": pageSize,
		}}).AllowDiskUse().All(&result)

	if err != nil {
		log.Error(fmt.Sprintf("find user pornMinute err ", err))
		return nil, err
	}
	return result, nil
}

func GetMintueDetailCount(username string, startTime time.Time, endTime time.Time, rType string) (int, error) {
	s := db.Stats.GetSession()
	defer s.Close()
	var result []*MinuteDetail
	err := s.DB(db.Stats.DB).C("porn_minute").Pipe([]bson.M{
		{
			"$match": bson.M{
				"user": username,
				"created_at": bson.M{
					"$gt": startTime, "$lt": endTime}},
		}, {
			"$project": bson.M{
				"detail": 1},
		}, {
			"$unwind": "$detail",
		}, {
			"$match": bson.M{
				"detail.res": rType, "detail.review": true},
		}, {
			"$limit": 1000,
		}}).All(&result)

	if err != nil {
		log.Error(fmt.Sprintf("find user pornMinute err ", err))
		return 0, err
	}
	return len(result), nil
}

func GetMintueReview(username string, startTime time.Time, endTime time.Time) ([]*PornMinute, error) {
	s := db.Stats.GetSession()
	defer s.Close()
	var result []*PornMinute
	err := s.DB(db.Stats.DB).C("porn_minute").Find(bson.M{
		"user":       username,
		"created_at": bson.M{"$gt": startTime, "$lt": endTime},
	}).Select(bson.M{"total": 1, "review": 1, "created_at": 1, "inke_review": 1}).Sort("created_at").All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
