// 参考: https://qiita.com/Sekky0905/items/fca9d9118ef23bf24791

package main

import (
	"html/template"
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

	http.HandleFunc("/", serverFunc)

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

func serverFunc(w http.ResponseWriter, r *http.Request) {
	var tpl *template.Template = template.Must(template.ParseFiles("pages/homePage.html"))

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}
