package invoice

import (
	ads "ManageCenter/service/model/addressmodel"
	invoiceapplymodel "ManageCenter/service/model/invoiceapplymodel"
	invoicemodel "ManageCenter/service/model/invoicemodel"
	vars "ManageCenter/service/vars"
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

func GetAllInvoiceApply(c *gin.Context) {

	invoiceapply, err := invoiceapplymodel.QueryAllInvoiceApply()
	if err != nil {
		log.Error(fmt.Sprintf("query user all invoice apply failed. err=%#v", err))
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

	apply := make([]*invoiceapplymodel.InvoiceApplyInfo, 0, len(invoiceapply))
	invoice := make([]*invoicemodel.InvoiceInfo, 0, len(invoiceapply))
	address := make([]*ads.AddressInfo, 0, len(invoiceapply))

	for _, res := range invoiceapply {
		apply = append(apply, res)

		im, err := invoicemodel.QueryOneInvoice(res.InvoiceId, res.Username)
		if err != nil {
			log.Error(fmt.Sprintf("query user invoice failed.err=%#v", err))
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
		invoice = append(invoice, im)

		ad, err := ads.QueryOneAddress(res.AddressId, res.Username)
		if err != nil {
			log.Error(fmt.Sprintf("query user address failed.err=%#v", err))
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
		address = append(address, ad)
	}

	rep := &Rep{
		Apply:   apply,
		Invoice: invoice,
		Address: address,
	}

	c.JSON(200, *rep)
}
