package roller

import (
	"fmt"
	"roller/pkg/interaction"
)

// Survey Performs a survey of the configuration's variables.
//
// keyFilter should provide a specific list of keys to survey, useful for new variables.
// newSurvey should indicate whether this is an entirely new survey (i.e. create a new template).
//
// Returns whether the survey is successful (true), or aborted (false).
func Survey(config *Config, keyFilter []string, newSurvey bool) bool {

	// Check if survey is skipped
	appArgs := interaction.ParseAppArgs()
	surveyAppArg, forcedSurvey := appArgs["survey"]
	if surveyAppArg == "skip" {
		return true
	}

	// Build the list of vars to be surveyed
	var vars map[string]ConfigVar

	if newSurvey || forcedSurvey {
		vars = config.Template.Vars
	} else {
		for _, key := range keyFilter {
			value, ok := config.Template.Vars[key]
			if ok {
				vars[key] = value
			}
		}
	}

	// Check we have something to survey, unless forced survey
	if !forcedSurvey && len(vars) == 0 {
		if newSurvey {
			fmt.Println("Skipped survey, template has no customisable variables.")
		} else {
			fmt.Println("Skipped survey, no new customisable variables.")
		}
		return true
	}

	// Inform the user of the survey
	if newSurvey {
		fmt.Println("This template has customisable variables, you will now be asked whether you want to change the default values.")
	} else {
		fmt.Println("New customisable variables for this project have been added, you will now be asked whether you want to change the default values.")
	}
	fmt.Println("")

	// Survey for answers
	for {

		fmt.Println("Press enter to keep the default value.")
		fmt.Println("")
		surveyVars(&vars)

		askNext := true
		for askNext {
			fmt.Println("Continue? (y = yes, n = abort, r = redo the survey)")
			outcome := interaction.PromptOrBlank("y", "n", "r")
			switch outcome {
			case "y":
				// Update config
				for key, value := range vars {
					config.Template.Vars[key] = value
				}
				return true
			case "n":
				return false
			case "r":
				askNext = false
			}
		}
	}
}

func surveyVars(vars *map[string]ConfigVar) {
	// Perform the survey
	for key, value := range *vars {
		value = surveyVar(key, value)
		(*vars)[key] = value
	}
}

func surveyVar(key string, value ConfigVar) (result ConfigVar) {
	fmt.Printf("Name: %s\n", key)
	if len(value.Description) > 0 {
		fmt.Printf("%s\n", value.Description)
	}
	fmt.Printf("Value [default: %s]: ", value.Value)

	var answer string
	fmt.Scanln(&answer)

	if len(answer) > 0 {
		value.Value = answer
	}
	return value
}
