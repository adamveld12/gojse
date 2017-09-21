package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var nonceStr = []byte("*nonce*")

func Mine(block []byte) []byte {
	hash := sha256.Sum256(block)
	if hash[0] == 0 && hash[1] == 0 {
		return hash[:]
	}

	return nil
}

type result struct {
	nonce     string
	hash      string
	hashCount int
}

type MineResult struct {
	Hash        string
	Nonce       string
	TotalHashes int
}

func MineBlock(block *Block) (MineResult, error) {
	count := runtime.NumCPU()
	results := make(chan result)
	quit := make(chan bool)
	stats := make(chan int)
	wg := &sync.WaitGroup{}

	for threads := count; threads > 0; threads-- {
		wg.Add(1)
		go asyncMine(*block, stats, results, quit, wg)
	}

	hps := 0
	totalHashes := 0
	started := time.Now()
	for {
		select {
		case r := <-results:
			close(quit)
			wg.Wait()
			close(results)
			close(stats)
			return MineResult{
				Hash:        r.hash,
				Nonce:       r.nonce,
				TotalHashes: totalHashes,
			}, nil
		case hashes := <-stats:
			hps += hashes
			totalHashes += hashes
			if time.Since(started) > time.Second {
				fmt.Printf("\t%d Hashes per second.\n", hps)
				started = time.Now()
				hps = 0
			}
		}
	}
}

func asyncMine(block Block, stats chan<- int, results chan<- result, quit <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	found := false
	hashCount := 0
	startNum := rand.Int63n(99999999)
	block.Nonce = "*nonce*"
	sourceText, _ := json.Marshal(block)
	buf := make([]byte, binary.MaxVarintLen64)

	for x := startNum; !found; x++ {
		select {
		case <-quit:
			return
		case stats <- hashCount:
			hashCount = 0
		default:
			hashCount++
			binary.PutVarint(buf, x)
			targetText := bytes.Replace(sourceText, nonceStr, buf, 1)
			if hash := Mine(targetText); hash != nil {
				select {
				case results <- result{nonce: fmt.Sprintf("%d", x), hash: fmt.Sprintf("%x", hash), hashCount: hashCount}:
				case <-quit:
					return
				}
				return
			}
		}
	}
}
