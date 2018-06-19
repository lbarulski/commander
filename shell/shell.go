package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

func main() {
	prompt := promptui.Select {
		Label: "Select Day",
		Items: []string{"Enter random pod from deployment", "Create a pod for me"},
		Templates: &promptui.SelectTemplates{
			Active: `▸ {{ . }}`, // gotemplate
			Inactive: `{{ . }}`,
		},
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
}