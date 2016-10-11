package notice

import (
	submitemail "ManageCenter/service/model/submit_email"
	vars "ManageCenter/service/vars"
	emailclient "ManageCenter/utils/emailclient"
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

type PromotionEmailParams struct {
	Info string `json:"info"`
}

func SendEmail(c *gin.Context) {
	var promotionEmailParams PromotionEmailParams
	if err := c.BindJSON(&promotionEmailParams); err != nil {
		log.Error(fmt.Sprintf("bind json error:%s", err))
		c.JSON(400, gin.H{
			"code":    vars.ErrParameter.Code,
			"message": vars.ErrParameter.Msg,
		})
		return
	}
	mail_info := promotionEmailParams.Info

	submitEmail, err := submitemail.QueryUserEmail()
	if err != nil {
		log.Error(fmt.Sprintf("query notice email failed.err=%#v", err))
		if err == submitemail.ErrEmailNotFound {
			c.JSON(400, gin.H{
				"code":    vars.ErrSubmitEmailNotFound.Code,
				"message": vars.ErrSubmitEmailNotFound.Msg,
			})
			return
		}
		c.JSON(400, gin.H{
			"code":    vars.ErrUserCursor.Code,
			"message": vars.ErrUserCursor.Msg,
		})
		return
	}

	if submitEmail != nil {
		for _, res := range submitEmail {
			mail := res.Email
			id := res.EmailId

			err := emailclient.SendPlainMail(mail_info, mail)
			if err != nil {
				log.Error("send promotion mail failed,email=" + res.Email + ",err=" + err.Error())
			}

			err = submitemail.ChangeEmailCount(id)
			if err != nil {
				log.Error("change promotion mail count,email=" + res.Email + ",err=" + err.Error())
			}
		}
	} else {
		log.Error("no email to promotion failed,Please inform the management side can add the e-mail address")
	}

	c.JSON(200, gin.H{
		"code":    0,
		"message": "success",
	})
}
