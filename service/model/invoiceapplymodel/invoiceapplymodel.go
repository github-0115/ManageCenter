package invoiceapplymodel

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
	ErrUserInvoiceApplyNotFound = errors.New("user invoice apply not found")
	ErrCursor                   = errors.New("Cursor err")
)

type InvoiceApplyInfo struct {
	ApplyId       string    `bson:"apply_id" binding:"required"`
	Username      string    `bson:"username" binding:"required"`
	InvoiceId     string    `bson:"invoice_id" binding:"required"`
	AddressId     string    `bson:"address_id" binding:"required"`
	LogisticsName string    `bson:"logistics_name" binding:"required"`
	TrackNum      string    `bson:"track_num" binding:"required"`
	Status        int64     `bson:"status" binding:"required"`
	Money         float64   `bson:"money" binding:"required"`
	CreatedAt     time.Time `bson:"created_at" binding:"required"`
}

func QueryInvoiceApplyCount(applyId string, username string) (int, error) {
	s := db.User.GetSession()
	defer s.Close()

	result, err := s.DB(db.User.DB).C("invoice_apply").Find(bson.M{
		"apply_id": applyId,
		"username": username,
	}).Count()
	if err != nil {
		log.Error(fmt.Sprintf("find user invoice_apply  err ", err))
		if err == mgo.ErrNotFound {
			return 0, ErrUserInvoiceApplyNotFound
		} else if err == mgo.ErrCursor {
			return 0, ErrCursor
		}
		return 0, err
	}
	return result, nil
}

func QueryOneInvoiceApply(applyId string, username string) (*InvoiceApplyInfo, error) {
	s := db.User.GetSession()
	defer s.Close()
	invoiceInfo := new(InvoiceApplyInfo)
	err := s.DB(db.User.DB).C("invoice_apply").Find(bson.M{
		"apply_id": applyId,
		"username": username,
	}).One(&invoiceInfo)
	if err != nil {
		log.Error(fmt.Sprintf("find user invoice_apply num err ", err))
		if err == mgo.ErrNotFound {
			return nil, ErrUserInvoiceApplyNotFound
		} else if err == mgo.ErrCursor {
			return nil, ErrCursor
		}
		return nil, err
	}
	return invoiceInfo, nil
}

func QueryAllInvoiceApply() ([]*InvoiceApplyInfo, error) {
	s := db.User.GetSession()
	defer s.Close()
	var result []*InvoiceApplyInfo
	err := s.DB(db.User.DB).C("invoice_apply").Find(nil).All(&result)
	if err != nil {
		log.Error(fmt.Sprintf("find user invoice_apply num err ", err))
		if err == mgo.ErrNotFound {
			return nil, ErrUserInvoiceApplyNotFound
		} else if err == mgo.ErrCursor {
			return nil, ErrCursor
		}
		return nil, err
	}
	return result, nil
}

func QueryUserInvoiceApply(username string) ([]*InvoiceApplyInfo, error) {
	s := db.User.GetSession()
	defer s.Close()
	var result []*InvoiceApplyInfo
	err := s.DB(db.User.DB).C("invoice_apply").Find(bson.M{
		"username": username,
	}).All(&result)
	if err != nil {
		log.Error(fmt.Sprintf("find user invoice_apply num err ", err))
		if err == mgo.ErrNotFound {
			return nil, ErrUserInvoiceApplyNotFound
		} else if err == mgo.ErrCursor {
			return nil, ErrCursor
		}
		return nil, err
	}
	return result, nil
}

func GetPagingUserInvoiceApply(username string, pageIndex int, pageSize int, sort string) ([]*InvoiceApplyInfo, int, error) {
	s := db.User.GetSession()
	defer s.Clone()
	var result []*InvoiceApplyInfo
	err := s.DB(db.User.DB).C("invoice_apply").Find(bson.M{
		"username": username,
	}).Sort(sort).Skip((pageIndex - 1) * pageSize).Limit(pageSize).All(&result)

	records, err := s.DB(db.User.DB).C("invoice_apply").Find(bson.M{
		"username": username,
	}).Count()
	if err != nil {
		log.Error(fmt.Sprintf("find user invoice_apply err", err))
		if err == mgo.ErrNotFound {

			return result, 0, ErrUserInvoiceApplyNotFound
		}
		return result, 0, err
	}
	return result, records, nil
}

func GetPagingInvoiceApply(pageIndex int, pageSize int, sort string) ([]*InvoiceApplyInfo, int, error) {
	s := db.User.GetSession()
	defer s.Clone()
	var result []*InvoiceApplyInfo
	err := s.DB(db.User.DB).C("invoice_apply").Find(nil).Sort(sort).Skip((pageIndex - 1) * pageSize).Limit(pageSize).All(&result)

	records, err := s.DB(db.User.DB).C("invoice_apply").Find(nil).Count()
	if err != nil {
		log.Error(fmt.Sprintf("find user invoice_apply err", err))
		if err == mgo.ErrNotFound {

			return result, 0, ErrUserInvoiceApplyNotFound
		}
		return result, 0, err
	}
	return result, records, nil
}

func UpdateInvoiceApply(applyId string, username string, key string, value int64) (bool, error) {
	s := db.User.GetSession()
	defer s.Close()
	err := s.DB(db.User.DB).C("invoice_apply").Update(bson.M{
		"apply_id": applyId,
		"username": username,
	}, bson.M{"$set": bson.M{key: value}})
	if err != nil {
		log.Error(fmt.Sprintf("update user invoice_apply info err ", err))
		if err == mgo.ErrNotFound {
			return false, ErrUserInvoiceApplyNotFound
		} else if err == mgo.ErrCursor {
			return false, ErrCursor
		}
		return false, err
	}
	return true, nil
}

func UpdateInvoiceApplyInfo(applyId string, username string, key string, value string) (bool, error) {
	s := db.User.GetSession()
	defer s.Close()
	err := s.DB(db.User.DB).C("invoice_apply").Update(bson.M{
		"apply_id": applyId,
		"username": username,
	}, bson.M{"$set": bson.M{key: value}})
	if err != nil {
		log.Error(fmt.Sprintf("update user invoice_apply info err ", err))
		if err == mgo.ErrNotFound {
			return false, ErrUserInvoiceApplyNotFound
		} else if err == mgo.ErrCursor {
			return false, ErrCursor
		}
		return false, err
	}
	return true, nil
}

func DeleteInvoiceApply(applyId string, username string) (bool, error) {
	s := db.User.GetSession()
	defer s.Close()
	err := s.DB(db.User.DB).C("invoice_apply").Remove(bson.M{
		"apply_id": applyId,
		"username": username,
	})
	if err != nil {
		log.Error(fmt.Sprintf("delete user invoice_apply err ", err))
		if err == mgo.ErrNotFound {
			return false, ErrUserInvoiceApplyNotFound
		} else if err == mgo.ErrCursor {
			return false, ErrCursor
		}
		return false, err
	}
	return true, nil
}
