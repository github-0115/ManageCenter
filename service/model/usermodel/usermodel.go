package usermodel

import (
	db "ManageCenter/service/db"
	"errors"
	"fmt"
	"time"

	log "github.com/inconshreveable/log15"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserColl struct {
	UserId      string    `bson:"user_id" binding:"required"`
	Username    string    `binding:"required" json:"username"`
	Password    string    `binding:"required" json:"password"`
	Email       string    `binding:"required" json:"email"`
	EmailStatus int64     `binding:"required" json:"email_status"`
	PhoneNum    string    `binding:"required" json:"phonenum"`
	Name        string    `binding:"required" json:"name"`
	Company     string    `bson:"company,omitempty" json:"company"`
	Reason      string    `bson:"reason" binding:"required" json:"reason"`
	Remark      string    `bson:"remark" json:"remark"`
	Status      int64     `binding:"required" json:"status"` //暂定： 0 审核通过  1 未审核  2 审核未通过
	Type        int64     `binding:"required" json:"type" `  //暂定： 0 免费  1 预付费  2 后付费  3 代理商
	Period      string    `json:"period"`                    // 结算类型， 月结、半年结等
	CreatedAt   time.Time `bson:"created_at" binding:"required" json:"created_at"`
}

type userStatus struct {
	Active  int64
	Ban     int64
	NotPass int64
}

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserCursor   = errors.New("Cursor err")
)

var UserStatus = &userStatus{0, 1, 2}

func (t *UserColl) Save() error {
	s := db.User.GetSession()
	defer s.Close()
	return s.DB(db.User.DB).C("user").Insert(&t)
}

func QueryUser(username string) (*UserColl, error) {
	coll := new(UserColl)
	s := db.User.GetSession()
	defer s.Close()
	err := s.DB(db.User.DB).C("user").Find(bson.M{
		"username": username,
	}).One(coll)
	if err != nil {
		log.Error(fmt.Sprintf("find user err ", err))
		if err == mgo.ErrNotFound {
			return nil, ErrUserNotFound
		} else if err == mgo.ErrCursor {
			return nil, ErrUserCursor
		}
		return nil, err
	}
	return coll, nil
}

func QueryUserNum(userId string, username string) (int, error) {
	s := db.User.GetSession()
	defer s.Close()
	count, err := s.DB(db.User.DB).C("user").Find(bson.M{
		"$or": []interface{}{
			bson.M{"user_id": userId},
			bson.M{"username": bson.M{"$regex": "^" + username + "$", "$options": "i"}},
		},
	}).Count()
	if err != nil {
		log.Error(fmt.Sprintf("find user num err ", err))
		if err == mgo.ErrNotFound {
			return 0, ErrUserNotFound
		} else if err == mgo.ErrCursor {
			return 0, ErrUserCursor
		}
		return 0, err
	}
	return count, nil
}

func QueryUserCount(key string, value string) (int, error) {
	s := db.User.GetSession()
	defer s.Close()
	count, err := s.DB(db.User.DB).C("user").Find(bson.M{
		key: value,
	}).Count()
	if err != nil {
		log.Error(fmt.Sprintf("find user num err ", err))
		if err == mgo.ErrNotFound {
			return 0, ErrUserNotFound
		} else if err == mgo.ErrCursor {
			return 0, ErrUserCursor
		}
		return 0, err
	}
	return count, nil
}

func QuerySubmitCount(startTime time.Time, endTime time.Time) ([]*UserColl, error) {
	s := db.User.GetSession()
	defer s.Close()
	var result []*UserColl
	err := s.DB(db.User.DB).C("user").Find(bson.M{
		"created_at": bson.M{"$gt": startTime, "$lt": endTime},
	}).All(&result)
	if err != nil {
		log.Error(fmt.Sprintf("find user num err ", err))
		if err == mgo.ErrNotFound {
			return nil, ErrUserNotFound
		} else if err == mgo.ErrCursor {
			return nil, ErrUserCursor
		}
		return nil, err
	}
	return result, nil
}

