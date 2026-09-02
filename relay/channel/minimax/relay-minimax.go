package minimax

import (
	"fmt"

	channelconstant "github.com/zhongxinxuhk/new-api-bov/constant"
	relaycommon "github.com/zhongxinxuhk/new-api-bov/relay/common"
	"github.com/zhongxinxuhk/new-api-bov/relay/constant"
	"github.com/zhongxinxuhk/new-api-bov/relaykit/types"
)

func GetRequestURL(info *relaycommon.RelayInfo) (string, error) {
	baseUrl := info.ChannelBaseUrl
	if baseUrl == "" {
		baseUrl = channelconstant.GetChannelBaseURL(channelconstant.ChannelTypeMiniMax)
	}
	switch info.RelayFormat {
	case types.RelayFormatClaude:
		return fmt.Sprintf("%s/anthropic/v1/messages", info.ChannelBaseUrl), nil
	default:
		switch info.RelayMode {
		case constant.RelayModeChatCompletions:
			return fmt.Sprintf("%s/v1/text/chatcompletion_v2", baseUrl), nil
		case constant.RelayModeImagesGenerations:
			return fmt.Sprintf("%s/v1/image_generation", baseUrl), nil
		case constant.RelayModeAudioSpeech:
			return fmt.Sprintf("%s/v1/t2a_v2", baseUrl), nil
		default:
			return "", fmt.Errorf("unsupported relay mode: %d", info.RelayMode)
		}
	}
}
