package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"crypto/sha256"

	cpu "github.com/shirou/gopsutil/cpu"
)

// Seeds the thread with a combination of time now and the pid
func getSeed() {
	pid := os.Getpid()
	var pid64 int64
	pid64 = int64(pid)
	timeNow := time.Now().UTC().UnixNano()
	seedInt := pid64 + timeNow
	rand.Seed(seedInt)
}

// Used to get a random int.
func getRandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// Actually get the random output
func getEntropy() ([]byte, error) {
	d, err := time.ParseDuration("1s")
	if err != nil {
		return nil, err
	}
	info, err := cpu.Percent(d, false)
	if err != nil {
		return nil, err
	}
	var infoTotal int
	for i := range info {
		usage := info[i]
		var intVal int
		intVal = int(usage)
		infoTotal = infoTotal + intVal
	}
	hash := sha256.New()
	infoString := strconv.Itoa(infoTotal)
	randInt := getRandomInt(0, 99999999)
	intString := strconv.Itoa(randInt)
	randString := intString + infoString
	hash.Write([]byte(randString))
	entropyValue := hash.Sum(nil)
	return entropyValue, nil
}

// Keep running until we exit
func main() {
	for {
		output, err := getEntropy()
		if err != nil {
			fmt.Println("Error occured getting entropy:", err)
			os.Exit(1)
		}
		fmt.Println(string(output))
		getSeed()
	}
}
