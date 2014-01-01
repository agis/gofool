package main

import (
	"bufio"
	"fmt"
	"hash/crc32"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

const NShard = 4

func main() {
	started := time.Now()

	re := regexp.MustCompile("\"uuid\":\"([[:xdigit:]]+)\"")

	f, err := os.Open("a.log")
	if err != nil {
		log.Fatal(err)
	}

	shards := [NShard]*os.File{}
	writers := [NShard]*bufio.Writer{}
	for i := 0; i < NShard; i++ {
		shards[i], err = os.Create("shard_" + strconv.Itoa(i))
		if err != nil {
			log.Fatal(err)
		}

		writers[i] = bufio.NewWriter(shards[i])
	}

	var match [][]byte

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if match = re.FindSubmatch(scanner.Bytes()); match != nil {
			writers[crc32.ChecksumIEEE(match[1])%NShard].Write(scanner.Bytes())
		}
	}

	elapsed := time.Since(started)

	fmt.Println(elapsed)
}
