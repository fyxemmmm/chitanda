package chitanda

import (
	"encoding/json"
	"log"
)

type Model interface {
	String() string
}

type Models string

func ToModels(v interface{}) Models {
	b, err := json.Marshal(v)
	if err != nil {
		log.Println(err)
	}
	return Models(b)
}
