package package_apply

import (
	packageapplymodel "ManageCenter/service/model/package_apply_model"
	usermodel "ManageCenter/service/model/usermodel"
	vars "ManageCenter/service/vars"
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

func GetOnePackageApply(c *gin.Context) {
	username := c.Query("username")
	id := c.Query("package_apply_id")

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

	packageapply, err := packageapplymodel.QueryOnePackageApply(id, username)
	if err != nil {
		log.Error(fmt.Sprintf("query user package apply count failed.username=%s, err=%#v", username, err))
		if err == packageapplymodel.ErrPackageApplyNotFound {
			c.JSON(400, gin.H{
				"code":    vars.ErrUserPakageApplyNotFound.Code,
				"message": vars.ErrUserPakageApplyNotFound.Msg,
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
		"code":  0,
		"apply": packageapply,
	})
	return
}
