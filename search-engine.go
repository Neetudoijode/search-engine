package search

import (
	"strings"
	"sync"
	"backend/parser"
)

var (
	inMemoryData []parser.Record
	mu           sync.RWMutex
)

func LoadData(records []parser.Record) {
	mu.Lock()
	defer mu.Unlock()
	inMemoryData = records
}

func Search(keyword string) []parser.Record {
	mu.RLock()
	defer mu.RUnlock()

	var results []parser.Record
	for _, record := range inMemoryData {
		if strings.Contains(strings.ToLower(record.Message), strings.ToLower(keyword)) {
			results = append(results, record)
		}
	}
	return results
}