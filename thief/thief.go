package thief

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func CheckJsonFilter(messages []byte) (bool, error) {
	var request map[string]interface{}
	flag := false
	err := json.Unmarshal(messages, &request)
	if err != nil {
		fmt.Printf("Ошибка при десериализации: %v", err)
		return false, err
	}
	r := reflect.TypeOf(tP)
	for k := range request {
		for i := 0; i < r.NumField(); i++ {
			if k == r.Field(i).Name {
				flag = true
				return flag, nil
			}
		}
	}
	return flag, err
}

type ProblemMessage struct {
	is_cyber      string `json:is_cyber`
	warning_type  string `json:warning_type`
	warning_start string `json:warning_start`
}

var tP ProblemMessage

var jsonMessage map[string]interface{}
