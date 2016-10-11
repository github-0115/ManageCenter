package keymodel

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
	TaskStatus        = TaskStatusStruct{0, 1, 2}
	ErrUserNotFound   = errors.New("user not found")
	ErrKeyNotFound    = errors.New("Key not found")
	ErrKeyNumNotFound = errors.New("Key num err")
	ErrUserCursor     = errors.New("Cursor err")
)

type TaskStatusStruct struct {
	Created int64
	Finish  int64
	FAILED  int64
}

type KeyColl struct {
	UserId    string    `bson:"user_id"`
	KeyId     string    `bson:"key_id"`
	AccessKey string    `bson:"access_key"`
	SecretKey string    `bson:"secret_key"`
	IsExcess  bool      `bson:"is_excess"` //是否超量
	IsActive  bool      `bson:"is_active"` //是否激活能用
	CreatedAt time.Time `bson:"created_at"`
}

type SignColl struct {
	UserId    string
	SecretKey string
	now       time.Time
}

func (t *KeyColl) Save() error {
	s := db.User.GetSession()
	defer s.Close()
	return s.DB(db.User.DB).C("key").Insert(&t)
}

func QueryKeyNum(userId string) (int, error) {
	s := db.User.GetSession()
	defer s.Close()
	count, err := s.DB(db.User.DB).C("key").Find(bson.M{
		"user_id": userId,
	}).Count()
	if err != nil {
		log.Error(fmt.Sprintf("find user key num err ", err))
		if err == mgo.ErrNotFound {
			return 0, ErrKeyNumNotFound
		} else if err == mgo.ErrCursor {
			return 0, ErrUserCursor
		}
		return 0, err
	}
	return count, nil
}

func QueryUserKey(userId string) ([]*KeyColl, error) {
	var coll []*KeyColl
	s := db.User.GetSession()
	defer s.Close()
	err := s.DB(db.User.DB).C("key").Find(bson.M{
		"user_id": userId,
	}).All(&coll)
	if err != nil {
		log.Error(fmt.Sprintf("find user key  err ", err))
		if err == mgo.ErrNotFound {
			return nil, ErrKeyNotFound
		} else if err == mgo.ErrCursor {
			return nil, ErrUserCursor
		}
		return nil, err
	}
	return coll, nil
}

func QueryKey(accessKey string, userId string) (*KeyColl, error) {
	coll := new(KeyColl)
	s := db.User.GetSession()
	defer s.Close()
	err := s.DB(db.User.DB).C("key").Find(bson.M{
		"access_key": accessKey,
		"user_id":    userId,
	}).One(coll)
	if err != nil {
		log.Error(fmt.Sprintf("find ak err ", err))
		if err == mgo.ErrNotFound {
			return nil, ErrKeyNotFound
		} else if err == mgo.ErrCursor {
			return nil, ErrUserCursor
		}

		return nil, err
	}
	return coll, nil
}

func DeactivateKey(id string) error {
	s := db.User.GetSession()
	defer s.Close()
	err := s.DB(db.User.DB).C("key").Update(
		bson.M{
			"$or": []bson.M{
				bson.M{"key_id": id},
				bson.M{"user_id": id},
			},
		},
		bson.M{"$set": bson.M{"is_active": false}})

	if err != nil {
		log.Error("deactivate key failed.id=", id)
		if err == mgo.ErrNotFound {
			return ErrKeyNotFound
		} else if err == mgo.ErrCursor {
			return ErrUserCursor
		}
	}
	return nil
}

func ActivateKey(id string) error {
	s := db.User.GetSession()
	defer s.Close()
	err := s.DB(db.User.DB).C("key").Update(
		bson.M{
			"$or": []bson.M{
				bson.M{"key_id": id},
				bson.M{"user_id": id},
			},
		},
		bson.M{"$set": bson.M{"is_active": true}})

	if err != nil {
		log.Error("activate key failed.keyId=", id)
		if err == mgo.ErrNotFound {
			return ErrKeyNotFound
		} else if err == mgo.ErrCursor {
			return ErrUserCursor
		}
	}
	return nil
}

func DeleteKey(keyId string) error {
	s := db.User.GetSession()
	defer s.Close()
	err := s.DB(db.User.DB).C("key").Remove(
		bson.M{"key_id": keyId})

	if err != nil {
		log.Error("delete key failed.keyId=", keyId)
		if err == mgo.ErrNotFound {
			return ErrKeyNotFound
		} else if err == mgo.ErrCursor {
			return ErrUserCursor
		}
	}
	return nil
}

func DeleteUserKey(id string) error {
	s := db.User.GetSession()
	defer s.Close()
	err := s.DB(db.User.DB).C("key").Remove(
		bson.M{"user_id": id})

	if err != nil {
		log.Error("delete key failed.keyId=", id)
		if err == mgo.ErrNotFound {
			return ErrKeyNotFound
		} else if err == mgo.ErrCursor {
			return ErrUserCursor
		}
	}
	return nil
}
