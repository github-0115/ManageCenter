package users

import (
	sms "ManageCenter/service/model/vcodesmsmodel"
	vars "ManageCenter/service/vars"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

func PotentialUser(c *gin.Context) {

	pageIndex, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("rows"))
	sorted := c.Query("sort")
	if err != nil {
		log.Error(fmt.Sprintf("strconv Atoi err%v", err))
		c.JSON(400, gin.H{
			"code":    vars.ErrParameter.Code,
			"message": vars.ErrParameter.Msg,
		})
		return
	}
	if sorted == "" {
		sorted = "-created_at"
	}

	result, records, err := sms.GetPagingPuser(int(pageIndex), int(pageSize))
	if err != nil {
		if err == sms.ErrVcodeNotFound {
			c.JSON(400, gin.H{
				"code":    vars.ErrPVcodeNotFound.Code,
				"message": vars.ErrPVcodeNotFound.Msg,
			})
			return
		}
		c.JSON(400, gin.H{
			"code":    vars.ErrPhoneNotFound.Code,
			"message": vars.ErrPhoneNotFound.Msg,
		})
		return
	}

	total := records/int(pageSize) + 1

	c.JSON(200, gin.H{
		"result":  result,
		"page":    pageIndex,
		"total":   total,
		"records": records,
	})
}
