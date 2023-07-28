package validation

import (
	"errors"
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
	newReqBody["key"] = words[1]

	var err interface{} = nil
	if words[0] == "SET" && length >= 3 {
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
	} else if words[0] == "GET" && length == 2 {

	} else if words[0] == "QPUSH" && length >= 3 {
		newReqBody["items"] = words[2]
		for _, el := range words[3:] {
			newReqBody["items"] = newReqBody["items"] + " " + el
		}
	} else if words[0] == "QPOP" && length == 2 {

	} else if words[0] == "BQPOP" && length == 3 {

	} else {
		return newReqBody, errorText
	}

	return newReqBody, nil
}
