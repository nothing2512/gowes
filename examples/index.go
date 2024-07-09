package examples

import (
	"fmt"
	"net/http"

	"github.com/nothing2512/gowes/server"
)

func runServer() {
	s := server.Init("11111111113333333333332222222222", "1112223334445556")
	s.OnCommand(func(m server.Message) {
		fmt.Println(m.Command, m.Message)
	})
	s.Start("0.0.0.0:8080")
}

func runHttp() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "examples/index.html")
	})
	fmt.Println("Listening HTTP On 0.0.0.0:3333")
	if err := http.ListenAndServe("0.0.0.0:3333", nil); err != nil {
		panic(err)
	}
}

func Run() {
	go runServer()
	runHttp()
}
