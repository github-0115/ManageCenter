package transfers

import (
	transfersmodel "ManageCenter/service/model/transfersmodel"
	usermodel "ManageCenter/service/model/usermodel"
	vars "ManageCenter/service/vars"
	"fmt"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

func GetUserTransfers(c *gin.Context) {
	username := c.Query("user_name")
	pageIndex, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("rows"))
	sorted := c.Query("sort")
	sord := c.Query("sord")
	if err != nil {
		log.Error(fmt.Sprintf("strconv Atoi err%v", err))
		c.JSON(400, gin.H{
			"code":    vars.ErrParameter.Code,
			"message": vars.ErrParameter.Msg,
		})
		return
	}
	if sord == "desc" {
		sorted = "-" + sorted
	}
	if sorted == "" {
		sorted = "-created_at"
	}

	_, err = usermodel.QueryUser(username)
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

	result, records, err := transfersmodel.GetPagingUserTransfers(username, pageIndex, pageSize, sorted)
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
		result = make([]*transfersmodel.TransfersInfo, 0, 0)
	}

	pagetotal := int(math.Ceil(float64(records) / float64(pageSize)))

	c.JSON(200, gin.H{
		"code":      0,
		"transfers": result,
		"total":     records,
		"page":      pageIndex,
		"pagetotal": pagetotal,
	})
}
