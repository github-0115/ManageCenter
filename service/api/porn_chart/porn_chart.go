package port_chart

import (
	pd "ManageCenter/service/model/porn_day"
	ph "ManageCenter/service/model/porn_hour"
	vars "ManageCenter/service/vars"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

type Rep struct {
	Time      []string `json:"time"`
	Total     []int64  `json:"total"`
	P         []int64  `json:"p"`
	S         []int64  `json:"s"`
	N         []int64  `json:"n"`
	PReview   []int64  `json:"p_review"`
	SReview   []int64  `json:"s_review"`
	NReview   []int64  `json:"n_review"`
	IsScreen  []int64  `json:"isScreen"`
	NotScreen []int64  `json:"notScreen"`
}

func GetPornChart(c *gin.Context) {
	magnitude := c.Query("magnitude")
	username := c.Query("username")
	reqIsAll := c.Query("isAll")
	var (
		startTime time.Time
		endTime   time.Time
	)
	d, _ := time.ParseDuration("-24h")
	endTime = time.Now()
	switch magnitude {
	case "day":
		startTime = endTime.Add(d * 1)
	case "week":
		startTime = endTime.Add(d * 7)
	case "month":
		startTime = endTime.Add(d * 30)
		if reqIsAll == "true" {
			getAllPornDay(c, startTime, endTime)
			return
		} else {
			getPornDay(c, username, startTime, endTime)
			return
		}
	default:
		c.JSON(400, gin.H{
			"code":    vars.ErrParameter.Code,
			"message": vars.ErrParameter.Msg,
		})
		return
	}
	if reqIsAll == "true" {
		getAllPornHour(c, startTime, endTime)
		return
	} else {
		getPornHour(c, username, startTime, endTime)
		return
	}
}

func getAllPornDay(c *gin.Context, startTime time.Time, endTime time.Time) {
	pornDay, err := pd.GetAllDayTotal(startTime, endTime)
	if err != nil {
		log.Error(fmt.Sprintf("query portday err%v", err))
		c.JSON(400, gin.H{
			"code":    vars.ErrPornDayNotFount.Code,
			"message": vars.ErrPornDayNotFount.Msg,
		})
		return
	}
	localTime := make([]string, 0, len(pornDay))
	total := make([]int64, 0, len(pornDay))
	p := make([]int64, 0, len(pornDay))
	s := make([]int64, 0, len(pornDay))
	n := make([]int64, 0, len(pornDay))
	pReview := make([]int64, 0, len(pornDay))
	sReview := make([]int64, 0, len(pornDay))
	nReview := make([]int64, 0, len(pornDay))
	isScreen := make([]int64, 0, len(pornDay))
	notScreen := make([]int64, 0, len(pornDay))
	for _, res := range pornDay {
		localTime = append(localTime, res.CreatedAt.Format("2006-01-02"))
		total = append(total, res.Total)
		p = append(p, res.P)
		s = append(s, res.S)
		n = append(n, res.N)
		pReview = append(pReview, res.PReview)
		sReview = append(sReview, res.SReview)
		nReview = append(nReview, res.NReview)
		isScreen = append(isScreen, res.IsScreen)
		notScreen = append(notScreen, res.NotScreen)
	}

	rep := &Rep{
		Time:      localTime,
		Total:     total,
		P:         p,
		S:         s,
		N:         n,
		PReview:   pReview,
		SReview:   sReview,
		NReview:   nReview,
		IsScreen:  isScreen,
		NotScreen: notScreen,
	}
	c.JSON(200, *rep)
}

func getPornDay(c *gin.Context, username string, startTime time.Time, endTime time.Time) {
	pornDay, err := pd.QueryUserPornDay(username, startTime, endTime)
	if err != nil {
		log.Error(fmt.Sprintf("query portday err%vusername=%v", err, username))
		c.JSON(400, gin.H{
			"code":    vars.ErrPornDayNotFount.Code,
			"message": vars.ErrPornDayNotFount.Msg,
		})
		return
	}
	localTime := make([]string, 0, len(pornDay))
	total := make([]int64, 0, len(pornDay))
	p := make([]int64, 0, len(pornDay))
	s := make([]int64, 0, len(pornDay))
	n := make([]int64, 0, len(pornDay))
	pReview := make([]int64, 0, len(pornDay))
	sReview := make([]int64, 0, len(pornDay))
	nReview := make([]int64, 0, len(pornDay))
	isScreen := make([]int64, 0, len(pornDay))
	notScreen := make([]int64, 0, len(pornDay))
	for _, res := range pornDay {
		localTime = append(localTime, res.CreatedAt.Format("2006-01-02"))
		total = append(total, res.Total)
		p = append(p, res.P)
		s = append(s, res.S)
		n = append(n, res.N)
		pReview = append(pReview, res.PReview)
		sReview = append(sReview, res.SReview)
		nReview = append(nReview, res.NReview)
		isScreen = append(isScreen, res.IsScreen)
		notScreen = append(notScreen, res.NotScreen)
	}

	rep := &Rep{
		Time:      localTime,
		Total:     total,
		P:         p,
		S:         s,
		N:         n,
		PReview:   pReview,
		SReview:   sReview,
		NReview:   nReview,
		IsScreen:  isScreen,
		NotScreen: notScreen,
	}
	c.JSON(200, *rep)
}

func getAllPornHour(c *gin.Context, startTime time.Time, endTime time.Time) {
	portHour, err := ph.GetAllHourTotal(startTime, endTime)
	if err != nil {
		log.Error(fmt.Sprintf("query porthour err%v"))
		c.JSON(400, gin.H{
			"code":    vars.ErrPornHourNotFount.Code,
			"message": vars.ErrPornHourNotFount.Msg,
		})
		return
	}
	localTime := make([]string, 0, len(portHour))
	total := make([]int64, 0, len(portHour))
	p := make([]int64, 0, len(portHour))
	s := make([]int64, 0, len(portHour))
	n := make([]int64, 0, len(portHour))
	pReview := make([]int64, 0, len(portHour))
	sReview := make([]int64, 0, len(portHour))
	nReview := make([]int64, 0, len(portHour))
	isScreen := make([]int64, 0, len(portHour))
	notScreen := make([]int64, 0, len(portHour))
	for _, res := range portHour {
		localTime = append(localTime, res.CreatedAt.Format("2006-01-02 15:04"))
		total = append(total, res.Total)
		p = append(p, res.P)
		s = append(s, res.S)
		n = append(n, res.N)
		pReview = append(pReview, res.PReview)
		sReview = append(sReview, res.SReview)
		nReview = append(nReview, res.NReview)
		isScreen = append(isScreen, res.IsScreen)
		notScreen = append(notScreen, res.NotScreen)
	}

	rep := &Rep{
		Time:      localTime,
		Total:     total,
		P:         p,
		S:         s,
		N:         n,
		PReview:   pReview,
		SReview:   sReview,
		NReview:   nReview,
		IsScreen:  isScreen,
		NotScreen: notScreen,
	}

	c.JSON(200, *rep)
}

func getPornHour(c *gin.Context, username string, startTime time.Time, endTime time.Time) {
	portHour, err := ph.QueryPornHour(username, startTime, endTime)
	if err != nil {
		log.Error(fmt.Sprintf("query porthour err%vusername=%v", err, username))
		c.JSON(400, gin.H{
			"code":    vars.ErrPornHourNotFount.Code,
			"message": vars.ErrPornHourNotFount.Msg,
		})
		return
	}
	localTime := make([]string, 0, len(portHour))
	total := make([]int64, 0, len(portHour))
	p := make([]int64, 0, len(portHour))
	s := make([]int64, 0, len(portHour))
	n := make([]int64, 0, len(portHour))
	pReview := make([]int64, 0, len(portHour))
	sReview := make([]int64, 0, len(portHour))
	nReview := make([]int64, 0, len(portHour))
	isScreen := make([]int64, 0, len(portHour))
	notScreen := make([]int64, 0, len(portHour))
	for _, res := range portHour {
		localTime = append(localTime, res.CreatedAt.Format("2006-01-02 15:04"))
		total = append(total, res.Total)
		p = append(p, res.P)
		s = append(s, res.S)
		n = append(n, res.N)
		pReview = append(pReview, res.PReview)
		sReview = append(sReview, res.SReview)
		nReview = append(nReview, res.NReview)
		isScreen = append(isScreen, res.IsScreen)
		notScreen = append(notScreen, res.NotScreen)
	}

	rep := &Rep{
		Time:      localTime,
		Total:     total,
		P:         p,
		S:         s,
		N:         n,
		PReview:   pReview,
		SReview:   sReview,
		NReview:   nReview,
		IsScreen:  isScreen,
		NotScreen: notScreen,
	}

	c.JSON(200, *rep)
}
