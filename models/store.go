package models

import (
	"encoding/csv"
	"os"
	"sync"
)

type Store struct {
	StoreId   string `json:"store_id`
	StoreName string `json:"store_name"`
	AreaCode  string `json:"area_code"`
}

var (
	storeFilePath string = "StoreMasterAssignment.csv"
	fileMutex            = sync.RWMutex{}
)

func SetStoreMasterFilePath(path string) {
	fileMutex.Lock()
	defer fileMutex.Unlock()

	storeFilePath = path
}
func StoreExists(storeID string) (bool, error) {
	fileMutex.RLock()
	defer fileMutex.RUnlock()

	file, err := os.Open(storeFilePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return false, err
		}

		if len(record) < 1 {
			continue // Skip invalid rows
		}

		if len(record) >= 3 && record[2] == storeID {
			return true, nil
		}
	}

	return false, nil
}
