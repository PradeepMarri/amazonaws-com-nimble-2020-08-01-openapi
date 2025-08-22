package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/amazonnimblestudio/mcp-server/config"
	"github.com/amazonnimblestudio/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func GetstreamingsessionstreamHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		sessionIdVal, ok := args["sessionId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: sessionId"), nil
		}
		sessionId, ok := sessionIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: sessionId"), nil
		}
		streamIdVal, ok := args["streamId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: streamId"), nil
		}
		streamId, ok := streamIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: streamId"), nil
		}
		studioIdVal, ok := args["studioId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: studioId"), nil
		}
		studioId, ok := studioIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: studioId"), nil
		}
		url := fmt.Sprintf("%s/2020-08-01/studios/%s/streaming-sessions/%s/streams/%s", cfg.BaseURL, sessionId, streamId, studioId)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.GetStreamingSessionStreamResponse
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateGetstreamingsessionstreamTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_2020-08-01_studios_studioId_streaming-sessions_sessionId_streams_streamId",
		mcp.WithDescription("<p>Gets a StreamingSessionStream for a streaming session.</p> <p>Invoke this operation to poll the resource after invoking <code>CreateStreamingSessionStream</code>.</p> <p>After the <code>StreamingSessionStream</code> changes to the <code>READY</code> state, the url property will contain a stream to be used with the DCV streaming client.</p>"),
		mcp.WithString("sessionId", mcp.Required(), mcp.Description("The streaming session ID.")),
		mcp.WithString("streamId", mcp.Required(), mcp.Description("The streaming session stream ID.")),
		mcp.WithString("studioId", mcp.Required(), mcp.Description("The studio ID. ")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetstreamingsessionstreamHandler(cfg),
	}
}
