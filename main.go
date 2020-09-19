package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("*.tmpl")

	r.GET("/", homeGETFunc)
	r.GET("/want", wantGETFunc)
	r.GET("/tt2020", tt2020GETFunc)

	r.Run(":8080")
}

// homeGETFunc は/にGETされたときの処理
func homeGETFunc(c *gin.Context) {
	c.HTML(200, "index.tmpl", nil)
}

// wantGETFunc は/wantにGETされたときの処理
func wantGETFunc(c *gin.Context) {
	c.Redirect(301, "https://www.amazon.co.jp/hz/wishlist/ls/3V56KMUKW9OT3")
}

// tt2020GETFunc は/tt2020にGETされたときの処理
func tt2020GETFunc(c *gin.Context) {
	c.Redirect(301, "https://twitter.com/takara2314/status/1285527855284105216")
}
