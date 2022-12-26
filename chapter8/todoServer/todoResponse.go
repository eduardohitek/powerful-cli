package main

import (
	"encoding/json"
	"time"

	"github.com/eduardohitek/powerful-cli/chapter2/todo"
)

type todoResponse struct {
	Results todo.List `json:"results"`
}

func (r *todoResponse) MarshalJSON() ([]byte, error) {
	resp := struct {
		Results      todo.List `json:"results,omitempty" bson:"results"`
		Date         int64     `json:"date,omitempty" bson:"date"`
		TotalResults int       `json:"total_results,omitempty" bson:"total_results"`
	}{
		Results:      r.Results,
		Date:         time.Now().Unix(),
		TotalResults: len(r.Results),
	}
	return json.Marshal(resp)
}
