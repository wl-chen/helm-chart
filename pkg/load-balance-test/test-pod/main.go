package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	res := make(map[string]int32)

	for i := 0; i < 1000; i++ {
		response, err := http.Get("http://www.baidu.com/pod/ip")
		if err != nil {
			fmt.Printf("get error: %v %s", err, err.Error())
			return
		}
		if response == nil || response.Body == nil {
			fmt.Printf("reponse %v nil", response)
		}
		defer response.Body.Close()
		body, err2 := ioutil.ReadAll(response.Body)
		if err2 != nil {
			fmt.Printf("ioutil read error %v %s", err2, err2.Error())
			return
		}
		// fmt.Printf("%s", string(body))

		uuid := string(body)
		res[uuid] = res[uuid] + 1
	}

	for pod, count := range res {
		fmt.Printf("visit pod %s %d times in %d sum", pod, count, 1000)
		fmt.Println()
	}
}