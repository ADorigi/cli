package config

type Configuration struct {
	OutputFormat                string            `json:"output_format"`
	ApiEndpoint                 string            `json:"api_endpoint"`
	UtilizationAnalyzerEndpoint string            `json:"utilization_analyzer_endpoint"`
	ApiKey                      string            `json:"api_key"`
	Integrations                map[string]string `json:"integrations"`
}

func NewConfiguration(
	outputFormat string,
	apiEndpoint string,
	utilizationAnalyzerEndpoint string,
	apiKey string,
) *Configuration {
	return &Configuration{
		OutputFormat:                outputFormat,
		ApiEndpoint:                 apiEndpoint,
		UtilizationAnalyzerEndpoint: utilizationAnalyzerEndpoint,
		ApiKey:                      apiKey,
		Integrations:                map[string]string{},
	}
}
