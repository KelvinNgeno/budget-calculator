package main

import (
	_ "embed"
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

//go:embed index.html
var indexHTML string

var tmpl = template.Must(template.New("index").Parse(indexHTML))

type PageData struct {
	Submitted     bool
	Income        float64
	Rent          float64
	Food          float64
	Transport     float64
	Other         float64
	TotalExpenses float64
	Remaining     float64
	Advice        string
	StatusClass   string
}

func formFloat(r *http.Request, name string) float64 {
	v, _ := strconv.ParseFloat(r.FormValue(name), 64)
	return v
}

func budgetHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	data := PageData{}

	if r.Method == http.MethodPost {
		data.Submitted = true
		data.Income = formFloat(r, "income")
		data.Rent = formFloat(r, "rent")
		data.Food = formFloat(r, "food")
		data.Transport = formFloat(r, "transport")
		data.Other = formFloat(r, "other")

		data.TotalExpenses = math.Round((data.Rent+data.Food+data.Transport+data.Other)*100) / 100
		data.Remaining = math.Round((data.Income-data.TotalExpenses)*100) / 100

		switch {
		case data.Remaining > 0:
			data.StatusClass = "positive"
			data.Advice = fmt.Sprintf("Good job! You can save %.2f KES this month.", data.Remaining)
		case data.Remaining == 0:
			data.StatusClass = "neutral"
			data.Advice = "You broke even. No savings this month."
		default:
			data.StatusClass = "negative"
			data.Advice = fmt.Sprintf("Warning! You are over budget by %.2f KES.", -data.Remaining)
		}
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func openBrowser(url string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}
	_ = cmd.Start()
}

func main() {
	const url = "http://localhost:8080"
	http.HandleFunc("/", budgetHandler)

	fmt.Println("Budget Calculator running at " + url)
	fmt.Println("Keep this window open while you use it. Press Ctrl+C to stop.")

	go func() {
		time.Sleep(500 * time.Millisecond)
		openBrowser(url)
	}()

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
