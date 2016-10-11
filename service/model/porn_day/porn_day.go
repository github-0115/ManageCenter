package porn_day

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
	PornDayNotFount = errors.New("result of pornDay not found")
)

type PornDay struct {
	User           string
	CreatedAt      time.Time `bson:"created_at"`
	Total          int64     `bson:"total"`
	P              int64     `bson:"p"`
	S              int64     `bson:"s"`
	N              int64     `bson:"n"`
	PReview        int64     `bson:"p_review"`
	SReview        int64     `bson:"s_review"`
	NReview        int64     `bson:"n_review"`
	PConfirm       int64     `bson:"p_confirm"`
	SConfirm       int64     `bson:"s_confirm"`
	NConfirm       int64     `bson:"n_confirm"`
	IsScreen       int64     `bson:"isScreen"`
	NotScreen      int64     `bson:"notScreen"`
	ScreenshotTota int64     `bson:"screenshot_tota"`
}

type PornDaySum struct {
	CreatedAt time.Time `bson:"_id"`
	Total     int64
	P         int64 `bson:"p"`
	S         int64 `bson:"s"`
	N         int64 `bson:"n"`
	PReview   int64 `bson:"p_review"`
	SReview   int64 `bson:"s_review"`
	NReview   int64 `bson:"n_review"`
	IsScreen  int64 `bson:"isScreen"`
	NotScreen int64 `bson:"notScreen"`
}

type UserDayTotalSum struct {
	CreatedAt time.Time `bson:"_id"`
	Total     int64
}

