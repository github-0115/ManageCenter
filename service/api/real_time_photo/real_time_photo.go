package real_time_photo

import (
	pornminute "ManageCenter/service/model/porn_minute"
	vars "ManageCenter/service/vars"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

type Rep struct {
	Time     []string  `json:"time"`
	Rate     []float32 `json:"rate"`
	RateInke []float32 `json:"rate_inke"`
}

func GetPhotoReviewRate(c *gin.Context) {
	username := c.Query("username")
	d, _ := time.ParseDuration("-60m")
	endTime := time.Now().UTC()
	startTime := endTime.Add(d * 24)
	mintueReview, err := pornminute.GetMintueReview(username, startTime, endTime)
	if err != nil {
		log.Error(fmt.Sprintf("query portMinute err%vusername=%v", err, username))
		c.JSON(400, gin.H{
			"code":    vars.ErrPornMinuteNotFount.Code,
			"message": vars.ErrPornMinuteNotFount.Msg,
		})
		return
	}
	localTime := make([]string, 0, len(mintueReview))
	rate := make([]float32, 0, len(mintueReview))
	rateInke := make([]float32, 0, len(mintueReview))
	for _, res := range mintueReview {
		localTime = append(localTime, res.CreatedAt.Format("2006-01-02 15:04"))
		var rateTemp float32
		if res.Total == 0 {
			continue
		} else {
			rateTemp = float32(res.Review) / float32(res.Total)
			rate = append(rate, rateTemp)
			rateInke = append(rateInke, float32(res.InkeReview)/float32(res.Total))
		}
	}

	rep := &Rep{
		Time:     localTime,
		Rate:     rate,
		RateInke: rateInke,
	}
	if rep == nil {
		c.JSON(400, gin.H{
			"code":    vars.ErrPornMinuteNotFount.Code,
			"message": vars.ErrPornMinuteNotFount.Msg,
		})
		return
	}
	c.JSON(200, *rep)
	return

}
