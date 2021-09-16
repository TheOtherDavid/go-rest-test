package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"time"

	c "github.com/TheOtherDavid/go-rest-test/config"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigName("config")

	viper.AddConfigPath("./config")

	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var configuration c.Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	fmt.Println("Reading variables...")

	fmt.Println("Target URL: ", configuration.Rest.Url)
	baseUrl := configuration.Rest.Url

	sleepTime := 3000 * time.Millisecond

	fmt.Println("Sending Call 0")
	sendRenameCommand(baseUrl, "0")

	time.Sleep(sleepTime)

	fmt.Println("Sending Call 1")
	sendRenameCommand(baseUrl, "1")

	time.Sleep(sleepTime)

	fmt.Println("Sending Call 2")
	sendRenameCommand(baseUrl, "2")

	time.Sleep(sleepTime)

	fmt.Println("Sending Call 3")
	sendRenameCommand(baseUrl, "3")

	time.Sleep(sleepTime)

	fmt.Println("Sending Call 4")
	sendRenameCommand(baseUrl, "4")

	time.Sleep(sleepTime)

	fmt.Println("Sending Call 5")
	sendRenameCommand(baseUrl, "5")

	time.Sleep(sleepTime)

	fmt.Println("Sending Call 6")
	sendRenameCommand(baseUrl, "6")

	time.Sleep(sleepTime)

	fmt.Println("Sending Call 7")
	sendRenameCommand(baseUrl, "7")

	time.Sleep(sleepTime)

	fmt.Println("Sending Call 8")
	sendRenameCommand(baseUrl, "8")

	time.Sleep(sleepTime)

	fmt.Println("Sending Call 9")
	sendRenameCommand(baseUrl, "9")

	time.Sleep(sleepTime)

	fmt.Println("Sending Call 10")
	sendRenameCommand(baseUrl, "10")

	fmt.Println("Complete")
}

func sendRenameCommand(baseUrl string, body string) error {
	jsonBody := []byte(body)

	fullUrl := baseUrl

	req, _ := http.NewRequest("POST", fullUrl, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	//Comment out the timeout to try to recreate timeout bugs.
	//client.Timeout = time.Second * 5

	resp, err := client.Do(req)
	if resp != nil {
		println("Response code " + strconv.Itoa(resp.StatusCode))
	} else {
		println("Error code " + err.Error())
	}
	return err
}
