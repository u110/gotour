package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/u110/gotour/daemon/counter"
	"io"
	"log"
	"net/http"
	"strconv"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var cnt = counter.SafeCounter{V: make(map[string]int)}

func main() {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		user := User{Name: "Taro", Age: cnt.Value("age")}
		return c.JSON(http.StatusCreated, user)
	})

	e.POST("/", func(c echo.Context) error {
		user := new(User)
		if err := c.Bind(user); err != nil {
			return err
		}
		cnt.Inc("age")
		user.Age = cnt.Value("age")
		return c.JSON(http.StatusCreated, user)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func handleParams(w http.ResponseWriter, r *http.Request) {

	// クエリパラメータ取得してみる
	fmt.Fprintf(w, "クエリ：%s\n", r.URL.RawQuery)

	// Bodyデータを扱う場合には、事前にパースを行う
	// r.ParseForm()
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	// Formデータを取得.
	form := r.PostForm
	fmt.Fprintf(w, "フォーム：\n%v\n", form)

	// または、クエリパラメータも含めて全部.
	params := r.Form
	fmt.Fprintf(w, "フォーム2：\n%v\n", params)

	fmt.Fprintf(w, "フォーム3：\n%v\n", r.FormValue("Age"))

	form = r.PostForm
	fmt.Fprintf(w, "Age：\n%v\n", form)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		user := User{Name: "Taro", Age: cnt.Value("age")}
		res, err := json.Marshal(user)
		if err != nil {
			log.Print("😇json marshal error.")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	case http.MethodPost:
		if req.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusBadRequest)
			return
		} else {
			log.Print("🙃application/json")
		}

		length, err := strconv.Atoi(req.Header.Get("Content-Length"))
		if err != nil {
			log.Print("😇Content-LengthがnilだからStatusInternalServerError")
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			log.Print("🙃lengthは", length, "バイトです。")
		}

		var user User
		buffer := make([]byte, length)

		_, err = req.Body.Read(buffer)

		if err != nil && err != io.EOF {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(buffer, &user)
		cnt.Inc("age")
		user.Age = cnt.Value("age")
		if err != nil {
			log.Print("😇Json Unmarshal Error ")
		}

		fmt.Fprintf(w, "%#v\n", user)
		w.WriteHeader(http.StatusOK)
	default:
		fmt.Fprintf(w, "Default ")
		w.WriteHeader(http.StatusOK)
	}
}
