package main

import "fmt"
import "net/http"
import "crypto/sha1"
import "io/ioutil"
import "log"
import "bufio"
import "os"
import "sync"

// Reads newline-separated URLs from stdin and computes their SHA1 checksum.
func main() {
	client := new(http.Client)
	wg := new(sync.WaitGroup)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		go fetchURL(scanner.Text(), client, wg)
		wg.Add(1)
	}

	wg.Wait()
}

// Fetches url and prints the SHA1 checksum of the response body.
func fetchURL(url string, cl *http.Client, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := cl.Get(url)
	if err != nil {
		log.Fatal(err, url)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s -> %x\n", url, sha1.Sum(body))
}
