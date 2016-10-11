package invoice

import (
	invoiceapplymodel "ManageCenter/service/model/invoiceapplymodel"
	user "ManageCenter/service/model/usermodel"
	vars "ManageCenter/service/vars"
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

type ChangeParams struct {
	Username      string `json:"username"`
	ApplyId       string `json:"apply_id"`
	LogisticsName string `json:"logistics_name"`
	TrackNum      string `json:"track_num"`
}

type IsSuccess struct {
	Success string
	Fail    string
}

var (
	changeSuccess = &IsSuccess{"用户发票邮寄申请处理成功！", "用户发票邮寄申请处理失败！"}
)

func ChangeInvoiceAply(c *gin.Context) {

	var changeParams ChangeParams
	if err := c.BindJSON(&changeParams); err != nil {
		log.Error("bind json error:" + err.Error())
		c.JSON(400, gin.H{
			"code":    vars.ErrParameter.Code,
			"message": vars.ErrParameter.Msg,
		})
		return
	}
	username := changeParams.Username
	applyId := changeParams.ApplyId
	logisticsName := changeParams.LogisticsName
	trackNum := changeParams.TrackNum

	if username == "" || applyId == "" || logisticsName == "" || trackNum == "" {
		log.Error(fmt.Sprintf("user change invioce apply Parameter format err (nil),username=%s ,applyId=%s ,logisticsName=%s ,trackNum=%s ,", username, applyId, logisticsName, trackNum))
		c.JSON(400, gin.H{
			"code":    vars.ErrParameterNil.Code,
			"message": vars.ErrParameterNil.Msg,
		})
		return
	}

	_, err := user.QueryUser(username)
	if err != nil {
		log.Error(fmt.Sprintf("query user failed. username=%s  err=%#v", username, err))
		if err == user.ErrUserNotFound {
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

	statusSuccess, err := invoiceapplymodel.UpdateInvoiceApply(applyId, username, "status", 1)
	if err != nil {
		if err == invoiceapplymodel.ErrUserInvoiceApplyNotFound {
			log.Error(fmt.Sprintf("update user invoice apply failed. username=%s  err=%#v", username, err))
			c.JSON(400, gin.H{
				"code":    vars.ErrUserInvoiceApplyNotfount.Code,
				"message": vars.ErrUserInvoiceApplyNotfount.Msg,
			})
			return
		}
		c.JSON(400, gin.H{
			"code":    vars.ErrUserCursor.Code,
			"message": vars.ErrUserCursor.Msg,
		})
		return
	}

	logisticsSuccess, err := invoiceapplymodel.UpdateInvoiceApplyInfo(applyId, username, "logistics_name", logisticsName)
	if err != nil {
		if err == invoiceapplymodel.ErrUserInvoiceApplyNotFound {
			log.Error(fmt.Sprintf("update user invoice apply failed. username=%s  err=%#v", username, err))
			c.JSON(400, gin.H{
				"code":    vars.ErrUserInvoiceApplyNotfount.Code,
				"message": vars.ErrUserInvoiceApplyNotfount.Msg,
			})
			return
		}
		c.JSON(400, gin.H{
			"code":    vars.ErrUserCursor.Code,
			"message": vars.ErrUserCursor.Msg,
		})
		return
	}

	trackSuccess, err := invoiceapplymodel.UpdateInvoiceApplyInfo(applyId, username, "track_num", trackNum)
	if err != nil {
		if err == invoiceapplymodel.ErrUserInvoiceApplyNotFound {
			log.Error(fmt.Sprintf("update user invoice apply failed. username=%s  err=%#v", username, err))
			c.JSON(400, gin.H{
				"code":    vars.ErrUserInvoiceApplyNotfount.Code,
				"message": vars.ErrUserInvoiceApplyNotfount.Msg,
			})
			return
		}
		c.JSON(400, gin.H{
			"code":    vars.ErrUserCursor.Code,
			"message": vars.ErrUserCursor.Msg,
		})
		return
	}

	if !statusSuccess || !logisticsSuccess || !trackSuccess {
		c.JSON(400, gin.H{
			"code":    vars.ErrUserInvoiceApplyChange.Code,
			"message": changeSuccess.Fail,
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    0,
		"message": changeSuccess.Success,
	})
}
