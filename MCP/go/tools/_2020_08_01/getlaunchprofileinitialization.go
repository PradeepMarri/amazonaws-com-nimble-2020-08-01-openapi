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

func GetlaunchprofileinitializationHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		launchProfileIdVal, ok := args["launchProfileId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: launchProfileId"), nil
		}
		launchProfileId, ok := launchProfileIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: launchProfileId"), nil
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
		if val, ok := args["launchProfileProtocolVersions"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("launchProfileProtocolVersions=%v", val))
		}
		if val, ok := args["launchPurpose"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("launchPurpose=%v", val))
		}
		if val, ok := args["platform"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("platform=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/2020-08-01/studios/%s/launch-profiles/%s/init#launchProfileProtocolVersions&launchPurpose&platform%s", cfg.BaseURL, launchProfileId, studioId, queryString)
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
		var result models.GetLaunchProfileInitializationResponse
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

func CreateGetlaunchprofileinitializationTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_2020-08-01_studios_studioId_launch-profiles_launchProfileId_init#launchProfileProtocolVersions&launchPurpose&platform",
		mcp.WithDescription("Get a launch profile initialization."),
		mcp.WithString("launchProfileId", mcp.Required(), mcp.Description("The ID of the launch profile used to control access from the streaming session.")),
		mcp.WithArray("launchProfileProtocolVersions", mcp.Required(), mcp.Description("The launch profile protocol versions supported by the client.")),
		mcp.WithString("launchPurpose", mcp.Required(), mcp.Description("The launch purpose.")),
		mcp.WithString("platform", mcp.Required(), mcp.Description("The platform where this Launch Profile will be used, either Windows or Linux.")),
		mcp.WithString("studioId", mcp.Required(), mcp.Description("The studio ID. ")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetlaunchprofileinitializationHandler(cfg),
	}
}
