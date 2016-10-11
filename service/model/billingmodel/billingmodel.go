package billingmodel

import (
	db "ManageCenter/service/db"
	"errors"
	"fmt"
	"time"

	log "github.com/inconshreveable/log15"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BillColl struct {
	BillingId       string    `binding:"required" bson:"billing_id" json:"billing_id"`
	Username        string    `binding:"required" json:"username"`
	UsedTotal       int64     `binding:"required" bson:"used_total" json:"used_total"`             //使用总量
	TotalFee        float64   `binding:"required" bson:"total_fee" json:"total_fee"`               //多少钱 /分
	PeriodStart     time.Time `binding:"required" bson:"period_start" json:"period_start"`         //结算周期-开始
	PeriodEnd       time.Time `binding:"required" bson:"period_end" json:"period_end"`             //结算周期-结束
	Status          int64     `binding:"required" bson:"status" json:"status"`                     //状态 0：待处理； 1：已处理
	BillType        string    `binding:"required" bson:"bill_type" json:"bill_type"`               //结算类型 月、季、半年、年
	TransactionTime time.Time `binding:"required" bson:"transaction_time" json:"transaction_time"` //交易时间
	CreatedAt       time.Time `bson:"created_at" binding:"required" json:"created_at"`
}

var (
	ErrBillNotFound = errors.New("bill not found")
	ErrBillCursor   = errors.New("Cursor err")
	ErrBillNum      = errors.New("find user Bill num err")
)

func (t *BillColl) Save() error {
	s := db.Manager.GetSession()
	defer s.Close()
	return s.DB(db.Manager.DB).C("bill").Insert(&t)
}

func QueryOneBill(id string, username string) (*BillColl, error) {
	s := db.Manager.GetSession()
	defer s.Close()
	collection := new(BillColl)
	err := s.DB(db.Manager.DB).C("bill").Find(bson.M{
		"billing_id": id,
		"username":   username,
	}).One(collection)
	if err != nil {
		log.Error(fmt.Sprintf("find user Bill  err ", err))
		if err == mgo.ErrNotFound {
			return nil, ErrBillNotFound
		} else if err == mgo.ErrCursor {
			return nil, ErrBillCursor
		}
		return nil, err
	}
	return collection, nil
}

func QueryUserBills(username string, status int64, utype string) ([]*BillColl, error) {
	s := db.Manager.GetSession()
	defer s.Close()
	var result []*BillColl
	err := s.DB(db.Manager.DB).C("bill").Find(bson.M{
		"username":  username,
		"status":    status,
		"bill_type": utype,
	}).All(&result)
	if err != nil {
		log.Error(fmt.Sprintf("find user Bill  err ", err))
		if err == mgo.ErrNotFound {
			return nil, ErrBillNotFound
		} else if err == mgo.ErrCursor {
			return nil, ErrBillCursor
		}
		return nil, err
	}
	return result, nil
}

func UpdateBillingStatus(id string, username string, t time.Time) error {
	s := db.Manager.GetSession()
	defer s.Close()

	err := s.DB(db.Manager.DB).C("bill").Update(bson.M{
		"billing_id": id,
		"username":   username,
	}, bson.M{"$set": bson.M{"status": 1, "transaction_time": t}})
	if err != nil {
		log.Error(fmt.Sprintf("update user Bill  err ", err))
		if err == mgo.ErrNotFound {
			return ErrBillNotFound
		} else if err == mgo.ErrCursor {
			return ErrBillCursor
		}
		return err
	}
	return nil
}

func GetPagingBills(pageIndex int, pageSize int, sort string) ([]*BillColl, int, error) {
	s := db.Manager.GetSession()
	defer s.Close()
	var result []*BillColl
	err := s.DB(db.Manager.DB).C("bill").Find(nil).Sort(sort).Skip((pageIndex - 1) * pageSize).Limit(pageSize).All(&result)
	records, err := s.DB(db.Manager.DB).C("bill").Find(nil).Count()
	if err != nil {
		log.Error(fmt.Sprintf("find paging bill err", err))
		if err == mgo.ErrNotFound {

			return result, 0, ErrBillNotFound
		}
		return result, 0, err
	}
	return result, records, nil
}
