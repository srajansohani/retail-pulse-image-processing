package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/srajansohani/image-process-service/models"
)

var mu sync.Mutex

func SubmitJob(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Count  int            `json:"count"`
		Visits []models.Visit `json:"visits"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, `{"error":"Invalid payload"}`, http.StatusBadRequest)
		return
	}

	if payload.Count != len(payload.Visits) {
		http.Error(w, `{"error":"mentioned counts does not match with no of visits"}`, http.StatusBadRequest)
		return
	}

	var job models.Job
	job.Visits = payload.Visits

	mu.Lock()
	job.ID = models.JobIDCounter
	models.JobIDCounter++
	job.Status = "ongoing"
	models.JobStore.Store(job.ID, job)
	mu.Unlock()

	go processJob(job) //so that the processing occurs in background

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"job_id": job.ID})
}

func GetJobInfo(w http.ResponseWriter, r *http.Request) {
	jobIDStr := r.URL.Query().Get("jobid")
	if jobIDStr == "" {
		http.Error(w, `{"error":"Missing jobid"}`, http.StatusBadRequest)
		return
	}
	jobID, err := strconv.Atoi(jobIDStr)
	if err != nil {
		http.Error(w, `{"error":"Invalid jobid"}`, http.StatusBadRequest)
		return
	}
	job, ok := models.JobStore.Load(jobID)
	if !ok {
		http.Error(w, `{"error":"Job not found"}`, http.StatusBadRequest)
		return
	}
	jobData, ok := job.(models.Job)
	if !ok {
		http.Error(w, `{"error":"Job not found"}`, http.StatusBadRequest)
		return
	}

	// fmt.Println(jobData)

	var response map[string]interface{}

	if jobData.Status == "failed" {
		response = map[string]interface{}{
			"status": "failed",
			"job_id": jobData.ID,
			"error":  jobData.Errors, // Include only the errors
		}
	} else {
		response = map[string]interface{}{
			"status": jobData.Status,
			"job_id": jobData.ID,
		}
	}
	json.NewEncoder(w).Encode(response)
}

func processJob(job models.Job) {
	var failedVisits []models.JobError
	for _, visit := range job.Visits {
		isStoreExist, _ := models.StoreExists(visit.StoreID)

		if isStoreExist {
			for _, imageURL := range visit.ImageURLs {
				_, err := ProcessImage(imageURL)
				if err != nil {
					failedVisits = append(failedVisits, models.JobError{
						StoreID: visit.StoreID,
						Error:   fmt.Sprintf("Failed to process image: %v", err),
					})
				}
			}
		} else {
			failedVisits = append(failedVisits, models.JobError{
				StoreID: visit.StoreID,
				Error:   "Store does not exist:",
			})
		}
	}

	job.Status = "completed"
	if len(failedVisits) > 0 {
		job.Status = "failed"
		job.Errors = failedVisits
	}

	models.JobStore.Store(job.ID, job)
}
