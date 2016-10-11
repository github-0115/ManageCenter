package bill

import (
	billmodel "ManageCenter/service/model/billingmodel"
	finalbillingmodel "ManageCenter/service/model/final_billmodel"
	usermodel "ManageCenter/service/model/usermodel"
	vars "ManageCenter/service/vars"
	"fmt"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

type TransfersParmars struct {
	Username        string `json:"username"`
	FinalBillId     string `json:"finalbill_id"`
	TransferBank    string `json:"transfer_bank"`    //转账银行
	TransferAccount string `json:"transfer_account"` //转账帐号
	TransferId      string `json:"transfer_id"`      //转账单号
}

var (
	t = [...]int{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}
)

func UpdateTransfersBill(c *gin.Context) {
	var transfersParmars TransfersParmars
	if err := c.BindJSON(&transfersParmars); err != nil {
		log.Error(" AddSpecialInvoice bind json error:" + err.Error())
		c.JSON(400, gin.H{
			"code":    vars.ErrBindJson.Code,
			"message": vars.ErrBindJson.Msg,
		})
		return
	}
	finalBillId := transfersParmars.FinalBillId
	username := transfersParmars.Username
	transfersId := transfersParmars.TransferId
	bank := transfersParmars.TransferBank
	bank_num := transfersParmars.TransferAccount

	if transfersId == "" || bank == "" || bank_num == "" {
		log.Info("update user bill transfers Parameter format err (nil)")
		c.JSON(400, gin.H{
			"code":    vars.ErrParameterNil.Code,
			"message": vars.ErrParameterNil.Msg,
		})
		return
	}

	if bank_num != "" {
		if isOk, _ := regexp.MatchString(`^[\d]{16}|[\d]{19}$`, bank_num); !isOk {
			log.Info("bankaccount format incorrect.bankaccount=" + bank_num)
			c.JSON(400, gin.H{
				"code":    vars.ErrBankCountFormat.Code,
				"message": vars.ErrBankCountFormat.Msg,
			})
			return
		}
		if !checkBankCards(bank_num) {
			log.Error(fmt.Sprintf("bankaccount invalid .bankaccount=%s", bank_num))
			c.JSON(400, gin.H{
				"code":    vars.ErrBankCountInvalid.Code,
				"message": vars.ErrBankCountInvalid.Msg,
			})
			return
		}
	}

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

	err = finalbillingmodel.UpdateTransfersBill(finalBillId, username, transfersId, bank, bank_num)
	if err != nil {
		log.Error(fmt.Sprintf("update user bill transfers info failed.username=%s, err=%#v", username, err))
		if err == finalbillingmodel.ErrBillNotFound {
			c.JSON(400, gin.H{
				"code":    vars.ErrUserBillsNotfount.Code,
				"message": vars.ErrUserBillsNotfount.Msg,
			})
			return
		}
		c.JSON(400, gin.H{
			"code":    vars.ErrUserCursor.Code,
			"message": vars.ErrUserCursor.Msg,
		})
		return
	}

	finalBillColl, err := finalbillingmodel.QueryOneBill(finalBillId, username)
	if err != nil {
		log.Error(fmt.Sprintf("query user finalbill failed. err=%#v", err))
		if err == finalbillingmodel.ErrBillNotFound {
			c.JSON(400, gin.H{
				"code":    vars.ErrUserBillsNotfount.Code,
				"message": vars.ErrUserBillsNotfount.Msg,
			})
			return
		}
		c.JSON(400, gin.H{
			"code":    vars.ErrUserCursor.Code,
			"message": vars.ErrUserCursor.Msg,
		})
		return
	}

	updateBillStatus(finalBillColl.BillingId, username)

	c.JSON(200, gin.H{
		"code":    0,
		"message": "update user bill transfers info success",
	})
}

func updateBillStatus(ids []string, username string) {
	for _, res := range ids {
		err := billmodel.UpdateBillingStatus(res, username, time.Now())
		if err != nil {
			log.Error(fmt.Sprintf("update monthBill Status  failed. err=%#v", err))
		}
	}
}

func checkBankCards(s string) bool {
	odd := len(s) & 1
	var sum int
	for i, c := range s {
		if c < '0' || c > '9' {
			return false
		}
		if i&1 == odd {
			sum += t[c-'0']
		} else {
			sum += int(c - '0')
		}
	}
	return sum%10 == 0
}
