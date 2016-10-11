package managermodel

import (
	db "ManageCenter/service/db"

	"gopkg.in/mgo.v2/bson"
)

type ManagerColl struct {
	ManagerId   string `bson:"manager_id" binding:"required"`
	Managername string `binding:"required"`
	Password    string `binding:"required"`
}

func (m *ManagerColl) Save() error {
	s := db.Manager.GetSession()
	defer s.Close()
	return s.DB(db.Manager.DB).C("manager").Insert(&ManagerColl{bson.NewObjectId().String(), m.Managername, m.Password})
}

func Get(managername string) (*ManagerColl, error) {
	collection := new(ManagerColl)
	s := db.Manager.GetSession()
	defer s.Close()
	err := s.DB(db.Manager.DB).C("manager").Find(bson.M{
		"managername": managername,
	}).One(collection)
	if err != nil {
		return nil, err
	}
	return collection, nil
}

func QueryManager(managerUsername string) (int, error) {
	s := db.Manager.GetSession()
	defer s.Close()
	count, err := s.DB(db.Manager.DB).C("manager").Find(bson.M{"managername": managerUsername}).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func DeleteManager(managerUsername string) error {
	s := db.Manager.GetSession()
	defer s.Close()
	err := s.DB(db.Manager.DB).C("manager").Remove(bson.M{"managername": managerUsername})
	if err != nil {
		return err
	}
	return nil
}
