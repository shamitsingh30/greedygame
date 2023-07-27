package validation

import (
	"fmt"
	"strconv"
	"strings"
)

func ValidateFunc(command *string) (map[string]interface{}, bool) {
	words := strings.Split(*command, " ")
	length := len(words)

	newReqBody := make(map[string]interface{})
	if length < 2 {
		return newReqBody, false
	}

	newReqBody["querytype"] = words[0]

	var err interface{} = nil
	if length >= 3 && words[0] == "SET" {
		newReqBody["key"] = words[1]
		newReqBody["value"] = words[2]
		if length == 3 {

		} else if length == 4 && (words[3] == "XX" || words[3] == "NX") {
			newReqBody["condition"] = words[3]
		} else if length == 5 && words[3] == "EX" {
			newReqBody["expiry_time"], err = strconv.ParseInt(words[4], 10, 32)
			if err != nil {
				return newReqBody, false
			}
		} else if length == 6 && (words[5] == "XX" || words[5] == "NX") {
			newReqBody["expiry_time"], err = strconv.ParseInt(words[4], 10, 32)
			if err != nil {
				return newReqBody, false
			}
			newReqBody["condition"] = words[5]
		} else {
			return newReqBody, false
		}
	}

	fmt.Println("Validation checkpoint", words)
	return newReqBody, true
}
