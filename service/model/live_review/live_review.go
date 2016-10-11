package live_review

import (
	db "ManageCenter/service/db"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type LiveReview struct {
	User      string
	CreatedAt time.Time `bson:"created_at"`
	Rate      float32   `bson:"rate"`
	Version   string    `bson:"version"`
}

func (t *LiveReview) Save() error {
	s := db.Stats.GetSession()
	defer s.Close()
	return s.DB(db.Stats.DB).C("live_review").Insert(&t)
}

func GetLiveReviewRate(username string, startTime time.Time, endTime time.Time) ([]*LiveReview, error) {
	s := db.Stats.GetSession()
	defer s.Close()
	var result []*LiveReview
	err := s.DB(db.Stats.DB).C("live_review").Find(bson.M{
		"user":       username,
		"created_at": bson.M{"$gt": startTime, "$lt": endTime},
	}).Select(bson.M{"rate": 1, "created_at": 1, "version": 1}).Sort("created_at").All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetVersion(username string, startTime time.Time, endTime time.Time) ([]string, error) {
	s := db.Stats.GetSession()
	defer s.Close()
	var result []string
	err := s.DB(db.Stats.DB).C("live_review").Find(bson.M{
		"user":       username,
		"created_at": bson.M{"$gt": startTime, "$lt": endTime},
	}).Distinct("version", &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
