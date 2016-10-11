package login_token

import (
	cfg "ManageCenter/config"
	manager "ManageCenter/service/model/managermodel"
	security "ManageCenter/utils/security"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

type LoginParams struct {
	Managername string `json:"managername"`
	Password    string `json:"password"`
}

func Post(c *gin.Context) {

	var loginParams LoginParams
	if err := c.BindJSON(&loginParams); err != nil {
		log.Error("bind json error:" + err.Error())
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	managername := loginParams.Managername
	password := loginParams.Password

	managerColl, err := manager.Get(managername)

	if err != nil {
		log.Error(fmt.Sprintf("[login_token] query user failed.managername=%s, err=%#v", managername, err))
		c.JSON(400, gin.H{
			"message": "query manager failed",
		})
		return
	}

	if !security.CheckPasswordHash(password, managerColl.Password) {
		log.Error(fmt.Sprintf("password error.managername=%s, password=%s, saved_password=%s", managername, password, managerColl.Password))
		c.JSON(401, gin.H{
			"message": "password error",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"manager_id":  managerColl.ManagerId,
		"exp":         time.Now().Add(time.Hour * time.Duration(cfg.Cfg.LoginTokenExpire)).Unix(),
		"managername": managerColl.Managername,
	})

	tokenStr, err := token.SignedString([]byte(cfg.Cfg.LoginSecret))
	if err != nil {
		log.Error("gen token failed. err=" + err.Error())
		c.JSON(400, gin.H{
			"message": "gen token failed.",
		})
		return
	}

	c.JSON(200, gin.H{
		"code":  0,
		"token": tokenStr,
	})

}
