package main

import(
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	// Week 03: ここに課題のコードを記述してください
	// 詳細な課題内容はLMSで確認してください
	fmt.Println("Week 03 課題")
	// 以下に実装してください
	http.HandleFunc("/webfortune", webfortune)
	http.ListenAndServe(":8080", nil)
}

func webfortune(w http.ResponseWriter, r *http.Request) {
	seed := time.Now().UnixNano()
	d := rand.New(rand.NewSource(seed))
	var omikuji=[4]string{"大吉","中吉","小吉","凶"}
	fmt.Fprintln(w,"今の運勢は", omikuji[d.Int31n(5)],"です")
}

