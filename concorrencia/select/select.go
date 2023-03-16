package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

func titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return c
}

func oMaisRapido(url1, url2, url3 string) string {
	c1 := titulo(url1)
	c2 := titulo(url2)
	c3 := titulo(url3)

	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam!"
		// default:
		//	return "Sem resposta ainda!"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://www.cod3r.com.br",
		"https://www.google.com",
		"https://www.youtube.com",
	)
	fmt.Println(campeao)
}
