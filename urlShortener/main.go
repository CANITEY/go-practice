package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

type Url map[string]string

func (u Url) addUrl(url string) (string) {
	address := RandStringRunes(4)
	u[address] = url
	return address
}

func (u Url) getUrl(address string) string {
	data, ok := u[address]
	if !ok {
		return "Not Found"
	}

	return data
}

func (u Url) ShowUrls() {
	for key, data := range u {
		fmt.Printf("http://127.0.0.1:8888/%v     %v\n", key, data)
	}
}




func main() {
	u := make(Url)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := strings.Split(r.URL.Path, "/")[1]
		res := u.getUrl(url)
		if res == "Not Found" {
			fmt.Fprintf(w, "%v", res)
		} else {
			fmt.Fprintf(w, "<script>document.location='%v'</script>", res)
		}
	})
	go http.ListenAndServe(":8888", nil)

	for {
		var address string
		fmt.Print("Enter url to be shortened >> ")
		fmt.Scanln(&address)
		if address == "000" {
			u.ShowUrls()
		} else {
			shortened := u.addUrl(address)
			fmt.Printf("http://127.0.0.1:8888/%v\n",shortened)
		}

	}
}
