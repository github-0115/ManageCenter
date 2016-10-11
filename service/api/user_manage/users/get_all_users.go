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

var mgoOpers = map[string]bool{
	"eq":    true,
	"ne":    true,
	"lt":    true,
	"lte":   true,
	"gt":    true,
	"gte":   true,
	"in":    true,
	"or":    true,
	"not":   true,
	"regex": true,
}

func GetUsers(c *gin.Context) {

	pageIndex, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("rows"))
	search := c.Query("search")
	sorted := c.Query("sort")
	sord := c.Query("sord")
	if sorted == "" {
		sorted = "-created_at"
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
		result  []user.UserColl
		records int
		total   int
	)
	if search == "true" {
		searchField := c.Query("searchField")
		searchString := c.Query("searchString")
		searchOper := c.Query("searchOper")
		if searchField == "" || searchString == "" || searchOper == "" {
			c.JSON(400, gin.H{
				"message": "参数不合法",
			})
			return
		}

		operLegal := mgoOpers[searchOper]
		if operLegal == false {
			c.JSON(400, gin.H{
				"message": "参数不合法",
			})
			return
		} else {
			searchOper = "$" + searchOper
		}
		result, records, err = user.SearchPagingUsers(pageIndex, pageSize, sorted, sord, searchField, searchString, searchOper)
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

	} else {
		result, records, err = user.GetPagingUsers(pageIndex, pageSize, sorted, sord)
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
	}
	total = int(math.Ceil(float64(records) / float64(pageSize)))

	c.JSON(200, gin.H{
		"result":  result,
		"page":    pageIndex,
		"total":   total,
		"records": records,
	})
}
