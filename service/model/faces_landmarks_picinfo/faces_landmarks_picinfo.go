package faces_landmarks_picinfo

import (
	db "ManageCenter/service/db"
	"fmt"
	log "github.com/inconshreveable/log15"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type PicLandmark struct {
	OssUrl        string                 `bson:"oss_url" binding:"required"`
	FaceLandmarks []*FaceLandmarksStruct `bson:"face_landmarks" binding:"required" json:"face_landmarks"`
	CreatedAt     time.Time              `bson:"created_at" binding:"required"`
}
type FaceLandmarksStruct struct {
	Bound     *BoundsStruct      `bson:"left_top" json:"bound"`
	Landmarks []*LandmarksStruct `bson:"landmarks" json:"landmarks"`
}
type BoundsStruct struct {
	LeftTop     *LeftTopStruct     `bson:"left_top" json:"lefttop"`
	RightBottom *RightBottomStruct `bson:"right_bottom" json:"rightbottom"`
}
type LandmarksStruct struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type LeftTopStruct struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}
type RightBottomStruct struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

func GetDetail(startTime time.Time, endTime time.Time, sord int, pageIndex int, pageSize int) ([]*PicLandmark, int, error) {
	s := db.Demo.GetSession()
	defer s.Close()
	var result []*PicLandmark
	var sortString string
	if sord == -1 {
		sortString = "-created_at"
	} else {
		sortString = "created_at"
	}
	err := s.DB(db.Demo.DB).C("faces_landmark_picinfo").Find(bson.M{
		"created_at": bson.M{
			"$gte": startTime, "$lt": endTime},
	}).Sort(sortString).Skip((pageIndex - 1) * pageSize).Limit(pageSize).All(&result)
	count, err := s.DB(db.Demo.DB).C("faces_landmark_picinfo").Find(bson.M{
		"created_at": bson.M{
			"$gte": startTime, "$lt": endTime},
	}).Count()
	if err != nil {
		log.Error(fmt.Sprintf("find faceslandmark detail err", err))
		return nil, 0, err
	}
	return result, count, nil
}
