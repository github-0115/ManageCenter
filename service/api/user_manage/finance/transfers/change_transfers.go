package transfers

import (
	transfersmodel "ManageCenter/service/model/transfersmodel"
	usermodel "ManageCenter/service/model/usermodel"
	vars "ManageCenter/service/vars"
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

type TransfersParams struct {
	Username    string `json:"username"`
	TransfersId string `json:"transfers_id"`
}

func ChangeTransfers(c *gin.Context) {
	var transfersParams TransfersParams
	if err := c.BindJSON(&transfersParams); err != nil {
		log.Error(fmt.Sprintf("bind json error:%s", err))
		c.JSON(400, gin.H{
			"code":    vars.ErrParameter.Code,
			"message": vars.ErrParameter.Msg,
		})
		return
	}
	username := transfersParams.Username
	transfersId := transfersParams.TransfersId

	if username == "" || transfersId == "" {
		c.JSON(400, gin.H{
			"code":    vars.ErrParameterNil.Code,
			"message": vars.ErrParameterNil.Msg,
		})
		return
	}

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

	err = transfersmodel.UpdateTransfersStatus(transfersId, username)
	if err != nil {
		log.Error(fmt.Sprintf("update user transfers failed.username=%s, err=%#v", username, err))
		if err == transfersmodel.ErrUserTransfersNotFound {
			c.JSON(400, gin.H{
				"code":    vars.ErrUserTransfersNotfount.Code,
				"message": vars.ErrUserTransfersNotfount.Msg,
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
		"message": "success",
	})
}
