package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	urlBaiduFanyiSug := "https://fanyi.baidu.com/sug"
	data := map[string]string{"kw": "hi"}
	js, _ := json.Marshal(data)
	res, err := http.Post(urlBaiduFanyiSug, "application/json", bytes.NewBuffer(js))
	if err != nil {
		fmt.Printf("send request failed, err=%s", err)
	} else {
		rjs, _ := ioutil.ReadAll(res.Body)
		fmt.Printf("%s\n", string(rjs))
	}
}
