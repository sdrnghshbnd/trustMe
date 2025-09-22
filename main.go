package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

// CPU Burner that maxes out one core
func cpuStresser(id int, duration time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("CPU Burner %d started on core\n", id)

	start := time.Now()
	for time.Since(start) < duration {
		// dummy calc to burn CPU
		for i := 0; i < 1000000; i++ {
			_ = i * i * i
		}
	}
	fmt.Printf("CPU Burner %d finished\n", id)
}

// RAM Burner that allocates large amounts of memory
func ramStresser(id int, duration time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("RAM Burner %d started\n", id)

	var memoryChunks [][]byte
	start := time.Now()

	for time.Since(start) < duration {
		// Allocate 100MB chunks
		chunk := make([]byte, 100*1024*1024)
		// Gen dummy data to fill the space
		for i := range chunk {
			chunk[i] = byte(i % 256)
		}
		memoryChunks = append(memoryChunks, chunk)
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Printf("RAM Consumer %d finished (allocated %d chunks)\n", id, len(memoryChunks))
}

// Disk Burner with large file operations
func diskStresser(id int, duration time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Disk Burner %d started\n", id)

	filename := fmt.Sprintf("sorry_%d.tmp", id)
	os.Remove(filename)

	data := make([]byte, 100*1024*1024) // 100MB

	// Dummy data
	for i := range data {
		data[i] = byte(i % 256)
	}

	start := time.Now()
	for time.Since(start) < duration {
		// Dummy write operations
		file, _ := os.Create(filename)
		for j := 0; j < 10; j++ {
			file.Write(data)
		}
		file.Close()

		// Dummy read operations
		file, _ = os.Open(filename)
		buffer := make([]byte, 100*1024*1024)
		for j := 0; j < 10; j++ {
			file.Read(buffer)
			file.Seek(0, 0)
		}
		file.Close()
	}

	fmt.Printf("Disk Burner %d finished\n", id)
}

// System Burner - uses everything for 1 minutes
func systemStressTest() {
	fmt.Println("ARRRRE YOUUUUUU READYYYYYYYYYYYYY ?\n")
	fmt.Println("APOCALYPSE NOW !\n\n\n")
	time.Sleep(3 * time.Second)

	duration := 1 * time.Minute
	numCPU := runtime.NumCPU()

	var wg sync.WaitGroup

	// CPU Burner for all cores
	for i := 0; i < numCPU; i++ {
		wg.Add(1)
		go cpuStresser(i+1, duration, &wg)
	}

	// RAM Burner 4 consumers
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go ramStresser(i+1, duration, &wg)
	}

	// Disk Burner with 20 threads
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go diskStresser(i+1, duration, &wg)
	}
	start := time.Now()

	wg.Wait()

	fmt.Printf("STRESS TEST COMPLETED in %v\n", time.Since(start))
	fmt.Println("Thank you for your Time, Don't run random .exe files anymore")
}

func main() {
	s := []string{
		"+ HELLO ! Are you ready for APOCALYPSE ???",
		"- Wait, What ?????",
		"+ YEAAAH, You shouldn't run that .exe file :",
		"- Okay i regret, STOP it !",
		"+ TOO LATE :))) HAHAHAHAHAHAHAHAHAHAAAAAAAAAAAAAAAAAAA"}

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
		time.Sleep(2 * time.Second)
	}
	fmt.Println()

	// FULL SYSTEM STRESS TEST
	systemStressTest()
}
