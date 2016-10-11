package porn_hour

import (
	db "ManageCenter/service/db"
	"errors"
	"time"

	"gopkg.in/mgo.v2/bson"
)

var (
	PornHourNotFount = errors.New("result of pornHour not found")
)

type PornHour struct {
	User      string
	CreatedAt time.Time `bson:"created_at"`
	Total     int64
	P         int64 `bson:"p"`
	S         int64 `bson:"s"`
	N         int64 `bson:"n"`
	PReview   int64 `bson:"p_review"`
	SReview   int64 `bson:"s_review"`
	NReview   int64 `bson:"n_review"`
	PConfirm  int64 `bson:"p_confirm"`
	SConfirm  int64 `bson:"s_confirm"`
	IsScreen  int64 `bson:"isScreen"`
	NotScreen int64 `bson:"notScreen"`
}

type PornHourTotal struct {
	CreatedAt time.Time `bson:"_id"`
	Total     int64     `bson:"total"`
}

type PornHourSum struct {
	CreatedAt time.Time `bson:"_id"`
	Total     int64
	P         int64 `bson:"p"`
	S         int64 `bson:"s"`
	N         int64 `bson:"n"`
	PReview   int64 `bson:"p_review"`
	SReview   int64 `bson:"s_review"`
	NReview   int64 `bson:"n_review"`
	IsScreen  int64 `bson:"isScreen"`
	NotScreen int64 `bson:"notScreen"`
}

func QueryPornHour(username string, startTime time.Time, endTime time.Time) ([]*PornHour, error) {
	s := db.Stats.GetSession()
	defer s.Close()
	var result []*PornHour
	err := s.DB(db.Stats.DB).C("porn_hour").Find(bson.M{
		"user":       username,
		"created_at": bson.M{"$gte": startTime, "$lt": endTime},
	}).Select(bson.M{"total": 1, "p": 1, "s": 1, "n": 1, "p_review": 1, "s_review": 1, "n_review": 1, "created_at": 1, "isScreen": 1, "notScreen": 1}).Sort("created_at").All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetUserHourTotal(username string, startTime time.Time, endTime time.Time) ([]*PornHourSum, error) {
	s := db.Stats.GetSession()
	defer s.Clone()
	var result []*PornHourSum
	err := s.DB(db.Stats.DB).C("porn_hour").Pipe([]bson.M{
		{
			"$match": bson.M{
				"user": username,
				"created_at": bson.M{
					"$gte": startTime, "$lt": endTime}},
		}, {
			"$group": bson.M{
				"_id": "$created_at",
				"total": bson.M{
					"$sum": "$total"},
				"p": bson.M{
					"$sum": "$p"},
				"s": bson.M{
					"$sum": "$s"},
				"n": bson.M{
					"$sum": "$n"},
				"p_review": bson.M{
					"$sum": "$p_review"},
				"s_review": bson.M{
					"$sum": "$s_review"},
				"isScreen": bson.M{
					"$sum": "$isScreen"},
				"notScreen": bson.M{
					"$sum": "$notScreen"},
			},
		}, {
			"$sort": bson.M{
				"_id": 1},
		}}).All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetAllHourTotal(startTime time.Time, endTime time.Time) ([]*PornHourSum, error) {
	s := db.Stats.GetSession()
	defer s.Clone()
	var result []*PornHourSum
	err := s.DB(db.Stats.DB).C("porn_hour").Pipe([]bson.M{
		{
			"$match": bson.M{
				"created_at": bson.M{
					"$gte": startTime, "$lt": endTime}},
		}, {
			"$group": bson.M{
				"_id": "$created_at",
				"total": bson.M{
					"$sum": "$total"},
				"p": bson.M{
					"$sum": "$p"},
				"s": bson.M{
					"$sum": "$s"},
				"n": bson.M{
					"$sum": "$n"},
				"p_review": bson.M{
					"$sum": "$p_review"},
				"s_review": bson.M{
					"$sum": "$s_review"},
				"isScreen": bson.M{
					"$sum": "$isScreen"},
				"notScreen": bson.M{
					"$sum": "$notScreen"},
			},
		}, {
			"$sort": bson.M{
				"_id": 1},
		}}).All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
