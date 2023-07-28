package validation

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func ValidateFunc(command *string) (map[string]string, error) {
	words := strings.Split(*command, " ")
	length := len(words)

	newReqBody := make(map[string]string)
	errorText := errors.New("invalid command")
	if length < 2 {
		return newReqBody, errorText
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
			_, err = strconv.Atoi(words[4])
			if err != nil {
				return newReqBody, errorText
			}
			newReqBody["expiry_time"] = words[4]
		} else if length == 6 && (words[5] == "XX" || words[5] == "NX") {
			_, err = strconv.Atoi(words[4])
			if err != nil {
				return newReqBody, errorText
			}
			newReqBody["expiry_time"] = words[4]
			newReqBody["condition"] = words[5]
		} else {
			return newReqBody, errorText
		}
	} else if length == 2 && words[0] == "GET" {
		newReqBody["querytype"] = "GET"
		newReqBody["key"] = words[1]
	}

	fmt.Println("Validation checkpoint", words)
	return newReqBody, nil
}