func CheckUsers(username string, status int64, remark string) (int64, error) {
	s := db.User.GetSession()
	defer s.Close()

	if status == 0 {
		err := s.DB(db.User.DB).C("user").Update(bson.M{"username": username}, bson.M{"$set": bson.M{"status": UserStatus.Active, "remark": remark}})
		if err != nil {
			log.Error(fmt.Sprintf("update err", err))
			if err == mgo.ErrNotFound {

				return UserStatus.Ban, mgo.ErrNotFound
			}
			return UserStatus.Ban, err
		}
		return UserStatus.Active, nil
	} else if status == 2 {
		err := s.DB(db.User.DB).C("user").Update(bson.M{"username": username}, bson.M{"$set": bson.M{"status": UserStatus.NotPass, "remark": remark}})
		if err != nil {
			log.Error(fmt.Sprintf("update err", err))
			if err == mgo.ErrNotFound {

				return UserStatus.Ban, mgo.ErrNotFound
			}
			return UserStatus.Ban, err
		}
		return UserStatus.NotPass, nil
	} else {
		err := s.DB(db.User.DB).C("user").Update(bson.M{"username": username}, bson.M{"$set": bson.M{"status": UserStatus.Ban, "remark": remark}})
		if err != nil {
			log.Error(fmt.Sprintf("update err", err))
			if err == mgo.ErrNotFound {

				return UserStatus.Active, ErrUserNotFound
			}
			return UserStatus.Active, err
		}
		return UserStatus.Ban, nil
	}
}

func ChangeUserInfo(user string, key string, value string) (bool, error) {
	s := db.User.GetSession()
	defer s.Close()

	err := s.DB(db.User.DB).C("user").Update(bson.M{"username": user}, bson.M{"$set": bson.M{key: value}})
	if err != nil {
		log.Error(fmt.Sprintf("update user err", err))
		if err == mgo.ErrNotFound {
			return false, ErrUserNotFound
		}
		return false, err
	}
	return true, nil
}

func ChangeUserRights(user string, key string, value string) (bool, error) {
	s := db.User.GetSession()
	defer s.Close()
	err := s.DB(db.User.DB).C("user").Update(bson.M{"username": user}, bson.M{"$set": bson.M{key: value}})
	if err != nil {
		log.Error(fmt.Sprintf("update user rights err", err))
		if err == mgo.ErrNotFound {
			return false, ErrUserNotFound
		}
		return false, err
	}
	return true, nil
}

func ChangeUserType(user string, value int64) (bool, error) {
	s := db.User.GetSession()
	defer s.Close()

	err := s.DB(db.User.DB).C("user").Update(bson.M{
		"$or": []bson.M{
			bson.M{"user_id": user},
			bson.M{"username": bson.M{"$regex": "^" + user + "$", "$options": "i"}},
		},
	}, bson.M{"$set": bson.M{"type": value}})
	if err != nil {
		log.Error(fmt.Sprintf("update user err", err))
		if err == mgo.ErrNotFound {

			return false, ErrUserNotFound
		}
		return false, err
	}
	return true, nil

}

func QueryEmailCount(key string, value string) (int, error) {
	s := db.User.GetSession()
	defer s.Close()
	count, err := s.DB(db.User.DB).C("user").Find(bson.M{
		key:            bson.M{"$regex": "^" + value + "$", "$options": "i"},
		"email_status": 1,
	}).Count()
	if err != nil {
		log.Error(fmt.Sprintf("find user num err ", err))
		if err == mgo.ErrNotFound {
			return 0, ErrUserNotFound
		} else if err == mgo.ErrCursor {
			return 0, ErrUserCursor
		}
		return 0, err
	}
	return count, nil
}

