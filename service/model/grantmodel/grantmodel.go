package grantmodel

import (
	db "ManageCenter/service/db"
	"errors"
	"fmt"
	"strings"
	"time"

	log "github.com/inconshreveable/log15"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Rights struct {
	Username  string    `binding:"required" json:"username"`
	Porn      int64     `bson:"porn" binding:"required" json:"porn"`
	Screen    int64     `bson:"screen" binding:"required" json:"screen"`
	Violence  int64     `bson:"violence" binding:"required" json:"violence"`
	CreatedAt time.Time `bson:"created_at" binding:"required" json:"created_at"`
}

var (
	ErrGrantNotFound = errors.New("user Grant not found")
	ErrGrantrCursor  = errors.New("Cursor err")
)

func (t *Rights) Save() error {
	s := db.User.GetSession()
	defer s.Close()
	return s.DB(db.User.DB).C("grant").Insert(&t)
}

func GetRights(user string) (*Rights, error) {
	s := db.User.GetSession()
	defer s.Close()

	rights := new(Rights)
	err := s.DB(db.User.DB).C("grant").Find(bson.M{"username": user}).One(&rights)
	if err != nil {
		return nil, err
	}
	return rights, nil
}

func UpdateRights(user string, key string, value int64) (bool, error) {
	s := db.User.GetSession()
	defer s.Close()

	err := s.DB(db.User.DB).C("grant").Update(bson.M{"username": user}, bson.M{"$set": bson.M{strings.ToLower(key): value}})
	if err != nil {
		log.Error(fmt.Sprintf("open user grant err ", err))
		if err == mgo.ErrNotFound {
			return false, ErrGrantNotFound
		} else if err == mgo.ErrCursor {
			return false, ErrGrantrCursor
		}
		return false, err
	}

	return true, err
}

func DeleteUserRights(username string) error {
	s := db.User.GetSession()
	defer s.Close()
	err := s.DB(db.User.DB).C("grant").Remove(bson.M{"username": username})
	if err != nil {
		log.Error(fmt.Sprintf("delete user grant err", err))
		if err == mgo.ErrNotFound {

			return ErrGrantrCursor
		}
		return err
	}
	return nil
}
