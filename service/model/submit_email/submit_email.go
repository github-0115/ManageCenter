package submit_email

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
	ErrEmailNotFound = errors.New("Email not found")
	ErrUserCursor    = errors.New("Cursor err")
)

type SubmitEmailColl struct {
	UserId    string    `bson:"user_id"`
	EmailId   string    `bson:"email_id"`
	Email     string    `bson:"email"`
	Count     int64     `bson:"count"`
	Status    int64     `bson:"status"` //默认0是开放
	CreatedAt time.Time `bson:"created_at"`
}

func QueryUserEmail() ([]*SubmitEmailColl, error) {
	s := db.User.GetSession()
	defer s.Close()
	var coll []*SubmitEmailColl
	err := s.DB(db.User.DB).C("submit_email").Find(bson.M{"status": 0}).All(&coll)
	if err != nil {
		log.Error(fmt.Sprintf("find email err ", err))
		if err == mgo.ErrNotFound {
			return nil, ErrEmailNotFound
		} else if err == mgo.ErrCursor {
			return nil, ErrUserCursor
		}

		return nil, err
	}
	return coll, nil
}

func ChangeEmailCount(id string) error {
	s := db.User.GetSession()
	defer s.Close()
	err := s.DB(db.User.DB).C("submit_email").Update(
		bson.M{"email_id": id}, bson.M{"$inc": bson.M{
			"count": 1,
		}})

	if err != nil {
		log.Error("change email failed.emailId=", id)
		if err == mgo.ErrNotFound {
			return ErrEmailNotFound
		} else if err == mgo.ErrCursor {
			return ErrUserCursor
		}
	}
	return nil
}
