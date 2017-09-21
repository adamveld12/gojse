package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

func main() {
	username := os.Getenv("GOJSE_USERNAME")
	password := os.Getenv("GOJSE_PASSWORD")

	if username == "" {
		fmt.Println("GOJSE_USERNAME envvar must be specified.")
		return
	}

	if password == "" {
		fmt.Println("GOJSE_PASSWORD envvar must be specified.")
		return
	}

	fmt.Printf("=========================\nGo JSE Miner\nMining JSE coin with %d threads\n=========================\n\n", runtime.NumCPU())

	user, err := Login(username, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Logged in as %s (user id: %d)\n", user.Email, user.UID)

	hashesInSession := 0
	for {
		block, err := Fetch()
		if err != nil {
			log.Printf("Could not retrieve block: %+v", err)
			continue
		}

		started := time.Now()
		fmt.Printf("New block (ID: %d) starting - %d hash(es) found this session\n", block.ID, hashesInSession)
		mineResult, err := MineBlock(block)
		if err != nil {
			log.Printf("An error occurred while mining %+v", err)
			continue
		}
		ended := time.Since(started)
		fmt.Printf("Block %d mined in %v seconds.\nhash: %s\nnonce: %s\ntotal hashes generated: %d\n\n",
			block.ID,
			ended.Seconds(),
			mineResult.Hash,
			mineResult.Nonce,
			mineResult.TotalHashes,
		)

		hashesInSession++
		if err := Submit(block, mineResult.Nonce, mineResult.Hash, user.UID); err != nil {
			log.Println("ERROR:", err.Error())
		}
	}
}
