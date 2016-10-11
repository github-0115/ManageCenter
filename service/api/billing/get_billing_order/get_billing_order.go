package get_billing_order

import (
	billingmodel "ManageCenter/service/model/billingmodel"
	//	usermodel "UserCenter/service/model/usermodel"
	vars "ManageCenter/service/vars"
	"fmt"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

//var mgoOpers = map[string]bool{
//	"eq":    true,
//	"ne":    true,
//	"lt":    true,
//	"lte":   true,
//	"gt":    true,
//	"gte":   true,
//	"in":    true,
//	"or":    true,
//	"not":   true,
//	"regex": true,
//}

func GetPagingBills(c *gin.Context) {
	pageIndex, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("rows"))
	//	search := c.Query("search")
	sorted := c.Query("sort")
	sord := c.Query("sord")
	if sorted == "" {
		sorted = "created_at"
	}
	if sord == "desc" {
		sorted = "-" + sorted
	}
	if err != nil {
		c.JSON(400, gin.H{
			"message": "参数不合法",
		})
		return
	}
	var (
		result  []*billingmodel.BillColl
		records int
		total   int
	)

	result, records, err = billingmodel.GetPagingBills(pageIndex, pageSize, sorted)
	if err != nil {
		log.Error(fmt.Sprintf("find err", err))
		if err == billingmodel.ErrBillNotFound {
			c.JSON(404, gin.H{
				"code":    vars.ErrBillNotFound.Code,
				"message": vars.ErrBillNotFound.Msg,
			})
			return
		}

		c.JSON(404, gin.H{
			"code":    vars.ErrParameter.Code,
			"message": vars.ErrParameter.Msg,
		})
		return
	}

	total = int(math.Ceil(float64(records) / float64(pageSize)))
	c.JSON(200, gin.H{
		"result":  result,
		"page":    pageIndex,
		"total":   total,
		"records": records,
	})
}
