// 参考: https://qiita.com/Sekky0905/items/fca9d9118ef23bf24791

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// 静的なフォルダの読み込み
	for _, folder := range staticFolder {
		http.Handle("/"+folder+"/",
			http.StripPrefix("/"+folder+"/",
				http.FileServer(http.Dir(folder+"/"))))
	}

	http.HandleFunc("/", serverFuncMain)
	http.HandleFunc("/tt2020", serverFuncTT2020)
	http.HandleFunc("/want", serverFuncWantList)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("ポートが設定されていないため、ポート番号%sで開きました。\n", port)
	} else {
		log.Printf("ポート番号%s開きました。\n", port)
	}

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func serverFuncMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "まだ準備中だよ")

	// var tpl *template.Template = template.Must(template.ParseFiles("pages/homePage.html"))

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// err := tpl.Execute(w, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func serverFuncTT2020(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://twitter.com/takara2314/status/1285527855284105216", 301)
}

func serverFuncWantList(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://www.amazon.co.jp/hz/wishlist/ls/3V56KMUKW9OT3", 301)
}
