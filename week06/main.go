package main
import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)

func main() {
	// Week 06: ここに課題のコードを記述してください
	// 詳細な課題内容はLMSで確認してください
	fmt.Println("Week 06 課題")
	// 以下に実装してください
	http.Handle("/", http.FileServer(http.Dir("public/")))
	http.HandleFunc("/bmi", bmicaluculate)
	fmt.Println("Launch server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}
}
func bmicaluculate(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	w.Header().Set("Content-Type", "text/plain; 		charset=utf-8")
	tall, _ := strconv.ParseFloat(r.FormValue("tall"), 64)
	weight, _ := strconv.ParseFloat(r.FormValue("weight"), 64)
	tallm := tall / 100.0
	bmi := math.Round((weight/(tallm*tallm))*10) / 10
	fmt.Fprint(w, "BMIは ")
	fmt.Fprintln(w, bmi)
}
