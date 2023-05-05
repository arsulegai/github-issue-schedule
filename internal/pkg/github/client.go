package github

import (
	ctx "context"
	"github-issue-schedule/internal/pkg/configs"
	"github-issue-schedule/internal/pkg/utils"
	"github.com/google/go-github/v33/github"
	"golang.org/x/oauth2"
)

// Client is the custom handler for all requests
type Client struct {
	Client  *github.Client
	Context ctx.Context
}

// GHClientInterface restricts possible operations
type GHClientInterface interface {
	GetProject(project configs.Project) *ProjectInstance
}

// NewClient creates a new instance of GitHub client
func NewClient() GHClientInterface {
	token := utils.GetEnvOrDefault(configs.GitHubToken, "")
	context := ctx.Background()
	if token == "" {
		return Client{
			Client:  github.NewClient(nil),
			Context: context,
		}
	}
	oauth2Client := oauth2.NewClient(context, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	))
	return Client{
		Client:  github.NewClient(oauth2Client),
		Context: context,
	}
}
