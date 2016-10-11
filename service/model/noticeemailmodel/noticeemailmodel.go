package noticeemailmodel

import (
	db "ManageCenter/service/db"
	"errors"
	"fmt"
	"time"

	log "github.com/inconshreveable/log15"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type NoticeEmailColl struct {
	EmailId   string    `bson:"email_id"`
	Email     string    `bson:"email"`
	CreatedAt time.Time `bson:"created_at"`
}

var (
	ErrEmailNotFound    = errors.New("Email not found")
	ErrEmailNumNotFound = errors.New("Email num err")
	ErrUserCursor       = errors.New("Cursor err")
)

func (t *NoticeEmailColl) Save() error {
	s := db.Manager.GetSession()
	defer s.Close()
	return s.DB(db.Manager.DB).C("notice_email").Insert(&t)
}

func QueryNoticeAllEmail() ([]*NoticeEmailColl, error) {
	s := db.Manager.GetSession()
	defer s.Close()
	var coll []*NoticeEmailColl
	err := s.DB(db.Manager.DB).C("notice_email").Find(nil).All(&coll)
	if err != nil {
		log.Error(fmt.Sprintf("find all notice  email  err ", err))
		if err == mgo.ErrNotFound {
			return nil, ErrEmailNotFound
		} else if err == mgo.ErrCursor {
			return nil, ErrUserCursor
		}
		return nil, err
	}
	return coll, nil
}

func QueryNoticeEmailNum(email string) (int, error) {
	s := db.Manager.GetSession()
	defer s.Close()
	count, err := s.DB(db.Manager.DB).C("notice_email").Find(bson.M{"email": email}).Count()
	if err != nil {
		log.Error(fmt.Sprintf("find notice email num err ", err))
		if err == mgo.ErrNotFound {
			return 0, ErrEmailNumNotFound
		} else if err == mgo.ErrCursor {
			return 0, ErrUserCursor
		}
		return 0, err
	}
	return count, nil
}

func QueryNoticeEmail(id string) (*NoticeEmailColl, error) {
	coll := new(NoticeEmailColl)
	s := db.Manager.GetSession()
	defer s.Close()
	err := s.DB(db.Manager.DB).C("notice_email").Find(bson.M{"email_id": id}).One(coll)
	if err != nil {
		log.Error(fmt.Sprintf("check notice email_id err ", err))
		if err == mgo.ErrNotFound {
			return nil, ErrEmailNotFound
		} else if err == mgo.ErrCursor {
			return nil, ErrUserCursor
		}
		return nil, err
	}
	return coll, nil
}

func DeleteNoticeEmail(email string) error {
	s := db.Manager.GetSession()
	defer s.Close()
	err := s.DB(db.Manager.DB).C("notice_email").Remove(bson.M{"email": email})

	if err != nil {
		log.Error("delete notice email failed.emailId=", email)
		if err == mgo.ErrNotFound {
			return ErrEmailNotFound
		} else if err == mgo.ErrCursor {
			return ErrUserCursor
		}
	}
	return nil
}
