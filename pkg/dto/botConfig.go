package dto

type BotInfo struct {
	Id                  string                  `json:"id,omitempty"`
	Name                string                  `json:"Name,omitempty"`
	Description         string                  `json:"description,omitempty"`
	Type                string                  `json:"type,omitempty"`
	Timeout             int                     `json:"timeout,omitempty"`
	Active              bool                    `json:"active,omitempty"`
	MaxAttempts         int                     `json:"maxAttempts,omitempty"`
	BotFeatures         BotFeaturesDto          `json:"botFeatures"`
	BotDefaultResponses []BotDefaultResponseDto `json:"botDefaultResponses"`
	Flows               []FlowDto               `json:"flows"`
	Testing             bool                    `json:"testing"`
}

type QuestionTextDto struct {
	Id       string `json:"id,omitempty"`
	LangCode string `json:"langCode"`
	Text     string `json:"text"`
}

type QuestionDto struct {
	Id                   string
	Name                 string
	StepNumber           int
	StepType             string
	MaxAttempts          int
	Url                  string
	ResLabel             string
	ViewStyle            string
	ListTitle            string
	Interactive          bool
	ValidationType       string
	ValidationExample    string
	ValidationFormat     string
	UrlService           bool
	Keywords             []string
	IdCampaign           int
	MimeType             string
	BulletPointFormat    string
	ValidationRegex      string
	SendChoices          bool
	Texts                []QuestionTextDto `json:"texts"`
	Choices              []ChoiceDto       `json:"choices"`
	FlowConfigs          []FlowConfigDto   `json:"flowConfigs"`
	ConnectorId          string
	AuxiliaryBotId       string
	VarsCrm              BotCrmDto              `json:"varsCrm"`
	ConnectorVarsMapping []ConnectorVariableDto `json:"connectorVarsMapping"`
	TypificationId       int64
	ResponseAsList       bool
	ResponseAttToShow    string
	ClosePoll            bool
}
type BotCrmDto struct {
	campaignId    string
	identificador string
	name          string
	lastName      string
	phone         string
	campo6        string
	campo7        string
	campo8        string
	campo9        string
	campo10       string
}

type BotFeaturesDto struct {
	VerifyMaxQueue              bool
	VerifyMaxQueueAtStart       bool
	VerifyOfftime               bool
	VerifyOfftimeAtStart        bool
	SendGoodbyeAutoanswer       bool
	TimeoutMessageActive        bool
	RetryLastMessage            bool
	NlpActive                   bool
	TimeoutAction               string
	MaxAttemptsAction           string
	NlpLowSimilarityQuestionId  string
	TimeoutTransferCampaign     int
	MaxAttemptsTransferCampaign int
	TimeoutTypificationId       int64
	MaxAttemptsTypificationId   int64
	NlpSimilarityConfigDto      *NlpSimilarityConfigDto `json:"nlpSimilarityConfig,omitempty"`
}

type BotDefaultResponseDto struct {
	Id                 string
	LangCode           string
	TimeoutMessage     string
	MaxAttemptsMessage string
	TransferMessage    string
	FinalMessage       string
	RetryMessage       string
}

type FlowDto struct {
	Id        string
	FlowNum   int
	Name      string
	Keywords  []string
	Questions []QuestionDto
}

type ChoiceDto struct {
	Id           string
	Name         string
	Type         string
	Sequence     int
	ImageUrl     string
	ServiceUrl   string
	InternalText string
	ParentId     string
	BulletPoint  string
	Texts        []ChoiceTextDto
	operandA     string
	operandB     string
	operator     string
}

type ConnectorVariableDto struct {
	Id       string
	Variable string
	Value    string
}

type ChoiceTextDto struct {
	Id       string
	LangCode string
	Text     string
	Title    string
	Subtitle string
}

type FlowConfigDto struct {
	Id         string
	ChoiceId   string
	QuestionId string
	FlowId     string
}
