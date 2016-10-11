package bill

import (
	user "ManageCenter/service/model/usermodel"
	vars "ManageCenter/service/vars"
	"ManageCenter/utils/smsclient"
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

type NotPassParams struct {
	Username string `json:"username"`
}

func NotPassBill(c *gin.Context) {
	var notPassParams NotPassParams
	if err := c.BindJSON(&notPassParams); err != nil {
		log.Error("bind json error:" + err.Error())
		c.JSON(400, gin.H{
			"code":    vars.ErrParameter.Code,
			"message": vars.ErrParameter.Msg,
		})
		return
	}
	username := notPassParams.Username

	userColl, err := user.QueryUser(username)
	if err != nil {
		log.Error(fmt.Sprintf("find err", err))
		if err == user.ErrUserNotFound {
			c.JSON(404, gin.H{
				"code":    vars.ErrUserNotFound.Code,
				"message": vars.ErrUserNotFound.Msg,
			})
			return
		}

		c.JSON(404, gin.H{
			"code":    vars.ErrUserCursor.Code,
			"message": vars.ErrUserCursor.Msg,
		})
		return
	}

	result := smsclient.SendBillSMS(userColl.PhoneNum, username)
	if result != "" {
		log.Error("send check bill sms failed." + result)
		c.JSON(200, gin.H{
			"code":    vars.ErrSendSMS.Code,
			"message": vars.ErrSendSMS.Msg,
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    0,
		"message": "bill not pass",
	})
}
