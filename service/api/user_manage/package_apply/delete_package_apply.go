package package_apply

import (
	packageapplymodel "ManageCenter/service/model/package_apply_model"
	usermodel "ManageCenter/service/model/usermodel"
	vars "ManageCenter/service/vars"
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

type DeletePackageApplyParams struct {
	PackageApplyId string `json:"package_apply_id"`
	Username       string `json:"username"`
}

func DeletePackageApply(c *gin.Context) {
	var deletePackageApplyParams DeletePackageApplyParams
	if err := c.BindJSON(&deletePackageApplyParams); err != nil {
		log.Error(fmt.Sprintf("update package apply parameter bindJson err=%s ", err))
		c.JSON(400, gin.H{
			"code":    vars.ErrParameter.Code,
			"message": vars.ErrParameter.Msg,
		})
		return
	}
	applyId := deletePackageApplyParams.PackageApplyId
	username := deletePackageApplyParams.Username

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

	err = packageapplymodel.DeletePackageApply(applyId, username)
	if err != nil {
		log.Error(fmt.Sprintf("delete user package apply count failed.username=%s, err=%#v", username, err))
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
		"code":    0,
		"message": "delete user package apply success",
	})
}
