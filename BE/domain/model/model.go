package model

import "time"

type Refl struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type Thread struct {
	Date  string `json:"date"`
	Refls []Refl `json:"refls"`
}

type RequestGetThread struct {
	Since time.Time `json:"since"`
	Until time.Time `json:"until"`
}

type ResponseGetThread struct {
	Data []Thread `json:"data"`
}
