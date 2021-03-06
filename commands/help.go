package commands

import (
	"sort"

	"github.com/AbGuthrie/goquery/config"
	"github.com/AbGuthrie/goquery/models"
	"github.com/AbGuthrie/goquery/utils"

	prompt "github.com/c-bata/go-prompt"
)

func help(api models.GoQueryAPI, config *config.Config, cmdline string) error {
	commandNames := make([]string, 0)
	for k, _ := range CommandMap {
		commandNames = append(commandNames, k)
	}

	sort.Strings(commandNames)
	helpRows := make([]map[string]string, 0)

	for _, commandName := range commandNames {
		helpRows = append(helpRows, map[string]string{
			"command":     commandName,
			"description": CommandMap[commandName].Help(),
		})
	}

	utils.PrettyPrintQueryResults(helpRows, config.PrintMode)
	return nil
}

func helpHelp() string {
	return "Show the help strings for all goquery commands"
}

func helpSuggest(cmdline string) []prompt.Suggest {
	return []prompt.Suggest{}
}
