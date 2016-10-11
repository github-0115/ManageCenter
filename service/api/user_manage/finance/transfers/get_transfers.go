package transfers

import (
	transfersmodel "ManageCenter/service/model/transfersmodel"
	usermodel "ManageCenter/service/model/usermodel"
	vars "ManageCenter/service/vars"
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

func GetTransfers(c *gin.Context) {
	username := c.Query("user_name")
	transfersId := c.Query("transfers_id")

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

	result, err := transfersmodel.QueryTransfers(transfersId, username)
	if err != nil {
		log.Error(fmt.Sprintf("query user transfers failed.username=%s, err=%#v", username, err))
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

	if result == nil {
		result = &transfersmodel.TransfersInfo{}
	}

	c.JSON(200, gin.H{
		"code":      0,
		"transfers": result,
	})
}
