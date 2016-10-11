package notice

import (
	noticeemailmodel "ManageCenter/service/model/noticeemailmodel"
	vars "ManageCenter/service/vars"
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

func AllNoticeEmail(c *gin.Context) {

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

	if noticeEmail == nil {
		noticeEmail = make([]*noticeemailmodel.NoticeEmailColl, 0, 0)
	}

	c.JSON(200, gin.H{
		"code":         0,
		"notice_email": noticeEmail,
	})
}
