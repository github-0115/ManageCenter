package api_view

import (
	pd "ManageCenter/service/model/porn_day"
	vars "ManageCenter/service/vars"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

type RepApi struct {
	Time        string  `json:"time"`
	Total       int64   `json:"total"`
	P           int64   `json:"p"`
	S           int64   `json:"s"`
	N           int64   `json:"n"`
	PReview     int64   `json:"p_review"`
	SReview     int64   `json:"s_review"`
	NReview     int64   `json:"n_review"`
	PRate       float32 `json:"p_rate"`
	SRate       float32 `json:"s_rate"`
	NRate       float32 `json:"n_rate"`
	IsScreen    int64   `json:"isScreen"`
	NotScreen   int64   `json:"notScreen"`
	PReviewRate float32 `json:"p_review_rate"`
	SReviewRate float32 `json:"s_review_rate"`
}

func GetApi(c *gin.Context) {

	pageIndex, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("rows"))
	pageSize = pageSize / 4
	username := c.Query("username")
	if err != nil {
		log.Error(fmt.Sprintf("strconv Atoi err%v", err))
		c.JSON(400, gin.H{
			"code":    vars.ErrParameter.Code,
			"message": vars.ErrParameter.Msg,
		})
		return
	}
	var (
		startTime time.Time
		endTime   time.Time
	)
	d, _ := time.ParseDuration("-24h")
	endTime = time.Now()
	startTime = endTime.Add(d * 90)
	if username == "" {
		pornDay, err := pd.GetPagingDaySum(pageIndex, pageSize, startTime, endTime)
		if err != nil {
			log.Error(fmt.Sprintf("query DaySum err%v", err))
			c.JSON(400, gin.H{
				"code":    vars.ErrPornDayNotFount.Code,
				"message": vars.ErrPornDayNotFount.Msg,
			})
			return
		}
		records, err := pd.GetPagingDaySumCount(startTime, endTime)
		if err != nil {
			log.Error(fmt.Sprintf("query DaySumCount err%v", err))
			c.JSON(400, gin.H{
				"code":    vars.ErrPornDayNotFount.Code,
				"message": vars.ErrPornDayNotFount.Msg,
			})
			return
		}
		repRes := make([]*RepApi, 0, len(pornDay))
		for _, res := range pornDay {
			ApiElement := &RepApi{
				Time:        res.CreatedAt.Format("2006-01-02"),
				Total:       res.Total,
				P:           res.P,
				S:           res.S,
				N:           res.N,
				PReview:     res.PReview,
				SReview:     res.SReview,
				NReview:     res.NReview,
				IsScreen:    res.IsScreen,
				NotScreen:   res.NotScreen,
				PRate:       float32(res.P) / float32(res.Total),
				SRate:       float32(res.S) / float32(res.Total),
				NRate:       float32(res.N) / float32(res.Total),
				PReviewRate: float32(res.PReview) / float32(res.Total),
				SReviewRate: float32(res.SReview) / float32(res.Total),
			}
			repRes = append(repRes, ApiElement)
		}
		total := int(math.Ceil(float64(records) / float64(pageSize)))
		c.JSON(200, gin.H{
			"code":    0,
			"result":  repRes,
			"page":    pageIndex,
			"total":   total,
			"records": records,
		})
	} else {
		pornDay, records, err := pd.GetUserPagingDaySum(username, pageIndex, pageSize, startTime, endTime)
		if err != nil {
			log.Error(fmt.Sprintf("query DaySum err%v", err))
			c.JSON(400, gin.H{
				"code":    vars.ErrPornDayNotFount.Code,
				"message": vars.ErrPornDayNotFount.Msg,
			})
			return
		}
		repRes := make([]*RepApi, 0, len(pornDay))
		for _, res := range pornDay {
			ApiElement := &RepApi{
				Time:        res.CreatedAt.Format("2006-01-02"),
				Total:       res.Total,
				P:           res.P,
				S:           res.S,
				N:           res.N,
				PReview:     res.PReview,
				SReview:     res.SReview,
				NReview:     res.NReview,
				IsScreen:    res.IsScreen,
				NotScreen:   res.NotScreen,
				PRate:       float32(res.P) / float32(res.Total),
				SRate:       float32(res.S) / float32(res.Total),
				NRate:       float32(res.N) / float32(res.Total),
				PReviewRate: float32(res.PReview) / float32(res.Total),
				SReviewRate: float32(res.SReview) / float32(res.Total),
			}
			repRes = append(repRes, ApiElement)
		}
		total := int(math.Ceil(float64(records) / float64(pageSize)))
		c.JSON(200, gin.H{
			"code":    0,
			"result":  repRes,
			"page":    pageIndex,
			"total":   total,
			"records": records,
		})
	}
}
