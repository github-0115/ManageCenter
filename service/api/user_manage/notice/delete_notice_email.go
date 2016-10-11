package notice

import (
	noticeemailmodel "ManageCenter/service/model/noticeemailmodel"
	vars "ManageCenter/service/vars"
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

func DeleteNoticeEmails(c *gin.Context) {

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

	err := noticeemailmodel.DeleteNoticeEmail(email)
	if err != nil {
		log.Error(fmt.Sprintf("delete notice email failed.err=%#v", err))
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

	c.JSON(200, gin.H{
		"code":    0,
		"message": "success",
	})
}
