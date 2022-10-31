package model

import "time"

type Reflection struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type DaysReflections struct {
	Date        string       `json:"date"`
	Reflections []Reflection `json:"reflections"`
}

type RequestGetReflection struct {
	Since time.Time `json:"since"`
	Until time.Time `json:"until"`
}

type ResponseGetReflection struct {
	Data []DaysReflections `json:"data"`
}
