package web

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type statistics struct {
	numbers []float64
	mean    float64
	median  float64
}

const form2 = `
	<html><body>
		<form action="/" method="POST">
			<label for="numbers">Numbers (comma or space-separated):</label><br>
			<input type="text" name="numbers" size="30"><br />
			<input type="submit" value="Calculate">
		</form>
	</html></body>
`

const error2 = `<p class="error">%s</p>`

var pageTop = ""
var pageBottom = ""

func WebStatisticsMain() {
	http.HandleFunc("/", homePage)
	if err := http.ListenAndServe(":8888", nil); err != nil {
		panic("failed to star server" + err.Error())
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := r.ParseForm()
	fmt.Fprint(w, pageTop, form2)
	if err != nil {
		fmt.Fprint(w, error2, err)
	} else {
		if numbers, message, ok := processRequest(r); ok {
			stats := getStats(numbers)
			fmt.Fprint(w, formatStats(stats))
		} else if message != "" {
			fmt.Fprint(w, error2, message)
		}
	}
	fmt.Fprint(w, pageBottom)

}

func processRequest(r *http.Request) ([]float64, string, bool) {
	var numbers []float64
	var text string
	if slice, found := r.Form["numbers"]; found && len(slice) > 0 {
		//处理如果网页中输入的是中文逗号
		if strings.Contains(slice[0], "&#65292") {
			text = strings.Replace(slice[0], "&#65292;", " ", -1)
		} else {
			text = strings.Replace(slice[0], ",", " ", -1)
		}
		for _, field := range strings.Fields(text) {
			if x, err := strconv.ParseFloat(field, 64); err != nil {
				return numbers, "'" + field + "' is invaild", false
			} else {
				numbers = append(numbers, x)
			}
		}
	}
	if len(numbers) == 0 {
		return numbers, "", false
	}
	return numbers, "", true
}

func getStats(numbers []float64) (stats statistics) {
	stats.numbers = numbers
	sort.Float64s(stats.numbers)
	stats.mean = sum(numbers) / float64(len(numbers))
	stats.median = median(numbers)
	return
}

func sum(numbers []float64) (total float64) {
	for _, x := range numbers {
		total += x
	}
	return
}

func median(numbers []float64) float64 {
	middle := len(numbers) / 2
	result := numbers[middle]
	if len(numbers)%2 == 0 {
		result = (result + numbers[middle-1]) / 2
	}
	return result
}

func formatStats(stats statistics) string {
	return fmt.Sprintf(`<table border="1">
		<tr><th colspan="2">Results</th></tr>
		<tr><td>Numbers</td><td>%v</td></tr>
		<tr><td>Count</td><td>%d</td></tr>
		<tr><td>Mean</td><td>%f</td></tr>
		<tr><td>Median</td><td>%f</td></tr>
		</table>`, stats.numbers, len(stats.numbers), stats.mean, stats.median)
}
