package main

import (
	"main/sender"
	yamlparser "main/yaml-parser"
	"os"
	"sync"
	"time"
)

func main() {
	configFilePath := os.Getenv("PINGGU_CONFIG_FILE")
	config, err := yamlparser.ParseYaml(configFilePath)
	if err != nil {
		panic(err)
	}
	numUsers := config.Users
	duration := config.Duration
	users := make([]sender.User, numUsers)
	for range numUsers {
		var user sender.NormalUser
		getRequests := make([]sender.Get, len(config.HttpRequests.Get))
		for _, getRequest := range config.HttpRequests.Get {
			getRequests = append(getRequests, sender.Get{
				Endpoint: getRequest.Endpoint,
			})
		}
		postRequests := make([]sender.Post, len(config.HttpRequests.Post))
		for _, postRequest := range config.HttpRequests.Post {
			postRequests = append(postRequests, sender.Post{
				Endpoint: postRequest.Endpoint,
				Body:     []byte(postRequest.Body),
			})
		}
		user.GetEndpoints = getRequests
		user.PostEndpoints = postRequests
		users = append(users, user)
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
						user.Action()
					},
				)
			}
			wg.Wait()
		}
	}
}
