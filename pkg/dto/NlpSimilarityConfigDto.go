package dto

type NlpSimilarityConfigDto struct {
	UserApiId                 int64                         `json:"userApiId"`
	NlpMediumSimilarityConfig *NlpMediumSimilarityConfigDto `json:"nlpMediumSimilarityConfig,omitempty"`
	NlpLowSimilarityConfig    *NlpLowSimilarityConfigDto    `json:"nlpLowSimilarityConfig,omitempty"`
}

type NlpMediumSimilarityConfigDto struct {
	ExitOptionActive bool   `json:"exitOptionActive"`
	OptionListTitle  string `json:"optionListTitle"`
	ExitOptionMsg    string `json:"exitOptionMsg"`
	ExitResponseMsg  string `json:"exitResponseMsg"`
}

type NlpLowSimilarityConfigDto struct {
	Retries  *int8  `json:"retries"`
	RetryMsg string `json:"retryMsg"`
	FinalMsg string `json:"finalMsg"`
}
