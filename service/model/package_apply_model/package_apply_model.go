package package_apply_model

import (
	db "ManageCenter/service/db"
	"errors"
	"fmt"
	"time"

	log "github.com/inconshreveable/log15"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PackageApplyColl struct {
	PackageApplyId string    `binding:"required" bson:"package_apply_id" json:"package_apply_id"`
	Username       string    `binding:"required" bson:"username" json:"username" `
	ContactName    string    `binding:"required" bson:"contact_name" json:"contact_name"`
	Phone          string    `binding:"required" bson:"phone" json:"phone"`
	UseTotal       int64     `binding:"required" bson:"use_total" json:"use_total"` //预计用量 单位：张
	Status         int64     `binding:"required" bson:"status" json:"status"`       //状态 0：待处理； 1：已处理；
	Remark         string    `binding:"required" bson:"remark" json:"remark"`       //备注
	CreatedAt      time.Time `bson:"created_at"`
}

var (
	ErrPackageApplyNotFound = errors.New("user package apply not found")
	ErrPackageApplyCursor   = errors.New("user package apply Cursor err")
)

func QueryOnePackageApply(id string, username string) (*PackageApplyColl, error) {
	s := db.Manager.GetSession()
	defer s.Close()
	collection := new(PackageApplyColl)
	err := s.DB(db.Manager.DB).C("package_apply").Find(bson.M{
		"package_apply_id": id,
		"username":         username,
	}).One(collection)
	if err != nil {
		log.Error(fmt.Sprintf("find user package apply err ", err))
		if err == mgo.ErrNotFound {
			return nil, ErrPackageApplyNotFound
		} else if err == mgo.ErrCursor {
			return nil, ErrPackageApplyCursor
		}
		return nil, err
	}
	return collection, nil
}

func QueryPagingPackageApplys(pageIndex int, pageSize int, sort string, status int64) ([]*PackageApplyColl, int, error) {
	s := db.Manager.GetSession()
	defer s.Close()
	var result []*PackageApplyColl
	err := s.DB(db.Manager.DB).C("package_apply").Find(bson.M{
		"status": status,
	}).Sort(sort).Skip((pageIndex - 1) * pageSize).Limit(pageSize).All(&result)

	records, err := s.DB(db.Manager.DB).C("package_apply").Find(bson.M{
		"status": status,
	}).Count()
	if err != nil {
		log.Error(fmt.Sprintf("find paging user package apply err", err))
		if err == mgo.ErrNotFound {
			return result, 0, ErrPackageApplyNotFound
		}
		return result, 0, err
	}
	return result, records, nil
}

func UpdatePagingPackageApply(id string, username string, remark string) error {
	s := db.Manager.GetSession()
	defer s.Close()

	err := s.DB(db.Manager.DB).C("package_apply").Update(bson.M{
		"package_apply_id": id,
		"username":         username,
	}, bson.M{"$set": bson.M{"status": 1, "remark": remark}})
	if err != nil {
		log.Error(fmt.Sprintf("update user package apply err ", err))
		if err == mgo.ErrNotFound {
			return ErrPackageApplyNotFound
		} else if err == mgo.ErrCursor {
			return ErrPackageApplyCursor
		}
		return err
	}
	return nil
}

func DeletePackageApply(id string, username string) error {
	s := db.Manager.GetSession()
	defer s.Close()

	err := s.DB(db.Manager.DB).C("package_apply").Remove(bson.M{
		"package_apply_id": id,
		"username":         username,
	})
	if err != nil {
		log.Error(fmt.Sprintf("delete user package apply err ", err))
		if err == mgo.ErrNotFound {
			return ErrPackageApplyNotFound
		} else if err == mgo.ErrCursor {
			return ErrPackageApplyCursor
		}
		return err
	}
	return nil
}
