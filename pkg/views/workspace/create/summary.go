// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package create

import (
	"errors"
	"fmt"
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	util "github.com/daytonaio/daytona/internal/util"
	"github.com/daytonaio/daytona/pkg/apiclient"
	"github.com/daytonaio/daytona/pkg/views"
)

type ProjectDetail string

const (
	Build              ProjectDetail = "Build"
	DevcontainerConfig ProjectDetail = "Devcontainer Config"
	Image              ProjectDetail = "Image"
	User               ProjectDetail = "User"
	EnvVars            ProjectDetail = "Env Vars"
	EMPTY_STRING                     = ""
	DEFAULT_PADDING                  = 21
)

type ProjectDefaults struct {
	BuildChoice          BuildChoice
	Image                *string
	ImageUser            *string
	DevcontainerFilePath string
}

type SummaryModel struct {
	lg          *lipgloss.Renderer
	styles      *Styles
	form        *huh.Form
	width       int
	quitting    bool
	name        string
	projectList []apiclient.CreateProjectDTO
	defaults    *ProjectDefaults
}

type SubmissionFormConfig struct {
	Name          *string
	SuggestedName string
	ExistingNames []string
	ProjectList   *[]apiclient.CreateProjectDTO
	Defaults      *ProjectDefaults
}

var configureCheck bool
var userCancelled bool
var ProjectsConfigurationChanged bool

// submission form object?

func RunSubmissionForm(name *string, suggestedName string, existingNames []string, projectList *[]apiclient.CreateProjectDTO, defaults *ProjectDefaults) error {
	configureCheck = false

	m := NewSummaryModel(name, suggestedName, existingNames, *projectList, defaults)

	if _, err := tea.NewProgram(m).Run(); err != nil {
		return err
	}

	if userCancelled {
		return errors.New("user cancelled")
	}

	if !configureCheck {
		return nil
	}

	if defaults.Image == nil || defaults.ImageUser == nil {
		return fmt.Errorf("default project entries are not set")
	}

	var err error
	ProjectsConfigurationChanged, err = ConfigureProjects(projectList, *defaults)
	if err != nil {
		return err
	}

	return RunSubmissionForm(name, suggestedName, existingNames, projectList, defaults)
}

func RenderSummary(name string, projectList []apiclient.CreateProjectDTO, defaults *ProjectDefaults) (string, error) {
	var output string
	if name == "" {
		output = views.GetStyledMainTitle("SUMMARY")
	} else {
		output = views.GetStyledMainTitle(fmt.Sprintf("SUMMARY - Workspace %s", name))
	}

	for _, project := range projectList {
		if project.Source == nil || project.Source.Repository == nil || project.Source.Repository.Url == nil {
			return "", fmt.Errorf("repository is required")
		}
	}

	output += "\n\n"

	for i := range projectList {
		if len(projectList) == 1 {
			output += fmt.Sprintf("%s - %s\n", lipgloss.NewStyle().Foreground(views.Green).Render("Project"), (*projectList[i].Source.Repository.Url))
		} else {
			output += fmt.Sprintf("%s - %s\n", lipgloss.NewStyle().Foreground(views.Green).Render(fmt.Sprintf("%s #%d", "Project", i+1)), (*projectList[i].Source.Repository.Url))
		}

		projectBuildChoice, choiceName := getProjectBuildChoice(projectList[i], defaults)
		output += renderProjectDetails(projectList[i], projectBuildChoice, choiceName)
		if i < len(projectList)-1 {
			output += "\n\n"
		}
	}

	return output, nil
}

func renderProjectDetails(project apiclient.CreateProjectDTO, buildChoice BuildChoice, choiceName string) string {
	output := projectDetailOutput(Build, choiceName)

	if buildChoice == DEVCONTAINER {
		if project.Build != nil {
			if project.Build.Devcontainer != nil {
				if project.Build.Devcontainer.DevContainerFilePath != nil {
					output += "\n"
					output += projectDetailOutput(DevcontainerConfig, *project.Build.Devcontainer.DevContainerFilePath)
				}
			}
		}
	} else {
		if project.Image != nil {
			if output != "" {
				output += "\n"
			}
			output += projectDetailOutput(Image, *project.Image)
		}

		if project.User != nil {
			if output != "" {
				output += "\n"
			}
			output += projectDetailOutput(User, *project.User)
		}
	}

	if project.EnvVars != nil && len(*project.EnvVars) > 0 {
		if output != "" {
			output += "\n"
		}

		var envVars string
		for key, val := range *project.EnvVars {
			envVars += fmt.Sprintf("%s=%s; ", key, val)
		}
		output += projectDetailOutput(EnvVars, strings.TrimSuffix(envVars, "; "))
	}

	return output
}

func projectDetailOutput(projectDetailKey ProjectDetail, projectDetailValue string) string {
	return fmt.Sprintf("\t%s%-*s%s", lipgloss.NewStyle().Foreground(views.Green).Render(string(projectDetailKey)), DEFAULT_PADDING-len(string(projectDetailKey)), EMPTY_STRING, projectDetailValue)
}

func getProjectBuildChoice(project apiclient.CreateProjectDTO, defaults *ProjectDefaults) (BuildChoice, string) {
	if project.Build == nil {
		if *project.Image == *defaults.Image && *project.User == *defaults.ImageUser {
			return NONE, "None"
		} else {
			return CUSTOMIMAGE, "Custom Image"
		}
	} else {
		if project.Build.Devcontainer != nil {
			return DEVCONTAINER, "Devcontainer"
		} else {
			return AUTOMATIC, "Automatic"
		}
	}
}

func NewSummaryModel(name *string, suggestedName string, existingNames []string, projectList []apiclient.CreateProjectDTO, defaults *ProjectDefaults) SummaryModel {
	m := SummaryModel{width: maxWidth}
	m.lg = lipgloss.DefaultRenderer()
	m.styles = NewStyles(m.lg)
	m.name = *name
	m.projectList = projectList
	m.defaults = defaults

	if *name == "" {
		*name = suggestedName
	}

	m.form = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Workspace name").
				Value(name).
				Key("name").
				Validate(func(str string) error {
					result, err := util.GetValidatedWorkspaceName(str)
					if err != nil {
						return err
					}
					for _, name := range existingNames {
						if name == result {
							return errors.New("name already exists")
						}
					}
					*name = result
					return nil
				}),
		),
	).WithShowHelp(false).WithTheme(views.GetCustomTheme())

	return m
}

func (m SummaryModel) Init() tea.Cmd {
	return m.form.Init()
}

func (m SummaryModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			userCancelled = true
			m.quitting = true
			return m, tea.Quit
		case "f10":
			m.quitting = true
			m.form.State = huh.StateCompleted
			configureCheck = true
			return m, tea.Quit
		}
	}

	var cmds []tea.Cmd

	// Process the form
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	if m.form.State == huh.StateCompleted {
		// Quit when the form is done.
		m.quitting = true
		cmds = append(cmds, tea.Quit)
	}

	return m, tea.Batch(cmds...)
}

func (m SummaryModel) View() string {
	if m.quitting {
		return ""
	}

	view := m.form.View() + configurationHelpLine

	if len(m.projectList) > 1 || len(m.projectList) == 1 && ProjectsConfigurationChanged {
		summary, err := RenderSummary(m.name, m.projectList, m.defaults)
		if err != nil {
			log.Fatal(err)
		}
		view = views.GetBorderedMessage(summary) + "\n" + view
	}

	return view
}
