package json

import (
	"encoding/json"
	"html"
	"io"

	"github.com/Meghee/kit/datacopy"
)

// DecodeAndEscapeHTML decodes JSON data from io.Reader and escape HTML characters from
// all the string proprties/values in it.
func DecodeAndEscapeHTML(r io.Reader, data interface{}) error {
	var res interface{}
	err := json.NewDecoder(r).Decode(&res)
	if err != nil {
		return err
	}
	switch res.(type) {
	case map[string]interface{}:
		res = parseMap(res.(map[string]interface{}))

	case []interface{}:
		res = parseSlice(res.([]interface{}))
	}
	err = datacopy.Copy(res, data)
	return err
}

func parseMap(data map[string]interface{}) (res map[string]interface{}) {
	res = make(map[string]interface{})
	for k, v := range data {
		switch v.(type) {
		case map[string]interface{}:
			res[k] = parseMap(v.(map[string]interface{}))

		case string:
			res[k] = html.EscapeString(v.(string))

		case []interface{}:
			res[k] = parseSlice(v.([]interface{}))

		default:
			res[k] = v
		}
	}
	return
}

func parseSlice(data []interface{}) (res []interface{}) {
	for _, v := range data {
		switch v.(type) {
		case map[string]interface{}:
			res = append(res, parseMap(v.(map[string]interface{})))

		case string:
			res = append(res, html.EscapeString(v.(string)))

		case []interface{}:
			res = append(res, parseSlice(v.([]interface{})))

		default:
			res = append(res, v)
		}
	}
	return
}
