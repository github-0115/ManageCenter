package console_detail

import (
	pornday "ManageCenter/service/model/porn_day"
	pornhour "ManageCenter/service/model/porn_hour"
	usermodel "ManageCenter/service/model/usermodel"
	vars "ManageCenter/service/vars"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

type Rep struct {
	Time   []string `json:"time"`
	Total  []int64  `json:"total"`
	Review []int64  `json:"review"`
}

type RepToatl struct {
	PReview        int64 `json:"p_review"`
	SReview        int64 `json:"s_review"`
	NReview        int64 `json:"n_review"`
	IsScreen       int64 `json:"isScreen"`
	NotScreen      int64 `json:"notScreen"`
	TodayTotal     int64 `json:"today_total"`
	YesterdayTotal int64 `json:"yesterday_total"`
	P              int64 `json:"p"`
	S              int64 `json:"s"`
	N              int64 `json:"n"`
}

func ConsoleDetail(c *gin.Context) {
	var (
		startTime time.Time
		endTime   time.Time
	)
	times, _ := time.ParseInLocation("2006-01-02", time.Now().Format("2006-01-02"), time.Local)
	d, _ := time.ParseDuration("24h")
	startTime = times
	endTime = times.Add(d * 1)
	tPornHour, err := pornhour.GetAllHourTotal(startTime, endTime)
	if err != nil {
		log.Error(fmt.Sprintf("query pornHour err%v", err))
		c.JSON(400, gin.H{
			"code":    vars.ErrPornHourNotFount.Code,
			"message": vars.ErrPornHourNotFount.Msg,
		})
		return
	}
	var (
		todayTotal     int64
		yesterdayTotal int64
		p              int64
		s              int64
		n              int64
		pReview        int64
		sReview        int64
		nReview        int64
		isScreen       int64
		notScreen      int64
		ytTotal        int64
	)
	localTimes := make([]string, 0, len(tPornHour))
	totals := make([]int64, 0, len(tPornHour))
	reviews := make([]int64, 0, len(tPornHour))
	for _, res := range tPornHour {
		localTimes = append(localTimes, res.CreatedAt.Format("15:04"))
		totals = append(totals, res.Total)
		reviews = append(reviews, res.PReview)
		reviews = append(reviews, res.SReview)
		reviews = append(reviews, res.NReview)

		todayTotal = todayTotal + res.Total
		p = p + res.P
		s = s + res.S
		n = n + res.N
		pReview = pReview + res.PReview
		sReview = sReview + res.SReview
		nReview = nReview + res.NReview
		isScreen = nReview + res.IsScreen
		notScreen = nReview + res.NotScreen
	}

	today_rep := &Rep{
		Time:   localTimes,
		Total:  totals,
		Review: reviews,
	}
	userCount, err := usermodel.QuerySubmitCount(startTime, endTime)
	if err != nil {
		log.Error(fmt.Sprintf("query Submit user Count failed. username=%s  err=%#v", err))
		if err == usermodel.ErrUserNotFound {
			c.JSON(400, gin.H{
				"code":    vars.ErrUserNotFound.Code,
				"message": vars.ErrUserNotFound.Msg,
			})
			return
		}
		c.JSON(400, gin.H{
			"code":    vars.ErrUserCursor.Code,
			"message": vars.ErrUserCursor.Msg,
		})
		return
	}

	dy, _ := time.ParseDuration("-24h")
	endTime = times
	startTime = times.Add(dy * 1)

	pornDay, err := pornday.QueryPornDay(startTime, endTime)
	if err != nil {
		log.Error(fmt.Sprintf("query pornday err%v", err))
		c.JSON(400, gin.H{
			"code":    vars.ErrPornDayNotFount.Code,
			"message": vars.ErrPornDayNotFount.Msg,
		})
		return
	}

	yPornHour, err := pornhour.GetAllHourTotal(startTime, endTime)
	if err != nil {
		log.Error(fmt.Sprintf("query pornHour err%v", err))
		c.JSON(400, gin.H{
			"code":    vars.ErrPornHourNotFount.Code,
			"message": vars.ErrPornHourNotFount.Msg,
		})
		return
	}
	ylocalTimes := make([]string, 0, len(yPornHour))
	ytotals := make([]int64, 0, len(yPornHour))
	yreviews := make([]int64, 0, len(yPornHour))
	for _, res := range yPornHour {
		ylocalTimes = append(ylocalTimes, res.CreatedAt.Format("15:04"))
		ytotals = append(ytotals, res.Total)
		yreviews = append(yreviews, res.PReview)
		yreviews = append(yreviews, res.SReview)
		yreviews = append(yreviews, res.NReview)

	}

	for i := 0; i < len(yPornHour); i++ {
		if i <= len(tPornHour) {
			yesterdayTotal += yPornHour[i].Total
		}
	}

	yesterday_rep := &Rep{
		Time:   ylocalTimes,
		Total:  ytotals,
		Review: yreviews,
	}

	rep := &RepToatl{
		PReview:        pReview,
		SReview:        sReview,
		NReview:        nReview,
		IsScreen:       isScreen,
		NotScreen:      notScreen,
		TodayTotal:     todayTotal,
		YesterdayTotal: yesterdayTotal,
		P:              p,
		S:              s,
		N:              n,
	}

	ytpornDay, err := pornday.QueryPornDay(startTime.Add(dy*1), endTime.Add(dy*1))
	if err != nil {
		log.Error(fmt.Sprintf("query pornday err%v", err))
		c.JSON(400, gin.H{
			"code":    vars.ErrPornDayNotFount.Code,
			"message": vars.ErrPornDayNotFount.Msg,
		})
		return
	}
	startUserList := make([]string, 0)
	for i := 0; i < len(ytpornDay); i++ {
		ytTotal += ytpornDay[i].Total
	}

	for i := 0; i < len(pornDay); i++ {
		var rate int64
		if len(ytpornDay) != 0 {
			rate = ytTotal / int64(len(ytpornDay))
		} else {
			rate = 0
		}
		if pornDay[i].Total > rate {
			startUserList = append(startUserList, pornDay[i].User)
		}
	}

	c.JSON(200, gin.H{
		"code":          0,
		"today_detail":  rep,
		"today_rep":     today_rep,
		"yesterday_rep": yesterday_rep,
		"submit_user":   len(userCount),
		"startUser":     len(startUserList),
		"startUserList": startUserList,
	})
	return
}
