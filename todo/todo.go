package todo

import (
	"strconv"
	"encoding/json"
	"fmt"
	"io/ioutil"
)	

type Item struct {
	Text     string
	Priority int
	position int
	Done bool
}

// ByPriority implements the Sort Interface for Item based on Priority & Position
type ByPriority []Item

func (s ByPriority) Len() int {
	return len(s)
}

func (s ByPriority) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByPriority) Less(i, j int) bool {

	if s[i].Done != s[j].Done {
		return s[i].Done
	}

	if s[i].Priority != s[j].Priority {
		return s[i].Priority < s[j].Priority
	}

	return s[i].position < s[j].position
}



func (item *Item) Label() string {
	return strconv.Itoa(item.position) + "."
}

func (item *Item) PrettyP() string {
	switch item.Priority {
	case 1:
		return "(1)"
	case 3:
		return "(3)"		
	default:
		return " "	
	}
}

func (item *Item) PrettyD() string{
	switch item.Done {
	case true:
		return "X"
	default:
		return " "
	}
}


func (item *Item) SetPriority(priority int) {

	switch priority {
	case 1:
		item.Priority = 1
	case 3:
		item.Priority = 3
	default:
		item.Priority = 2
	}
}

func ReadItems(fileName string) ([]Item, error) {

	b, err := ioutil.ReadFile(fileName)

	if err != nil {
		return []Item{}, err
	}

	var items []Item
	err = json.Unmarshal(b, &items)

	if err != nil {
		return []Item{}, err
	}

	for k, _ := range items {
		items[k].position = k+1
	}
	return items, nil
}

func SaveItems(fileName string, items []Item) error {

	b, err := json.Marshal(items)
	if err != nil {
		return err
	}

	fmt.Println(string(b))

	err = ioutil.WriteFile(fileName, b, 0644)

	if err != nil {
		return err
	}

	return nil
}
