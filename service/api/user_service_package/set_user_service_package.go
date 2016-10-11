package user_service_package

import (
	servpkgmodel "ManageCenter/service/model/servicepackagemodel"
	user "ManageCenter/service/model/usermodel"
	vars "ManageCenter/service/vars"
	emailclient "ManageCenter/utils/emailclient"
	redisclient "ManageCenter/utils/redisclient"
	"ManageCenter/utils/smsclient"
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

type SetUSPParams struct {
	Username    string `json:"username"`
	UserType    int64  `json:"usertype"`
	Price       int64  `json:"price"`
	MonthTotal  int64  `json:"month_total"`
	Concurrency int64  `json:"concurrency"`
	FreeDay     int64  `json:"freeday"`
	Period      string `json:"period"`
}

func SetUserServicePackage(c *gin.Context) {
	var setUSPParams SetUSPParams
	if err := c.BindJSON(&setUSPParams); err != nil {
		log.Error("bind json error:" + err.Error())
		c.JSON(400, gin.H{
			"code":    vars.ErrParameter.Code,
			"message": vars.ErrParameter.Msg,
		})
		return
	}

	username := setUSPParams.Username
	usertype := setUSPParams.UserType
	month_total := setUSPParams.MonthTotal
	concurrency := setUSPParams.Concurrency
	price := setUSPParams.Price
	freeday := setUSPParams.FreeDay
	period := setUSPParams.Period

	daily_amount := month_total / 30
	var default_daily int64
	userServPkg, err := servpkgmodel.GetUserServicePackage(username)
	if err != nil {
		log.Error(fmt.Sprintf("query user service package failed. username=%s  err=%#v", username, err))
		freeUserPackage, err := servpkgmodel.GetInterfaceConfig()
		if err != nil {
			log.Error(fmt.Sprintf("find interface config err ", err))
			c.JSON(400, gin.H{
				"code":    vars.ErrUserServicePackageNotFound.Code,
				"message": vars.ErrUserServicePackageNotFound.Msg,
			})
			return
		} else {
			default_daily = freeUserPackage.FreeDayTotal
		}
	} else {
		default_daily = userServPkg.DailyAmount
	}

	// new daily amount takes effect immediately
	if daily_amount != default_daily {
		err = redisclient.SetAPIDailyUserTotal(username, int(daily_amount))
		log.Info(fmt.Sprintf("set user service package %v, %v, %v", username, daily_amount, userServPkg.DailyAmount))

		if err != nil {
			log.Error(fmt.Sprintf("set api daily user total err:%v, username:%v", err, username))
		}
	}

	log.Info(fmt.Sprintf("set user service package %v, %v, %v, %v, %v", username, month_total, daily_amount, concurrency, period))
	usercoll, err := user.QueryUser(username)
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

	switch usertype {
	case 0:
		_, err := servpkgmodel.UpdateFreeUserServicePackage(username, usertype, daily_amount, freeday)
		if err != nil {
			log.Error(fmt.Sprintf("set user service package err ", err))
			c.JSON(400, gin.H{
				"code":    vars.ErrUserServicePackageNotFound.Code,
				"message": vars.ErrUserServicePackageNotFound.Msg,
			})
			return
		}

	case 1:
		c.JSON(200, gin.H{
			"code":    0,
			"message": "this feature is not yet open",
		})
		return

	case 2:
		_, err := servpkgmodel.UpdatePostPaidUserServicePackage(username, usertype, month_total, daily_amount, price, concurrency, period)
		if err != nil {
			log.Error(fmt.Sprintf("set user service package err ", err))
			c.JSON(400, gin.H{
				"code":    vars.ErrUserServicePackageNotFound.Code,
				"message": vars.ErrUserServicePackageNotFound.Msg,
			})
			return
		}

	case 3:
		_, err := servpkgmodel.UpdateAgentsUserServicePackage(username, usertype, 1e15, 3.33333e13, 10, period) //3.4e13
		if err != nil {
			log.Error(fmt.Sprintf("set user service package err ", err))
			c.JSON(400, gin.H{
				"code":    vars.ErrUserServicePackageNotFound.Code,
				"message": vars.ErrUserServicePackageNotFound.Msg,
			})
			return
		}

	default:
		c.JSON(200, gin.H{
			"code":    0,
			"message": "user type err",
		})
		return
	}

	if usertype != usercoll.Type {
		_, err := user.ChangeUserType(username, usertype)
		if err != nil {
			log.Error(fmt.Sprintf("ChangeUserInfo failed.  username=%s usertype=%s, err=%#v", username, usertype, err))
			if err == user.ErrUserNotFound {
				c.JSON(400, gin.H{
					"code":    vars.ErrUserNotFound.Code,
					"message": vars.ErrUserNotFound.Msg,
				})
				return
			}
			c.JSON(400, gin.H{
				"code":    vars.ErrOther.Code,
				"message": vars.ErrOther.Msg,
			})
			return
		}

		utype := changeUserTypeToString(usertype)

		result := smsclient.SendPackageSMS(usercoll.PhoneNum, username, utype)
		if result != "" {
			log.Error(fmt.Sprintf("send user package apply deal sms failed.", result))
		}

		if usercoll.Email != "" && usercoll.EmailStatus == 1 {
			email_body := "尊敬的" + username + "，您的账户已经成功变更成" + utype + "，有任何问题，请与我们联系：bd@deepir.com"
			err := emailclient.SendPackageApplyMail(email_body, usercoll.Email)
			if err != nil {
				log.Error(fmt.Sprintf("send user package apply deal email failed.", err))
			}
		}

	}

	c.JSON(200, gin.H{
		"code":    0,
		"message": "set user service success",
	})
	return
}

func changeUserTypeToString(usertype int64) string {
	var utype string
	switch usertype {
	case 0:
		utype = "免费用户"
	case 1:
		utype = "预付费用户"
	case 2:
		utype = "后付费用户"
	case 3:
		utype = "代理商"
	default:
		utype = "免费用户"
	}
	return utype
}
