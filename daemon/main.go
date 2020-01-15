package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	counter := SafeCounter{v: make(map[string]int)}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		helloHandler(w, r, counter)
	})
	http.ListenAndServe(":8080", nil)
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

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func helloHandler(w http.ResponseWriter, req *http.Request, counter SafeCounter) {
	switch req.Method {
	case http.MethodGet:
		user := User{Name: "Taro", Age: counter.Value("age")}
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
		counter.Inc("age")
		user.Age = counter.Value("age")
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
