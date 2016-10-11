package smsclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	log "github.com/inconshreveable/log15"
)

const rawUrl = "http://api.synhey.com/json?key=dqvtnz&secret=4wwXBZ6b&"

var (
	checksuccessTemplate = "尊敬的 %s，您的账户已经通过审核，您可以登录开发者平台使用密钥进行开发接入了。"
	billnotpassTemplate  = "尊敬的 %s，我们无法核实您的结算信息，请确保填写的信息正确。"
	packageapplyTemplate = "尊敬的 %s，您的账户已经成功变更成 %s，有任何问题，请与我们联系：bd@deepir.com"
)

func SendVerifySMS(phone string, username string) string {
	var msg = fmt.Sprintf(checksuccessTemplate, username)
	log.Info("msg = " + msg)

	smsUrl := rawUrl + "to=" + phone + "&text=" + url.QueryEscape(msg)
	log.Info("smsUrl = " + smsUrl)
	log.Info("after escape, smsUrl = " + smsUrl)

	resp, err := http.Get(smsUrl)
	if err != nil {
		log.Error("http get failed." + err.Error())
		return err.Error()
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("resp read error." + err.Error())
		return err.Error()
	}

	log.Info(string(body))
	return ""
}

func SendBillSMS(phone string, username string) string {
	var msg = fmt.Sprintf(billnotpassTemplate, username)
	log.Info("msg = " + msg)

	smsUrl := rawUrl + "to=" + phone + "&text=" + url.QueryEscape(msg)
	log.Info("smsUrl = " + smsUrl)
	log.Info("after escape, smsUrl = " + smsUrl)

	resp, err := http.Get(smsUrl)
	if err != nil {
		log.Error("http get failed." + err.Error())
		return err.Error()
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("resp read error." + err.Error())
		return err.Error()
	}

	log.Info(string(body))
	return ""
}

func SendPackageSMS(phone string, username string, utype string) string {
	var msg = fmt.Sprintf(packageapplyTemplate, username, utype)
	log.Info("msg = " + msg)

	smsUrl := rawUrl + "to=" + phone + "&text=" + url.QueryEscape(msg)
	log.Info("smsUrl = " + smsUrl)
	log.Info("after escape, smsUrl = " + smsUrl)

	resp, err := http.Get(smsUrl)
	if err != nil {
		log.Error("http get failed." + err.Error())
		return err.Error()
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("resp read error." + err.Error())
		return err.Error()
	}

	log.Info(string(body))
	return ""
}
