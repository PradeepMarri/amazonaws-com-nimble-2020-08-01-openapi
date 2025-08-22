package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/amazonnimblestudio/mcp-server/config"
	"github.com/amazonnimblestudio/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func CreatelaunchprofileHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		// Create properly typed request body using the generated schema
		var requestBody map[string]interface{}
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/2020-08-01/studios/%s/launch-profiles", cfg.BaseURL, studioId)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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
		var result models.CreateLaunchProfileResponse
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

func CreateCreatelaunchprofileTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_2020-08-01_studios_studioId_launch-profiles",
		mcp.WithDescription("Create a launch profile."),
		mcp.WithString("X-Amz-Client-Token", mcp.Description("Unique, case-sensitive identifier that you provide to ensure the idempotency of the request. If you don’t specify a client token, the Amazon Web Services SDK automatically generates a client token and uses it for the request to ensure idempotency.")),
		mcp.WithString("studioId", mcp.Required(), mcp.Description("The studio ID. ")),
		mcp.WithString("description", mcp.Description("Input parameter: A human-readable description of the launch profile.")),
		mcp.WithArray("ec2SubnetIds", mcp.Required(), mcp.Description("Input parameter: Specifies the IDs of the EC2 subnets where streaming sessions will be accessible from. These subnets must support the specified instance types. ")),
		mcp.WithArray("launchProfileProtocolVersions", mcp.Required(), mcp.Description("Input parameter: The version number of the protocol that is used by the launch profile. The only valid version is \"2021-03-31\".")),
		mcp.WithString("name", mcp.Required(), mcp.Description("Input parameter: The name for the launch profile.")),
		mcp.WithObject("streamConfiguration", mcp.Required(), mcp.Description("Input parameter: Configuration for streaming workstations created using this launch profile.")),
		mcp.WithArray("studioComponentIds", mcp.Required(), mcp.Description("Input parameter: Unique identifiers for a collection of studio components that can be used with this launch profile.")),
		mcp.WithObject("tags", mcp.Description("Input parameter: A collection of labels, in the form of key-value pairs, that apply to this resource.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CreatelaunchprofileHandler(cfg),
	}
}
