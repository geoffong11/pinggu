# About
A program that aims to simulate multiple users (good and bad actors).

Many times, scale is something that is not really we can do at home.
Hence this program helps to do so. It is a Golang program that
simulates multiple users.

Focuses on HTTP APIs (GET, POST, DELETE)

WARNINGS: 
- Since this is a very simple program designed to test the concurrency and the timing of the
  website, the timing output from this program will be influenced by network connectivity. A better measure
  will be to check the duration from the app itself.
- You may have to turn off rate limiting to see the full effects of concurrency

# Goals
1) Create multiple users through multiple concurrent goroutines
2) Track the timing between requests and responses, to see if there is any degradation (can use otel)
3) Create bad actor users if specified. But what is bad actor users?
    - A bad actor user is one that sends malformed requests
    - Simulate in the real world when network breaks down


# What it do
1) Read the yaml config file
2) Send requests based on the config file
3) After the duration (in seconds) is up, return the results

# Input
A sample yaml file is given below:
```
users: 3
duration: 5 # in seconds
requests:
  GET:
    - endpoint: http://localhost:3000/jobs
    - endpoint: http://localhost:3000/job/1
  POST:
    - endpoint: http://localhost:3000/locations/add_locations
      body: |
        {
          "location": "Australia"
        }
    - endpoint: http://localhost:3000/jobs/1/edit_job
      body: |
        {
          "pay": 10000
        }

```

Put your yaml config file in the location specified by the environment variable `PINGGU_CONFIG_FILE`
