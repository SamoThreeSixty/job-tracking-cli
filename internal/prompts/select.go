package prompts

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

func YesNoPrompt(label string, yesFirst ...bool) bool {

    // If caller didn't pass a second argument, default to true
    def := false
    if len(yesFirst) > 0 {
        def = yesFirst[0]
    }

    items := []string{"Yes", "No"}
    if !def {
        items = []string{"No", "Yes"}
    }

	prompt := promptui.Select{
		Label: label,
		Items: items,
	}

	_, result, err := prompt.Run() 
	if err != nil {
		fmt.Printf("Prompt failed: %v\n", err)
		return false
	}

	return result == "Yes"
}