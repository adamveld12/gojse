package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	loginURL        = "https://jsecoin.com/server/login/"
	blockRequestURL = "https://jsecoin.com/server/request/"
	submitHashURL   = "https://jsecoin.com/server/submit/"
)

type Block struct {
	ID           int    `json:"block"`
	Nonce        string `json:"nonce"`
	PreviousHash string `json:"previousHash"`
	Hash         string `json:"hash"`
	Version      string `json:"version"`
	Server       string `json:"server"`
	Difficulty   int    `json:"difficulty"`
	Data         string `json:"data"`
	Frequency    int    `json:"frequency"`
	StartTime    int64  `json:"startTime"`
	Size         int    `json:"size"`
}

type User struct {
	UID   int    `json:"uid"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type submitStruct struct {
	BlockID int    `json:"block"`
	Hash    string `json:"hash"`
	Nonce   string `json:"nonce"`
	UserID  int    `json:"pubid"`
	SiteID  string `json:"siteid"`
	Uniq    string `json:"uniq,omitempty"`
}

func Login(email, password string) (*User, error) {
	payload := fmt.Sprintf("{ \"email\": \"%s\", \"password\": \"%s\", \"initial\":1}", email, password)
	req, err := http.NewRequest("POST", loginURL, bytes.NewBufferString(payload))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, fmt.Errorf("Could not create request for logging in: %+v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Failed making request for login data: %+v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Response code received was not 200 but %d - %s", res.StatusCode, res.Status)
	}
	decoder := json.NewDecoder(res.Body)

	loginRes := User{}
	if err := decoder.Decode(&loginRes); err != nil {
		return nil, fmt.Errorf("Could not read response body: %+v", err)
	}

	return &loginRes, nil
}

func Fetch() (*Block, error) {
	res, err := http.PostForm(blockRequestURL, url.Values(map[string][]string{
		"o": []string{"1"},
	}))

	if err != nil {
		return nil, fmt.Errorf("Failed making request for block data: %+v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Response code received was not 200 but %d - %s", res.StatusCode, res.Status)
	}

	decoder := json.NewDecoder(res.Body)

	block := Block{}
	if err := decoder.Decode(&block); err != nil {
		return nil, fmt.Errorf("Could not decode block json: %+v", err)
	}

	return &block, nil
}

func Save() error {
	return nil

}

func Submit(block *Block, nonce, hash string, uid int) error {
	submission := submitStruct{
		block.ID,
		hash,
		nonce,
		uid,
		"Platform Mining",
		"",
	}

	payload, _ := json.Marshal(submission)
	res, err := http.DefaultClient.PostForm(submitHashURL, map[string][]string{
		"o": []string{fmt.Sprintf("%s", payload)},
	})

	if err != nil {
		return fmt.Errorf("Could not complete request to submit block: %+v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Unexpected http response code, got %d - %+v", res.StatusCode, res.Status)
	}

	// the response just says thanks
	/*
		resPayload, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("Could not read response: %+v", err)
		}
	*/

	return nil
}
