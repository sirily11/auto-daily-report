package report

import (
	"auto-daily-report/src/repositories/github"
	"bytes"
	"context"
	"github.com/google/logger"
	openai "github.com/sashabaranov/go-openai"
	"html/template"
)

type ServiceInterface interface {
	GenerateReport() (string, error)
}

type Service struct {
	client          *openai.Client
	issueRepository github.IssueRepositoryInterface
}

func NewService(client *openai.Client, issueRepository github.IssueRepositoryInterface) ServiceInterface {
	return &Service{
		client:          client,
		issueRepository: issueRepository,
	}
}

func (s Service) GenerateReport() (string, error) {
	data, err := s.issueRepository.FindAllUnreportedIssues()
	if err != nil {
		return "", err
	}

	systemMessageTemplate := `
	Write a human readable report based on the following data similar to the following template:
	今天我们这边做了这些内容：
	  1. a修复了手机端字体大小匹配和弹窗大小适配问题；添加了倒计时是否可见的功能
	  2. 空投平台中，a在开始写新的ui，写了navigation bar里面的button
	  3. b在写空投平台的后端，写了空投平台的后端接口
	所有的关闭的问题都是修复了的意思。
	Data:
	**Issues closed today:**
	{{range .}}
	---
	Title: {{.Issue.Title}}
	Body: {{.Issue.Body}}
	Repository: {{.Repository.FullName}}
	Assignee: {{.Issue.Assignee.Login}}
	{{end}}
	`

	tmpl, err := template.New("systemMessage").Parse(systemMessageTemplate)
	if err != nil {
		logger.Error(err)
		return "", err
	}

	var result bytes.Buffer
	// execute the template to outputMessage
	err = tmpl.Execute(&result, data)

	logger.Info("Generating report using GPT-3 Turbo")
	gptResp, err := s.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4Turbo0125,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: result.String(),
				},
			},
		},
	)

	if err != nil {
		logger.Error(err)
		return "", err
	}

	return gptResp.Choices[0].Message.Content, nil
}
