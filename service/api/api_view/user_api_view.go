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

type Rep struct {
	User        string  `json:"user"`
	CreatedAt   string  `json:"time"`
	Total       int64   `json:"total"`
	P           int64   `json:"p"`
	S           int64   `json:"s"`
	N           int64   `json:"n"`
	PReview     int64   `json:"p_review"`
	SReview     int64   `json:"s_review"`
	NReview     int64   `json:"n_review"`
	IsScreen    int64   `json:"isScreen"`
	NotScreen   int64   `json:"notScreen"`
	PRate       float32 `json:"p_rate"`
	SRate       float32 `json:"s_rate"`
	NRate       float32 `json:"n_rate"`
	PReviewRate float32 `json:"p_review_rate"`
	SReviewRate float32 `json:"s_review_rate"`
}

func GetUserApi(c *gin.Context) {
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
	pageSize = pageSize / 4
	username := c.Query("username")
	var (
		startTime time.Time
		endTime   time.Time
		result    []*pd.PornDay
		records   int
	)
	d, _ := time.ParseDuration("-24h")
	endTime = time.Now()
	startTime = endTime.Add(d * 90)
	result, records, err = pd.GetUserPagingDaySum(username, pageIndex, pageSize, startTime, endTime)
	if err != nil {
		log.Error(fmt.Sprintf("query pornday err %v", err))
		c.JSON(400, gin.H{
			"code":    vars.ErrPornDayNotFount.Code,
			"message": vars.ErrPornDayNotFount.Msg,
		})
		return
	}

	rep := make([]*Rep, 0, len(result))

	for _, item := range result {
		h := &Rep{
			User:        item.User,
			CreatedAt:   item.CreatedAt.Format("2006-01-02"),
			Total:       item.Total,
			P:           item.P,
			S:           item.S,
			N:           item.N,
			PReview:     item.PReview,
			SReview:     item.SReview,
			NReview:     item.NReview,
			IsScreen:    item.IsScreen,
			NotScreen:   item.NotScreen,
			PRate:       float32(item.P) / float32(item.Total),
			SRate:       float32(item.S) / float32(item.Total),
			NRate:       float32(item.N) / float32(item.Total),
			PReviewRate: float32(item.PReview) / float32(item.Total),
			SReviewRate: float32(item.SReview) / float32(item.Total),
		}
		rep = append(rep, h)
	}
	total := int(math.Ceil(float64(records) / float64(pageSize)))
	c.JSON(200, gin.H{
		"code":    0,
		"total":   total,
		"result":  rep,
		"page":    pageIndex,
		"records": records,
	})
}