func GetAllBillingUsers() ([]UserColl, error) {
	s := db.User.GetSession()
	defer s.Close()
	var result []UserColl
	// 已审核的、付费的用户
	err := s.DB(db.User.DB).C("user").Find(bson.M{"type": bson.M{"$gt": 0}, "status": 0}).All(&result)
	if err != nil {
		log.Error(fmt.Sprintf("find err", err))
		if err == mgo.ErrNotFound {

			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return result, nil
}

func GetPagingPaidUsers(pageIndex int, pageSize int, sort string, sord string) ([]UserColl, int, error) {
	s := db.User.GetSession()
	defer s.Close()
	var result []UserColl
	err := s.DB(db.User.DB).C("user").Find(bson.M{"type": 1}).Sort(sort).Skip((pageIndex - 1) * pageSize).Limit(pageSize).All(&result)
	records, err := s.DB(db.User.DB).C("user").Find(bson.M{"type": 1}).Count()
	if err != nil {
		log.Error(fmt.Sprintf("find err", err))
		if err == mgo.ErrNotFound {

			return result, 0, ErrUserNotFound
		}
		return result, 0, err
	}
	return result, records, nil
}

func GetPagingUsers(pageIndex int, pageSize int, sort string, sord string) ([]UserColl, int, error) {
	s := db.User.GetSession()
	defer s.Close()
	var result []UserColl
	err := s.DB(db.User.DB).C("user").Find(nil).Sort(sort).Skip((pageIndex - 1) * pageSize).Limit(pageSize).All(&result)
	records, err := s.DB(db.User.DB).C("user").Find(nil).Count()
	if err != nil {
		log.Error(fmt.Sprintf("find err", err))
		if err == mgo.ErrNotFound {

			return result, 0, ErrUserNotFound
		}
		return result, 0, err
	}
	return result, records, nil
}

func SearchPagingUsers(pageIndex int, pageSize int, sort string, sord string, searchField string, searchString string, searchOper string) ([]UserColl, int, error) {
	s := db.User.GetSession()
	defer s.Close()
	var result []UserColl
	err := s.DB(db.User.DB).C("user").Find(bson.M{searchField: bson.M{searchOper: searchString}}).Sort(sort).Skip((pageIndex - 1) * pageSize).Limit(pageSize).All(&result)
	records, err := s.DB(db.User.DB).C("user").Find(bson.M{searchField: bson.M{searchOper: searchString}}).Count()
	if err != nil {
		if err == mgo.ErrNotFound {

			return result, 0, ErrUserNotFound
		}
		return result, 0, err
	}
	return result, records, nil
}

func StatSearchPagingUsers(pageIndex int, pageSize int, sort string, sord string, searchString string) ([]*UserColl, int, error) {
	s := db.User.GetSession()
	defer s.Close()
	var result []*UserColl
	err := s.DB(db.User.DB).C("user").Find(bson.M{
		"status": 0,
		"$or": []interface{}{
			bson.M{"username": bson.M{"$regex": searchString}},
			bson.M{"company": bson.M{"$regex": searchString}},
		}}).Sort(sort).Skip((pageIndex - 1) * pageSize).Limit(pageSize).All(&result)
	records, err := s.DB(db.User.DB).C("user").Find(bson.M{
		"status": 0,
		"$or": []interface{}{
			bson.M{"username": bson.M{"$regex": searchString}},
			bson.M{"company": bson.M{"$regex": searchString}},
		}}).Count()
	if err != nil {
		if err == mgo.ErrNotFound {

			return result, 0, ErrUserNotFound
		}
		return result, 0, err
	}
	return result, records, nil
}

func DeleteUser(id string, username string) error {
	s := db.User.GetSession()
	defer s.Close()
	err := s.DB(db.User.DB).C("user").Remove(bson.M{"username": username})
	if err != nil {
		log.Error(fmt.Sprintf("delete user  err", err))
		if err == mgo.ErrNotFound {

			return ErrUserNotFound
		}
		return err
	}
	return nil
}
