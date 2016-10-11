package vcodesmsmodel

import (
	db "ManageCenter/service/db"
	"errors"

	log "github.com/inconshreveable/log15"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var ErrVcodeNotFound = errors.New("Potential vcode not found")
var ErrPhoneNotFound = errors.New("Potential phone not found")

type VcodeSMSCntColl struct {
	Type  string `bson:"type" binding:"required" json:"type"`
	Phone string `bson:"phone" binding:"required" json:"phone"`
	Count int    `bson:"count" binding:"required" json:"count"`
}

func GetAllVcodePhone() ([]VcodeSMSCntColl, error) {
	s := db.User.GetSession()
	defer s.Close()
	var result []VcodeSMSCntColl
	iter := s.DB(db.User.DB).C("sms_count").Find(bson.M{
		"type": "vcode",
	}).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		log.Error("get all vcode phone failed. err=" + err.Error())
		if err == mgo.ErrNotFound {

			return result, ErrVcodeNotFound
		}
		return result, err
	}
	return result, nil

}

func GetPagingPuser(pageIndex int, pageSize int) ([]VcodeSMSCntColl, int, error) {
	s := db.User.GetSession()
	defer s.Close()
	var result []VcodeSMSCntColl
	err := s.DB(db.User.DB).C("sms_count").Find(bson.M{
		"type": "vcode",
	}).Skip((pageIndex - 1) * pageSize).Limit(pageSize).All(&result)
	if err != nil {
		log.Error("get all vcode phone failed. err=" + err.Error())
		if err == mgo.ErrNotFound {

			return result, 0, ErrPhoneNotFound
		}
		return result, 0, err

	}
	records, _ := s.DB(db.User.DB).C("sms_count").Find(nil).Count()

	return result, records, nil
}

func RmVcodePhoneDoc(phone string) error {
	s := db.User.GetSession()
	defer s.Close()
	err := s.DB(db.User.DB).C("sms_count").Remove(bson.M{
		"type":  "vcode",
		"phone": phone,
	})
	if err != nil {
		log.Error("rm vcode phone document failed. err=" + err.Error())
		return err
	}
	return nil
}
