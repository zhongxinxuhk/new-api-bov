package service

import (
	relaycommon "github.com/zhongxinxuhk/new-api-bov/relay/common"
	"github.com/zhongxinxuhk/new-api-bov/relaykit/dto"
	"github.com/zhongxinxuhk/new-api-bov/relaykit/relayconvert"
)

func NormalizeCacheCreationSplit(totalTokens int, tokens5m int, tokens1h int) (int, int) {
	return relayconvert.NormalizeCacheCreationSplit(totalTokens, tokens5m, tokens1h)
}

func StreamResponseOpenAI2Claude(openAIResponse *dto.ChatCompletionsStreamResponse, info *relaycommon.RelayInfo) []*dto.ClaudeResponse {
	return relayconvert.StreamResponseOpenAI2Claude(openAIResponse, info)
}

func ResponseOpenAI2Claude(openAIResponse *dto.OpenAITextResponse, info *relaycommon.RelayInfo) *dto.ClaudeResponse {
	return relayconvert.ResponseOpenAI2Claude(openAIResponse, info)
}

func ResponseOpenAI2Gemini(openAIResponse *dto.OpenAITextResponse, info *relaycommon.RelayInfo) *dto.GeminiChatResponse {
	return relayconvert.ResponseOpenAI2Gemini(openAIResponse, info)
}

func StreamResponseOpenAI2Gemini(openAIResponse *dto.ChatCompletionsStreamResponse, info *relaycommon.RelayInfo) *dto.GeminiChatResponse {
	return relayconvert.StreamResponseOpenAI2Gemini(openAIResponse, info)
}
