package models

import "log"

func ShowErrorInLog(err error) {
	log.Println("something wrong", err)
	// but we can return json
	// msg := models.Message{ Message: "something wrong")
	// json.NewEncoder(writer).Encode(msg)
}
