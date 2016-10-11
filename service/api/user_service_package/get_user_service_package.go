package user_service_package

import (
	pornday "ManageCenter/service/model/porn_day"
	servpkgmodel "ManageCenter/service/model/servicepackagemodel"
	vars "ManageCenter/service/vars"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

type Rep struct {
	UserType    int64  `json:"usertype"` //暂定： 0 免费  1 预付费  2 后付费  3 代理商
	Price       int64  `json:"price"`
	Overage     int64  `json:"overage"`
	MonthTotal  int64  `json:"month_total"`
	DailyAmount int64  `json:"daily_amount"`
	Concurrency int64  `json:"concurrency"`
	FreeDay     int64  `json:"freeday"`
	Period      string `json:"period"`
}

func GetFirstOfMonth(t time.Time) time.Time {
	loc := t.Location()
	year, month, _ := t.Date()
	firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, loc)
	return firstOfMonth
}

func GetUserServicePackage(c *gin.Context) {
	username := c.Query("username")

	var (
		overage     int64
		monthTotal  int64
		concurrency int64
		dailyamount int64
		billType    string
		usertype    int64
		price       int64
		freeday     int64
	)

	userPackage, err := servpkgmodel.GetUserServicePackage(username)
	if err != nil {
		log.Error(fmt.Sprintf("find user service package err ", err))
		freeUserPackage, err := servpkgmodel.GetInterfaceConfig()
		if err != nil {
			log.Error(fmt.Sprintf("find interface config err ", err))
			c.JSON(400, gin.H{
				"code":    vars.ErrUserServicePackageNotFound.Code,
				"message": vars.ErrUserServicePackageNotFound.Msg,
			})
			return
		} else {
			dailyamount = freeUserPackage.FreeDayTotal
			monthTotal = freeUserPackage.FreeDayTotal * 30
			concurrency = freeUserPackage.FreeConcurrency
			price = freeUserPackage.DefaultPrice
			billType = "月结"
			usertype = 1
		}
	} else {
		usertype = userPackage.UserType
		price = userPackage.Price
		freeday = userPackage.FreeDay
		monthTotal = userPackage.MonthTotal
		dailyamount = userPackage.DailyAmount
		concurrency = userPackage.Concurrency
		billType = userPackage.Period
	}

	var (
		firstOfMonth time.Time
		now          time.Time
	)

	now = time.Now()

	firstOfMonth = GetFirstOfMonth(now)

	dayTotalSumArr, err := pornday.GetUserDayTotalSum(username, firstOfMonth, now)
	var used_total int64
	for _, day := range dayTotalSumArr {
		used_total += day.Total
	}

	if usertype == 0 {
		overage = monthTotal - used_total
	} else {
		overage = dailyamount*30 - used_total
	}

	rep := &Rep{
		Overage:     overage,
		MonthTotal:  monthTotal,
		DailyAmount: dailyamount,
		Concurrency: concurrency,
		Period:      billType,
		UserType:    usertype,
		Price:       price,
		FreeDay:     freeday,
	}
	if rep == nil {
		c.JSON(400, gin.H{
			"code":    vars.ErrUserServicePackageNotFound.Code,
			"message": vars.ErrUserServicePackageNotFound.Msg,
		})
		return
	}
	c.JSON(200, *rep)
	return
}
