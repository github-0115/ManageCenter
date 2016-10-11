package main

import (
	apidetailendpoint "ManageCenter/service/api/api_detail"
	apidetailscreen "ManageCenter/service/api/api_detail_screen"
	apiviewendpoint "ManageCenter/service/api/api_view"
	billing "ManageCenter/service/api/billing/get_billing_order"
	consoledetailendpoint "ManageCenter/service/api/console_detail"
	demodetail "ManageCenter/service/api/demo_detail"
	interfaceconfig "ManageCenter/service/api/interface_config"
	logintoken "ManageCenter/service/api/login_token"
	pornchartendpoint "ManageCenter/service/api/porn_chart"
	realtimemiss "ManageCenter/service/api/real_time_miss"
	realtimephoto "ManageCenter/service/api/real_time_photo"
	realtimeroom "ManageCenter/service/api/real_time_room"
	bill "ManageCenter/service/api/user_manage/finance/bill"
	transfers "ManageCenter/service/api/user_manage/finance/transfers"
	grant "ManageCenter/service/api/user_manage/grant"
	invoice "ManageCenter/service/api/user_manage/invoice"
	notice "ManageCenter/service/api/user_manage/notice"
	packageapply "ManageCenter/service/api/user_manage/package_apply"
	users "ManageCenter/service/api/user_manage/users"
	userservicepackage "ManageCenter/service/api/user_service_package"
	controller "ManageCenter/service/controller"
	middleware "ManageCenter/service/middleware"
	"flag"
	"fmt"
	"time"

	initmanager "ManageCenter/init"

	"github.com/gin-gonic/gin"
	log "github.com/inconshreveable/log15"
	"github.com/itsjamie/gin-cors"
)

var (
	port  = flag.Int64("p", 80, "port")
	debug = flag.Bool("d", false, "debug model")
)

