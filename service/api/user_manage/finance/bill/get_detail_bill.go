package bill

import (
	billingmodel "ManageCenter/service/model/billingmodel"
	finalbillingmodel "ManageCenter/service/model/final_billmodel"
	usermodel "ManageCenter/service/model/usermodel"
	vars "ManageCenter/service/vars"
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

func GetDetailBill(c *gin.Context) {
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

	detailBill := make([]*billingmodel.BillColl, 0, len(finalBillColl.BillingId))

	for _, res := range finalBillColl.BillingId {
		billColl, err := billingmodel.QueryOneBill(res, username)
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

		detailBill = append(detailBill, billColl)
	}

	c.JSON(200, gin.H{
		"code": 0,
		"bill": detailBill,
	})
	return
}
