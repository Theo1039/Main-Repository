 workerPool_Practice

A simple **Go (Golang) worker pool** implementation using **goroutines**, **channels**, and **sync.WaitGroup**.
This project is for learning and practicing concurrency patterns in Go.

---

   Features

 Worker pool pattern
 Goroutines for concurrent execution
 Channels for job distribution and results
 `sync.WaitGroup` for synchronization
 Clean and beginner-friendly code

---

 Project Structure


workerPool_Practice/
├── main.go
├── go.mod
├── go.sum
└── README.md


---

   How It Works

1. A **jobs channel** sends tasks to workers
2. Multiple **worker goroutines** listen for jobs
3. Each worker processes a job and sends results
4. A **WaitGroup** ensures all workers finish before exit

This pattern is commonly used in real-world Go applications for:

  Background processing
  Task queues
  Parallel workloads

---

   How to Run

Make sure Go is installed:

...bash
go version


Run the program:

...bash
go run main.go
---

---

  Requirements

 Go 1.20+

---

   Learning Goals

 Understand goroutines
 Learn channel communication
 Practice synchronization using WaitGroups
 Write idiomatic Go concurrency code

---

    License

This project is for learning purposes.

---

  Author

  Theophilus Kpurugbara**
Learning Go & backend development 
