package invoice

import (
	ads "ManageCenter/service/model/addressmodel"
	invoiceapplymodel "ManageCenter/service/model/invoiceapplymodel"
	invoicemodel "ManageCenter/service/model/invoicemodel"
	usermodel "ManageCenter/service/model/usermodel"
	vars "ManageCenter/service/vars"
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

func GetOneInvoiceApply(c *gin.Context) {
	username := c.Query("user_name")
	applyId := c.Query("apply_id")

	_, err := usermodel.QueryUser(username)
	if err != nil {
		log.Error(fmt.Sprintf("query user failed.username=%s, err=%#v", username, err))
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

	invoiceapply, err := invoiceapplymodel.QueryOneInvoiceApply(applyId, username)
	if err != nil {
		log.Error(fmt.Sprintf("query user invoice apply failed.username=%s, err=%#v", username, err))
		if err == invoiceapplymodel.ErrUserInvoiceApplyNotFound {
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

	invoice, err := invoicemodel.QueryOneInvoice(invoiceapply.InvoiceId, invoiceapply.Username)
	if err != nil {
		log.Error(fmt.Sprintf("query user invoice failed.username=%s, err=%#v", username, err))
		if err == invoicemodel.ErrUserInvoiceNotFound {
			c.JSON(400, gin.H{
				"code":    vars.ErrUserInvoiceNotfount.Code,
				"message": vars.ErrUserInvoiceNotfount.Msg,
			})
			return
		}
		c.JSON(400, gin.H{
			"code":    vars.ErrUserCursor.Code,
			"message": vars.ErrUserCursor.Msg,
		})
		return
	}

	address, err := ads.QueryOneAddress(invoiceapply.AddressId, invoiceapply.Username)
	if err != nil {
		log.Error(fmt.Sprintf("query user address failed.username=%s, err=%#v", username, err))
		if err == ads.ErrUserAdsNotFound {
			c.JSON(400, gin.H{
				"code":    vars.ErrUserAdsNotfount.Code,
				"message": vars.ErrUserAdsNotfount.Msg,
			})
			return
		}
		c.JSON(400, gin.H{
			"code":    vars.ErrUserCursor.Code,
			"message": vars.ErrUserCursor.Msg,
		})
		return
	}

	c.JSON(200, gin.H{
		"code":         0,
		"invoiceapply": invoiceapply,
		"invoice":      invoice,
		"address":      address,
	})
}
