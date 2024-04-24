package github

import (
	github2 "auto-daily-report/src/dto/github"
	"auto-daily-report/src/repositories/github"
)

type ServiceInterface interface {
	FetchAllUnreportedIssues() ([]github2.IssueWebhookDTO, error)
	InsetIssue(issue github2.IssueWebhookDTO) error
}

type Service struct {
	githubIssueRepository github.IssueRepositoryInterface
}

func NewService(githubIssueRepository github.IssueRepositoryInterface) ServiceInterface {
	return &Service{
		githubIssueRepository: githubIssueRepository,
	}
}

func (s Service) FetchAllUnreportedIssues() ([]github2.IssueWebhookDTO, error) {
	return s.githubIssueRepository.FindAllUnreportedIssues()
}

func (s Service) InsetIssue(issue github2.IssueWebhookDTO) error {
	return s.githubIssueRepository.InsertIssue(issue)
}
