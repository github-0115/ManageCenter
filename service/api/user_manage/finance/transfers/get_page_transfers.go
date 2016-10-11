package transfers

import (
	transfersmodel "ManageCenter/service/model/transfersmodel"
	vars "ManageCenter/service/vars"
	"fmt"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

func GetPageTransfers(c *gin.Context) {
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

	result, records, err := transfersmodel.GetPagingTransfers(pageIndex, pageSize, sorted)
	if err != nil {
		log.Error(fmt.Sprintf("query user transfers failed, err=%#v", err))
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

	pagetotal := int(math.Ceil(float64(records) / float64(pageSize)))

	if result == nil {
		result = make([]*transfersmodel.TransfersInfo, 0, 0)
	}

	c.JSON(200, gin.H{
		"code":      0,
		"transfers": result,
		"total":     records,
		"page":      pageIndex,
		"pagetotal": pagetotal,
	})
}
