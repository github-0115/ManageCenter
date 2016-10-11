package bill

import (
	finalbillingmodel "ManageCenter/service/model/final_billmodel"
	usermodel "ManageCenter/service/model/usermodel"
	vars "ManageCenter/service/vars"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

func GetOneBill(c *gin.Context) {
	username := c.Query("username")
	id := c.Query("finalbill_id")

	_, err := usermodel.QueryUser(username)
	if err != nil {
		log.Error(fmt.Sprintf("query user failed.username=%s, err=%#v", username, err))
		if err == usermodel.ErrUserNotFound {
			c.JSON(400, gin.H{
				"code":    vars.ErrUserNotFound.Code,
				"message": vars.ErrUserNotFound.Msg,
			})
			return
		}
		c.JSON(400, gin.H{
			"code":    vars.ErrUserCursor.Code,
			"message": vars.ErrUserCursor.Msg,
		})
		return
	}

	finalBillColl, err := finalbillingmodel.QueryOneBill(id, username)
	if err != nil {
		log.Error(fmt.Sprintf("query user finalbill failed. err=%#v", err))
		if err == finalbillingmodel.ErrBillNotFound {
			c.JSON(400, gin.H{
				"code":    vars.ErrUserBillsNotfount.Code,
				"message": vars.ErrUserBillsNotfount.Msg,
			})
			return
		}
		c.JSON(400, gin.H{
			"code":    vars.ErrUserCursor.Code,
			"message": vars.ErrUserCursor.Msg,
		})
		return
	}

	ms := months(finalBillColl.PeriodStart, finalBillColl.PeriodEnd)
	var isDays int
	switch finalBillColl.BillType {
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

	c.JSON(200, gin.H{
		"code":     0,
		"bill":     finalBillColl,
		"billdate": isDays,
	})
	return
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
