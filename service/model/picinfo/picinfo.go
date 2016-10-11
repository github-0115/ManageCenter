package picinfo

import (
	db "ManageCenter/service/db"
	"fmt"
	log "github.com/inconshreveable/log15"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Pic struct {
	OssUrl    string    `bson:"oss_url" binding:"required"`
	Result    *Result   `bson:"result" binding:"required"`
	CreatedAt time.Time `bson:"created_at" binding:"required"`
}

type Result struct {
	Porn *Porn `json:"porn"`
}
type Porn struct {
	Rate   float32 `json:"rate"`
	Res    string  `json:"res"`
	Review bool    `json:"review"`
}

func GetDetail(startTime time.Time, endTime time.Time, sord int, pageIndex int, pageSize int) ([]*Pic, int, error) {
	s := db.Demo.GetSession()
	defer s.Close()
	var result []*Pic
	var sortString string
	if sord == -1 {
		sortString = "-created_at"
	} else {
		sortString = "created_at"
	}
	err := s.DB(db.Demo.DB).C("picinfo").Find(bson.M{
		"created_at": bson.M{
			"$gte": startTime, "$lt": endTime},
	}).Sort(sortString).Skip((pageIndex - 1) * pageSize).Limit(pageSize).All(&result)
	count, err := s.DB(db.Demo.DB).C("picinfo").Find(bson.M{
		"created_at": bson.M{
			"$gte": startTime, "$lt": endTime},
	}).Count()
	if err != nil {
		log.Error(fmt.Sprintf("find detail err ", err))
		return nil, 0, err
	}
	return result, count, nil
}
