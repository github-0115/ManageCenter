package redisclient

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	cfg "ManageCenter/config"

	log "github.com/inconshreveable/log15"
	"gopkg.in/redis.v3"
)

const verifySMSReqIPKey = "verify_sms_req_ip:"
const verifyEmailKey = "verify_email_ip:"
const verifyVcodeKey = "verify_vcode:"
const loginKey = "login_ip:"
const consumeKey = "consume_ip:"
const countKey = "detail_count:"
const checkEmailKey = "check_email_str:"

var redisAddr = cfg.DBCfg.RedisBroker.Host + ":" + strconv.Itoa(cfg.DBCfg.RedisBroker.Port)
var RedisNotFound = errors.New("redis not found err")
var client = redis.NewClient(&redis.Options{
	Addr:     redisAddr,
	Password: cfg.DBCfg.RedisBroker.Password, // no password set
	DB:       cfg.DBCfg.RedisBroker.DB,       // use default DB
})

// move to a stand-alone module
var apiDailyRedisAddr = cfg.DBCfg.APIDailyRedisBroker.Host + ":" + strconv.Itoa(cfg.DBCfg.APIDailyRedisBroker.Port)

//var APIDailyRedisNotFound = errors.New("api daily redis not found err")
var apiDailyClient = redis.NewClient(&redis.Options{
	Addr:     apiDailyRedisAddr,
	Password: cfg.DBCfg.APIDailyRedisBroker.Password, // no password set
	DB:       cfg.DBCfg.APIDailyRedisBroker.DB,       // use default DB
})

func IsIPReqVerifySMS(ip string) (yes bool, err error) {
	yes, err = client.Exists(verifySMSReqIPKey + ip).Result()
	log.Info(fmt.Sprintf("IsIPReqVerifySMS. ip=%s, yes=%v, err=%#v.", ip, yes, err))
	return yes, err
}

func SetVerifySMSReqIP(ip string) (err error) {
	err = client.Set(verifySMSReqIPKey+ip, 1, time.Duration(60)*time.Second).Err()
	log.Info(fmt.Sprintf("SetVerifySMSReqIP. ip=%s, err=%#v.", ip, err))
	if err != nil {
		log.Error(fmt.Sprintf("verify code exists error: %v", err))
		return err
	}
	return nil
}

func IsVerifyVcodeExists(phone string) (bool, error) {
	res, err := client.Exists(verifyVcodeKey + phone).Result()
	log.Info(fmt.Sprintf("IsVerifyVcodeExists. phone=%s, yes=%s, err=%#v.", phone, res, err))
	if err != nil {
		log.Error(fmt.Sprintf("verify code exists error: %v", err))
		return false, err
	}
	return res, nil
}

func SetEmailVcode(email string, vcode string) (err error) {
	err = client.Set(verifyEmailKey+email, vcode, time.Duration(1800)*time.Second).Err()
	if err != nil {
		log.Error(fmt.Sprintf("SetemailVcode. email=%s, vcode=%s,err=%#v.", email, vcode, err))
		return err
	}

	return err
}

func GetEmailVcode(email string) (vcode string) {
	vcode, err := client.Get(verifyEmailKey + email).Result()
	if err != nil {
		if err == redis.Nil {
			log.Error(fmt.Sprintf("GetemailVcode. email=%s, vcode=%s, err=%#v.", email, vcode, err))
			return ""
		}
		return "error"
	}

	return vcode
}

func SetVerifyVcode(phone string, vcode string) (err error) {
	err = client.Set(verifyVcodeKey+phone, vcode, time.Duration(60)*time.Second).Err()
	if err != nil {
		log.Error(fmt.Sprintf("SetVerifyVcode. phone=%s, vcode=%s,err=%#v.", phone, vcode, err))
		return err
	}

	return err
}

func GetVerifyVcode(phone string) (vcode string) {
	vcode, err := client.Get(verifyVcodeKey + phone).Result()
	if err == redis.Nil {
		log.Error(fmt.Sprintf("GetVerifyVcode. phone=%s, vcode=%s, err=%#v.", phone, vcode, err))
		return ""
	} else if err != nil {
		return "error"
	}

	return vcode
}

