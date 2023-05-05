package github

import "github-issue-schedule/internal/pkg/configs"

// ProjectInstance has all the information needed
// about a project
type ProjectInstance struct {
	project configs.Project
	client  Client
}

// GetProject returns the project object
func (g Client) GetProject(project configs.Project) *ProjectInstance {
	return &ProjectInstance{
		project: project,
		client:  g,
	}
}
