package github

import (
	"auto-daily-report/src/config/constants/Database"
	"auto-daily-report/src/config/constants/environments"
	"auto-daily-report/src/dto/github"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/logger"
	"net/http"
)

type IssueRepositoryInterface interface {
	// FindAllUnreportedIssues returns all unreported issues
	FindAllUnreportedIssues() ([]github.IssueWebhookDTO, error)
	InsertIssue(issue github.IssueWebhookDTO) error
}

type IssueRepository struct {
	client         *resty.Client
	collectionName string
}

func NewIssueRepository() IssueRepositoryInterface {
	client := resty.New()
	client.BaseURL = environments.MongoDBDataAPIEndpoint
	client.SetHeaders(map[string]string{
		"apiKey": environments.MongoDBDataAPIKey,
	})
	return &IssueRepository{
		client:         client,
		collectionName: "issues",
	}
}

func (r *IssueRepository) FindAllUnreportedIssues() ([]github.IssueWebhookDTO, error) {
	resp, err := r.client.R().SetBody(DataAPIRequest{
		DataSource: environments.MongoDBDataSource,
		Database:   Database.DatabaseName,
		Collection: r.collectionName,
		Filter: map[string]interface{}{
			"reported": map[string]interface{}{
				"$ne": true,
			},
		},
	}).SetResult(map[string]any{}).Post("/action/find")
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		logger.Error(resp)
		return nil, fmt.Errorf(resp.String())
	}

	var data []github.IssueWebhookDTO
	resultPtr := resp.Result().(*map[string]any)
	result := *resultPtr
	jsonResult, _ := json.Marshal(result["documents"])
	err = json.Unmarshal(jsonResult, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *IssueRepository) InsertIssue(issue github.IssueWebhookDTO) error {
	req := DataAPIRequest{
		DataSource: environments.MongoDBDataSource,
		Database:   Database.DatabaseName,
		Collection: r.collectionName,
		Document:   issue,
	}
	resp, err := r.client.R().SetBody(req).Post("/action/insertOne")
	if err != nil {
		logger.Error(err)
		return err
	}

	if resp.StatusCode() != http.StatusCreated {
		logger.Error(resp)
		return fmt.Errorf(resp.String())
	}

	return nil
}
