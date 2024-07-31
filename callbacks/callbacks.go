package callbacks

import (
	"context"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/schema"
)

// Handler is the interface that allows for hooking into specific parts of an
// LLM application.
//
// LLMs Chain's Lifecycle
//
//nolint:all
type Handler interface {
	// Handle Text
	HandleText(ctx context.Context, text string)
	HandleLLMStart(ctx context.Context, prompts []string)
	HandleLLMGenerateContentStart(ctx context.Context, ms []llms.MessageContent)
	HandleLLMGenerateContentEnd(ctx context.Context, res *llms.ContentResponse)
	HandleLLMError(ctx context.Context, err error)
	// What To Do When Chain Is Begin.
	HandleChainStart(ctx context.Context, inputs map[string]any)
	// What To Do When Chain Is Over.
	HandleChainEnd(ctx context.Context, outputs map[string]any)
	HandleChainError(ctx context.Context, err error)
	HandleToolStart(ctx context.Context, input string)
	HandleToolEnd(ctx context.Context, output string)
	HandleToolError(ctx context.Context, err error)
	HandleAgentAction(ctx context.Context, action schema.AgentAction)
	HandleAgentFinish(ctx context.Context, finish schema.AgentFinish)
	HandleRetrieverStart(ctx context.Context, query string)
	HandleRetrieverEnd(ctx context.Context, query string, documents []schema.Document)
	HandleStreamingFunc(ctx context.Context, chunk []byte)
}

// HandlerHaver is an interface used to get callbacks handler.
type HandlerHaver interface {
	GetCallbackHandler() Handler
}
