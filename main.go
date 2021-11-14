package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/procyon-projects/chrono"
)

// GitCommit the current git commit
// will be injected during build
var GitCommit string

// Version
var Version string

// HumanVersion version with commit
var HumanVersion = fmt.Sprintf("%s-(%s)", Version, GitCommit)

func main() {

	fmt.Printf("Hello world %s\n", HumanVersion)
	// job := func() {
	// 	t := time.Now()
	// 	fmt.Println("Time's up! @", t.UTC())
	// }

	// onetimejob := func() {
	// 	t := time.Now()
	// 	fmt.Println("One time! @", t.UTC())
	// }
	now := time.Now()
	taskScheduler := chrono.NewDefaultTaskScheduler()

	_, err := taskScheduler.Schedule(func(ctx context.Context) {
		log.Print("One-Shot Task")
	}, chrono.WithStartTime(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second()+1))

	if err == nil {
		log.Print("Task has been scheduled successfully.")
	}

	time.Sleep(30 * time.Second)
	// // Keep the program from not exiting.
	// runtime.Goexit()

}
