package notice

import (
	noticeemailmodel "ManageCenter/service/model/noticeemailmodel"
	vars "ManageCenter/service/vars"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
	"github.com/satori/go.uuid"
)

type NoticeEmailsParams struct {
	Email string `json:"email"`
}

var (
	emailsPattern  = "^[_a-z0-9-]+(\\.[_a-z0-9-]+)*@[a-z0-9-]+(\\.[a-z0-9-]+)*(\\.[a-z]{2,4})$"
	emailsRegex, _ = regexp.Compile(emailsPattern)
)

func NoticeEmail(c *gin.Context) {
	var noticeEmailsParams NoticeEmailsParams
	if err := c.BindJSON(&noticeEmailsParams); err != nil {
		log.Error(fmt.Sprintf("bind json error:%s", err))
		c.JSON(400, gin.H{
			"code":    vars.ErrParameter.Code,
			"message": vars.ErrParameter.Msg,
		})
		return
	}
	emailstr := noticeEmailsParams.Email
	emailstr = strings.Replace(emailstr, " ", "", -1)
	emailstr = strings.Replace(emailstr, "\n", "", -1)
	email := strings.Split(emailstr, ";")

	if email != nil {

		var retEmail []string
		for i := 0; i < len(email); i++ {
			if (i > 0 && email[i-1] == email[i]) || len(email[i]) == 0 {
				continue
			}
			if !emailsRegex.MatchString(email[i]) {
				log.Info("email format incorrect.email=" + email[i])
				c.JSON(400, gin.H{
					"code":    vars.ErrEmailFormat.Code,
					"message": vars.ErrEmailFormat.Msg,
				})
				return
			}
			retEmail = append(retEmail, email[i])
		}

		noticeEmail, err := noticeemailmodel.QueryNoticeAllEmail()
		if err != nil {
			log.Error(fmt.Sprintf("query notice email failed.err=%#v", err))
			if err == noticeemailmodel.ErrEmailNotFound {
				c.JSON(200, gin.H{
					"code":    vars.ErrNoticeEmailNotFound.Code,
					"message": vars.ErrNoticeEmailNotFound.Msg,
				})
				return
			}
			c.JSON(200, gin.H{
				"code":    vars.ErrUserCursor.Code,
				"message": vars.ErrUserCursor.Msg,
			})
			return
		}

		err = DeleteNoticeEmail(noticeEmail)
		if err != nil {
			log.Error(fmt.Sprintf("delete notice email failed.err=%#v", err))
			c.JSON(400, gin.H{
				"code":    vars.ErrUserCursor.Code,
				"message": vars.ErrUserCursor.Msg,
			})
			return
		}

		err = SaveNoticeEmail(retEmail)
		if err != nil {
			log.Error(fmt.Sprintf("save notice email failed.email=%s, err=%#v", err))
			c.JSON(400, gin.H{
				"code":    vars.ErrUserCursor.Code,
				"message": vars.ErrUserCursor.Msg,
			})
			return
		}
	}

	c.JSON(200, gin.H{
		"code":    0,
		"message": "success",
	})
}

func SaveNoticeEmail(emails []string) error {

	for _, res := range emails {
		email_id := uuid.NewV4().String()
		emailColl, err := noticeemailmodel.QueryNoticeEmail(email_id)
		if err != nil {
			if err != noticeemailmodel.ErrEmailNotFound {
				return noticeemailmodel.ErrUserCursor
			}
		}

		if emailColl != nil {
			email_id = uuid.NewV4().String()
		}

		sbemail := noticeemailmodel.NoticeEmailColl{
			EmailId:   email_id,
			Email:     res,
			CreatedAt: time.Now(),
		}
		err = sbemail.Save()
		if err != nil {
			log.Error("notice email save failed,err=" + err.Error())
			sbemail.Save()
		}
	}

	return nil
}

func DeleteNoticeEmail(noticeEmails []*noticeemailmodel.NoticeEmailColl) error {
	if noticeEmails != nil {
		for _, res := range noticeEmails {
			err := noticeemailmodel.DeleteNoticeEmail(res.Email)
			if err != nil {
				if err == noticeemailmodel.ErrEmailNotFound {
					return noticeemailmodel.ErrEmailNotFound
				}
				return noticeemailmodel.ErrUserCursor
			}
		}
	}
	return nil
}
