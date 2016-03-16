package main

import (
	"bytes"
	"fmt"
)

type Task struct {
	Title       string
	Description string
	Finished    bool
	Ignored     bool
	Parent      *Task
	Children    []*Task
}

func (task *Task) AddChild(child *Task) {
	task.Children = append(task.Children, child)
}

func (task *Task) String() string {
	buffer := bytes.NewBufferString("")

	fmt.Fprintf(buffer, "%v -> %v\n", task.Title, task.Description)
	for _, child := range task.Children {
		if child.Finished {
			fmt.Fprintf(buffer, "\t[x] %v -> %v\n", child.Title, child.Description)
		} else {
			fmt.Fprintf(buffer, "\t[ ] %v -> %v\n", child.Title, child.Description)
		}
	}
	return buffer.String()
}

//createExample returns an example Task
func createExample() Task {
	parent := Task{
		Title:       "Checklist webapp",
		Description: "Write a webscale checklist management app."}
	child1 := Task{
		Title:       "GitHub Repository",
		Description: "Create a github project, because it's nice.",
		Finished:    true}
	child2 := Task{
		Title:       "Prototype",
		Description: "Create a first prototype"}
	child3 := Task{
		Title:       "Design",
		Description: "Create some nice js and css, for much bling bling"}
	child4 := Task{
		Title:       "Java-Port",
		Description: "Port the webapp to java + tomcat and introduce an AbstractFactoryDesignPattern",
		Ignored:     true}
	parent.Children = []*Task{&child1, &child2, &child3, &child4}

	return parent
}

func main() {
	example := createExample()
	fmt.Printf(example.String())
}
