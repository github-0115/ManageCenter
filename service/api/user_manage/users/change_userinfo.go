package users

import (
	servpkgmodel "ManageCenter/service/model/servicepackagemodel"
	user "ManageCenter/service/model/usermodel"
	vars "ManageCenter/service/vars"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

type ChangeParams struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Company  string `json:"company"`
	UserType int64  `json:"usertype"`
}

type IsSuccess struct {
	Success string
	Fail    string
}

var (
	changeSuccess = &IsSuccess{"用户信息修改成功！", "用户信息修改失败！"}
)

func ChangeUserInfo(c *gin.Context) {

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
	phone := changeParams.Phone
	email := changeParams.Email
	name := changeParams.Name
	company := changeParams.Company
	usertype := changeParams.UserType

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

	if email != "" {
		if !emailRegex.MatchString(email) {
			log.Info("email format incorrect.email=" + email)
			c.JSON(400, gin.H{
				"code":    vars.ErrEmailFormat.Code,
				"message": vars.ErrEmailFormat.Msg,
			})
			return
		}

		existUserNum, err := user.QueryEmailCount("email", email)
		if err != nil {
			log.Error(fmt.Sprintf("query user failed. username=%s email=%s, err=%#v", username, email, err))
			if err != user.ErrUserNotFound {
				c.JSON(400, gin.H{
					"code":    vars.ErrUserCursor.Code,
					"message": vars.ErrUserCursor.Msg,
				})
				return
			}
		}
		if existUserNum >= 1 {
			if !strings.EqualFold(email, usercoll.Email) {
				log.Error(fmt.Sprintf("user num >= 1, %i", existUserNum))
				c.JSON(400, gin.H{
					"code":    vars.ErrSignEmail.Code,
					"message": vars.ErrSignEmail.Msg,
				})
				return
			}
		}

		_, err = user.ChangeUserInfo(username, "email", email)
		if err != nil {
			log.Error(fmt.Sprintf("ChangeUserInfo failed. username=%s email=%s, err=%#v", username, email, err))
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
	}

	if phone != "86" {
		if !phoneRegex.MatchString(phone) {
			log.Info("phone format incorrect.phone=" + phone)
			c.JSON(400, gin.H{
				"code":    vars.ErrParameterFormat.Code,
				"message": vars.ErrParameterFormat.Msg,
			})
			return
		}

		phoneUserNum, err := user.QueryUserCount("phonenum", phone)
		if err != nil {
			log.Error(fmt.Sprintf("query user failed.phone=%s, err=%#v", phone, err))
			if err != user.ErrUserNotFound {
				c.JSON(400, gin.H{
					"code":    vars.ErrUserCursor.Code,
					"message": vars.ErrUserCursor.Msg,
				})
				return
			}
			c.JSON(400, gin.H{
				"code":    vars.ErrOther.Code,
				"message": err,
			})
			return
		}
		if phoneUserNum >= 1 {
			if !strings.EqualFold(email, usercoll.Email) {
				log.Error(fmt.Sprintf("user phone num >= 1, %i", phoneUserNum))
				c.JSON(400, gin.H{
					"code":    vars.ErrSignPhone.Code,
					"message": vars.ErrSignPhone.Msg,
				})
				return
			}
		}

		_, err = user.ChangeUserInfo(username, "phonenum", phone)
		if err != nil {
			log.Error(fmt.Sprintf("ChangeUserInfo failed. username=%s phone=%s, err=%#v", username, phone, err))
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
	}

	if name != "" {
		_, err := user.ChangeUserInfo(username, "name", name)
		if err != nil {
			log.Error(fmt.Sprintf("ChangeUserInfo failed. username=%s name=%s, err=%#v", username, name, err))
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

	}

	if company != "" {
		_, err := user.ChangeUserInfo(username, "company", company)
		if err != nil {
			log.Error(fmt.Sprintf("ChangeUserInfo failed.  username=%s company=%s, err=%#v", username, company, err))
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

	}

	if usertype >= 0 && usertype < 3 {
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

		if usertype == 0 {
			free_user_package, err := servpkgmodel.GetInterfaceConfig()
			if err != nil {
				log.Error(fmt.Sprintf("find interface config err ", err))
				c.JSON(400, gin.H{
					"code":    vars.ErrUserServicePackageNotFound.Code,
					"message": vars.ErrUserServicePackageNotFound.Msg,
				})
				return
			} else {
				month_total := free_user_package.FreeDayTotal * 30
				concurrency := free_user_package.FreeConcurrency
				period := "月结"

				_, err = servpkgmodel.UpdateUserServicePackage(username, int64(month_total), int64(month_total/30), int64(concurrency), period)
				if err != nil {
					log.Error(fmt.Sprintf("set user service package err ", err))
					c.JSON(400, gin.H{
						"code":    vars.ErrUserServicePackageNotFound.Code,
						"message": vars.ErrUserServicePackageNotFound.Msg,
					})
					return
				}
			}

		} else {
			_, err := servpkgmodel.UpdateUserServicePackage(username, 1e15, 3.4e13, 10, "月结")
			if err != nil {
				log.Error(fmt.Sprintf("set user service package err ", err))
				c.JSON(400, gin.H{
					"code":    vars.ErrUserServicePackageNotFound.Code,
					"message": vars.ErrUserServicePackageNotFound.Msg,
				})
				return
			}
		}

	}

	c.JSON(200, gin.H{
		"code":    0,
		"message": changeSuccess.Success,
	})

}
