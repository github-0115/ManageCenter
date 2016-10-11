package final_billmodel

import (
	db "ManageCenter/service/db"
	"errors"
	"fmt"
	"time"

	log "github.com/inconshreveable/log15"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type FinalBillColl struct {
	FinalBillId     string    `binding:"required" bson:"finalbill_id" json:"finalbill_id"`
	Username        string    `binding:"required" json:"username"`
	BillingId       []string  `binding:"required" bson:"billing_id" json:"billing_id"`
	UsedTotal       int64     `binding:"required" bson:"used_total" json:"used_total"`             //使用总量
	TotalFee        float64   `binding:"required" bson:"total_fee" json:"total_fee"`               //多少钱 /分
	PeriodStart     time.Time `binding:"required" bson:"period_start" json:"period_start"`         //结算周期-开始
	PeriodEnd       time.Time `binding:"required" bson:"period_end" json:"period_end"`             //结算周期-结束
	FavorRange      float64   `binding:"required" bson:"favor_range" json:"favor_range"`           //优惠区间 1、0.9、0.85、0.75、0.7、0.65
	PaidAmount      float64   `binding:"required" bson:"paid_amount" json:"paid_amount"`           //实际需要支付多少钱 /分
	Status          int64     `binding:"required" bson:"status" json:"status"`                     //状态 0：待处理； 1：已处理； 2：已完结
	Remark          string    `binding:"required" bson:"remark" json:"remark"`                     //备注
	BillType        string    `binding:"required" bson:"bill_type" json:"bill_type"`               //结算类型 月、季、半年、年
	TransferBank    string    `binding:"required" bson:"transfer_bank" json:"transfer_bank"`       //转账银行
	TransferAccount string    `binding:"required" bson:"transfer_account" json:"transfer_account"` //转账帐号
	TransferId      string    `binding:"required" bson:"transfer_id" json:"transfer_id"`           //转账单号
	TransactionTime time.Time `binding:"required" bson:"transaction_time" json:"transaction_time"` //交易时间
	CreatedAt       time.Time `bson:"created_at" binding:"required" json:"created_at"`
}

type billStatus struct {
	NotHandle int64
	Handle    int64
	Settle    int64
}

var (
	ErrBillNotFound = errors.New("bill not found")
	ErrBillCursor   = errors.New("Cursor err")
	ErrBillNum      = errors.New("find user Bill num err")
	BillStatus      = &billStatus{0, 1, 2}
)

func (t *FinalBillColl) Save() error {
	s := db.Manager.GetSession()
	defer s.Close()
	return s.DB(db.Manager.DB).C("final_bill").Insert(&t)
}

func QueryOneBill(id string, username string) (*FinalBillColl, error) {
	s := db.Manager.GetSession()
	defer s.Close()
	collection := new(FinalBillColl)
	err := s.DB(db.Manager.DB).C("final_bill").Find(bson.M{
		"finalbill_id": id,
		"username":     username,
	}).One(collection)
	if err != nil {
		log.Error(fmt.Sprintf("find user final bill  err ", err))
		if err == mgo.ErrNotFound {
			return nil, ErrBillNotFound
		} else if err == mgo.ErrCursor {
			return nil, ErrBillCursor
		}
		return nil, err
	}
	return collection, nil
}

func QueryUserBill(username string, status int64, utype string) (*FinalBillColl, error) {
	s := db.Manager.GetSession()
	defer s.Close()
	collection := new(FinalBillColl)
	err := s.DB(db.Manager.DB).C("final_bill").Find(bson.M{
		"username":  username,
		"status":    status,
		"bill_type": utype,
	}).One(collection)
	if err != nil {
		log.Error(fmt.Sprintf("find user final bill  err ", err))
		if err == mgo.ErrNotFound {
			return nil, ErrBillNotFound
		} else if err == mgo.ErrCursor {
			return nil, ErrBillCursor
		}
		return nil, err
	}
	return collection, nil
}

func DeleteUserBill(username string, status int64, utype string) error {
	s := db.Manager.GetSession()
	defer s.Close()

	err := s.DB(db.Manager.DB).C("final_bill").Remove(bson.M{
		"username":  username,
		"status":    status,
		"bill_type": utype,
	})
	if err != nil {
		log.Error(fmt.Sprintf("delete user final bill  err ", err))
		if err == mgo.ErrNotFound {
			return ErrBillNotFound
		} else if err == mgo.ErrCursor {
			return ErrBillCursor
		}
		return err
	}
	return nil
}

func GetNotSettleBills() ([]*FinalBillColl, error) {
	s := db.Manager.GetSession()
	defer s.Close()
	var result []*FinalBillColl
	err := s.DB(db.Manager.DB).C("final_bill").Find(bson.M{
		"$or": []bson.M{
			bson.M{"status": BillStatus.Handle},
			bson.M{"status": BillStatus.NotHandle},
		},
	}).All(&result)

	if err != nil {
		log.Error(fmt.Sprintf("find not settle final bill err", err))
		if err == mgo.ErrNotFound {
			return nil, ErrBillNotFound
		}
		return nil, err
	}
	return result, nil
}

func GetPagingBills(pageIndex int, pageSize int, sort string, status int64) ([]*FinalBillColl, int, error) {
	s := db.Manager.GetSession()
	defer s.Close()
	var result []*FinalBillColl
	err := s.DB(db.Manager.DB).C("final_bill").Find(bson.M{
		"status": status,
	}).Sort(sort).Skip((pageIndex - 1) * pageSize).Limit(pageSize).All(&result)

	records, err := s.DB(db.Manager.DB).C("final_bill").Find(bson.M{
		"status": status,
	}).Count()
	if err != nil {
		log.Error(fmt.Sprintf("find paging final bill err", err))
		if err == mgo.ErrNotFound {
			return result, 0, ErrBillNotFound
		}
		return result, 0, err
	}
	return result, records, nil
}

func UpdateFinalBillInfo(id string, username string, status int64, ids []string, total int64, total_fee float64, t time.Time) error {
	s := db.Manager.GetSession()
	defer s.Close()

	err := s.DB(db.Manager.DB).C("final_bill").Update(bson.M{
		"finalbill_id": id,
		"username":     username,
		"status":       status,
	}, bson.M{"$set": bson.M{"billing_id": ids, "used_total": total, "total_fee": total_fee, "paid_amount": total_fee, "period_end": t}})
	if err != nil {
		log.Error(fmt.Sprintf("update user  final bill favorable  err ", err))
		if err == mgo.ErrNotFound {
			return ErrBillNotFound
		} else if err == mgo.ErrCursor {
			return ErrBillCursor
		}
		return err
	}
	return nil
}

func UpdateFavorBill(id string, username string, count float64) error {
	s := db.Manager.GetSession()
	defer s.Close()

	err := s.DB(db.Manager.DB).C("final_bill").Update(bson.M{
		"finalbill_id": id,
		"username":     username,
	}, bson.M{"$set": bson.M{"paid_amount": count, "status": 1}})
	if err != nil {
		log.Error(fmt.Sprintf("update user  final bill favorable  err ", err))
		if err == mgo.ErrNotFound {
			return ErrBillNotFound
		} else if err == mgo.ErrCursor {
			return ErrBillCursor
		}
		return err
	}
	return nil
}

func UpdateTransfersBill(id string, username string, tno string, bank string, bankNum string) error {
	s := db.Manager.GetSession()
	defer s.Close()

	err := s.DB(db.Manager.DB).C("final_bill").Update(bson.M{
		"finalbill_id": id,
		"username":     username,
	}, bson.M{"$set": bson.M{"transfer_id": tno, "transfer_bank": bank, "transfer_account": bankNum, "status": 2}})
	if err != nil {
		log.Error(fmt.Sprintf("update user  final bill favorable  err ", err))
		if err == mgo.ErrNotFound {
			return ErrBillNotFound
		} else if err == mgo.ErrCursor {
			return ErrBillCursor
		}
		return err
	}
	return nil
}
