package main

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

func Mine(block []byte) []byte {
	hash := sha256.Sum256(block)
	if hash[0] == 0 && hash[1] == 0 {
		return hash[:]
	}

	return nil
}

type work struct {
	block []byte
	nonce string
}

type result struct {
	nonce     string
	hash      string
	hashCount int
}

func MineBlock(block *Block) (string, string, error) {
	// max # of threads
	count := runtime.NumCPU()
	found := false
	startNum := rand.Intn(99999999)
	workQueue := make(chan work)
	results := make(chan result)
	quit := make(chan bool)
	wg := &sync.WaitGroup{}

	for threads := count; threads > 0; threads-- {
		wg.Add(1)
		go asyncMine(workQueue, results, quit, wg)
	}

	for x := startNum; !found; x++ {
		block.Nonce = fmt.Sprintf("%d", x)
		targetText, _ := json.Marshal(block)
		select {
		case r := <-results:
			close(quit)
			close(workQueue)
			close(results)
			wg.Wait()
			return r.hash, r.nonce, nil
		case workQueue <- work{targetText, block.Nonce}:
		}
	}

	return "", "", errors.New("FUC")
}

func asyncMine(workQueue <-chan work, results chan<- result, quit <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	totalHashes := 0

	for work := range workQueue {
		select {
		case <-quit:
			return
		default:
			totalHashes++
			if hash := Mine(work.block); hash != nil {
				results <- result{nonce: work.nonce, hash: fmt.Sprintf("%x", hash), hashCount: totalHashes}
				return
			}
		}
	}
}
