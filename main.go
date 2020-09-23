package main

import (
	bingwallpaper "bing-wallpaper-jsonp/gin_handelr"

	"github.com/gin-gonic/gin"
)

const (
	//gin response json data
	serverErrorMessage string = "Server internal error, please contact the administrator zzerding@foxmail.com"
	//url
	urlBingServer string = "https://cn.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1"
)

//response
func ginResponse(c *gin.Context) {
	data, err := bingwallpaper.GetCache(urlBingServer)
	if err != nil {
		c.JSON(500, gin.H{"msg": serverErrorMessage})
	}
	c.JSONP(200, data)
}

func main() {
	r := gin.Default()

	r.GET("/bing-wallpaper-jsonp", ginResponse)
	r.Run(":80")
}
