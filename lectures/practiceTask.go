package main

import (
	"context"
	"fmt"
	"time"
)

// Task: Implement and Test a Simulated Database Query with Context in Go
// Objective:
//
// Understand the use of context.Context in Go for managing request-scoped values, deadlines, and cancellations.
// Implement a simulated database query function that respects context cancellation and timeouts.
// Write tests to verify that the simulated database query function correctly handles context timeouts and cancellations.
// Requirements:
//
// Implement the Simulated Database Query Function:
//
// Your function, named QueryDatabase, must accept a parameter of type context.Context.
// The function simulates a database query operation that lasts for a fixed duration (e.g., 3 seconds).
// If the context is cancelled or reaches its deadline before the query completes, the function should immediately stop and return an error indicating the cancellation.
// If the query completes successfully before the cancellation or timeout, print "Query successful" and return nil.
//
// Use the Query Function with a Context Timeout:
// In your main function, create a context with a timeout shorter than the query's duration to simulate a timeout scenario.
// Call the QueryDatabase function with this context.
// Handle and print the error if the query is cancelled or times out.
//
// Testing Your Implementation:
//
// Write a test in a separate file to verify that your QueryDatabase function behaves as expected.
// Your test should cover two scenarios:
// The query completes successfully before the context times out.
// The query is cancelled due to a context timeout.
// Use the testing package to implement your test cases and verify the correct behavior of your function.
func main() {
	ctx := context.Background()
	ctxVal, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	QueryDatabase(ctxVal)

}

func QueryDatabase(ctx context.Context) {
	ctxM, cancel := context.WithTimeout(ctx, 4*time.Second)
	defer cancel()
	select {
	case <-ctx.Done():
		fmt.Println("Query successful")

	case <-ctxM.Done():
		fmt.Print(ctx.Err())
		fmt.Println("  Timeout")

	}

}
