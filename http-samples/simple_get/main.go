package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

// google.comにリクエストを送信しレスポンスのボディを取得
func main() {
	// リクエストの送信処理
	resp, err := http.Get("http://google.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// ボディの読み取り
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile("./output.html", []byte(body), os.ModeAppend)
}
