package engine

import (
	"fmt"
	"time"
)

// report (Português): Monta um relatório de execução ao final do teste
func (e *Engine) report(firstDataTime, timeDuration time.Duration, frameworkName string) {
	fmt.Printf("Framework name: %v\n", frameworkName)
	fmt.Printf("Data size: %v\n", e.sizeOfData)
	fmt.Printf("Events size: %v\n", e.sizeOfEvents)
	fmt.Printf("First data load time: %v\n", firstDataTime)
	fmt.Printf("Execution time: %v\n", timeDuration)
	fmt.Printf("Events list:\n")
	fmt.Printf("  set all cache: %v\n", e.totalSetAllCache)
	fmt.Printf("  set one key: %v\n", e.totalSetOne)
	fmt.Printf("  set invalidate one key: %v\n", e.totalInvalidateKey)
	fmt.Printf("  set invalidate all data: %v\n", e.totalInvalidateAll)
	fmt.Printf("  set invalidate all data: %v\n", e.totalInvalidateAll)
	fmt.Printf("  get all: %v\n", e.totalGetAll)
	fmt.Printf("  get key: %v\n\n", e.totalGetKey)

}
