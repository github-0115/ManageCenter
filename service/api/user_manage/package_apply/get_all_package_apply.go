package package_apply

import (
	packageapplymodel "ManageCenter/service/model/package_apply_model"
	vars "ManageCenter/service/vars"
	"fmt"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

func GetAllPackageApply(c *gin.Context) {
	pageIndex, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("rows"))
	status, err := strconv.Atoi(c.Query("status"))
	sorted := c.Query("sort")

	if sorted == "" {
		sorted = "created_at"
	}

	if err != nil {
		log.Error(fmt.Sprintf("strconv Atoi err%v", err))
		c.JSON(400, gin.H{
			"code":    vars.ErrParameter.Code,
			"message": vars.ErrParameter.Msg,
		})
		return
	}

	packageapply, records, err := packageapplymodel.QueryPagingPackageApplys(pageIndex, pageSize, sorted, int64(status))
	if err != nil {
		log.Error(fmt.Sprintf("query user package apply failed err=%#v", err))
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

	if packageapply == nil {
		packageapply = make([]*packageapplymodel.PackageApplyColl, 0, 0)
	}

	pagetotal := int(math.Ceil(float64(records) / float64(pageSize)))

	c.JSON(200, gin.H{
		"code":      0,
		"bill":      packageapply,
		"total":     records,
		"page":      pageIndex,
		"pagetotal": pagetotal,
	})
	return
}
