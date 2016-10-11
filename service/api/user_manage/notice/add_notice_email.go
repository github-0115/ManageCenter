package notice

import (
	noticeemailmodel "ManageCenter/service/model/noticeemailmodel"
	vars "ManageCenter/service/vars"
	"fmt"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
	"github.com/satori/go.uuid"
)

type NoticeEmailParams struct {
	Email string `json:"email"`
}

var (
	emailPattern  = "^[_a-z0-9-]+(\\.[_a-z0-9-]+)*@[a-z0-9-]+(\\.[a-z0-9-]+)*(\\.[a-z]{2,4})$"
	emailRegex, _ = regexp.Compile(emailPattern)
)

func AddNoticeEmail(c *gin.Context) {
	var noticeEmailParams NoticeEmailParams
	if err := c.BindJSON(&noticeEmailParams); err != nil {
		log.Error(fmt.Sprintf("bind json error:%s", err))
		c.JSON(400, gin.H{
			"code":    vars.ErrParameter.Code,
			"message": vars.ErrParameter.Msg,
		})
		return
	}
	email := noticeEmailParams.Email

	if email != "" {
		if !emailRegex.MatchString(email) {
			log.Info("email format incorrect.email=" + email)
			c.JSON(400, gin.H{
				"code":    vars.ErrEmailFormat.Code,
				"message": vars.ErrEmailFormat.Msg,
			})
			return
		}

		existNum, err := noticeemailmodel.QueryNoticeEmailNum(email)
		if err != nil {
			log.Error(fmt.Sprintf("query notice email failed.email=%s, err=%#v", email, err))
			if err != noticeemailmodel.ErrEmailNumNotFound {
				c.JSON(400, gin.H{
					"code":    vars.ErrUserCursor.Code,
					"message": vars.ErrUserCursor.Msg,
				})
				return
			}
		}
		if existNum >= 1 {
			log.Error(fmt.Sprintf("notice email num >= 1, %i", existNum))
			c.JSON(400, gin.H{
				"code":    vars.ErrNoticeEmailExist.Code,
				"message": vars.ErrNoticeEmailExist.Msg,
			})
			return
		}
	}

	email_id := uuid.NewV4().String()
	emailColl, err := noticeemailmodel.QueryNoticeEmail(email_id)
	if err != nil {
		if err != noticeemailmodel.ErrEmailNotFound {
			c.JSON(200, gin.H{
				"code":    vars.ErrUserCursor.Code,
				"message": vars.ErrUserCursor.Msg,
			})
			return
		}
	}
	if emailColl != nil {
		c.JSON(200, gin.H{
			"code":    vars.ErrNoticeEmailExist.Code,
			"message": vars.ErrNoticeEmailExist.Msg,
		})
		return
	}
	sbemail := noticeemailmodel.NoticeEmailColl{
		EmailId:   email_id,
		Email:     email,
		CreatedAt: time.Now(),
	}
	err = sbemail.Save()
	if err != nil {
		log.Error("notice email save failed,err=" + err.Error())
		sbemail.Save()
	}

	c.JSON(200, gin.H{
		"code":    0,
		"message": "success",
	})
}
