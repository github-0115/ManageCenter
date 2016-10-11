package real_time_photo

import (
	missrate "ManageCenter/service/model/miss_rate"
	vars "ManageCenter/service/vars"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

type Rep struct {
	Type    string  `json:"type"`
	Version string  `json:"version"`
	Result  *Result `json:"result"`
}

type Result struct {
	Miss  []int64   `json:"miss"`
	Total []int64   `json:"total"`
	Time  []string  `json:"time"`
	Rate  []float32 `json:"rate"`
}

func GetMissRate(c *gin.Context) {
	endTime := time.Now().UTC()
	startTime := endTime.Add(time.Hour * (-24) * 30)
	missRateMin, err := missrate.GetMissRate("min", startTime, endTime)
	if err != nil {
		log.Error(fmt.Sprintf("query missRate err%v", err))
		c.JSON(400, gin.H{
			"code":    vars.ErrMissRateNotFound.Code,
			"message": vars.ErrMissRateNotFound.Msg,
		})
		return
	}
	missRateMax, err := missrate.GetMissRate("max", startTime, endTime)
	if err != nil {
		log.Error(fmt.Sprintf("query missRate err%v", err))
		c.JSON(400, gin.H{
			"code":    vars.ErrMissRateNotFound.Code,
			"message": vars.ErrMissRateNotFound.Msg,
		})
		return
	}
	reps := make([]*Rep, 0, 4)
	addResult(&reps, missRateMin, "min")
	addResult(&reps, missRateMax, "max")
	if len(reps) <= 0 {
		c.JSON(400, gin.H{
			"code":    vars.ErrMissRateNotFound.Code,
			"message": vars.ErrMissRateNotFound.Msg,
		})
		return
	}
	c.JSON(200, reps)
	return

}

func addResult(reps *[]*Rep, missRate []*missrate.MissStruct, missType string) {
	localTime := make([]string, 0, len(missRate))
	total := make([]int64, 0, len(missRate))
	rate62 := make([]float32, 0, len(missRate))
	miss62 := make([]int64, 0, len(missRate))
	rate63 := make([]float32, 0, len(missRate))
	miss63 := make([]int64, 0, len(missRate))
	for _, resv := range missRate {
		localTime = append(localTime, resv.Day.Format("2006-01-02 15:04"))
		total = append(total, resv.Total)
		rate62 = append(rate62, resv.V62)
		miss62 = append(miss62, resv.V62Miss)
		rate63 = append(rate63, resv.V63)
		miss63 = append(miss63, resv.V63Miss)
	}
	result62 := &Result{
		Time:  localTime,
		Rate:  rate62,
		Miss:  miss62,
		Total: total,
	}
	rep62 := &Rep{
		Version: "v6_2",
		Type:    missType,
		Result:  result62,
	}
	result63 := &Result{
		Time:  localTime,
		Rate:  rate63,
		Miss:  miss63,
		Total: total,
	}
	rep63 := &Rep{
		Version: "v6_3",
		Type:    missType,
		Result:  result63,
	}
	*reps = append(*reps, rep62)
	*reps = append(*reps, rep63)
}
