package cmd

import (
	"sort"
	"os"
	"text/tabwriter"
	"fmt"
	"log"
	"github.com/spf13/cobra"
	"github.com/venkyvb/todo-cli/todo"
)

var (
	showDone bool 
	showAll bool
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the ToDo items",
	Long: `Allows you to list the ToDO items that you have to work on`,
	Run: listRun,
}

func listRun(cmd *cobra.Command, args []string) {
	
	items, err := todo.ReadItems(dataFile)

	sort.Sort(todo.ByPriority(items))

	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ',0)	
	if err != nil {
		log.Printf("%v", err)
	}

	for _, v := range items {
		
		if showAll || v.Done == showDone {
			fmt.Fprintln(w, v.Label() + "\t" + v.PrettyD() + "\t" + v.PrettyP() + "\t" + v.Text + "\t")
		}

	}

	w.Flush()
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&showDone, "done", "d", false, "Show 'Done' items")
	listCmd.Flags().BoolVarP(&showAll, "all", "a", false, "Show all items, including 'Done'")
}
