package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"examples.com/interfaces/note"
	"examples.com/interfaces/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {

	printSomething(1)
	printNumber(1.2)

	title, content := getNoteData()
	todoText := getUserInput("To do text:")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	note, err := note.New(title, content)

	if err != nil {
		fmt.Print(err)
		return
	}

	err = outputData(note)

	if err != nil {
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}
}

func printSomething(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("Integer: ", value)
	case float64:
		fmt.Println("Float: ", value)
	case string:
		fmt.Println("String: ", value)
	}

}

func printNumber(value any) {
	intVal, ok := value.(int)

	if ok {
		intVal = intVal + 1
		fmt.Println(intVal)
	}

	floatVal, ok := value.(float64)

	if ok {
		floatVal = floatVal + 1
		fmt.Println(floatVal)
	}

}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed")
		return err
	}

	fmt.Println("Saving the code succeeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
