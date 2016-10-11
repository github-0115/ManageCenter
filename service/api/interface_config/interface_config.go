package interface_config

import (
	servpkgmodel "ManageCenter/service/model/servicepackagemodel"
	vars "ManageCenter/service/vars"
	"fmt"
	//	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

func GetInterfaceConfig(c *gin.Context) {
	interface_config, err := servpkgmodel.GetInterfaceConfig()
	if err != nil {
		log.Error(fmt.Sprintf("find interface config err ", err))
		c.JSON(400, gin.H{
			"code":    vars.ErrInterfaceConfigNotFound.Code,
			"message": vars.ErrInterfaceConfigNotFound.Msg,
		})
		return
	}

	c.JSON(200, gin.H{
		"default_month_total": interface_config.DefaultMonthTotal,
		"default_concurrency": interface_config.DefaultConcurrency,
		"free_day_total":      interface_config.FreeDayTotal,
		"free_concurrency":    interface_config.FreeConcurrency,
		"default_price":       interface_config.DefaultPrice,
		"validity_time":       interface_config.ValidityTime,
	})

}

type SetICParams struct {
	DefaultPrice       int64 `json:"default_price"`
	DefaultMonthTotal  int64 `json:"default_month_total"`
	DefaultConcurrency int64 `json:"default_concurrency"`
	ValidityTime       int64 `json:"validity_time"`
	FreeDayTotal       int64 `json:"free_day_total"`
	FreeConcurrency    int64 `json:"free_concurrency"`
}

func SetInterfaceConfig(c *gin.Context) {

	var setICParams SetICParams
	if err := c.BindJSON(&setICParams); err != nil {
		log.Error("bind json error:" + err.Error())
		c.JSON(400, gin.H{
			"code":    vars.ErrParameter.Code,
			"message": vars.ErrParameter.Msg,
		})
		return
	}
	default_price := setICParams.DefaultPrice
	default_month_total := setICParams.DefaultMonthTotal
	default_concurrency := setICParams.DefaultConcurrency
	validity_time := setICParams.ValidityTime
	free_day_total := setICParams.FreeDayTotal
	free_concurrency := setICParams.FreeConcurrency

	log.Info(fmt.Sprintf("default %v %v %v %v, free: %v %v", default_month_total, default_concurrency, default_price, validity_time, free_day_total, free_concurrency))

	_, err := servpkgmodel.UpdateInterfaceConfig(int64(default_month_total), int64(default_concurrency), int64(free_day_total), int64(free_concurrency), default_price, validity_time)
	if err != nil {
		log.Error(fmt.Sprintf("find interface config err ", err))
		c.JSON(400, gin.H{
			"code":    vars.ErrInterfaceConfigNotFound.Code,
			"message": vars.ErrInterfaceConfigNotFound.Msg,
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 0,
	})

}
