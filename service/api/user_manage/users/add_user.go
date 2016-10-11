package users

import (
	grant "ManageCenter/service/model/grantmodel"
	keymodel "ManageCenter/service/model/keymodel"
	servpkgmodel "ManageCenter/service/model/servicepackagemodel"
	usermodel "ManageCenter/service/model/usermodel"
	walletmodel "ManageCenter/service/model/walletmodel"

	vars "ManageCenter/service/vars"
	security "ManageCenter/utils/security"
	"fmt"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
	"github.com/satori/go.uuid"
)

type AddParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Company  string `json:"company"`
}

var (
	usernamePattern   = "\\A[a-zA-Z0-9_-]{2,16}\\z"
	emailPattern      = "^[_a-z0-9-]+(\\.[_a-z0-9-]+)*@[a-z0-9-]+(\\.[a-z0-9-]+)*(\\.[a-z]{2,4})$"
	phonePattern      = "\\A86(13[0-9]|14[57]|15[012356789]|17[0678]|18[0-9])[0-9]{8}\\z"
	phoneRegex, _     = regexp.Compile(phonePattern)
	usernameRegexx, _ = regexp.Compile(usernamePattern)
	emailRegex, _     = regexp.Compile(emailPattern)
)

func AddUser(c *gin.Context) {

	var addParams AddParams
	if err := c.BindJSON(&addParams); err != nil {
		log.Error("bind json error:" + err.Error())
		c.JSON(400, gin.H{
			"code":    vars.ErrParameter.Code,
			"message": vars.ErrParameter.Msg,
		})
		return
	}
	username := addParams.Username
	password := addParams.Password
	phone := addParams.Phone
	email := addParams.Email
	name := addParams.Name
	company := addParams.Company

	if !usernameRegexx.MatchString(username) {
		log.Info("username format incorrect.username=" + username)
		c.JSON(400, gin.H{
			"code":    vars.ErrUsernameFormat.Code,
			"message": vars.ErrUsernameFormat.Msg,
		})
		return
	}

	user_id := uuid.NewV4().String()
	existUserNum, err := usermodel.QueryUserNum(user_id, username)
	if err != nil {
		log.Error(fmt.Sprintf("query user failed.username=%s, err=%#v", username, err))
		if err == usermodel.ErrUserNotFound {
			log.Error(fmt.Sprintf(username + "可用"))
		} else if err == usermodel.ErrUserCursor {
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
	if existUserNum >= 1 {
		log.Error(fmt.Sprintf("user num >= 1, %i", existUserNum, ", user_id = %v", user_id))
		c.JSON(400, gin.H{
			"code":    vars.ErrUserExist.Code,
			"message": vars.ErrUserExist.Msg,
		})
		return
	}
	savedPassword := security.GeneratePasswordHash(password)

	if email != "" {
		if !emailRegex.MatchString(email) {
			log.Info("email format incorrect.email=" + email)
			c.JSON(400, gin.H{
				"code":    vars.ErrEmailFormat.Code,
				"message": vars.ErrEmailFormat.Msg,
			})
			return
		}
		existUserNum, err := usermodel.QueryUserCount("email", email)
		if err != nil {
			log.Error(fmt.Sprintf("query user failed.email=%s, err=%#v", email, err))
			if err == usermodel.ErrUserNotFound {
				log.Error(fmt.Sprintf(username + "可用"))
			} else if err == usermodel.ErrUserCursor {
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
		if existUserNum >= 1 {
			log.Error(fmt.Sprintf("user email num >= 1, %i", existUserNum))
			c.JSON(400, gin.H{
				"code":    vars.ErrSignEmail.Code,
				"message": vars.ErrSignEmail.Msg,
			})
			return
		}
	}

	if !phoneRegex.MatchString(phone) {
		log.Info("phone format incorrect.phone=" + phone)
		c.JSON(400, gin.H{
			"code":    vars.ErrParameterFormat.Code,
			"message": vars.ErrParameterFormat.Msg,
		})
		return
	}
	phoneUserNum, err := usermodel.QueryUserCount("phonenum", phone)
	if err != nil {
		log.Error(fmt.Sprintf("query user failed.phone=%s, err=%#v", phone, err))
		if err == usermodel.ErrUserNotFound {
			log.Error(fmt.Sprintf(phone + "可用"))
		} else if err == usermodel.ErrUserCursor {
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
		log.Error(fmt.Sprintf("user phone num >= 1, %i", phoneUserNum))
		c.JSON(400, gin.H{
			"code":    vars.ErrSignPhone.Code,
			"message": vars.ErrSignPhone.Msg,
		})
		return
	}

	user := &usermodel.UserColl{
		UserId:    user_id,
		Username:  username,
		Password:  savedPassword,
		Email:     email,
		PhoneNum:  phone,
		Name:      name,
		Company:   company,
		Remark:    "",
		Status:    0,
		Type:      2,
		CreatedAt: time.Now(),
	}
	err = user.Save()
	if err != nil {
		log.Error("user info save err " + err.Error())
		c.JSON(400, gin.H{
			"code":    vars.ErrUserSave.Code,
			"message": vars.ErrUserSave.Msg,
		})
		return
	}

	wallet := walletmodel.Wallet{
		Username:       username,
		Balance:        0.0,
		Total:          0.0,
		InvoicedAmount: 0.0,
		Status:         walletmodel.WalletEnable,
		UpdatedAt:      time.Now(),
		CreatedAt:      time.Now(),
	}

	err = wallet.Save()
	if err != nil {
		log.Error("user wallet save err " + err.Error())

		usererr := usermodel.DeleteUser(user_id, username)
		if usererr != nil {
			log.Error(fmt.Sprintf("delete user err", usererr))
			if usererr != usermodel.ErrUserNotFound {
				c.JSON(200, gin.H{
					"code":    vars.ErrUserCursor.Code,
					"message": vars.ErrUserCursor.Msg,
				})
				return
			}

		}

		c.JSON(400, gin.H{
			"code":    vars.ErrSignSave.Code,
			"message": vars.ErrSignSave.Msg,
		})
		return
	}

	keyId := uuid.NewV4().String()
	accessKey := uuid.NewV4().String()
	secretkey := uuid.NewV4().String()

	keyColl := &keymodel.KeyColl{
		UserId:    user_id,
		KeyId:     keyId,
		AccessKey: accessKey,
		SecretKey: secretkey,
		IsActive:  true,
		IsExcess:  false,
		CreatedAt: time.Now(),
	}

	err = keyColl.Save()
	if err != nil {
		log.Error("user keyColl save err " + err.Error())

		usererr := usermodel.DeleteUser(user_id, username)
		if usererr != nil {
			log.Error(fmt.Sprintf("delete user err", usererr))
			if usererr != usermodel.ErrUserNotFound {
				c.JSON(200, gin.H{
					"code":    vars.ErrUserCursor.Code,
					"message": vars.ErrUserCursor.Msg,
				})
				return
			}
		}

		walleterr := walletmodel.DeleteWallet(username)
		if walleterr != nil {
			log.Error(fmt.Sprintf("delete user err", walleterr))
			if walleterr != walletmodel.ErrUserWalletNotFound {
				c.JSON(200, gin.H{
					"code":    vars.ErrUserCursor.Code,
					"message": vars.ErrUserCursor.Msg,
				})
				return
			}

		}

		c.JSON(400, gin.H{
			"code":    vars.ErrSignSave.Code,
			"message": vars.ErrSignSave.Msg,
		})
		return
	}

	rights := &grant.Rights{
		Username:  username,
		Porn:      1,
		Screen:    0,
		Violence:  0,
		CreatedAt: time.Now(),
	}
	err = rights.Save()
	if err != nil {
		log.Error(fmt.Sprintf(" user grant err", err))
	}

	userServicePackage := &servpkgmodel.UserServicePackage{
		Username:    username,
		UserType:    2,
		MonthTotal:  1e15,
		DailyAmount: 3.4e13,
		Concurrency: 10,
		Period:      "月结",
		CreatedAt:   time.Now(),
	}

	err = userServicePackage.Save()
	if err != nil {
		log.Error(fmt.Sprintf(" save user service package err", err))
	}

	c.JSON(200, gin.H{
		"code":    0,
		"message": "用户 " + username + " 添加成功！",
	})

}
