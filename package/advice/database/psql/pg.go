package psql

import "askUs/v1/package/advice"

type AdviceData struct {
}

func Init() advice.DB {
	return &AdviceData{}
}
