package main

import (
	billingmodel "ManageCenter/service/model/billingmodel"
	finalbillingmodel "ManageCenter/service/model/final_billmodel"
	pornday "ManageCenter/service/model/porn_day"
	servpkgmodel "ManageCenter/service/model/servicepackagemodel"
	usermodel "ManageCenter/service/model/usermodel"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	log "github.com/inconshreveable/log15"
)

func GetFirstAndEndOfLastMonth(t time.Time) (time.Time, time.Time) {
	t = t.AddDate(0, -1, 0)
	loc := t.Location()
	year, month, _ := t.Date()
	firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, loc)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	lastOfMonth = time.Date(lastOfMonth.Year(), lastOfMonth.Month(), lastOfMonth.Day(), 23, 59, 59, 0, lastOfMonth.Location())
	return firstOfMonth, lastOfMonth
}

func RoundUp(input float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * input
	round = math.Ceil(digit)
	newVal = round / pow
	return newVal
}

func CalTotalFee(used_total float64, confirmed_total float64, review_total float64) (float64, float64, float64) {
	//var totalFee float64
	var favorRange float64
	var paidAmount float64
	used_total /= 10000
	confirmed_total /= 10000
	review_total /= 10000
	//totalFee = 25*confirmed_total + 6.25*review_total

	switch {
	case used_total <= 300:
		favorRange = 1.0
		paidAmount = 25*confirmed_total + 6.25*review_total
	case 300 < used_total && used_total <= 1500:
		favorRange = 0.9
		paidAmount = 22.5*confirmed_total + 5.63*review_total
	case 1500 < used_total && used_total <= 3000:
		favorRange = 0.85
		paidAmount = 21.25*confirmed_total + 5.31*review_total
	case 3000 < used_total && used_total <= 6000:
		favorRange = 0.75
		paidAmount = 18.75*confirmed_total + 4.69*review_total
	case 6000 < used_total && used_total <= 14000:
		favorRange = 0.7
		paidAmount = 17.5*confirmed_total + 4.38*review_total
	case 14000 < used_total:
		favorRange = 0.65
		paidAmount = 16.25*confirmed_total + 4.06*review_total
	}

	return RoundUp(paidAmount, 2), favorRange, RoundUp(paidAmount, 2)

}

/*每月初自动生成上个月账单*/
func Billing() {
	now := time.Now()

	first, end := GetFirstAndEndOfLastMonth(now)
	// 已审核的、付费的用户
	billing_users, err := usermodel.GetAllBillingUsers()
	if err != nil {
		log.Error(fmt.Sprintf("time Parse err%v", err))
		return

	}

	for _, user := range billing_users {
		username := user.Username
		var billType = "月结"
		userPackage, err := servpkgmodel.GetUserServicePackage(username)
		if err != nil {
			log.Error(fmt.Sprintf("find user service package err ", err))
		} else {
			billType = userPackage.Period
		}

		day_total_sum_arr, err := pornday.GetUserDayTotalInPeriod(username, first, end)
		var used_total int64
		var confirmed_total int64
		var review_total int64
		for _, day := range day_total_sum_arr {
			used_total += day.Total
			confirmed_total = confirmed_total + day.P + day.S + day.N
			review_total = review_total + day.PReview + day.SReview + day.NReview
		}
		totalFee, _, _ := CalTotalFee(float64(used_total), float64(confirmed_total), float64(review_total))
		created_at := user.CreatedAt
		start := first
		if first.Before(created_at) {
			start = created_at
		}

		timestamp := time.Now().Unix()
		tm := time.Unix(timestamp, 0)

		monthBill := &billingmodel.BillColl{
			BillingId:   tm.Format("20060102030405") + strconv.Itoa(rand.Intn(8999)+1000),
			Username:    username,
			UsedTotal:   used_total,
			TotalFee:    totalFee,
			PeriodStart: start,
			PeriodEnd:   end,
			Status:      0,
			BillType:    billType,
			CreatedAt:   now,
		}
		err = monthBill.Save()
		if err != nil {
			log.Error(fmt.Sprintf("save user month bill err:%v, username:%v", err, username))
			continue
		}

		finalBill(monthBill)
	}
}

