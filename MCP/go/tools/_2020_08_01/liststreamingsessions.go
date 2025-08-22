package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/amazonnimblestudio/mcp-server/config"
	"github.com/amazonnimblestudio/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func ListstreamingsessionsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		studioIdVal, ok := args["studioId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: studioId"), nil
		}
		studioId, ok := studioIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: studioId"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["createdBy"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("createdBy=%v", val))
		}
		if val, ok := args["nextToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("nextToken=%v", val))
		}
		if val, ok := args["ownedBy"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ownedBy=%v", val))
		}
		if val, ok := args["sessionIds"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sessionIds=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/2020-08-01/studios/%s/streaming-sessions%s", cfg.BaseURL, studioId, queryString)
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
		var result models.ListStreamingSessionsResponse
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

func CreateListstreamingsessionsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_2020-08-01_studios_studioId_streaming-sessions",
		mcp.WithDescription("Lists the streaming sessions in a studio."),
		mcp.WithString("createdBy", mcp.Description("Filters the request to streaming sessions created by the given user.")),
		mcp.WithString("nextToken", mcp.Description("The token for the next set of results, or null if there are no more results.")),
		mcp.WithString("ownedBy", mcp.Description("Filters the request to streaming session owned by the given user")),
		mcp.WithString("sessionIds", mcp.Description("Filters the request to only the provided session IDs.")),
		mcp.WithString("studioId", mcp.Required(), mcp.Description("The studio ID. ")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ListstreamingsessionsHandler(cfg),
	}
}
