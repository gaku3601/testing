package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fetchUserData(createUrl(6))
}

// リモートからUser情報を取得する
// [引数]
// id: 指定するID
// [戻値]
func fetchUserData(url string) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%v", string(body))
}

func createUrl(id int) string {
	return fmt.Sprintf("http://fg-69c8cbcd.herokuapp.com/user/%d", id)
}
