package input

import (
	"github.com/adorigi/checkctl/pkg/config"
	"github.com/charmbracelet/huh"
)

func GetConfigurationFromForm() (*config.Configuration, error) {

	var apiKey string
	var apiEndpoint string
	var utilizationAnalyzerEndpoint string
	var outputFormat string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Enter API key").Value(&apiKey),
			huh.NewInput().Title("Enter API endpoint").Value(&apiEndpoint),
			huh.NewInput().Title("Enter utilization analyzer endpoint").Value(&utilizationAnalyzerEndpoint),
			huh.NewSelect[string]().
				Title("Select output format").
				Options(
					huh.NewOption("JSON", "json"),
				).
				Value(&outputFormat),
		),
	)

	err := form.Run()
	if err != nil {
		return nil, err
	}

	return config.NewConfiguration(
		outputFormat,
		apiEndpoint,
		utilizationAnalyzerEndpoint,
		apiKey,
	), nil

}
