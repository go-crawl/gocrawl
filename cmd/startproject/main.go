package startproject

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var StartprojectCmd = &cobra.Command{
	Use:   "startproject <project_name>",
	Short: "Create a new crawling project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		projectPath := filepath.Join(projectName)
		subProjectPath := filepath.Join(projectName, projectName)

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

		fmt.Println("Project", projectName, "created.")
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
