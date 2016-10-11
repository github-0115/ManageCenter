package servicepackagemodel

import (
	db "ManageCenter/service/db"
	"errors"

	"time"

	"fmt"

	log "github.com/inconshreveable/log15"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserServicePackage struct {
	Username    string    `binding:"required" json:"username"`
	UserType    int64     `bson:"usertype" binding:"required" json:"usertype"` //暂定： 0 免费  1 预付费  2 后付费  3 代理商
	Price       int64     `bson:"price" binding:"required" json:"price"`
	MonthTotal  int64     `bson:"month_total" binding:"required" json:"month_total"`
	DailyAmount int64     `bson:"daily_amount" binding:"required" json:"daily_amount"`
	Concurrency int64     `bson:"concurrency" binding:"required" json:"concurrency"`
	Period      string    `bson:"period" binding:"required" json:"period"` //结算类型 月、季、半年、年
	FreeDay     int64     `bson:"freeday" binding:"required" json:"freeday"`
	CreatedAt   time.Time `bson:"created_at" binding:"required" json:"created_at"`
}

type InterfaceConfig struct {
	DefaultPrice       int64 `bson:"default_price" binding:"required" json:"default_price"`
	DefaultMonthTotal  int64 `bson:"default_month_total" binding:"required" json:"default_month_total"`
	DefaultConcurrency int64 `bson:"default_concurrency" binding:"required" json:"default_concurrency"`
	FreeDayTotal       int64 `bson:"free_day_total" binding:"required" json:"free_day_total"`
	FreeConcurrency    int64 `bson:"free_concurrency" binding:"required" json:"free_concurrency"`
	ValidityTime       int64 `bson:"validity_time" binding:"required" json:"validity_time"`
}

var (
	ErrUserServicePackageNotFound = errors.New("user's service package not found")
	ErrUserServicePackageCursor   = errors.New("user service package Cursor err")
	ErrInterfaceConfigNotFound    = errors.New("interface config not found")
	ErrInterfaceConfigCursor      = errors.New("interface config Cursor err")
)

var (
	Paid_daily_amount = 100000000000
	zeroAmount        = 0
)

func (t *UserServicePackage) Save() error {
	s := db.Manager.GetSession()
	defer s.Close()
	return s.DB(db.Manager.DB).C("userservpkg").Insert(&t)
}

func GetUserServicePackage(username string) (*UserServicePackage, error) {
	coll := new(UserServicePackage)
	s := db.Manager.GetSession()
	defer s.Close()
	err := s.DB(db.Manager.DB).C("userservpkg").Find(bson.M{
		"username": username,
	}).One(coll)
	if err != nil {
		log.Error(fmt.Sprintf("find user service package err ", err))
		if err == mgo.ErrNotFound {
			return nil, ErrUserServicePackageNotFound
		} else if err == mgo.ErrCursor {
			return nil, ErrUserServicePackageCursor
		}
		return nil, err
	}
	return coll, nil
}

func GetAllUserServicePackages() ([]*UserServicePackage, error) {
	var results []*UserServicePackage
	s := db.Manager.GetSession()
	defer s.Close()
	err := s.DB(db.Manager.DB).C("userservpkg").Find(nil).All(&results)
	if err != nil {
		log.Error(fmt.Sprintf("find user service package err ", err))
		if err == mgo.ErrNotFound {
			return nil, ErrUserServicePackageNotFound
		} else if err == mgo.ErrCursor {
			return nil, ErrUserServicePackageCursor
		}
		return nil, err
	}
	return results, nil

}

func UpdateFreeUserServicePackage(user string, usertype int64, daily_amount int64, freeday int64) (bool, error) {
	s := db.Manager.GetSession()
	defer s.Close()

	_, err := s.DB(db.Manager.DB).C("userservpkg").Upsert(bson.M{"username": user},
		bson.M{"$set": bson.M{"usertype": usertype, "month_total": zeroAmount, "daily_amount": daily_amount, "price": zeroAmount, "concurrency": zeroAmount, "freeday": freeday}})
	if err != nil {
		log.Error(fmt.Sprintf("update frss user service package err ", err))
		if err == mgo.ErrNotFound {
			return false, ErrUserServicePackageNotFound
		} else if err == mgo.ErrCursor {
			return false, ErrUserServicePackageCursor
		}
		return false, err
	}
	return true, nil
}

func UpdatePostPaidUserServicePackage(user string, usertype int64, month_total int64, daily_amount int64, price int64, concurrency int64, period string) (bool, error) {
	s := db.Manager.GetSession()
	defer s.Close()

	_, err := s.DB(db.Manager.DB).C("userservpkg").Upsert(bson.M{"username": user},
		bson.M{"$set": bson.M{"usertype": usertype, "month_total": month_total, "daily_amount": daily_amount, "price": price, "concurrency": concurrency, "period": period}})
	if err != nil {
		log.Error(fmt.Sprintf("update frss user service package err ", err))
		if err == mgo.ErrNotFound {
			return false, ErrUserServicePackageNotFound
		} else if err == mgo.ErrCursor {
			return false, ErrUserServicePackageCursor
		}
		return false, err
	}
	return true, nil
}

func UpdateAgentsUserServicePackage(user string, usertype int64, month_total int64, daily_amount int64, concurrency int64, period string) (bool, error) {
	s := db.Manager.GetSession()
	defer s.Close()

	_, err := s.DB(db.Manager.DB).C("userservpkg").Upsert(bson.M{"username": user},
		bson.M{"$set": bson.M{"usertype": usertype, "month_total": month_total, "daily_amount": daily_amount, "concurrency": concurrency, "period": period}})
	if err != nil {
		log.Error(fmt.Sprintf("update frss user service package err ", err))
		if err == mgo.ErrNotFound {
			return false, ErrUserServicePackageNotFound
		} else if err == mgo.ErrCursor {
			return false, ErrUserServicePackageCursor
		}
		return false, err
	}
	return true, nil
}

func UpdateUserServicePackage(user string, month_total int64, daily_amount int64, concurrency int64, period string) (bool, error) {
	s := db.Manager.GetSession()
	defer s.Close()

	_, err := s.DB(db.Manager.DB).C("userservpkg").Upsert(bson.M{"username": user},
		bson.M{"$set": bson.M{"month_total": month_total, "concurrency": concurrency},
			"setOnInsert": bson.M{"created_at": time.Now()}})
	if err != nil {
		log.Error(fmt.Sprintf("update user service package month_total err ", err))
		if err == mgo.ErrNotFound {
			return false, ErrUserServicePackageNotFound
		} else if err == mgo.ErrCursor {
			return false, ErrUserServicePackageCursor
		}
		return false, err
	}

	_, err = s.DB(db.Manager.DB).C("userservpkg").Upsert(
		bson.M{
			"$and": []bson.M{
				bson.M{"username": user},
				bson.M{"daily_amount": bson.M{"$ne": Paid_daily_amount}},
			}},
		bson.M{"$set": bson.M{"daily_amount": daily_amount, "concurrency": concurrency}})
	if err != nil {
		log.Error(fmt.Sprintf("update user service package daily_amount err ", err))
		if err == mgo.ErrNotFound {
			return false, ErrUserServicePackageNotFound
		} else if err == mgo.ErrCursor {
			return false, ErrUserServicePackageCursor
		}
		return false, err
	}

	return true, err
}

func DeleteUserInterfaceConfig(username string) error {
	s := db.Manager.GetSession()
	defer s.Close()
	err := s.DB(db.Manager.DB).C("userservpkg").Remove(bson.M{"username": username})
	if err != nil {
		log.Error(fmt.Sprintf("delete user interfaceconfig  err", err))
		if err == mgo.ErrNotFound {

			return ErrInterfaceConfigNotFound
		}
		return err
	}
	return nil
}

func (t *InterfaceConfig) Save() error {
	s := db.Manager.GetSession()
	defer s.Close()
	return s.DB(db.Manager.DB).C("interfaceconfig").Insert(&t)
}

func GetInterfaceConfig() (*InterfaceConfig, error) {
	coll := new(InterfaceConfig)
	s := db.Manager.GetSession()
	defer s.Close()
	err := s.DB(db.Manager.DB).C("interfaceconfig").Find(bson.M{}).One(coll)
	if err != nil {
		log.Error(fmt.Sprintf("find user service package err ", err))
		if err == mgo.ErrNotFound {
			return nil, ErrInterfaceConfigNotFound
		} else if err == mgo.ErrCursor {
			return nil, ErrInterfaceConfigCursor
		}
		return nil, err
	}
	return coll, nil
}

func UpdateInterfaceConfig(default_month_total int64, default_concurrency int64, free_day_total int64, free_concurrency int64, price int64, validity_time int64) (bool, error) {
	s := db.Manager.GetSession()
	defer s.Close()

	_, err := s.DB(db.Manager.DB).C("interfaceconfig").Upsert(bson.M{},
		bson.M{"$set": bson.M{
			"default_month_total": default_month_total,
			"default_concurrency": default_concurrency,
			"default_price":       price,
			"validity_time":       validity_time,
			"free_day_total":      free_day_total,
			"free_concurrency":    free_concurrency}})
	if err != nil {
		log.Error(fmt.Sprintf("update interface config err ", err))
		if err == mgo.ErrNotFound {
			return false, ErrInterfaceConfigNotFound
		} else if err == mgo.ErrCursor {
			return false, ErrInterfaceConfigCursor
		}
		return false, err
	}

	return true, err
}