func GetUserDayTotalSum(username string, startTime time.Time, endTime time.Time) ([]*UserDayTotalSum, error) {
	s := db.Stats.GetSession()
	defer s.Close()
	var result []*UserDayTotalSum
	err := s.DB(db.Stats.DB).C("porn_day").Pipe([]bson.M{
		{
			"$match": bson.M{
				"user":       username,
				"created_at": bson.M{"$gt": startTime, "$lt": endTime}},
		},
		{
			"$group": bson.M{
				"_id":   "$created_at",
				"total": bson.M{"$sum": "$total"},
			},
		},
		{"$sort": bson.M{"_id": 1}}}).All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetUserDayTotalInPeriod(username string, startTime time.Time, endTime time.Time) ([]*PornDaySum, error) {
	s := db.Stats.GetSession()
	defer s.Close()
	var result []*PornDaySum
	err := s.DB(db.Stats.DB).C("porn_day").Pipe([]bson.M{
		{
			"$match": bson.M{
				"user":       username,
				"created_at": bson.M{"$gt": startTime, "$lt": endTime}},
		}, {
			"$group": bson.M{
				"_id": "$created_at",
				"total": bson.M{
					"$sum": "$total"},
				"p": bson.M{
					"$sum": "$p"},
				"s": bson.M{
					"$sum": "$s"},
				"n": bson.M{
					"$sum": "$n"},
				"p_review": bson.M{
					"$sum": "$p_review"},
				"s_review": bson.M{
					"$sum": "$s_review"},
				"n_review": bson.M{
					"$sum": "$n_review"},
				"isScreen": bson.M{
					"$sum": "$isScreen"},
				"notScreen": bson.M{
					"$sum": "$notScreen"},
			},
		}, {
			"$sort": bson.M{
				"_id": 1},
		}}).All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetUserDayTotal(username string, startTime time.Time, endTime time.Time) ([]*PornDaySum, error) {
	s := db.Stats.GetSession()
	defer s.Close()
	var result []*PornDaySum
	err := s.DB(db.Stats.DB).C("porn_day").Pipe([]bson.M{
		{
			"$match": bson.M{
				"user": username},
		}, {
			"$group": bson.M{
				"_id": "$created_at",
				"total": bson.M{
					"$sum": "$total"},
				"p": bson.M{
					"$sum": "$p"},
				"s": bson.M{
					"$sum": "$s"},
				"n": bson.M{
					"$sum": "$n"},
				"p_review": bson.M{
					"$sum": "$p_review"},
				"s_review": bson.M{
					"$sum": "$s_review"},
				"n_review": bson.M{
					"$sum": "$n_review"},
				"isScreen": bson.M{
					"$sum": "$isScreen"},
				"notScreen": bson.M{
					"$sum": "$notScreen"},
			},
		}, {
			"$sort": bson.M{
				"_id": 1},
		}}).All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func QueryUserPornDay(username string, startTime time.Time, endTime time.Time) ([]*PornDay, error) {
	s := db.Stats.GetSession()
	defer s.Close()
	var result []*PornDay
	err := s.DB(db.Stats.DB).C("porn_day").Find(bson.M{
		"user":       username,
		"created_at": bson.M{"$gt": startTime, "$lt": endTime},
	}).Select(bson.M{"total": 1, "p": 1, "s": 1, "n": 1, "p_review": 1, "s_review": 1, "n_review": 1, "created_at": 1, "isScreen": 1, "notScreen": 1}).Sort("created_at").All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func QueryPornDay(startTime time.Time, endTime time.Time) ([]*PornDay, error) {
	s := db.Stats.GetSession()
	defer s.Close()
	var result []*PornDay
	err := s.DB(db.Stats.DB).C("porn_day").Find(bson.M{
		"created_at": bson.M{"$gt": startTime, "$lt": endTime},
	}).Select(bson.M{"user": 1, "total": 1}).Sort("-total").All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetPagingDayHisory(username string, pageIndex int, pageSize int, sort string) ([]*PornDay, int, error) {
	s := db.Stats.GetSession()
	defer s.Close()
	var result []*PornDay
	err := s.DB(db.Stats.DB).C("porn_day").Find(bson.M{
		"user": username,
	}).Sort(sort).Skip((pageIndex - 1) * pageSize).Limit(pageSize).All(&result)

	records, err := s.DB(db.Stats.DB).C("porn_day").Find(bson.M{
		"user": username,
	}).Count()
	if err != nil {
		log.Error(fmt.Sprintf("find user pornday err", err))
		if err == mgo.ErrNotFound {

			return result, 0, PornDayNotFount
		}
		return result, 0, err
	}
	return result, records, nil
}

func GetAllDayTotal(startTime time.Time, endTime time.Time) ([]*PornDaySum, error) {
	s := db.Stats.GetSession()
	defer s.Close()
	var result []*PornDaySum
	err := s.DB(db.Stats.DB).C("porn_day").Pipe([]bson.M{
		{
			"$match": bson.M{
				"created_at": bson.M{
					"$gt": startTime, "$lt": endTime}},
		}, {
			"$group": bson.M{
				"_id": "$created_at",
				"total": bson.M{
					"$sum": "$total"},
				"p": bson.M{
					"$sum": "$p"},
				"s": bson.M{
					"$sum": "$s"},
				"n": bson.M{
					"$sum": "$n"},
				"p_review": bson.M{
					"$sum": "$p_review"},
				"s_review": bson.M{
					"$sum": "$s_review"},
				"n_review": bson.M{
					"$sum": "$n_review"},
				"isScreen": bson.M{
					"$sum": "$isScreen"},
				"notScreen": bson.M{
					"$sum": "$notScreen"},
			},
		}, {
			"$sort": bson.M{
				"_id": 1},
		}}).All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetUserPagingDaySum(username string, pageIndex int, pageSize int, startTime time.Time, endTime time.Time) ([]*PornDay, int, error) {
	s := db.Stats.GetSession()
	defer s.Close()
	var result []*PornDay
	err := s.DB(db.Stats.DB).C("porn_day").Find(bson.M{
		"user": username,
		"created_at": bson.M{
			"$gt": startTime, "$lt": endTime},
	}).Sort("-created_at").Skip((pageIndex - 1) * pageSize).Limit(pageSize).All(&result)

	records, err := s.DB(db.Stats.DB).C("porn_day").Find(bson.M{
		"user": username,
		"created_at": bson.M{
			"$gt": startTime, "$lt": endTime},
	}).Count()
	if err != nil {
		return nil, 0, err
	}
	return result, records, nil
}

func SearchUserPagingDaySum(pageIndex int, pageSize int, sort string, startTime time.Time, endTime time.Time, searchField string, searchString string,
	searchOper string) ([]*PornDay, int, error) {

	s := db.Stats.GetSession()
	defer s.Close()
	var result []*PornDay
	err := s.DB(db.Stats.DB).C("porn_day").Find(bson.M{
		searchField: bson.M{
			searchOper: searchString},
		"created_at": bson.M{
			"$gt": startTime, "$lt": endTime},
	}).Sort(sort).Skip((pageIndex - 1) * pageSize).Limit(pageSize).All(&result)

	records, err := s.DB(db.Stats.DB).C("porn_day").Find(bson.M{
		searchField: bson.M{
			searchOper: searchString},
		"created_at": bson.M{
			"$gt": startTime, "$lt": endTime},
	}).Count()
	if err != nil {
		return nil, 0, err
	}
	return result, records, nil
}

func GetPagingDaySum(pageIndex int, pageSize int, startTime time.Time, endTime time.Time) ([]*PornDaySum, error) {
	s := db.Stats.GetSession()
	defer s.Close()
	var result []*PornDaySum
	err := s.DB(db.Stats.DB).C("porn_day").Pipe([]bson.M{
		{
			"$match": bson.M{
				"created_at": bson.M{
					"$gt": startTime, "$lt": endTime}},
		}, {
			"$group": bson.M{
				"_id": "$created_at",
				"total": bson.M{
					"$sum": "$total"},
				"p": bson.M{
					"$sum": "$p"},
				"s": bson.M{
					"$sum": "$s"},
				"n": bson.M{
					"$sum": "$n"},
				"isScreen": bson.M{
					"$sum": "$isScreen"},
				"notScreen": bson.M{
					"$sum": "$notScreen"},
				"p_review": bson.M{
					"$sum": "$p_review"},
				"s_review": bson.M{
					"$sum": "$s_review"},
				"n_review": bson.M{
					"$sum": "$n_review"},
			},
		}, {
			"$sort": bson.M{
				"_id": -1},
		}, {
			"$skip": (pageIndex - 1) * pageSize,
		}, {
			"$limit": pageSize,
		}}).All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetPagingDaySumCount(startTime time.Time, endTime time.Time) (int, error) {
	s := db.Stats.GetSession()
	defer s.Close()
	var result []*PornDaySum
	err := s.DB(db.Stats.DB).C("porn_day").Pipe([]bson.M{
		{
			"$match": bson.M{
				"created_at": bson.M{
					"$gt": startTime, "$lt": endTime}},
		}, {
			"$group": bson.M{
				"_id": "$created_at",
			},
		}}).All(&result)

	if err != nil {
		return 0, err
	}

	return len(result), nil
}
