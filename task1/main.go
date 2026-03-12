package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type RepositoryInfo struct {
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	StargazersCount int       `json:"stargazers_count"`
	ForksCount      int       `json:"forks_count"`
	CreationDate    time.Time `json:"created_at"`
}

func (r RepositoryInfo) String() string {
	return fmt.Sprintf("Repository %s:\n"+
		"Description: %s\n"+
		"Stars: %d\n"+
		"Forks: %d\n"+
		"Created at: %s\n",
		r.Name,
		r.Description,
		r.StargazersCount,
		r.ForksCount,
		r.CreationDate.Format("02.01.2006"))
}

func ParseInput(input string) (string, error) {
	rawPath := input
	if strings.HasPrefix(input, "http") {
		u, err := url.Parse(input)
		if err != nil {
			return "", fmt.Errorf("url parsing failed:\n%w", err)
		}
		if u.Host != "github.com" && u.Host != "www.github.com" {
			return "", fmt.Errorf("unsupported host %s, only github.com is supported", u.Host)
		}
		rawPath = u.Path
	}
	cleanPath := strings.Trim(rawPath, "/")
	cleanPath = strings.TrimSuffix(cleanPath, ".git")
	cleanPath = strings.Trim(cleanPath, "/")

	parts := strings.Split(cleanPath, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", fmt.Errorf("invalid repository format: expected <owner/repo>, got '%s'", rawPath)
	}
	return cleanPath, nil
}

func GetRepositoryInfo(apiEndpointPath string) (*RepositoryInfo, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	apiUrl := "https://api.github.com/repos/" + apiEndpointPath
	resp, err := client.Get(apiUrl)
	if err != nil {
		return nil, fmt.Errorf("network error:\n%w", err)
	}
	defer func() {
		err := resp.Body.Close()
		fmt.Println("Warning: failed to close response body", err)
	}()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("api error:\n%s", resp.Status)
	}

	var repositoryInfo RepositoryInfo
	err = json.NewDecoder(resp.Body).Decode(&repositoryInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to decode json:\n%w", err)
	}
	return &repositoryInfo, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("invalid input arguments\nUsage: %s <github-url-or-repo/owner>\n", os.Args[0])
		return
	}
	repositoryPath, err := ParseInput(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	repositoryInfo, err := GetRepositoryInfo(repositoryPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(repositoryInfo)
}
