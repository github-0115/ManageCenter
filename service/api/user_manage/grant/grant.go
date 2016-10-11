package grant

import (
	grant "ManageCenter/service/model/grantmodel"
	vars "ManageCenter/service/vars"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
)

type GrantParams struct {
	Username string `json:"username"`
	Porn     int64  `json:"porn"`
	Screen   int64  `json:"screen"`
	Violence int64  `json:"violence"`
}

func GetGrant(c *gin.Context) {

	username := c.Query("username")

	if username == "" {
		c.JSON(400, gin.H{
			"code":    vars.ErrParameter.Code,
			"message": vars.ErrParameter.Msg,
		})
		return
	}

	usergrant, err := grant.GetRights(username)
	if err != nil {
		log.Error(fmt.Sprintf("update user=%s grant error:%s:", username, err))
		if err == grant.ErrGrantNotFound {
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
		}
		c.JSON(400, gin.H{
			"code":    vars.ErrUserGrant.Code,
			"message": vars.ErrUserGrant.Msg,
		})
		return
	}

	c.JSON(200, gin.H{
		"code":     0,
		"porn":     usergrant.Porn,
		"screen":   usergrant.Screen,
		"violence": usergrant.Violence,
	})
}

func ChangeGrant(c *gin.Context) {

	var grantParams GrantParams
	if err := c.BindJSON(&grantParams); err != nil {
		log.Error(fmt.Sprintf("bind json error:%s", err))
		c.JSON(400, gin.H{
			"code":    vars.ErrParameter.Code,
			"message": vars.ErrParameter.Msg,
		})
		return
	}

	username := grantParams.Username
	porn := grantParams.Porn
	screen := grantParams.Screen
	viloence := grantParams.Violence

	if username == "" {
		c.JSON(400, gin.H{
			"code":    vars.ErrParameter.Code,
			"message": vars.ErrParameter.Msg,
		})
		return
	}

	_, err := grant.UpdateRights(username, "porn", porn)
	if err != nil {
		log.Error(fmt.Sprintf("update user=%s grant error:%s:", username, err))
		if err == grant.ErrGrantNotFound {
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
		}
		c.JSON(400, gin.H{
			"code":    vars.ErrUserGrant.Code,
			"message": vars.ErrUserGrant.Msg,
		})
		return
	}

	_, err = grant.UpdateRights(username, "screen", screen)
	if err != nil {
		log.Error(fmt.Sprintf("update user=%s grant error:%s:", username, err))
		c.JSON(400, gin.H{
			"code":    vars.ErrUserGrant.Code,
			"message": vars.ErrUserGrant.Msg,
		})
		return
	}

	_, err = grant.UpdateRights(username, "viloence", viloence)
	if err != nil {
		log.Error(fmt.Sprintf("update user=%s grant error:%s:", username, err))
		c.JSON(400, gin.H{
			"code":    vars.ErrUserGrant.Code,
			"message": vars.ErrUserGrant.Msg,
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    0,
		"message": "grant update success",
	})
}
