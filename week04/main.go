package main
import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Week 04: ここに課題のコードを記述してください
	// 詳細な課題内容はLMSで確認してください
	fmt.Println("Week 04 課題")
	// 以下に実装してください
	http.HandleFunc("/info", info) //追加

	http.ListenAndServe(":8080", nil)
}

// 一部ヘッダを取得する例．追加
func info(w http.ResponseWriter, r *http.Request) {
	var h map[string][]string
	h = r.Header
	jst, _ := time.LoadLocation("Asia/Tokyo")
	fmt.Fprint(w, "現在の時刻は", (time.Now().In(jst)).Format("15:04"),"で、")
	fmt.Fprintln(w, "利用しているブラウザは", h["User-Agent"][0], "ですね")
}
