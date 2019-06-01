package cmd

import (
	"sort"
	"log"
	"strconv"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/venkyvb/todo-cli/todo"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Aliases: []string{"do"},
	Short: "Mark your ToDo as done",
	Long: `Helps to reach all caught up stage :)`,
	Run: runDone,
}

func runDone(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(dataFile)
	if err != nil {
		return
	}

	i, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(args[0], " is not a valid label\n", err)
	}

	if i > 0 && i < len(items) {
		items[i-1].Done = true
		fmt.Printf("%q %v\n", items[i-1].Text, " marked as Done")

		sort.Sort(todo.ByPriority(items))
		todo.SaveItems(dataFile, items)
	} else {
		log.Println(i, " doesnt match any of the items")
	}

}

func init() {
	rootCmd.AddCommand(doneCmd)
}
