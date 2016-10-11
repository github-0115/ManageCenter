package main

import (
	usermodel "ManageCenter/service/model/usermodel"

	servpkgmodel "ManageCenter/service/model/servicepackagemodel"
	redisclient "ManageCenter/utils/redisclient"
	"fmt"
	"time"

	log "github.com/inconshreveable/log15"
)

func setAPIDaily(results []*servpkgmodel.UserServicePackage) {
	yeardayOfNow := time.Now().YearDay()
	for _, userServicePackage := range results {
		username := userServicePackage.Username
		dailyAmount := int(userServicePackage.DailyAmount)
		user, err := usermodel.QueryUser(username)
		if err != nil {
			log.Error(fmt.Sprintf("user: %v not exist", username))
		}
		userType := user.Type
		log.Info(fmt.Sprintf("username: %v, dailyAmout: %v, userType: %v", username, dailyAmount, userType))

		if userType == 0 {
			// 免费用户
			createdAt := userServicePackage.CreatedAt
			yeardayOfCreated := createdAt.YearDay()
			log.Info(fmt.Sprintf("created_at: %v", createdAt))

			if yeardayOfNow-yeardayOfCreated >= 30 {
				err := redisclient.SetAPIDailyUserTotal(username, 0)
				log.Info(fmt.Sprintf("set api daily user total 0, yearday_of_now: %v, yearday_of_created: %v", yeardayOfNow, yeardayOfCreated))

				if err != nil {
					log.Error(fmt.Sprintf("set api daily user total err:%v, username:%v ", err, username))
				}
				continue
			}
		}
		err = redisclient.SetAPIDailyUserTotal(username, dailyAmount)
		log.Info(fmt.Sprintf("set api daily user total redis, username: %v, dailyAmout: %v", username, dailyAmount))

		if err != nil {
			log.Error(fmt.Sprintf("set api daily user total err: %v ", err))
		}

	}
}

func main() {
	//	fmt.Println("hello world")
	log.Info(fmt.Sprintf("daily api user total starts at %v", time.Now().String()))
	results, err := servpkgmodel.GetAllUserServicePackages()
	log.Info(fmt.Sprintf("results: %v", results))

	if err != nil {
		log.Error(fmt.Sprintf("get all user service package err: %v ", err))

		return
	}

	setAPIDaily(results)
}
