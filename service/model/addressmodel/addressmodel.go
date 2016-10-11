package addressmodel

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
	ErrUserAdsNotFound = errors.New("user ads not found")
	ErrNumNotFound     = errors.New("find user ads num err")
	ErrAdsCursor       = errors.New("Cursor err")
)

type AddressInfo struct {
	AddressId string    `bson:"address_id"`
	Username  string    `bson:"username"`
	Province  string    `bson:"province"`
	City      string    `bson:"city"`
	District  string    `bson:"district"`
	Address   string    `bson:"address"`
	ZipCode   string    `bson:"zipcode"`
	PhoneNum  string    `bson:"phonenum"`
	Name      string    `bson:"name"`
	TelNum    string    `bson:"telnum"`
	Status    int64     `bson:"status"` // 0 一般 1 默认 2 删除不可用
	CreatedAt time.Time `bson:"created_at"`
}

func QueryAddressCount(addressId string, username string) (int, error) {
	s := db.User.GetSession()
	defer s.Close()

	result, err := s.DB(db.User.DB).C("address").Find(bson.M{
		"address_id": addressId,
		"username":   username,
		"status":     bson.M{"$ne": 2},
	}).Count()
	if err != nil {
		log.Error(fmt.Sprintf("find user ads  err ", err))
		if err == mgo.ErrNotFound {
			return 0, ErrUserAdsNotFound
		} else if err == mgo.ErrCursor {
			return 0, ErrAdsCursor
		}
		return 0, err
	}
	return result, nil
}

func QueryOneAddress(addressId string, username string) (*AddressInfo, error) {
	s := db.User.GetSession()
	defer s.Close()
	addressInfo := new(AddressInfo)
	err := s.DB(db.User.DB).C("address").Find(bson.M{
		"address_id": addressId,
		"username":   username,
	}).One(&addressInfo)
	if err != nil {
		log.Error(fmt.Sprintf("find user ads  err =%s", err))
		if err == mgo.ErrNotFound {
			return nil, ErrUserAdsNotFound
		} else if err == mgo.ErrCursor {
			return nil, ErrAdsCursor
		}
		return nil, err
	}
	return addressInfo, nil
}