func ready() {

	flag.Parse()
	if *debug == false {
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	ready()
	initmanager.InitManager()
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "POST",
		RequestHeaders:  "Origin, Authorization, Content-Type, Access-Control-Allow-Headers, LoginToken",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false}))

	r.LoadHTMLFiles(
		"templates/manager/login.html",
		"templates/manager/index.html",
		"templates/manager/alluser.html",
		"templates/manager/potentialuser.html",
		"templates/manager/datastat.html",
		"templates/manager/alldata.html",
		"templates/manager/detaildata.html",
		"templates/manager/picreview.html",
		"templates/manager/searchuser.html",
		"templates/manager/apiset.html",
		"templates/footer.html",
		"templates/header.html",
		"templates/script.html",
		"templates/chart.html",
		"templates/nav.html",
		"templates/manager/lookimg.html",
		"templates/manager/demo_looking.html",
		"templates/manager/emailset.html",
		"templates/manager/invoices.html",
		"templates/manager/biling.html",
		"templates/manager/billdetail.html",
		"templates/manager/packageapply.html",
	)
	r.Static("assets", "templates/manager/assets")
	r.Static("layer", "templates/manager/layer")
	r.Static("js", "templates/manager/js")
	r.Static("css", "templates/manager/css")
	r.Static("images", "templates/manager/images")
	r.GET("/", controller.Redirect)
	r.GET("/login", controller.Login)
	r.GET("/index", controller.Index)
	r.GET("/alluser", controller.Alluser)
	r.GET("/potential", controller.Potential)
	r.GET("/datastat", controller.Datastat)
	r.GET("/alldata", controller.Alldata)
	r.GET("/searchuser", controller.Searchuser)
	r.GET("/detail", controller.Detail)
	r.GET("/picreview", controller.Picreview)
	r.GET("/apiset", controller.Apiset)
	r.GET("/biling", controller.Biling)
	r.GET("/invoices", controller.Invoices)
	r.GET("/emailset", controller.Emailset)
	r.POST("/login_token", logintoken.Post)
	r.GET("/lookimg", controller.Lookimg)
	r.GET("/billdetail", controller.Billdetail)
	r.GET("/packageapply", controller.Packageapply)
	r.GET("/manager/demo_looking", controller.DemoLooking)
	authorized := r.Group("/")
	authorized.Use(middleware.AuthToken)
	{
		authorized.GET("/console_detail", consoledetailendpoint.ConsoleDetail)
		authorized.GET("/potential_users", users.PotentialUser)
		authorized.GET("/paging_users", users.GetUsers)
		authorized.POST("/check_users", users.Check)
		authorized.PUT("/userinfo", users.ChangeUserInfo)
		authorized.POST("/add_users", users.AddUser)
		authorized.POST("/notice_email", notice.NoticeEmail)
		authorized.POST("/email_promotion", notice.SendEmail)
		authorized.POST("/add_notice_email", notice.AddNoticeEmail)
		authorized.GET("/get_all_notice_email", notice.AllNoticeEmail)
		authorized.DELETE("/delete_notice_email", notice.DeleteNoticeEmails)
		authorized.PUT("/usergrant", grant.ChangeGrant)
		authorized.GET("/get_usergrant", grant.GetGrant)
		authorized.GET("/invoiceapply", invoice.GetOneInvoiceApply)
		authorized.GET("/user_invoiceapply", invoice.GetUserInvoiceApply)
		authorized.GET("/user_page_invoiceapply", invoice.GetPagenvoiceApply)
		authorized.GET("/all_invoiceapply", invoice.GetAllInvoiceApply)
		authorized.GET("/page_invoiceapply", invoice.GetPagenvoiceApply)
		authorized.PUT("/invoiceapply", invoice.ChangeInvoiceAply)
		authorized.GET("/one_packageapply", packageapply.GetOnePackageApply)
		authorized.GET("/all_packageapply", packageapply.GetAllPackageApply)
		authorized.PUT("/packageapply", packageapply.UpdatePackageApply)
		authorized.DELETE("/packageapply", packageapply.DeletePackageApply)
		authorized.GET("/stat_search_users", users.StatSearchUsers)
		authorized.GET("/get_interface_config", interfaceconfig.GetInterfaceConfig)
		authorized.POST("/set_interface_config", interfaceconfig.SetInterfaceConfig)
		authorized.GET("/get_user_service", userservicepackage.GetUserServicePackage)
		authorized.POST("/set_user_service", userservicepackage.SetUserServicePackage)
		authorized.GET("/demo/detail", demodetail.GetDemoDetail)
	}

	apiAuthorized := r.Group("/api")
	apiAuthorized.Use(middleware.AuthToken)
	{
		apiAuthorized.GET("/porn/chart", pornchartendpoint.GetPornChart)
		apiAuthorized.GET("/porn/timechart", pornchartendpoint.GetTimePornChart)
		apiAuthorized.GET("/view", apiviewendpoint.GetApi)
		apiAuthorized.GET("/user_view", apiviewendpoint.GetUserApi)
		apiAuthorized.GET("/photo_rate", realtimephoto.GetPhotoReviewRate)
		apiAuthorized.GET("/room_rate", realtimeroom.GetRoomReviewRate)
		apiAuthorized.GET("/miss_rate", realtimemiss.GetMissRate)
		apiAuthorized.GET("/porn/detail", apidetailendpoint.GetPornDetail)
		apiAuthorized.GET("/porn/screen_detail", apidetailscreen.GetScreenDetail)
	}

	billAuthorized := r.Group("/bill")
	//billAuthorized.Use(middleware.AuthToken)
	{
		billAuthorized.GET("/paging_bills", billing.GetPagingBills)
		billAuthorized.GET("/transfers", transfers.GetTransfers)
		billAuthorized.PUT("/transfers", transfers.ChangeTransfers)
		billAuthorized.GET("/user_transfers", transfers.GetUserTransfers)
		billAuthorized.GET("/all_transfers", transfers.GetPageTransfers)
		billAuthorized.GET("/one_bills", bill.GetOneBill)
		billAuthorized.GET("/all_bills", bill.GetAllBills)
		billAuthorized.GET("/detail_bills", bill.GetDetailBill)
		billAuthorized.PUT("/favor_bills", bill.UpdateFavorBill)
		billAuthorized.PUT("/transfers_bills", bill.UpdateTransfersBill)
		billAuthorized.POST("/not_pass_bills", bill.NotPassBill)
		billAuthorized.GET("/not_settle", bill.NotSettle)
	}

	if gin.IsDebugging() {
		debugAuthorized := r.Group("/")
		debugAuthorized.Use(middleware.AuthToken)
		{
			debugAuthorized.DELETE("/delete_users", users.Delete)
		}
	}

	p := fmt.Sprintf(":%d", *port)
	log.Info("listen port " + p)

	r.Run(p)
}
