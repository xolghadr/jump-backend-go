package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/rand"
)

func simpleTask() string {
	time.Sleep(1 * time.Second)
	return "result"
}

func TestSimple(t *testing.T) {
	fResult := Async(simpleTask)
	assert.False(t, fResult.Done.Load())
	result := fResult.Await()
	assert.Equal(t, "result", result)
	assert.True(t, fResult.Done.Load())
}

func handleTimeout(d time.Duration, t *testing.T) func() {
	start := time.Now()
	return func() {
		elapsed := time.Since(start)
		if elapsed > d {
			t.FailNow()
		}
	}
}

func TestMultiple(t *testing.T) {
	defer handleTimeout(1100*time.Millisecond, t)()

	fResult1 := Async(simpleTask)
	fResult2 := Async(simpleTask)

	res1 := fResult1.Await()
	res2 := fResult2.Await()

	assert.Equal(t, "result", res1)
	assert.Equal(t, "result", res2)

	// they should run concurrently and done in about a second
}

func TestCombine(t *testing.T) {
	fResult1 := Async(simpleTask)
	fResult2 := Async(simpleTask)

	combinedFResult := CombineFutureResults(fResult1, fResult2)

	// first item
	select {
	case <-time.After(1100 * time.Millisecond):
		t.FailNow()

	case res := <-combinedFResult.ResultChan:
		assert.Equal(t, "result", res)
	}

	// second item should be availble fast
	select {
	case <-time.After(100 * time.Millisecond):
		t.FailNow()

	case res := <-combinedFResult.ResultChan:
		assert.Equal(t, "result", res)
	}
}

func TestTimeout(t *testing.T) {
	fResult := AsyncWithTimeout(simpleTask, 700*time.Millisecond)

	select {
	case <-time.After(800 * time.Millisecond):
		t.FailNow()

	case res := <-fResult.ResultChan:
		assert.Equal(t, "timeout", res) // timeout is reached before 800ms
	}
}
func TestUselessTimeout(t *testing.T) {
	fResult := AsyncWithTimeout(simpleTask, 1700*time.Millisecond)

	select {
	case <-time.After(2000 * time.Millisecond):
		t.FailNow()

	case res := <-fResult.ResultChan:
		assert.Equal(t, "result", res) // timeout is reached before 2000ms
	}
}
func TestMassiveCombine(t *testing.T) {
	// Create a large number of async tasks
	numTasks := 1000
	futures := make([]*FutureResult, numTasks)

	for i := 0; i < numTasks/2; i++ {
		futures[i] = Async(func() string {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
			return fmt.Sprintf("result-%d", i)
		})
	}
	for i := numTasks / 2; i < numTasks; i++ {
		futures[i] = AsyncWithTimeout(func() string {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
			return fmt.Sprintf("result-%d", i)
		}, 120*time.Millisecond)
	}

	// Combine all futures
	combinedFuture := CombineFutureResults(futures...)

	// Collect and verify results
	results := make([]string, 0, numTasks)
	for i := 0; i < numTasks; i++ {
		select {
		case res := <-combinedFuture.ResultChan:
			results = append(results, res)
		case <-time.After(2 * time.Second):
			t.Fatalf("Timeout waiting for result %d", i)
		}
	}

	// Verify we got all results
	assert.Equal(t, numTasks, len(results))

	// Optional: Check that all results are unique or match expected format
	uniqueResults := make(map[string]bool)
	for _, res := range results {
		uniqueResults[res] = true
	}
	assert.Equal(t, numTasks, len(uniqueResults))
}

func TestDone(t *testing.T) {
	defer handleTimeout(3*time.Second, t)()
	n := 10
	results := make([]*FutureResult, 0, n)
	for i := 0; i < n/2; i++ {
		results = append(results, Async(simpleTask))
		results = append(results, AsyncWithTimeout(simpleTask, 1200*time.Millisecond))
	}
	for _, f := range results {
		assert.False(t, f.Done.Load())
	}
	for _, f := range results {
		_ = f.Await()
		assert.True(t, f.Done.Load())
	}
}
