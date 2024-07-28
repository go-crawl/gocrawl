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
		err := os.Mkdir(projectName, 0755)
		if err != nil {
			fmt.Println("Error creating project directory:", err)
			return
		}
		fmt.Println("Project", projectName, "created.")
		// Create additional directories as needed
		os.Mkdir(filepath.Join(projectName, "cmd"), 0755)
		os.Mkdir(filepath.Join(projectName, "internal"), 0755)
		os.Mkdir(filepath.Join(projectName, "configs"), 0755)
		os.Mkdir(filepath.Join(projectName, "scripts"), 0755)
		os.Mkdir(filepath.Join(projectName, "test"), 0755)
	},
}
