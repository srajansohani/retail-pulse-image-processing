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
	storeCache    map[string]bool
)

func SetStoreMasterFilePath(path string) {
	fileMutex.Lock()
	defer fileMutex.Unlock()

	storeFilePath = path
}

func LoadStores() error {
	fileMutex.Lock()
	defer fileMutex.Unlock()

	file, err := os.Open(storeFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	cache := make(map[string]bool)
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return err
		}

		if len(record) >= 3 {
			cache[record[2]] = true // record[2] is the StoreID based on previous logic
		}
	}

	storeCache = cache
	return nil
}

func StoreExists(storeID string) (bool, error) {
	fileMutex.RLock()
	defer fileMutex.RUnlock()

	if storeCache == nil {
		return false, nil // or error? Let's say false if not loaded
	}

	exists := storeCache[storeID]
	return exists, nil
}
