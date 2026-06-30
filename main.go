package main

import (
	"fmt"
	"main/sender"
	yamlparser "main/yaml-parser"
	"os"
	"sync"
	"time"
)

func main() {
	var configFilePath string
	configFilePath, exists := os.LookupEnv("PINGGU_CONFIG_FILE")
	if !exists {
		configFilePath = "./config.yaml"
	}
	config, err := yamlparser.ParseYaml(configFilePath)
	if err != nil {
		panic(err)
	}
	numUsers := config.Users
	duration := config.Duration
	users := make([]sender.User, numUsers)
	for i := range numUsers {
		var user sender.NormalUser
		getRequests := make([]sender.Get, len(config.HttpRequests.Get))
		for i, getRequest := range config.HttpRequests.Get {
			getRequests[i] = sender.Get{
				Endpoint: getRequest.Endpoint,
			}
		}
		postRequests := make([]sender.Post, len(config.HttpRequests.Post))
		for i, postRequest := range config.HttpRequests.Post {
			postRequests[i] = sender.Post{
				Endpoint: postRequest.Endpoint,
				Body:     []byte(postRequest.Body),
			}
		}
		user.GetEndpoints = getRequests
		user.PostEndpoints = postRequests
		users[i] = user
	}
	timer := time.NewTimer(time.Duration(duration) * time.Second)

	defer timer.Stop()
	var wg sync.WaitGroup
	for {
		select {
		case <-timer.C:
			return
		default:
			for _, user := range users {
				wg.Go(
					func() {
						res, err := user.Action()
						if err != nil {
							panic(err)
						}
						fmt.Println(string(res))
					},
				)
			}
			wg.Wait()
		}
	}
}