/*根据用户结算方式自动集成final账单*/
func finalBill(bill *billingmodel.BillColl) {

	switch bill.BillType {

	case "月结":
		oneBill(bill)

	case "季结":
		mergeBill(bill)

	case "半年结":
		mergeBill(bill)

	case "年结":
		mergeBill(bill)

	default:
		oneBill(bill)
	}

}

/*月结用户集成账单*/
func oneBill(bill *billingmodel.BillColl) {

	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	finalbillInfo := &finalbillingmodel.FinalBillColl{
		FinalBillId: tm.Format("20060102030405") + strconv.Itoa(rand.Intn(8999)+1000),
		Username:    bill.Username,
		BillingId:   []string{bill.BillingId},
		UsedTotal:   bill.UsedTotal,
		BillType:    bill.BillType,
		PeriodStart: bill.PeriodStart,
		PeriodEnd:   bill.PeriodEnd,
		TotalFee:    bill.TotalFee, //多少钱 /分
		PaidAmount:  bill.TotalFee, //实际多少钱 /分
		FavorRange:  1,             //优惠
		Status:      0,
		Remark:      "待处理",
		CreatedAt:   time.Now(),
	}

	err := finalbillInfo.Save()
	if err != nil {
		log.Error(fmt.Sprintf("save user final bill save failed.username=%s, err=%#v", bill.Username, err))
		finalbillInfo.Save()
	}
}

/*季、半年、年结用户集成账单*/
func mergeBill(bill *billingmodel.BillColl) {
	finalUserOldBill, err := finalbillingmodel.QueryUserBill(bill.Username, bill.Status, bill.BillType)
	if err != nil {
		log.Error(fmt.Sprintf("query user finalbill failed.username=%s, err=%#v", bill.Username, err))
	}

	if finalUserOldBill == nil {
		timestamp := time.Now().Unix()
		tm := time.Unix(timestamp, 0)

		finalUserOldBill = &finalbillingmodel.FinalBillColl{
			FinalBillId: tm.Format("20060102030405") + strconv.Itoa(rand.Intn(8999)+1000),
			Username:    bill.Username,
			BillingId:   []string{bill.BillingId},
			UsedTotal:   bill.UsedTotal,
			BillType:    bill.BillType,
			PeriodStart: bill.PeriodStart,
			PeriodEnd:   bill.PeriodEnd,
			TotalFee:    bill.TotalFee, //多少钱 /分
			PaidAmount:  bill.TotalFee, //实际多少钱 /分
			FavorRange:  1,             //优惠
			Status:      0,
			Remark:      "待处理",
			CreatedAt:   time.Now(),
		}
	} else {
		finalUserOldBill.BillingId = append(finalUserOldBill.BillingId, bill.BillingId)
		finalUserOldBill.UsedTotal += bill.UsedTotal
		finalUserOldBill.TotalFee += bill.TotalFee
		finalUserOldBill.PaidAmount += bill.TotalFee
		finalUserOldBill.PeriodEnd = bill.PeriodEnd
	}

	err = finalbillingmodel.DeleteUserBill(bill.Username, bill.Status, bill.BillType)
	if err != nil {
		log.Error(fmt.Sprintf("delete user final bill save failed.username=%s, err=%#v", bill.Username, err))
		finalbillingmodel.DeleteUserBill(bill.Username, bill.Status, bill.BillType)
	}

	err = finalUserOldBill.Save()
	if err != nil {
		log.Error(fmt.Sprintf("save user final bill save failed.username=%s, err=%#v", bill.Username, err))
		finalUserOldBill.Save()
	}
}

func main() {
	Billing()
}
