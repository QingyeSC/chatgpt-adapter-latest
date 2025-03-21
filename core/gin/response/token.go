package response

import (
    "chatgpt-adapter/core/logger"
    "github.com/pkoukk/tiktoken-go"
)

func CalcTokens(content string) int {
    encoder, err := tiktoken.GetEncoding("cl100k_base") // GPT-4 使用的编码
    if err != nil {
        logger.Error(err)
        return 0
    }
    
    tokens := encoder.Encode(content, nil, nil)
    return len(tokens)
}

func CalcUsageTokens(content string, previousTokens int) map[string]interface{} {
    tokens := CalcTokens(content)
    return map[string]interface{}{
        "prompt_tokens":     previousTokens,
        "completion_tokens": tokens,
        "total_tokens":      previousTokens + tokens,
    }
}
