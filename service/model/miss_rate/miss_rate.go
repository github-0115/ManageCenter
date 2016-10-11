package miss_rate

import (
	db "ManageCenter/service/db"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type MissStruct struct {
	V63Miss   int64   `bson:"v6_3_miss"`
	V62Miss   int64   `bson:"v6_2_miss"`
	V63       float32 `bson:"v6_3"`
	V62       float32 `bson:"v6_2"`
	Total     int64
	Day       time.Time
	Type      string    `bson:"type"`
	CreatedAt time.Time `bson:"created_at"`
}

func GetMissRate(missType string, startTime time.Time, endTime time.Time) ([]*MissStruct, error) {
	s := db.Stats.GetSession()
	defer s.Close()
	var result []*MissStruct
	err := s.DB(db.Stats.DB).C("miss_rate").Find(bson.M{
		"type": missType,
		"day":  bson.M{"$gt": startTime, "$lt": endTime},
	}).Sort("day").All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
