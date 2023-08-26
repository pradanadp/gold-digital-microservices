package helper

import "github.com/teris-io/shortid"

func GenerateId() string {
	sid, _ := shortid.New(1, shortid.DefaultABC, 2342)

	id, err := sid.Generate()
	if err != nil {
		return ""
	}

	return id
}
