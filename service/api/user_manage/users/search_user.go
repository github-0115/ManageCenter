package users

import (
	user "ManageCenter/service/model/usermodel"
	vars "ManageCenter/service/vars"
	"fmt"
	"math"

	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

func StatSearchUsers(c *gin.Context) {
	pageIndex, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("rows"))

	searchString := c.Query("query")
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
		result  []*user.UserColl
		records int
		total   int
	)
	result, records, err = user.StatSearchPagingUsers(pageIndex, pageSize, sorted, sord, searchString)
	if err != nil {
		log.Error(fmt.Sprintf("find err", err))
		if err == user.ErrUserNotFound {
			c.JSON(404, gin.H{
				"code":    vars.ErrUserNotFound.Code,
				"message": vars.ErrUserNotFound.Msg,
			})
			return
		}

		c.JSON(404, gin.H{
			"code":    vars.ErrUserCursor.Code,
			"message": vars.ErrUserCursor.Msg,
		})
		return
	}
	if result == nil {
		result = make([]*user.UserColl, 0, len(result))
	}
	total = int(math.Ceil(float64(records) / float64(pageSize)))
	c.JSON(200, gin.H{
		"result":  result,
		"page":    pageIndex,
		"total":   total,
		"records": records,
	})
}
