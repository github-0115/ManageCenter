package middleware

import (
	"ManageCenter/config"
	configmodel "ManageCenter/service/model/config"
	vars "ManageCenter/service/vars"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

func AuthToken(c *gin.Context) {
	tokenStr := c.Request.Header.Get("LoginToken")
	if tokenStr == "" {
		c.JSON(401, gin.H{
			"code":    vars.ErrNeedToken.Code,
			"message": vars.ErrNeedToken.Msg,
		})
		c.Abort()
		return
	}
	_, ok := configmodel.BanTokens[tokenStr]
	if ok {
		c.JSON(401, gin.H{
			"code":    vars.ErrInvalidToken.Code,
			"message": vars.ErrInvalidToken.Msg,
		})
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Cfg.LoginSecret), nil
	})

	if err != nil || !token.Valid {
		log.Error(fmt.Sprintf("[AUTH] token error: ", err, " token: ", token))
		c.JSON(401, gin.H{
			"code":    vars.ErrInvalidToken.Code,
			"message": vars.ErrInvalidToken.Msg,
		})
		c.Abort()
		return
	}

	claims := token.Claims.(jwt.MapClaims)

	if exp, ok := claims["exp"].(float64); !ok || exp <= 0.0 {
		c.JSON(401, gin.H{
			"code":    vars.ErrIncompleteToken.Code,
			"message": vars.ErrIncompleteToken.Msg,
		})
		c.Abort()
		return
	}

	user_id, ok := claims["manager_id"].(string)
	username, ok := claims["managername"].(string)

	if !ok {

		c.JSON(401, gin.H{
			"message": "username in token must be a string",
		})
		c.Abort()
		return
	} else if username == "" {
		c.JSON(401, gin.H{
			"message": "token must contain a username",
		})
		c.Abort()
		return
	}

	c.Set("user_id", user_id)
	c.Set("username", username)
	c.Next()
}
