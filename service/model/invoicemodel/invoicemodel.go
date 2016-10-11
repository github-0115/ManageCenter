package invoicemodel

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
	ErrUserInvoiceNotFound = errors.New("user invoice not found")
	ErrCursor              = errors.New("Cursor err")
)

type InvoiceInfo struct {
	InvoiceId     string    `bson:"invoice_id" binding:"required"`
	Username      string    `bson:"username" binding:"required"`
	Type          int64     `bson:"type" binding:"required"`
	Title         string    `bson:"title" binding:"required"`
	RegisterNum   string    `bson:"registernum" binding:"required"`
	BankName      string    `bson:"bankname" binding:"required"`
	BankAccount   string    `bson:"bankaccount" binding:"required"`
	BankAddress   string    `bson:"bankaddress" binding:"required"`
	TelNum        string    `bson:"telnum" binding:"required"`
	RegisterPhoto string    `bson:"registerphoto" binding:"required"`
	TaxPhoto      string    `bson:"taxphoto" binding:"required"`
	TaxpayerPhoto string    `bson:"taxpayerphoto" binding:"required"`
	Status        int64     `bson:"status" binding:"required"` //是否可用： 0 一般  1 默认 2 不可用
	CreatedAt     time.Time `bson:"created_at" binding:"required"`
}

func QueryInvoiceCount(invoiceId string, username string) (int, error) {
	s := db.User.GetSession()
	defer s.Close()

	result, err := s.DB(db.User.DB).C("invoice").Find(bson.M{
		"invoice_id": invoiceId,
		"username":   username,
	}).Count()
	if err != nil {
		log.Error(fmt.Sprintf("find user invoice  err ", err))
		if err == mgo.ErrNotFound {
			return 0, ErrUserInvoiceNotFound
		} else if err == mgo.ErrCursor {
			return 0, ErrCursor
		}
		return 0, err
	}
	return result, nil
}

func QueryOneInvoice(invoiceId string, username string) (*InvoiceInfo, error) {
	s := db.User.GetSession()
	defer s.Close()
	invoiceInfo := new(InvoiceInfo)
	err := s.DB(db.User.DB).C("invoice").Find(bson.M{
		"invoice_id": invoiceId,
		"username":   username,
	}).One(&invoiceInfo)
	if err != nil {
		log.Error(fmt.Sprintf("find user invoice num err ", err))
		if err == mgo.ErrNotFound {
			return nil, ErrUserInvoiceNotFound
		} else if err == mgo.ErrCursor {
			return nil, ErrCursor
		}
		return nil, err
	}
	return invoiceInfo, nil
}

func DeleteInvoice(invoiceId string, username string) (bool, error) {
	s := db.User.GetSession()
	defer s.Close()
	err := s.DB(db.User.DB).C("invoice").Update(bson.M{
		"invoice_id": invoiceId,
		"username":   username,
	}, bson.M{"$set": bson.M{"status": 2}})
	if err != nil {
		log.Error(fmt.Sprintf("delete user invoice err ", err))
		if err == mgo.ErrNotFound {
			return false, ErrUserInvoiceNotFound
		} else if err == mgo.ErrCursor {
			return false, ErrCursor
		}
		return false, err
	}
	return true, nil
}
