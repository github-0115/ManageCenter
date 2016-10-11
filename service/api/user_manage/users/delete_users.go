package users

import (
	finalbillingmodel "ManageCenter/service/model/final_billmodel"
	grant "ManageCenter/service/model/grantmodel"
	keymodel "ManageCenter/service/model/keymodel"
	servpkgmodel "ManageCenter/service/model/servicepackagemodel"
	user "ManageCenter/service/model/usermodel"
	walletmodel "ManageCenter/service/model/walletmodel"
	vars "ManageCenter/service/vars"
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

type DeleteParams struct {
	Username string `json:"username"`
}

func Delete(c *gin.Context) {

	var deleteParams DeleteParams
	if err := c.BindJSON(&deleteParams); err != nil {
		log.Error("bind json error:" + err.Error())
		c.JSON(400, gin.H{
			"code":    vars.ErrParameter.Code,
			"message": vars.ErrParameter.Msg,
		})
		return
	}
	username := deleteParams.Username

	userColl, err := user.QueryUser(username)
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

	finalBillColl, err := finalbillingmodel.QueryUserBill(username, 0, userColl.Period)
	if err != nil {
		log.Error(fmt.Sprintf("query user finalbill failed. err=%#v", err))
		if err != finalbillingmodel.ErrBillNotFound {
			c.JSON(400, gin.H{
				"code":    vars.ErrUserCursor.Code,
				"message": vars.ErrUserCursor.Msg,
			})
			return
		}
	}
	if finalBillColl != nil {
		log.Error(fmt.Sprintf("user finalbill not Settlement. err=%#v", err))
		c.JSON(400, gin.H{
			"code":    vars.ErrUserBillsNotSettlement.Code,
			"message": vars.ErrUserBillsNotSettlement.Msg,
		})
		return
	}

	//删除用户
	err = user.DeleteUser(userColl.UserId, username)
	if err != nil {
		log.Error(fmt.Sprintf("delete user err", err))
		if err == user.ErrUserNotFound {
			c.JSON(200, gin.H{
				"code":    vars.ErrUserNotFound.Code,
				"message": vars.ErrUserNotFound.Msg,
			})
			return
		}

		c.JSON(200, gin.H{
			"code":    vars.ErrUserCursor.Code,
			"message": vars.ErrUserCursor.Msg,
		})
		return
	}
	//删除用户wallet
	err = walletmodel.DeleteWallet(username)
	if err != nil {
		log.Error(fmt.Sprintf("DeleteWallet not found, user = %s", username))
		if err == walletmodel.ErrUserWalletNotFound {
			c.JSON(200, gin.H{
				"code":    vars.ErrUserWalletNotfount.Code,
				"message": vars.ErrUserWalletNotfount.Msg,
			})
			return
		}
		c.JSON(200, gin.H{
			"code":    vars.ErrUserCursor.Code,
			"message": vars.ErrUserCursor.Msg,
		})
		return
	}

	//删除用户ak\sk
	err = keymodel.DeleteUserKey(userColl.UserId)
	if err != nil {
		log.Error(fmt.Sprintf("DeactivateKey key not found, keyId = %s", userColl.UserId))
		if err == keymodel.ErrKeyNotFound {
			c.JSON(200, gin.H{
				"code":    vars.ErrKeyNotFound.Code,
				"message": vars.ErrKeyNotFound.Msg,
			})
			return
		}
		c.JSON(200, gin.H{
			"code":    vars.ErrUserCursor.Code,
			"message": vars.ErrUserCursor.Msg,
		})
		return
	}
	//删除用户权限
	err = grant.DeleteUserRights(username)
	if err != nil {
		log.Error(fmt.Sprintf("delete user=%s grant error:%s:", username, err))
		c.JSON(400, gin.H{
			"code":    vars.ErrUserGrant.Code,
			"message": vars.ErrUserGrant.Msg,
		})
		return
	}
	//删除用户服务配置
	err = servpkgmodel.DeleteUserInterfaceConfig(username)
	if err != nil {
		log.Error(fmt.Sprintf("delete interface config err ", err))
		c.JSON(400, gin.H{
			"code":    vars.ErrUserServicePackageNotFound.Code,
			"message": vars.ErrUserServicePackageNotFound.Msg,
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    0,
		"message": "用户 " + username + " 删除成功！",
	})

}
