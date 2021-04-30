package templatefuncs

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var FuncMap = template.FuncMap{
	"sliceLen": func(data interface{}) int {
		dataJSON, _ := json.Marshal(data)
		var newData []interface{}
		json.Unmarshal(dataJSON, &newData)
		return len(newData)
	},
	"addInts": func(ints ...int) int {
		total := int(0)
		for _, v := range ints {
			total += v
		}
		return total
	},
	"multiplyInts": func(ints ...int) int {
		total := int(0)
		for _, v := range ints {
			total *= v
		}
		return total
	},
	"roundFloat": func(num float64) float64 {
		return math.Ceil(num)
	},
	"numberFormat": func(num float64) string {
		str := fmt.Sprintf("%.0f", num)
		re := regexp.MustCompile("(\\d+)(\\d{3})")
		for n := ""; n != str; {
			n = str
			str = re.ReplaceAllString(str, "$1,$2")
		}
		return str
	},
	"dateFormat": func(dateStr string) string {
		date, err := time.Parse(time.RFC3339, dateStr)
		if err != nil {
			return ""
		}
		yy, mm, dd := date.Date()
		return strconv.Itoa(yy) + " - " + mm.String() + " - " + strconv.Itoa(dd)
	},
	"jsonEncode": func(data interface{}) string {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return ""
		}
		return strings.ReplaceAll(string(jsonData), "\"", "'")
	},
	"StrToUpper":  strings.ToUpper,
	"StrToLower":  strings.ToLower,
	"StrContains": strings.Contains,
	"StrIndex":    strings.Index,
}
