package faces_picinfo

import (
	db "ManageCenter/service/db"
	"fmt"
	log "github.com/inconshreveable/log15"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Pic struct {
	OssUrl     string              `bson:"oss_url" binding:"required"`
	FaceBounds []*FaceBoundsStruct `bson:"face_bounds" binding:"required" json:"face_bounds"`
	CreatedAt  time.Time           `bson:"created_at" binding:"required"`
}
type FaceBoundsStruct struct {
	LeftTop     *LeftTopStruct     `bson:"left_top" json:"lefttop"`
	RightBottom *RightBottomStruct `bson:"right_bottom" json:"rightbottom"`
}

type LeftTopStruct struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}
type RightBottomStruct struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
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
	err := s.DB(db.Demo.DB).C("faces_picinfo").Find(bson.M{
		"created_at": bson.M{
			"$gte": startTime, "$lt": endTime},
	}).Sort(sortString).Skip((pageIndex - 1) * pageSize).Limit(pageSize).All(&result)
	count, err := s.DB(db.Demo.DB).C("faces_picinfo").Find(bson.M{
		"created_at": bson.M{
			"$gte": startTime, "$lt": endTime},
	}).Count()
	if err != nil {
		log.Error(fmt.Sprintf("find faces detail err", err))
		return nil, 0, err
	}
	return result, count, nil
}
