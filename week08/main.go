package main

import (
	"fmt"
	"math"
	"net/http"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	// Week 08: ここに課題のコードを記述してください
	// 詳細な課題内容はLMSで確認してください

	fmt.Println("Week 08 課題")

	// 以下に実装してください
	fmt.Printf("Go version: %s\n", runtime.Version())
	http.Handle("/", http.FileServer(http.Dir("public/")))
	http.HandleFunc("/average", avehandler)

	fmt.Println("Launch server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}
}
func avehandler(w http.ResponseWriter, r *http.Request) {
	var sum, tt int
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	tokuten := strings.Split(r.FormValue("dd"), ",")
	fmt.Println(tokuten)
	tokuten2 := []int{}
	for i := range tokuten {
		tt, _ = strconv.Atoi(tokuten[i])
		tokuten2 = append(tokuten2, tt)
		sum += tt
	}
	average := math.Round((float64(sum) / float64(len(tokuten)))*10)/10
	fmt.Fprintln(w, "平均点は", average, "点")
	var number [10]int
	var i int
	for i = 0; i < len(tokuten); i++ {
		if tokuten2[i] >= 90 {
			number[0] = number[0] + 1
		} else if tokuten2[i] >= 80 {
			number[1] = number[1] + 1
		} else if tokuten2[i] >= 70 {
			number[2] = number[2] + 1
		} else if tokuten2[i] >= 60 {
			number[3] = number[3] + 1
		} else if tokuten2[i] >= 50 {
			number[4] = number[4] + 1
		} else if tokuten2[i] >= 40 {
			number[5] = number[5] + 1
		} else if tokuten2[i] >= 30 {
			number[6] = number[6] + 1
		} else if tokuten2[i] >= 20 {
			number[7] = number[7] + 1
		} else if tokuten2[i] >= 10 {
			number[8] = number[8] + 1
		} else if tokuten2[i] >= 0 {
			number[9] = number[9] + 1
		}
	}
	var j int
	for j = 0; j < len(number); j++ {
		fmt.Fprintln(w, 90-j*10, "点以上は，", number[j], "人です．")
	}
	fmt.Print(sum)
}