func SetPhone(ip string, phone string) (err error) {
	err = client.Set(consumeKey+ip, phone, time.Duration(600)*time.Second).Err()
	if err != nil {
		log.Error(fmt.Sprintf("consumeKey. ip=%s, err=%#v.", ip, err))
		return err
	}
	return nil
}

func GetPhone(ip string) (string, error) {
	phone, err := client.Get(consumeKey + ip).Result()
	if err == redis.Nil {
		log.Error(fmt.Sprintf("consumeKey: ip=%s,  err=%#v.", ip, err))
		return "", RedisNotFound
	} else if err != nil {

		return "", err
	}

	return phone, nil
}

func SetLoginCount(ip string, count int) (err error) {
	err = client.Set(loginKey+ip, count, time.Duration(600)*time.Second).Err()
	if err != nil {
		log.Error(fmt.Sprintf("loginKey: ip=%s, count=%d, err=%#v.", ip, count, err))
		return err
	}

	return nil
}

func GetLoginCount(ip string) (string, error) {
	count, err := client.Get(loginKey + ip).Result()
	if err == redis.Nil {
		log.Error(fmt.Sprintf("loginKey: ip=%s, count=%d, err=%#v.", ip, count, err))
		return "", RedisNotFound
	} else if err != nil {

		return "", err
	}

	return count, nil
}

func LoginConut(ip string) (int, error) {
	countStr, err := GetLoginCount(ip)
	if err != nil {
		log.Error("get count failed. err=" + err.Error())
		if err == RedisNotFound {
			e := SetLoginCount(ip, 1)
			if e != nil {
				log.Error("set count failed. err=" + e.Error())
				return 0, e
			}
			return 1, RedisNotFound
		}
		return 0, err
	}
	count, strerr := strconv.Atoi(countStr)
	if strerr != nil {
		log.Error("str to int  count failed. err=" + strerr.Error())
	}
	if count == 0 {
		return 1, nil
	}

	err = SetLoginCount(ip, count+1)
	if err != nil {
		log.Error("set count failed. err=" + err.Error())
		return 0, err
	}
	return count, nil
}

func SetDetailCount(countId string, count int) error {
	err := client.Set(countKey+countId, count, time.Duration(4)*time.Hour).Err()
	if err != nil {
		log.Error(fmt.Sprintf("countKey: countId=%s, count=%d, err=%#v.", countId, count, err))
		return err
	}
	return nil
}

func GetDetailCount(countId string) (int, error) {
	countStr, err := client.Get(countKey + countId).Result()
	if err == redis.Nil {
		log.Error(fmt.Sprintf("countKey: countId=%s, countStr=%s, err=%#v.", countId, countStr, err))
		return -1, RedisNotFound
	} else if err != nil {

		return -1, err
	}
	count, err := strconv.Atoi(countStr)
	if err != nil {
		log.Error("str to int  count failed. err=" + err.Error())
		return -1, nil
	}
	return count, nil
}

func SetCheckEmailStr(username string, checkStr string) error {
	err := client.Set(checkEmailKey+username, checkStr, time.Duration(24)*time.Hour).Err()
	if err != nil {
		log.Error(fmt.Sprintf("checkEmailKey: username=%s, checkStr=%s, err=%#v.", username, checkStr, err))
		return err
	}
	return nil
}

func GetCheckEmailStr(username string) (string, error) {
	checkStr, err := client.Get(checkEmailKey + username).Result()
	if err == redis.Nil {
		log.Error(fmt.Sprintf("checkEmailKey: username=%s, checkStr=%s, err=%#v.", username, checkStr, err))
		return "", RedisNotFound
	} else if err != nil {

		return "", err
	}
	return checkStr, nil
}

func SetAPIDailyUserTotal(username string, daily_total int) error {
	keySurfix := "_total"
	err := apiDailyClient.Set(username+keySurfix, daily_total, time.Duration(24)*time.Hour).Err()
	if err != nil {
		log.Error(fmt.Sprintf("set api daily user total: username=%s, daily_total=%s, err=%#v.", username, daily_total, err))
		return err
	}
	return nil
}
