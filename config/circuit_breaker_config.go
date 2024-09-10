package config

import (
	"github.com/sony/gobreaker/v2"
	"log"
	"time"
)

type StateChange struct {
	name string
	from gobreaker.State
	to   gobreaker.State
}

// var cb *gobreaker.CircuitBreaker[[]byte]
var stateChange StateChange

func CircuitBreakerConfig() *gobreaker.CircuitBreaker[[]byte] {
	settings := gobreaker.Settings{
		Name:        "key-service-cb",
		MaxRequests: 10,                  // Maximum number of requests allowed to pass through when half-open
		Interval:    60 * time.Second,    // Cyclic period of the closed state (clear internal Counts)
		Timeout:     30 * time.Second,    // Period of the open state (after which becomes half-open)
		ReadyToTrip: readyToTripFunction, // Custom function to determine if the circuit breaker should trip
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			// Log state changes
			log.Printf("Circuit breaker state changed: %s -> %s\n", from, to)
		},
	}
	return gobreaker.NewCircuitBreaker[[]byte](settings)

}

// readyToTripFunction is a custom function to determine if the circuit breaker should trip
func readyToTripFunction(counts gobreaker.Counts) bool {
	// Returns true if the number of consecutive failures is more than 5
	//return counts.ConsecutiveFailures > 5
	//fmt.Println(counts.Requests)
	failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
	return counts.Requests >= 50 && failureRatio >= 0.7

}

func CustomCircuitBreakerConfig() *gobreaker.CircuitBreaker[bool] {
	var customSt gobreaker.Settings
	customSt.Name = "cb"
	customSt.MaxRequests = 3
	customSt.Interval = time.Duration(30) * time.Second
	customSt.Timeout = time.Duration(90) * time.Second
	customSt.ReadyToTrip = func(counts gobreaker.Counts) bool {
		numReqs := counts.Requests
		failureRatio := float64(counts.TotalFailures) / float64(numReqs)

		//counts.clear() // no effect on customCB.counts

		return numReqs >= 3 && failureRatio >= 0.6
	}
	customSt.OnStateChange = func(name string, from gobreaker.State, to gobreaker.State) {
		stateChange = StateChange{name, from, to}
	}

	return gobreaker.NewCircuitBreaker[bool](customSt)
}
