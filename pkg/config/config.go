package config

type Configuration struct {
	OutputFormat string
	Endpoint     string
	ApiKey       string
}

func NewConfiguration(
	outputFormat string,
	endpoint string,
	apiKey string,
) *Configuration {
	return &Configuration{
		OutputFormat: outputFormat,
		Endpoint:     endpoint,
		ApiKey:       apiKey,
	}
}
