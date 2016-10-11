package bill

import (
	finalbillingmodel "ManageCenter/service/model/final_billmodel"
	usermodel "ManageCenter/service/model/usermodel"
	vars "ManageCenter/service/vars"
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

type FavorParams struct {
	Username    string  `json:"username"`
	FinalBillId string  `json:"finalbill_id"`
	PaidAmount  float64 `json:"paid_amount"`
}

func UpdateFavorBill(c *gin.Context) {

	var favorParams FavorParams
	if err := c.BindJSON(&favorParams); err != nil {
		log.Error(fmt.Sprintf("bind json error:%s", err))
		c.JSON(400, gin.H{
			"code":    vars.ErrParameter.Code,
			"message": vars.ErrParameter.Msg,
		})
		return
	}
	username := favorParams.Username
	finalBillId := favorParams.FinalBillId
	paidAmount := favorParams.PaidAmount

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

	err = finalbillingmodel.UpdateFavorBill(finalBillId, username, paidAmount)
	if err != nil {
		log.Error(fmt.Sprintf("update user finalbill paidamount failed. err=%#v", err))
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

	c.JSON(200, gin.H{
		"code":    0,
		"message": "update bill favor successs",
	})
}
