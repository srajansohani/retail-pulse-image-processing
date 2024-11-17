package models

import "sync"

type Status string

const (
	Pending   Status = "pending"
	Ongoing   Status = "ongoing"
	Completed Status = "completed"
	Failed    Status = "failed"
)

type Job struct {
	ID     int        `json:"id"`
	Status Status     `json:"status`
	Visits []Visit    `json:"visits"`
	Errors []JobError `json:"errors,omitempty"`
}

type Visit struct {
	StoreID   string   `json:"store_id"`
	ImageURLs []string `json:"image_url"`
	VisitTime string   `json:"visit_time"`
}

type JobError struct {
	StoreID string `json:"store_id"`
	Error   string `json:"error"`
}

var JobStore = sync.Map{}
var JobIDCounter = 1
