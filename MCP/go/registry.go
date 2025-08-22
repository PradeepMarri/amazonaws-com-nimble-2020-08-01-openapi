package main

import (
	"github.com/amazonnimblestudio/mcp-server/config"
	"github.com/amazonnimblestudio/mcp-server/models"
	tools__2020_08_01 "github.com/amazonnimblestudio/mcp-server/tools/_2020_08_01"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools__2020_08_01.CreateDeletestudiomemberTool(cfg),
		tools__2020_08_01.CreateGetstudiomemberTool(cfg),
		tools__2020_08_01.CreateListstudiosTool(cfg),
		tools__2020_08_01.CreateCreatestudioTool(cfg),
		tools__2020_08_01.CreateDeletestreamingimageTool(cfg),
		tools__2020_08_01.CreateGetstreamingimageTool(cfg),
		tools__2020_08_01.CreateUpdatestreamingimageTool(cfg),
		tools__2020_08_01.CreateCreatestreamingsessionstreamTool(cfg),
		tools__2020_08_01.CreateDeletelaunchprofileTool(cfg),
		tools__2020_08_01.CreateGetlaunchprofileTool(cfg),
		tools__2020_08_01.CreateUpdatelaunchprofileTool(cfg),
		tools__2020_08_01.CreateListlaunchprofilemembersTool(cfg),
		tools__2020_08_01.CreatePutlaunchprofilemembersTool(cfg),
		tools__2020_08_01.CreateListstreamingsessionsTool(cfg),
		tools__2020_08_01.CreateCreatestreamingsessionTool(cfg),
		tools__2020_08_01.CreateStopstreamingsessionTool(cfg),
		tools__2020_08_01.CreateListtagsforresourceTool(cfg),
		tools__2020_08_01.CreateTagresourceTool(cfg),
		tools__2020_08_01.CreateGetlaunchprofiledetailsTool(cfg),
		tools__2020_08_01.CreateUntagresourceTool(cfg),
		tools__2020_08_01.CreateCreatestreamingimageTool(cfg),
		tools__2020_08_01.CreateListstreamingimagesTool(cfg),
		tools__2020_08_01.CreateGetstreamingsessionbackupTool(cfg),
		tools__2020_08_01.CreateListstreamingsessionbackupsTool(cfg),
		tools__2020_08_01.CreateUpdatelaunchprofilememberTool(cfg),
		tools__2020_08_01.CreateDeletelaunchprofilememberTool(cfg),
		tools__2020_08_01.CreateGetlaunchprofilememberTool(cfg),
		tools__2020_08_01.CreateDeletestudiocomponentTool(cfg),
		tools__2020_08_01.CreateGetstudiocomponentTool(cfg),
		tools__2020_08_01.CreateUpdatestudiocomponentTool(cfg),
		tools__2020_08_01.CreateGetstreamingsessionstreamTool(cfg),
		tools__2020_08_01.CreateListstudiocomponentsTool(cfg),
		tools__2020_08_01.CreateCreatestudiocomponentTool(cfg),
		tools__2020_08_01.CreateListeulaacceptancesTool(cfg),
		tools__2020_08_01.CreateAccepteulasTool(cfg),
		tools__2020_08_01.CreateStartstreamingsessionTool(cfg),
		tools__2020_08_01.CreateGeteulaTool(cfg),
		tools__2020_08_01.CreateUpdatestudioTool(cfg),
		tools__2020_08_01.CreateDeletestudioTool(cfg),
		tools__2020_08_01.CreateGetstudioTool(cfg),
		tools__2020_08_01.CreateGetlaunchprofileinitializationTool(cfg),
		tools__2020_08_01.CreateStartstudiossoconfigurationrepairTool(cfg),
		tools__2020_08_01.CreateListeulasTool(cfg),
		tools__2020_08_01.CreateListlaunchprofilesTool(cfg),
		tools__2020_08_01.CreateCreatelaunchprofileTool(cfg),
		tools__2020_08_01.CreateListstudiomembersTool(cfg),
		tools__2020_08_01.CreatePutstudiomembersTool(cfg),
		tools__2020_08_01.CreateGetstreamingsessionTool(cfg),
		tools__2020_08_01.CreateDeletestreamingsessionTool(cfg),
	}
}
