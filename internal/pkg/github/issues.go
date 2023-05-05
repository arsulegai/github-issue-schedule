package github

import (
	"errors"
	"fmt"
	"github-issue-schedule/internal/pkg/configs"
	"github.com/google/go-github/v33/github"
	"log"
	"net/http"
)

// the issues provide utility methods on top
// of the ProjectInstance

// CreateIssue will create a GitHub issue
// and add maintainers as pointed
func (p *ProjectInstance) CreateIssue(schedule configs.Schedule) error {
	ghClient := p.client.Client
	issueRequest := github.IssueRequest{
		Title:    &schedule.Title,
		Body:     &schedule.Description,
		Assignee: &p.project.Maintainers,
	}
	issue, response, err := ghClient.Issues.Create(p.client.Context, p.project.GitHubOrg, p.project.GitHubRepo, &issueRequest)
	if err != nil {
		return err
	}
	log.Printf("response: %v", response)
	if response.StatusCode != http.StatusCreated {
		return errors.New(fmt.Sprintf("could not create the issue: %+v", response))
	}
	// created an issue
	log.Printf("created the issue %+v", issue.URL)
	return nil
}
