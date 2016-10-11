package bill

import (
	finalbillingmodel "ManageCenter/service/model/final_billmodel"
	vars "ManageCenter/service/vars"
	"fmt"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

func GetAllBills(c *gin.Context) {
	pageIndex, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("rows"))
	status, err := strconv.Atoi(c.Query("status"))
	sorted := c.Query("sort")
	sord := c.Query("sord")
	if sorted == "" {
		sorted = "-created_at"
	}
	if sord == "desc" {
		sorted = "-" + sorted
	}
	if err != nil {
		log.Error(fmt.Sprintf("strconv Atoi err%v", err))
		c.JSON(400, gin.H{
			"code":    vars.ErrParameter.Code,
			"message": vars.ErrParameter.Msg,
		})
		return
	}

	result, records, err := finalbillingmodel.GetPagingBills(pageIndex, pageSize, sorted, int64(status))
	if err != nil {
		log.Error(fmt.Sprintf("query user finalbill failed. err=%#v", err))
		if err == finalbillingmodel.ErrBillNotFound {
			c.JSON(400, gin.H{
				"code":    vars.ErrUserBillsNotfount.Code,
				"message": vars.ErrUserBillsNotfount.Msg,
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
		result = make([]*finalbillingmodel.FinalBillColl, 0, 0)
	}

	pagetotal := int(math.Ceil(float64(records) / float64(pageSize)))

	c.JSON(200, gin.H{
		"code":      0,
		"bill":      result,
		"total":     records,
		"page":      pageIndex,
		"pagetotal": pagetotal,
	})
	return
}
