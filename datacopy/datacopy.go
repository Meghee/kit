package datacopy

import "encoding/json"

// Copy copies data from one variable type to another.
//
// for example we might want to copy data from a
//
// map[string]interface{}{
// 	"username": "wisdommatt",
// 	"phone": "080838484855",
// } type
//
// to a
//
// struct{
// 	Username string `json:"username"`
// 	Phone string `json:"phone"`
// 	Password string `json:"password"`
// }{}
//
// it the case above it copies only the username and phone data from the map[string]interface{}
// type without touching the password.
//
// Note: this function encodes data1 to json and decode the json value of data1 into data2.
func Copy(data1, data2 interface{}) error {
	jsonData, err := json.Marshal(data1)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonData, data2)
	return err
}
