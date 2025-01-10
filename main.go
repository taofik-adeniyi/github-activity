package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

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

// X-GitHub-Api-Version

func getUrl(u string) string {
	return fmt.Sprintf("https://api.github.com/users/%v/events", u)
}
func main() {
	fmt.Println("Welcome to Github Tracker")
	userName, err := getUserName()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	if userName == "" {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("github username: ", userName)

	events, err := fetchData(userName)
	if err != nil {
		fmt.Println("Error", err.Error())
	}
	for _, event := range events {
		event.displayData()
	}
}

func getUserInput() []string {
	data := os.Args
	return data
}

func getUserName() (string, error) {
	tData := getUserInput()
	if len(tData) >= 2 {
		return tData[1], nil
	} else {
		return "", errors.New("error reading input")
	}
}

func fetchData(u string) ([]Event, error) {

	url := getUrl(u)
	fmt.Println(url)
	events := []Event{}

	resp, err := http.Get(url)
	if err != nil {
		return []Event{}, err
	}
	respByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Event{}, err
	}
	err = json.Unmarshal(respByte, &events)
	if err != nil {
		return []Event{}, err
	}
	resp.Body.Close()
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

func connect(url string, token *string) {
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Fatalf("Error creating request %v", err)
	}
	//Add custom headers
	tokenBearer := fmt.Sprintf("Bearer %v", token)
	req.Header.Add("Authorization", tokenBearer)
	req.Header.Add("clientId", "payfonte")

	// Add query params
	query := req.URL.Query()
	query.Add("eventType", "PushEvent")
	req.URL.RawQuery = query.Encode()

	// send out the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer res.Body.Close()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// Print the response status and body
	fmt.Println("Response Status:", res.Status)
	fmt.Println("Response Body:", string(body))
}
