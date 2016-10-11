package transfersmodel

import (
	db "ManageCenter/service/db"
	"errors"
	"fmt"
	"time"

	log "github.com/inconshreveable/log15"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TransfersInfo struct {
	TransfersId  string    `bson:"transfers_id" binding:"required"`
	Username     string    `bson:"username" binding:"required"`
	TransfersrNo string    `bson:"transfers_no" binding:"required"` //转账单号
	Bank         string    `bson:"bank" binding:"required"`         //转账银行
	BankNum      string    `bson:"bank_num" binding:"required"`     //转账帐号
	Status       int64     `bson:"status" binding:"required"`       //审核状态 0 未审核 1 审核成功 2 审核不通过
	CreatedAt    time.Time `bson:"created_at" binding:"required"`
}

var (
	ErrUserTransfersNotFound = errors.New("user Transfers not found")
	ErrNumNotFound           = errors.New("find user Transfers num err")
	ErrTransfersCursor       = errors.New("Cursor err")
)

func (t *TransfersInfo) Save() error {
	s := db.User.GetSession()
	defer s.Close()
	return s.DB(db.User.DB).C("transfers").Insert(&t)
}

func UpdateTransfers(id string, key string, value string) error {
	s := db.User.GetSession()
	defer s.Close()

	err := s.DB(db.User.DB).C("transfers").Update(bson.M{"transfers_id": id}, bson.M{"$set": bson.M{key: value}})
	if err != nil {
		return err
	}
	return nil
}

func QueryTransfersCount(id string, username string) (int, error) {
	s := db.User.GetSession()
	defer s.Close()

	result, err := s.DB(db.User.DB).C("transfers").Find(bson.M{
		"transfers_id": id,
		"username":     username,
	}).Count()
	if err != nil {
		log.Error(fmt.Sprintf("find user transfers  err ", err))
		if err == mgo.ErrNotFound {
			return 0, ErrUserTransfersNotFound
		} else if err == mgo.ErrCursor {
			return 0, ErrTransfersCursor
		}
		return 0, err
	}
	return result, nil
}

func QueryTransfers(id string, username string) (*TransfersInfo, error) {
	s := db.User.GetSession()
	defer s.Close()
	collection := new(TransfersInfo)
	err := s.DB(db.User.DB).C("transfers").Find(bson.M{
		"transfers_id": id,
		"username":     username,
	}).One(collection)
	if err != nil {
		return nil, err
	}
	return collection, nil
}

func UpdateTransfersStatus(id string, username string) error {
	s := db.User.GetSession()
	defer s.Close()
	err := s.DB(db.User.DB).C("transfers").Update(bson.M{
		"transfers_id": id,
		"username":     username,
	}, bson.M{"$set": bson.M{"status": 1}})
	if err != nil {
		return err
	}
	return nil
}

func QueryPageTransfers(username string, startTime time.Time, endTime time.Time, pageIndex int, pageSize int) ([]*TransfersInfo, int, error) {
	s := db.User.GetSession()
	defer s.Close()
	var result []*TransfersInfo
	err := s.DB(db.User.DB).C("transfers").Find(bson.M{
		"username": username,
		"created_at": bson.M{
			"$gte": startTime, "$lt": endTime},
	}).Sort("-created_at").Skip((pageIndex - 1) * pageSize).Limit(pageSize).All(&result)
	records, err := s.DB(db.DBUri).C("transfers").Find(bson.M{
		"username": username,
		"created_at": bson.M{
			"$gte": startTime, "$lt": endTime},
	}).Count()
	if err != nil {
		return nil, 0, err
	}
	return result, records, nil
}

func GetPagingUserTransfers(username string, pageIndex int, pageSize int, sort string) ([]*TransfersInfo, int, error) {
	s := db.User.GetSession()
	defer s.Clone()
	var result []*TransfersInfo
	err := s.DB(db.User.DB).C("transfers").Find(bson.M{
		"username": username,
	}).Sort(sort).Skip((pageIndex - 1) * pageSize).Limit(pageSize).All(&result)

	records, err := s.DB(db.User.DB).C("transfers").Find(bson.M{
		"username": username,
	}).Count()
	if err != nil {
		log.Error(fmt.Sprintf("find user transfers err", err))
		if err == mgo.ErrNotFound {

			return result, 0, ErrUserTransfersNotFound
		}
		return result, 0, err
	}
	return result, records, nil
}

func GetPagingTransfers(pageIndex int, pageSize int, sort string) ([]*TransfersInfo, int, error) {
	s := db.User.GetSession()
	defer s.Clone()
	var result []*TransfersInfo
	err := s.DB(db.User.DB).C("transfers").Find(nil).Sort(sort).Skip((pageIndex - 1) * pageSize).Limit(pageSize).All(&result)

	records, err := s.DB(db.User.DB).C("transfers").Find(nil).Count()
	if err != nil {
		log.Error(fmt.Sprintf("find user transfers err", err))
		if err == mgo.ErrNotFound {

			return result, 0, ErrUserTransfersNotFound
		}
		return result, 0, err
	}
	return result, records, nil
}
