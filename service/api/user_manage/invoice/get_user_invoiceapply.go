package invoice

import (
	ads "ManageCenter/service/model/addressmodel"
	invoiceapplymodel "ManageCenter/service/model/invoiceapplymodel"
	invoicemodel "ManageCenter/service/model/invoicemodel"
	usermodel "ManageCenter/service/model/usermodel"
	vars "ManageCenter/service/vars"
	"fmt"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

type Rep struct {
	Apply   []*invoiceapplymodel.InvoiceApplyInfo `json:"apply"`
	Invoice []*invoicemodel.InvoiceInfo           `json:"invoice"`
	Address []*ads.AddressInfo                    `json:"address"`
}

func GetUserInvoiceApply(c *gin.Context) {
	username := c.Query("user_name")
	pageIndex, err := strconv.Atoi(c.Query("page"))
	pageSize, err := strconv.Atoi(c.Query("rows"))
	sorted := c.Query("sort")
	sord := c.Query("sord")
	if err != nil {
		log.Error(fmt.Sprintf("strconv Atoi err%v", err))
		c.JSON(400, gin.H{
			"code":    vars.ErrParameter.Code,
			"message": vars.ErrParameter.Msg,
		})
		return
	}
	if sord == "desc" {
		sorted = "-" + sorted
	}
	if sorted == "" {
		sorted = "-created_at"
	}

	_, err = usermodel.QueryUser(username)
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

	invoiceapply, records, err := invoiceapplymodel.GetPagingUserInvoiceApply(username, pageIndex, pageSize, sorted)
	if err != nil {
		log.Error(fmt.Sprintf("query user all invoice apply failed.username=%s, err=%#v", username, err))
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
		invoice = append(invoice, im)

		ad, err := ads.QueryOneAddress(res.AddressId, res.Username)
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
		address = append(address, ad)
	}

	pagetotal := int(math.Ceil(float64(records) / float64(pageSize)))

	rep := &Rep{
		Apply:   apply,
		Invoice: invoice,
		Address: address,
	}

	c.JSON(200, gin.H{
		"code":      0,
		"total":     records,
		"history":   *rep,
		"page":      pageIndex,
		"pagetotal": pagetotal,
	})
}
