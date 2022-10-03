package main

import (
	"log"
	"net/http"
)

func main() {
	// クライアントとリクエストの作成
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, "https://google.com", nil)
	if err != nil {
		log.Fatal(err)
	}

	// ヘッダの設定
	req.Header.Add("test", `no-meaning`)

	// クライアントを使ってリクエストの送信
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}
