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
	var wg sync.WaitGroup
	for _, user := range users {
		wg.Go(
			func() {
				timer := time.NewTimer(time.Duration(duration) * time.Second)
				userActivity := []string{}
				timing := []int{}
				defer timer.Stop()
				for {
					select {
					case <-timer.C:
						fmt.Println(timing)
						return
					default:
						start := time.Now()
						res, err := user.Action()
						elapsed := time.Since(start)
						if err != nil {
							panic(err)
						}
						timing = append(timing, int(elapsed.Milliseconds()))
						userActivity = append(userActivity, string(res))
					}
				}
			},
		)
	}
	wg.Wait()
}
