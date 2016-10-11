package real_time_room

import (
	livereview "ManageCenter/service/model/live_review"
	vars "ManageCenter/service/vars"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

type Rep struct {
	Version string  `json:"version"`
	Result  *Result `json:"result"`
}

type Result struct {
	Time []string  `json:"time"`
	Rate []float32 `json:"rate"`
}

func GetRoomReviewRate(c *gin.Context) {
	username := c.Query("username")
	d, _ := time.ParseDuration("-60m")
	endTime := time.Now().UTC()
	startTime := endTime.Add(d * 24)
	liveReview, err := livereview.GetLiveReviewRate(username, startTime, endTime)
	if err != nil {
		log.Error(fmt.Sprintf("query liveReview err%vusername=%v", err, username))
		c.JSON(400, gin.H{
			"code":    vars.ErrLiveReveiwNotFount.Code,
			"message": vars.ErrLiveReveiwNotFount.Msg,
		})
		return
	}
	versions, err := livereview.GetVersion(username, startTime, endTime)
	if err != nil {
		log.Error(fmt.Sprintf("query liveReview err%vusername=%v", err, username))
		c.JSON(400, gin.H{
			"code":    vars.ErrLiveReveiwNotFount.Code,
			"message": vars.ErrLiveReveiwNotFount.Msg,
		})
		return
	}
	reps := make([]*Rep, 0, len(versions))
	for _, resv := range versions {
		localTime := make([]string, 0, len(liveReview))
		rate := make([]float32, 0, len(liveReview))
		for _, resl := range liveReview {
			if resl.Version == resv {
				localTime = append(localTime, resl.CreatedAt.Format("2006-01-02 15:04"))
				rate = append(rate, resl.Rate)
			}
		}
		result := &Result{
			Time: localTime,
			Rate: rate,
		}
		rep := &Rep{
			Version: resv,
			Result:  result,
		}
		reps = append(reps, rep)
	}

	if len(reps) <= 0 {
		c.JSON(400, gin.H{
			"code":    vars.ErrLiveReveiwNotFount.Code,
			"message": vars.ErrLiveReveiwNotFount.Msg,
		})
		return
	}
	c.JSON(200, reps)
	return

}
