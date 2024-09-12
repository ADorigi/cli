package config

type Configuration struct {
	OutputFormat string `json:"outputformat"`
	Endpoint     string `json:"endpoint"`
	ApiKey       string `json:apikey`
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
