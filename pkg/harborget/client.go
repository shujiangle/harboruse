package harborget

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	baseURL    string
	username   string
	password   string
	httpClient *http.Client
}

type Project struct {
	ProjectID    int       `json:"project_id"`
	Name         string    `json:"name"`
	CreationTime time.Time `json:"creation_time"`
}

type Repository struct {
	RepositoryID int       `json:"id"`
	Name         string    `json:"name"`
	CreationTime time.Time `json:"creation_time"`
}

type Repositorytag struct {
	Name string `json:"name"`
}

func NewClient(baseURL, username, password string) *Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &Client{
		baseURL:    baseURL,
		username:   username,
		password:   password,
		httpClient: &http.Client{Transport: tr},
		//httpClient: &http.Client{},
	}
}

func (c *Client) GetProjects() ([]Project, error) {
	apiURL := fmt.Sprintf("%s/api/projects", c.baseURL)
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.username, c.password)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get projects: status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var projects []Project
	err = json.Unmarshal(body, &projects)
	//err = json.NewDecoder(resp.Body).Decode(&projects)
	if err != nil {
		return nil, err
	}

	//fmt.Println(projects)
	return projects, nil
}

func (c *Client) GetRepositories(projectID int) ([]Repository, error) {
	apiURL := fmt.Sprintf("%s/api/repositories?project_id=%d", c.baseURL, projectID)
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.username, c.password)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get repositories: status code %d, response: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var repositories []Repository
	err = json.Unmarshal(body, &repositories)
	if err != nil {
		return nil, err
	}

	return repositories, nil
}

func (c *Client) GetRepositoriesTag(repositoriestag string) ([]Repositorytag, error) {
	apiURL := fmt.Sprintf("%s/api/repositories/%v/tags", c.baseURL, repositoriestag)
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.username, c.password)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get projects: status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var tag []Repositorytag
	err = json.Unmarshal(body, &tag)
	//err = json.NewDecoder(resp.Body).Decode(&projects)
	if err != nil {
		return nil, err
	}

	//fmt.Println(projects)
	return tag, nil
}
