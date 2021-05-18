package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			resp, err := http.Get("http://127.0.0.1:8080/cartApi/findAll?user_id=3")
			bytes, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				fmt.Println("ioutil.ReadAll err=",err)
				return
			}
			fmt.Println(string(bytes))
			wg.Done()
		}()
	}
	wg.Wait()
}
