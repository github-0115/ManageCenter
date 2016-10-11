package main

import (
	finalbillingmodel "ManageCenter/service/model/final_billmodel"
	keymodel "ManageCenter/service/model/keymodel"
	usermodel "ManageCenter/service/model/usermodel"
	"fmt"
	"time"

	log "github.com/inconshreveable/log15"
)

func main() {

	notSettleBills, err := finalbillingmodel.GetNotSettleBills()
	if err != nil {
		log.Error(fmt.Sprintf("find not settle final bill err", err))
		if err == finalbillingmodel.ErrBillNotFound {
			log.Info(fmt.Sprintf("user final bill all settle"))
			return
		}
		return
	}

	if notSettleBills != nil {
		HandelNotSettleBills(notSettleBills)
	}
}

func HandelNotSettleBills(bills []*finalbillingmodel.FinalBillColl) {
	for _, res := range bills {
		if isSettleDay(res) == 1 {
			DeactivateUserKey(res.Username)
		}
	}
}

func DeactivateUserKey(username string) {

	userColl, err := usermodel.QueryUser(username)
	if err != nil {
		log.Error("user not found, err = %s", err)
	}

	err = keymodel.DeactivateKey(userColl.UserId)
	if err != nil {
		log.Error(fmt.Sprintf("find user key err, User = %s", username))
	}

}

func isSettleDay(bill *finalbillingmodel.FinalBillColl) int {
	ms := months(bill.PeriodStart, bill.PeriodEnd)
	var isDays int
	switch bill.BillType {
	case "月结":
		if ms == 1 {
			isDays = 1
		}
	case "季结":
		if ms == 3 {
			isDays = 1
		}
	case "半年结":
		if ms == 6 {
			isDays = 1
		}
	case "年结":
		if ms == 12 {
			isDays = 1
		}
	default:
		isDays = 0
	}
	return isDays
}

func months(t1, t2 time.Time) int {
	y1, m1, _ := t1.Date()
	y2, m2, _ := t2.Date()

	return (y2-y1)*12 + (monthToInt(m2) - monthToInt(m1)) + 1
}

func monthToInt(m time.Month) int {
	switch m {
	case time.January:
		return 1
	case time.February:
		return 2
	case time.March:
		return 3
	case time.April:
		return 4
	case time.May:
		return 5
	case time.June:
		return 6
	case time.July:
		return 7
	case time.August:
		return 8
	case time.September:
		return 9
	case time.October:
		return 10
	case time.November:
		return 11
	case time.December:
		return 12

	default:
		return 1
	}
	return 1
}
