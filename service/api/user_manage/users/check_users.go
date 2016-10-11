package users

import (
	keymodel "ManageCenter/service/model/keymodel"
	servpkgmodel "ManageCenter/service/model/servicepackagemodel"
	redisclient "ManageCenter/utils/redisclient"

	user "ManageCenter/service/model/usermodel"
	vars "ManageCenter/service/vars"
	"ManageCenter/utils/smsclient"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
	"github.com/satori/go.uuid"
)

type CheckParams struct {
	Username string `json:"username"`
	Remark   string `json:"remark"`
	Status   int64  `json:"status"`
}

var freeday = 30

func Check(c *gin.Context) {

	var checkParams CheckParams
	if err := c.BindJSON(&checkParams); err != nil {
		log.Error("bind json error:" + err.Error())
		c.JSON(400, gin.H{
			"code":    vars.ErrParameter.Code,
			"message": vars.ErrParameter.Msg,
		})
		return
	}
	username := checkParams.Username
	status := checkParams.Status
	remark := checkParams.Remark

	userColl, err := user.QueryUser(username)
	isBan, err := user.CheckUsers(username, status, remark)
	if err != nil {
		log.Error(fmt.Sprintf("update err", err))
		if err == user.ErrUserNotFound {
			c.JSON(404, gin.H{
				"code":    vars.ErrUserNotFound.Code,
				"message": vars.ErrUserNotFound.Msg,
			})
			return
		}

		c.JSON(404, gin.H{
			"code":    vars.ErrParameter.Code,
			"message": vars.ErrParameter.Msg,
		})
		return
	}

	if isBan == 1 {
		err = deactivateKey(userColl.UserId)
		if err != nil {
			log.Error(fmt.Sprintf("update  deactivateKey err", err))
		}
		if userColl.Type != 0 {
			err = changeUserType(userColl.UserId, 0)
			if err != nil {
				log.Error(fmt.Sprintf("update changeUserType err", err))
			}
		}

	} else if isBan == 2 {
		err = deleteUserKey(userColl.UserId)
		if err != nil {
			log.Error(fmt.Sprintf("update deleteUserKey err", err))
		}
		if userColl.Type != 0 {
			err = changeUserType(userColl.UserId, 0)
			if err != nil {
				log.Error(fmt.Sprintf("update changeUserType err", err))
			}
		}

	} else {
		err = activateKey(userColl.UserId)
		if err != nil {
			log.Error(fmt.Sprintf("update  activateKey err", err))
		}
		if userColl.Type == 0 {
			err = changeUserType(userColl.UserId, 2)
			if err != nil {
				log.Error(fmt.Sprintf("update changeUserType err", err))
			}

			err = servicePackage(username, userColl.Type)
			if err != nil {
				log.Error(fmt.Sprintf("set user service package err ", err))
			}
		}

		result := smsclient.SendVerifySMS(userColl.PhoneNum, username)
		if result != "" {
			log.Error("send check sms failed." + result)
		}
	}

	c.JSON(200, gin.H{
		"code":    0,
		"message": isBan,
	})

}

func servicePackage(username string, utype int64) error {
	userpkg, err := servpkgmodel.GetUserServicePackage(username)
	var day_total int64
	if err != nil {
		log.Error(fmt.Sprintf("find user service package err ", err))
		if err == servpkgmodel.ErrUserServicePackageNotFound {

			free_user_package, err := servpkgmodel.GetInterfaceConfig()
			if err != nil {

				return servpkgmodel.ErrInterfaceConfigNotFound

			} else {
				day_total = free_user_package.FreeDayTotal

				_, err := servpkgmodel.UpdateFreeUserServicePackage(username, utype, day_total, int64(freeday))
				if err != nil {
					log.Error(fmt.Sprintf("set user service package err ", err))
					if err == servpkgmodel.ErrUserServicePackageNotFound {
						return servpkgmodel.ErrUserServicePackageNotFound
					}
					return err
				}
				return err
			}
			return err
		}
		return err
	} else {
		switch userpkg.UserType {
		case 0:
			day_total = userpkg.DailyAmount
		case 1:
			day_total = userpkg.DailyAmount
		case 2:
			day_total = userpkg.MonthTotal / 30
		case 3:
			day_total = userpkg.DailyAmount
		default:
			day_total = userpkg.DailyAmount
		}

	}

	err = redisclient.SetAPIDailyUserTotal(username, int(day_total))
	if err != nil {
		log.Error(fmt.Sprintf("set api daily user total err:%v, username:%v", err, username))
	}

	return nil
}

func deactivateKey(id string) error {
	err := keymodel.DeactivateKey(id)
	if err != nil {
		log.Error(fmt.Sprintf("DeactivateKey key not found, keyId = %s", id))
		if err == keymodel.ErrKeyNotFound {
			return keymodel.ErrKeyNotFound
		}
		return err
	}
	return nil
}

func deleteUserKey(id string) error {
	err := keymodel.DeleteUserKey(id)
	if err != nil {
		log.Error(fmt.Sprintf("deleteUserKey key not found, keyId = %s", id))
		if err == keymodel.ErrKeyNotFound {
			return keymodel.ErrKeyNotFound
		}
		return err
	}
	return nil
}

func activateKey(id string) error {
	err := keymodel.ActivateKey(id)
	if err != nil {
		log.Error(fmt.Sprintf("ActivateKey key not found, keyId = %s", id))
		if err == keymodel.ErrKeyNotFound {
			keyId := uuid.NewV4().String()
			accessKey := uuid.NewV4().String()
			secretkey := uuid.NewV4().String()

			keyColl := &keymodel.KeyColl{
				UserId:    id,
				KeyId:     keyId,
				AccessKey: accessKey,
				SecretKey: secretkey,
				IsActive:  true,
				IsExcess:  false,
				CreatedAt: time.Now(),
			}
			err = keyColl.Save()
			if err != nil {
				log.Error(fmt.Sprintf("keyColl.Save fail = %s", err))
				return err
			}
			return nil
		} else {
			return err
		}
	}
	return nil
}

func changeUserType(id string, utype int64) error {
	_, err := user.ChangeUserType(id, utype)
	if err != nil {
		log.Error("change user type not found, user_id = " + id)
		if err == user.ErrUserNotFound {

			return user.ErrUserNotFound

		} else if err == user.ErrUserCursor {

			return user.ErrUserCursor
		}
		return err
	}
	return nil
}
