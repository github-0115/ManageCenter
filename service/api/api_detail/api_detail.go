package api_detail

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
	Review    bool    `json:"review"`
	P         float32 `json:"p"`
	S         float32 `json:"s"`
	N         float32 `json:"n"`
	Url       string  `json:"url"`
	CreatedAt string  `json:"created_at"`
}

func GetPornDetail(c *gin.Context) {
	username := c.Query("username")
	reqTime, err := time.ParseInLocation("2006-01-02", c.Query("time"), time.Local)
	if err != nil {
		log.Error(fmt.Sprintf("time Parse err%v", err))
		c.JSON(400, gin.H{
			"code":    vars.ErrParameter.Code,
			"message": vars.ErrParameter.Msg,
		})
		return
	}
	endTime := reqTime.Add(24 * time.Hour)
	reqType := c.Query("type")
	reqReview := c.Query("review")
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
		sord   int
		review bool
		isAll  bool
	)
	if reqSord == "asc" {
		sord = 1
	} else {
		sord = -1
	}
	if reqReview == "false" {
		review = false
	} else {
		review = true
	}
	if reqIsAll == "true" {
		isAll = true
	} else {
		isAll = false
	}
	var (
		detail  []*detailmodel.Detail
		records int = 0
	)

	if isAll {
		detail, err = detailmodel.GetAllDetail(username, reqTime, endTime, reqType, reqSort, sord, pageIndex, pageSize)
		countId := username + reqType + reqIsAll + reqTime.Format("2006-01-02")
		records, err = redisclient.GetDetailCount(countId)
		if err != nil {
			log.Error(fmt.Sprintf(" redis GetDetailCount err ", err))
			records, err := detailmodel.GetAllDetailCount(username, reqTime, endTime, reqType)
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
	} else {
		detail, err = detailmodel.GetDetail(username, reqTime, endTime, reqType, review, reqSort, sord, pageIndex, pageSize)
		countId := username + reqType + reqIsAll + reqReview + reqTime.Format("2006-01-02")
		records, err = redisclient.GetDetailCount(countId)
		if err != nil {
			log.Error(fmt.Sprintf(" redis GetDetailCount err ", err))
			records, err = detailmodel.GetDetailCount(username, reqTime, endTime, reqType, review)
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
	}
	repList := make([]*Rep, 0, len(detail))
	for _, res := range detail {
		rep := &Rep{
			Res:       res.Result.Porn.Res,
			Review:    res.Result.Porn.Review,
			P:         res.Result.Tags.P,
			S:         res.Result.Tags.S,
			N:         res.Result.Tags.N,
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
