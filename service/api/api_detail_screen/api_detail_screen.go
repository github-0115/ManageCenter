package api_detail_screen

import (
	config "ManageCenter/config"
	detailmodel "ManageCenter/service/model/detailmodel"
	vars "ManageCenter/service/vars"
	redisclient "ManageCenter/utils/redisclient"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
	"math"
	"strconv"
	"time"
)

type Rep struct {
	Res       string  `json:"res"`
	IsScreen  float32 `json:"isScreen"`
	NotScreen float32 `json:"notScreen"`
	Url       string  `json:"url"`
	CreatedAt string  `json:"created_at"`
}

func GetScreenDetail(c *gin.Context) {
	username := c.Query("username")
	reqTime, err := time.Parse("2006-01-02", c.Query("time"))
	if err != nil {
		log.Error(fmt.Sprintf("time Parse err%v", err))
		c.JSON(400, gin.H{
			"code":    vars.ErrParameter.Code,
			"message": vars.ErrParameter.Msg,
		})
		return
	}
	endTime := reqTime.Add(24 * time.Hour)
	reqSord := c.Query("sord")
	reqSort := c.Query("sort")
	reqIsAll := c.Query("isAll")
	pageIndex, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("rows"))
	if err != nil {
		log.Error(fmt.Sprintf("strconv Atoi err%v", err))
		c.JSON(400, gin.H{
			"code":    vars.ErrParameter.Code,
			"message": vars.ErrParameter.Msg,
		})
		return
	}
	var (
		sord    int
		records int = 0
	)

	if reqSord == "asc" {
		sord = 1
	} else {
		sord = -1
	}
	detail, err := detailmodel.GetScreenDetail(username, reqTime, endTime, reqSort, sord, pageIndex, pageSize)
	countId := username + reqIsAll + reqTime.Format("2006-01-02")
	records, err = redisclient.GetDetailCount(countId)
	if err != nil {
		log.Error(fmt.Sprintf(" redis GetDetailCount err ", err))
		records, err = detailmodel.GetScreenDetailCount(username, reqTime, endTime)
		if err != nil {
			log.Error(fmt.Sprintf("find detail err ", err))
			c.JSON(400, gin.H{
				"code":    vars.ErrDetialNotFound.Code,
				"message": vars.ErrDetialNotFound.Msg,
			})
			return
		}
		err = redisclient.SetDetailCount(countId, records)
		if err != nil {
			log.Error(fmt.Sprintf("redis SetDetailCount err ", err))
		}
	}
	repList := make([]*Rep, 0, len(detail))
	for _, res := range detail {
		rep := &Rep{
			Res:       res.Result.Screenshot.Res,
			IsScreen:  res.Result.Tags.IsScreen,
			NotScreen: res.Result.Tags.NotScreen,
			Url:       config.Cfg.Domain + res.Result.Md5,
			CreatedAt: res.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		repList = append(repList, rep)
	}
	total := int(math.Ceil(float64(records) / float64(pageSize)))

	c.JSON(200, gin.H{
		"code":    0,
		"result":  repList,
		"page":    pageIndex,
		"total":   total,
		"records": records,
	})

}
