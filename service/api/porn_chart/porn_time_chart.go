package port_chart

import (
	vars "ManageCenter/service/vars"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

func GetTimePornChart(c *gin.Context) {
	username := c.Query("username")
	start := c.Query("start")
	end := c.Query("end")
	reqIsAll := c.Query("isAll")

	startTime, _ := time.ParseInLocation("2006-01-02", start, time.Local)
	endtime, _ := time.ParseInLocation("2006-01-02", end, time.Local)
	d, _ := time.ParseDuration("+24h")
	endTime := endtime.Add(d * 1)

	if endtime.Before(startTime) {
		log.Error(fmt.Sprintf("endTime=%s 在 startTime=%s之前!", endTime, startTime))
		c.JSON(400, gin.H{
			"code":    vars.ErrParameter.Code,
			"message": vars.ErrParameter.Msg,
		})
		return
	}

	if endTime.Sub(startTime).Hours() > 240 {
		if reqIsAll == "true" {
			getAllPornDay(c, startTime, endTime)
			return
		} else {
			getPornDay(c, username, startTime, endTime)
			return
		}
	}
	if reqIsAll == "true" {
		getAllPornHour(c, startTime, endTime)
		return
	} else {
		getPornHour(c, username, startTime, endTime)
		return
	}
}
