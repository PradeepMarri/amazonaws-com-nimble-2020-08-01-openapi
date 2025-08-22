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

func ListstreamingsessionbackupsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["nextToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("nextToken=%v", val))
		}
		if val, ok := args["ownedBy"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ownedBy=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/2020-08-01/studios/%s/streaming-session-backups%s", cfg.BaseURL, studioId, queryString)
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
		var result models.ListStreamingSessionBackupsResponse
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

func CreateListstreamingsessionbackupsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_2020-08-01_studios_studioId_streaming-session-backups",
		mcp.WithDescription("Lists the backups of a streaming session in a studio."),
		mcp.WithString("nextToken", mcp.Description("The token for the next set of results, or null if there are no more results.")),
		mcp.WithString("ownedBy", mcp.Description("The user ID of the user that owns the streaming session.")),
		mcp.WithString("studioId", mcp.Required(), mcp.Description("The studio ID. ")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ListstreamingsessionbackupsHandler(cfg),
	}
}
