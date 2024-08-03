package startproject

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var StartprojectCmd = &cobra.Command{
	Use:   "startproject <project_name> [project_dir]",
	Short: "Create a new crawling project",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		var projectDir string
		if len(args) == 2 {
			projectDir = args[1]
		} else {
			projectDir = projectName
		}

		projectPath := filepath.Join(projectDir)
		subProjectPath := filepath.Join(projectDir, projectName)

		// Create project directory
		err := os.MkdirAll(subProjectPath, 0755)
		if err != nil {
			fmt.Println("Error creating project directory:", err)
			return
		}

		// Create additional files and directories
		createFile(filepath.Join(projectPath, "gocrawl.cfg"), generateConfigContent(projectName))
		createFile(filepath.Join(subProjectPath, "items.go"), generateGoFileContent("items"))
		createFile(filepath.Join(subProjectPath, "middlewares.go"), generateGoFileContent("middlewares"))
		createFile(filepath.Join(subProjectPath, "pipelines.go"), generateGoFileContent("pipelines"))
		createFile(filepath.Join(subProjectPath, "settings.go"), generateGoFileContent("settings"))

		fmt.Println("Project", projectName, "created at", projectPath)
	},
}

func createFile(path string, content string) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	file.WriteString(content)
}

func generateConfigContent(projectName string) string {
	return fmt.Sprintf(`# Automatically created by: gocrawl startproject
#
# For more information about the [deploy] section see:
# https://demo.com

[settings]
default = %s.settings

[deploy]
#url = http://localhost:7800/
project = %s
`, projectName, projectName)
}

func generateGoFileContent(fileType string) string {
	return fmt.Sprintf(`package %s

// TODO: Add %s implementation
`, fileType, fileType)
}
