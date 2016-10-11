package walletmodel

import (
	db "ManageCenter/service/db"
	"errors"
	"fmt"
	"time"

	log "github.com/inconshreveable/log15"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Wallet struct {
	Username       string       `bson:"username" `
	Balance        float64      `bson:"balance" `
	Total          float64      `bson:"total" `
	InvoicedAmount float64      `bson:"invoiced_amount" `
	Status         WalletStatus `bson:"status" `
	UpdatedAt      time.Time    `bson:"updated_at"`
	CreatedAt      time.Time    `bson:"created_at"`
}

type WalletStatus int64

const (
	WalletEnable WalletStatus = iota
	WalletDisable
	WalletArrears
)

var ErrUserWalletNotFound = errors.New("user Wallet not found")

func Temp(walletStatus WalletStatus) string {
	var status string
	switch walletStatus {
	case WalletEnable:
		status = "enable"
	case WalletDisable:
		status = "disable"
	case WalletArrears:
		status = "arrears"
	}
	return status
}

func (t *Wallet) Save() error {
	s := db.User.GetSession()
	defer s.Close()
	return s.DB(db.User.DB).C("wallet").Insert(&t)
}

func QueryWallet(username string) (*Wallet, error) {
	s := db.User.GetSession()
	defer s.Close()
	result := new(Wallet)
	err := s.DB(db.User.DB).C("wallet").Find(bson.M{
		"username": username,
	}).One(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func DeleteWallet(username string) error {
	s := db.User.GetSession()
	defer s.Close()
	err := s.DB(db.User.DB).C("wallet").Remove(bson.M{"username": username})
	if err != nil {
		log.Error(fmt.Sprintf("delete user wallet err", err))
		if err == mgo.ErrNotFound {

			return ErrUserWalletNotFound
		}
		return err
	}
	return nil
}
