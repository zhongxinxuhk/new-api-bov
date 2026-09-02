// Package reasoning re-exports the pure model-name effort-suffix helpers,
// which moved to the conversion kit (relaykit/relayconvert/reasoning) as part
// of the relaykit extraction. Host code keeps importing this path unchanged.
package reasoning

import (
	kitreasoning "github.com/zhongxinxuhk/new-api-bov/relaykit/relayconvert/reasoning"
	"github.com/zhongxinxuhk/new-api-bov/setting/model_setting"
)

var (
	EffortSuffixes           = kitreasoning.EffortSuffixes
	OpenAIEffortSuffixes     = kitreasoning.OpenAIEffortSuffixes
	DeepSeekV4EffortSuffixes = kitreasoning.DeepSeekV4EffortSuffixes
)

var (
	TrimEffortSuffixWithSuffixes  = kitreasoning.TrimEffortSuffixWithSuffixes
	ParseDeepSeekV4ThinkingSuffix = kitreasoning.ParseDeepSeekV4ThinkingSuffix
	TrimGeminiThinkingSuffix      = kitreasoning.TrimGeminiThinkingSuffix
)

// ParseOpenAIReasoningEffortFromModelSuffix applies the host effort-tail
// whitelist so real model IDs such as qwen-max are not treated as aliases.
func ParseOpenAIReasoningEffortFromModelSuffix(modelName string) (string, string) {
	return kitreasoning.ParseOpenAIReasoningEffortFromModelSuffix(modelName, model_setting.ShouldPreserveEffortTail)
}
