package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// HTMLテンプレートの読み込み (同じディレクトリで実行する扱い)
	r.LoadHTMLGlob("./dist/*.html")
	// 静的ファイルの読み込み
	r.Static("/dist", "./dist")
	r.Static("/resources", "./resources")

	r.GET("/", homeGETFunc)
	r.GET("/about", aboutGETFunc)
	r.GET("/skills", skillsGETFunc)
	r.GET("/works", worksGETFunc)
	r.GET("/history", historyGETFunc)
	r.GET("/favorites", favoritesGETFunc)
	r.GET("/lab", labGETFunc)
	r.GET("/links", linksGETFunc)
	r.GET("/contact", contactGETFunc)
	r.GET("/want", wantGETFunc)
	r.GET("/tt2020", tt2020GETFunc)

	r.Run(":8080")
}

// homeGETFunc → / [GET]
func homeGETFunc(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

// aboutGETFunc → /about [GET]
func aboutGETFunc(c *gin.Context) {
	c.HTML(200, "about.html", nil)
}

// skillsGETFunc → /skills [GET]
func skillsGETFunc(c *gin.Context) {
	c.HTML(200, "skills.html", nil)
}

// worksGETFunc → /works [GET]
func worksGETFunc(c *gin.Context) {
	c.HTML(200, "works.html", nil)
}

// historyGETFunc → /history [GET]
func historyGETFunc(c *gin.Context) {
	c.HTML(200, "history.html", nil)
}

// favoritesGETFunc → /favorites [GET]
func favoritesGETFunc(c *gin.Context) {
	c.HTML(200, "favorites.html", nil)
}

// labGETFunc → /lab [GET]
func labGETFunc(c *gin.Context) {
	c.HTML(200, "lab.html", nil)
}

// linksGETFunc → /links [GET]
func linksGETFunc(c *gin.Context) {
	c.HTML(200, "links.html", nil)
}

// contactGETFunc → /contact [GET]
func contactGETFunc(c *gin.Context) {
	c.HTML(200, "contact.html", nil)
}

// wantGETFunc → /want [GET]
func wantGETFunc(c *gin.Context) {
	c.Redirect(301, "https://www.amazon.co.jp/hz/wishlist/ls/3V56KMUKW9OT3")
}

// tt2020GETFunc → /tt2020 [GET]
func tt2020GETFunc(c *gin.Context) {
	c.Redirect(301, "https://twitter.com/takara2314/status/1285527855284105216")
}
