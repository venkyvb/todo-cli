package cmd

import (
	"log"
	"github.com/spf13/cobra"
	"github.com/venkyvb/todo-cli/todo"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a ToDO",
	Long: `add will create a new ToDo item`,
	Run: addRun,
}

var priority int

func addRun(cmd *cobra.Command, args []string) {

	items, err := todo.ReadItems(dataFile)

	if err != nil {
		log.Println("Error reading items")
	}

	for _, v := range args {
		item := todo.Item{Text: v}
		item.SetPriority(priority)
		items = append(items, item)
	}

	todo.SaveItems(dataFile, items)
} 

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority: 1,2,3")
}
