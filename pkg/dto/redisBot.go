package dto

type RedisBot struct {
	Token     string  `json:"token,omitempty"`
	BotConfig BotInfo `json:"bot_config"`
}
