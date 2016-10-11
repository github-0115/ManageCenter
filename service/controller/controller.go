package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Redirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/index")
}

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", "")
}

func Alluser(c *gin.Context) {
	c.HTML(http.StatusOK, "/alluser", "")
}
func Potential(c *gin.Context) {
	c.HTML(http.StatusOK, "/potential", "")
}
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "/index", "")
}
func Datastat(c *gin.Context) {
	c.HTML(http.StatusOK, "/datastat", "")
}
func Alldata(c *gin.Context) {
	c.HTML(http.StatusOK, "/alldata", "")
}
func Detail(c *gin.Context) {
	c.HTML(http.StatusOK, "/detail", "")
}
func Picreview(c *gin.Context) {
	c.HTML(http.StatusOK, "/picreview", "")
}

func Searchuser(c *gin.Context) {
	c.HTML(http.StatusOK, "/searchuser", "")
}
func Apiset(c *gin.Context) {

	c.HTML(http.StatusOK, "/apiset", "")
}

func Lookimg(c *gin.Context) {
	c.HTML(http.StatusOK, "/lookimg", "")
}

func DemoLooking(c *gin.Context) {
	c.HTML(http.StatusOK, "manager/demo_looking", "")
}
func Emailset(c *gin.Context) {
	c.HTML(http.StatusOK, "/emailset", "")
}
func Invoices(c *gin.Context) {
	c.HTML(http.StatusOK, "/invoices", "")
}

func Biling(c *gin.Context) {
	c.HTML(http.StatusOK, "/biling", "")
}
func Billdetail(c *gin.Context) {
	c.HTML(http.StatusOK, "/billdetail", "")
}
func Packageapply(c *gin.Context) {
	c.HTML(http.StatusOK, "/packageapply", "")
}

