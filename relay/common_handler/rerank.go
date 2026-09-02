package common_handler

import (
	"io"
	"net/http"

	"github.com/zhongxinxuhk/new-api-bov/common"
	"github.com/zhongxinxuhk/new-api-bov/constant"
	"github.com/zhongxinxuhk/new-api-bov/logger"
	"github.com/zhongxinxuhk/new-api-bov/relay/channel/xinference"
	relaycommon "github.com/zhongxinxuhk/new-api-bov/relay/common"
	"github.com/zhongxinxuhk/new-api-bov/relaykit/dto"
	"github.com/zhongxinxuhk/new-api-bov/relaykit/types"
	"github.com/zhongxinxuhk/new-api-bov/service"

	"github.com/gin-gonic/gin"
)

func RerankHandler(c *gin.Context, info *relaycommon.RelayInfo, resp *http.Response) (*dto.Usage, *types.NewAPIError) {
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, types.NewOpenAIError(err, types.ErrorCodeReadResponseBodyFailed, http.StatusInternalServerError)
	}
	service.CloseResponseBodyGracefully(resp)
	logger.LogDebug(c, "reranker response body: %s", responseBody)
	var jinaResp dto.RerankResponse
	if info.ChannelType == constant.ChannelTypeXinference {
		var xinRerankResponse xinference.XinRerankResponse
		err = common.Unmarshal(responseBody, &xinRerankResponse)
		if err != nil {
			return nil, types.NewOpenAIError(err, types.ErrorCodeBadResponseBody, http.StatusInternalServerError)
		}
		jinaRespResults := make([]dto.RerankResponseResult, len(xinRerankResponse.Results))
		for i, result := range xinRerankResponse.Results {
			respResult := dto.RerankResponseResult{
				Index:          result.Index,
				RelevanceScore: result.RelevanceScore,
			}
			if info.ReturnDocuments {
				var document any
				if result.Document != nil {
					if doc, ok := result.Document.(string); ok {
						if doc == "" {
							document = info.Documents[result.Index]
						} else {
							document = doc
						}
					} else {
						document = result.Document
					}
				}
				respResult.Document = document
			}
			jinaRespResults[i] = respResult
		}
		jinaResp = dto.RerankResponse{
			Results: jinaRespResults,
			Usage: dto.Usage{
				PromptTokens: info.GetEstimatePromptTokens(),
				TotalTokens:  info.GetEstimatePromptTokens(),
			},
		}
	} else {
		err = common.Unmarshal(responseBody, &jinaResp)
		if err != nil {
			return nil, types.NewOpenAIError(err, types.ErrorCodeBadResponseBody, http.StatusInternalServerError)
		}
		jinaResp.Usage.PromptTokens = jinaResp.Usage.TotalTokens
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(http.StatusOK, jinaResp)
	return &jinaResp.Usage, nil
}
