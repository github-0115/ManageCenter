package config

import (
	_ "ManageCenter/service/log"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	log "github.com/inconshreveable/log15"
)

var (
	Cfg   Config
	DBCfg DBConfig
)

type Config struct {
	LoginSecret      string   `json:"login_secret"`
	ApiSecret        string   `json:"api_secret"`
	LoginTokenExpire int      `json:"login_token_expire"`
	ManagerUsername  string   `json:"manager_username"`
	ManagerPassword  string   `json:"manager_password"`
	ApiTokenExpire   int      `json:"api_token_expire"`
	LogDir           string   `json:"log_dir"`
	Domain           string   `json:"domain"`
	Mail             *MailCfg `json:"mail"`
}

type MailCfg struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

type DBConfig struct {
	RedisBroker           *RedisBrokerCfg `json:"redis"`
	UserCenterMongoTask   *MongoCfg       `json:"uc_mongo_task"`
	ManageCenterMongoTask *MongoCfg       `json:"mc_mongo_task"`
	StatsMongo            *MongoCfg       `json:"stats_mongo"`
	DemoMongo             *MongoCfg       `json:"demo_mongo"`
	APIDailyRedisBroker   *RedisBrokerCfg `json:"api_daily_redis"`
}

type RedisBrokerCfg struct {
	Host     string `json:"host"`
	DB       int64  `json:"db"`
	Port     int    `json:"port"`
	Password string `json:"password"`
}

type MongoCfg struct {
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	DB       string `json:"db"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func (s *MongoCfg) String() string {
	pwd, user := "", ""
	if s.User != "" && s.Password != "" {
		pwd = s.Password + "@"
		user = s.User + ":"
	}
	return fmt.Sprintf("mongodb://%s%s%s:%d/%s", user, pwd, s.Host, s.Port, s.DB)
}

func (s *RedisBrokerCfg) String() string {
	pwd, user := "", ""
	if s.Password != "" {
		pwd = s.Password + "@"
	}
	return fmt.Sprintf("redis://%s%s%s:%d/%s", user, pwd, s.Host, s.Port, s.DB)
}

func init() {
	log.Info("init config files")

	readCfg()

	log.Info("init sys config finish", log.Ctx{
		"cfg": Cfg,
	})

	readDBCfg()
	log.Info("init db config finish", log.Ctx{
		"DB config": DBCfg,
	})
}

func readCfg() {
	file, e := ioutil.ReadFile("./config.json")
	if e != nil {
		log.Error(fmt.Sprintf("read config error, e=%#v", e))
		os.Exit(1)
	}
	err := json.Unmarshal(file, &Cfg)
	if err != nil {
		log.Error(fmt.Sprintf("config not json format, e=%#v", err))
		panic(e)
	}
}

func readDBCfg() {
	file, e := ioutil.ReadFile("./db_config.json")
	if e != nil {
		log.Error(fmt.Sprintf("read db config error, e=%#v", e))
		panic(e)
	}
	err := json.Unmarshal(file, &DBCfg)
	if err != nil {
		log.Error(fmt.Sprintf("db config not json format, e=%#v", err))
		panic(e)
	}
}
