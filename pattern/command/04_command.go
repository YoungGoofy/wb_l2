package main

import (
	"fmt"
	"strings"
)

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
Позволяет инкапсулировать запросы в объекты. Таким образом, мы можем передавать запросы как аргументы в другие объекты,
позволяя нам параметризовать клиентские объекты с различными запросами.
*/

type Command interface {
	Execute()
}

type AddCommand struct {
	Text   string
	Editor *TextEditor
}

func (ac *AddCommand) Execute() {
	ac.Editor.AddText(ac.Text)
}

type RemoveCommand struct {
	Text   string
	Editor *TextEditor
}

func (rc *RemoveCommand) Execute() {
	rc.Editor.RemoveText(rc.Text)
}

type TextEditor struct {
	Text string
}

func (te *TextEditor) AddText(text string) {
	te.Text += text
}

func (te *TextEditor) RemoveText(text string) {
	te.Text = strings.Replace(te.Text, text, "", -1)
}

type Invoker struct {
	commands []Command
}

func (i *Invoker) AddCommand(c Command) {
	i.commands = append(i.commands, c)
}

func (i *Invoker) ExecuteCommand() {
	for _, command := range i.commands {
		command.Execute()
	}
}

func main() {
	editor := &TextEditor{}
	addCommand := &AddCommand{Text: "hello", Editor: editor}
	newAddCommand := &AddCommand{Text: "world", Editor: editor}
	removeCommand := &RemoveCommand{Text: "world", Editor: editor}

	invoker := Invoker{}
	invoker.AddCommand(addCommand)
	invoker.AddCommand(newAddCommand)
	invoker.AddCommand(removeCommand)

	invoker.ExecuteCommand()

	fmt.Println(editor.Text)
}
