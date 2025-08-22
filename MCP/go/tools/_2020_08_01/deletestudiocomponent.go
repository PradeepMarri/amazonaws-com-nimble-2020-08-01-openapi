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

func DeletestudiocomponentHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		studioComponentIdVal, ok := args["studioComponentId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: studioComponentId"), nil
		}
		studioComponentId, ok := studioComponentIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: studioComponentId"), nil
		}
		studioIdVal, ok := args["studioId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: studioId"), nil
		}
		studioId, ok := studioIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: studioId"), nil
		}
		url := fmt.Sprintf("%s/2020-08-01/studios/%s/studio-components/%s", cfg.BaseURL, studioComponentId, studioId)
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")
		if val, ok := args["X-Amz-Client-Token"]; ok {
			req.Header.Set("X-Amz-Client-Token", fmt.Sprintf("%v", val))
		}

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
		var result models.DeleteStudioComponentResponse
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

func CreateDeletestudiocomponentTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_2020-08-01_studios_studioId_studio-components_studioComponentId",
		mcp.WithDescription("Deletes a studio component resource."),
		mcp.WithString("X-Amz-Client-Token", mcp.Description("Unique, case-sensitive identifier that you provide to ensure the idempotency of the request. If you don’t specify a client token, the Amazon Web Services SDK automatically generates a client token and uses it for the request to ensure idempotency.")),
		mcp.WithString("studioComponentId", mcp.Required(), mcp.Description("The studio component ID.")),
		mcp.WithString("studioId", mcp.Required(), mcp.Description("The studio ID. ")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DeletestudiocomponentHandler(cfg),
	}
}
