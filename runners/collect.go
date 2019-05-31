package runners

import (
	"sync"

	"github.com/cha87de/kvmtop/models"
)

// InitializeCollect starts the periodic collect calls
func InitializeCollect(wg *sync.WaitGroup) {
	for {
		// wait with execution for lookup routine
		_, ok := <-lookupDone
		if !ok {
			wg.Done()
			return
		}
		Collect()
	}
}

// Collect runs one collect cycle to measure frequently changing metrics
func Collect() {
	// initialize models
	/*if models.Collection.Domains.Length() <= 0 {
		// wait for lookup to create domains
		return
	}*/

	// run collectors
	models.Collection.Collectors.Map.Range(func(_, collectorRaw interface{}) bool {
		collector := collectorRaw.(models.Collector)
		go collector.Collect()
		return true
	})
}
