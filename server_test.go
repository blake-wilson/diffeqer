package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func init() {
	go main()
	time.Sleep(time.Second)
}

func Test_NoError(t *testing.T) {

	data := &diffeqOpts{
		Timestep:    0.01,
		InitialTime: 0,
		FinalTime:   20,
		Expression:  "x*2",
		Method:      "euler",
	}

	body, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", "http://localhost:8080", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	assert.NoError(t, err)
	defer resp.Body.Close()
}
