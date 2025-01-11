package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var eventTypes []string = []string{"WatchEvent", "CommitCommentEvent", "CreateEvent", "DeleteEvent", "ForkEvent", "GollumEvent", "IssueCommentEvent", "IssuesEvent", "MemberEvent", "PublicEvent", "PullRequestReviewEvent", "PullRequestReviewCommentEvent", "PullRequestReviewThreadEvent", "PushEvent", "ReleaseEvent", "SponsorshipEvent", "WatchEvent"}

type Actor struct {
	Id           int    `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
	GravatarId   string `json:"gravatar_id"`
	Url          string `json:"url"`
	AvatarUrl    string `json:"avatar_url"`
}
type Repo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}
type Payload struct {
	Action       string    `json:"action"`
	RepositoryId int       `json:"repository_id"`
	PushId       int       `json:"push_id"`
	Size         int       `json:"size"`
	DistinctSize int       `json:"distinct_size"`
	Ref          string    `json:"ref"`
	Head         string    `json:"head"`
	Before       string    `json:"before"`
	Commits      []Commits `json:"commits"`
}
type Author struct {
}
type Commits struct {
	Sha      string `json:"sha"`
	Distinct bool   `json:"distinct"`
	Message  string `json:"message"`
	Author   Author `json:"author"`
	Url      string `json:"url"`
}
type Event struct {
	Id        string    `json:"id"`
	Type      string    `json:"type"`
	Actor     Actor     `json:"actor"`
	Repo      Repo      `json:"repo"`
	Payload   Payload   `json:"payload"`
	Public    bool      `json:"public"`
	CreatedAt time.Time `json:"created_at"`
}
type GithubReqError struct {
	Message          string `json:"message"`
	DocumentationUrl string `json:"documentation_url"`
	Status           string `json:"status"`
}

// X-GitHub-Api-Version

func getUrl(u string) string {
	return fmt.Sprintf("https://api.github.com/users/%v/events", u)
}
func main() {
	fmt.Println("Welcome to Github Event Activity Tracker")
	fmt.Println("")
	userName, err := getUserName()
	uInputs := getUserInput()
	filterType := uInputs[2]
	if filterType != "" {
		isValidEvent := false
		for _, eventType := range eventTypes {
			if eventType == filterType {
				isValidEvent = true
				break
			}
		}

		if !isValidEvent {
			fmt.Printf("Invalid event type passed: %s\n", filterType)
			os.Exit(1)
		}
	}

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	if userName == "" {
		fmt.Printf("Error: %v\n", err)
		return
	}

	events, err := fetchData(userName)
	if err != nil {
		fmt.Println("Error", err.Error())
	}

	gEvents, err := filterEvents(events, filterType)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, event := range gEvents {
		event.displayData()
	}

}

func filterEvents(events []Event, eType string) ([]Event, error) {
	var filteredEvent []Event
	for _, value := range events {
		fmt.Println(value.Type, eType)
		if value.Type == eType {
			filteredEvent = append(filteredEvent, value)
		}
	}
	if len(filteredEvent) < 1 {
		return nil, fmt.Errorf("Event of type: %v  does not exists", eType)
	}
	return filteredEvent, nil
}

func getUserInput() []string {
	data := os.Args
	return data
}

func getUserName() (string, error) {
	tData := getUserInput()
	if len(tData) >= 3 {
		return tData[1], nil
	} else {
		return "", errors.New("error reading input")
	}
}

func fetchData(u string) ([]Event, error) {

	url := getUrl(u)
	var events []Event

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %w", err)
	}
	defer resp.Body.Close()

	respByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check if the response is empty
	if len(respByte) == 0 {
		return nil, fmt.Errorf("empty response from GitHub API")
	}

	err = json.Unmarshal(respByte, &events)
	if err != nil {
		var githubReqError GithubReqError

		if err := json.Unmarshal(respByte, &githubReqError); err != nil {
			return nil, fmt.Errorf("failed to unmarshal response: %w", err)
		}

		// Log the GitHub API error details
		fmt.Printf("GitHub API Error:\nStatus: %v\nMessage: %s\nDocumentation: %s\n",
			githubReqError.Status, githubReqError.Message, githubReqError.DocumentationUrl)

		return nil, fmt.Errorf("GitHub API error: %s", githubReqError.Message)

	}

	return events, nil
}

func (e Event) displayData() {
	fmt.Printf("- Pushed %v commits to %v \n", len(e.Payload.Commits), e.Repo.Name)
	if e.Type == "IssuesEvent" || e.Type == "IssueCommentEvent" {
		fmt.Printf("- Opened a new issue in %v\n", e.Repo.Name)
	}
	fmt.Printf("- Starred %s \n", e.Repo.Name)
	fmt.Printf("...\n")
}

// func counter() func() int {
// 	count := 0
// 	return func() int {
// 		count++
// 		return count
// 	}
// }

// func connect(url string, token *string) {
// 	req, err := http.NewRequest("POST", url, nil)
// 	if err != nil {
// 		log.Fatalf("Error creating request %v", err)
// 	}
// 	//Add custom headers
// 	tokenBearer := fmt.Sprintf("Bearer %v", token)
// 	req.Header.Add("Authorization", tokenBearer)
// 	req.Header.Add("clientId", "payfonte")

// 	// Add query params
// 	query := req.URL.Query()
// 	query.Add("eventType", "PushEvent")
// 	req.URL.RawQuery = query.Encode()

// 	// send out the request
// 	client := &http.Client{}
// 	res, err := client.Do(req)
// 	if err != nil {
// 		log.Fatalf("Error sending request: %v", err)
// 	}
// 	defer res.Body.Close()

// 	// Read the response body
// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		log.Fatalf("Error reading response body: %v", err)
// 	}

// 	// Print the response status and body
// 	fmt.Println("Response Status:", res.Status)
// 	fmt.Println("Response Body:", string(body))
// }
