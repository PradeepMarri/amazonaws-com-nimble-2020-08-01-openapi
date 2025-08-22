package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// StudioEncryptionConfiguration represents the StudioEncryptionConfiguration schema from the OpenAPI specification
type StudioEncryptionConfiguration struct {
	Keyarn interface{} `json:"keyArn,omitempty"`
	Keytype interface{} `json:"keyType"`
}

// GetLaunchProfileResponse represents the GetLaunchProfileResponse schema from the OpenAPI specification
type GetLaunchProfileResponse struct {
	Launchprofile interface{} `json:"launchProfile,omitempty"`
}

// NewLaunchProfileMember represents the NewLaunchProfileMember schema from the OpenAPI specification
type NewLaunchProfileMember struct {
	Principalid interface{} `json:"principalId"`
	Persona interface{} `json:"persona"`
}

// UpdateLaunchProfileMemberResponse represents the UpdateLaunchProfileMemberResponse schema from the OpenAPI specification
type UpdateLaunchProfileMemberResponse struct {
	Member interface{} `json:"member,omitempty"`
}

// GetEulaRequest represents the GetEulaRequest schema from the OpenAPI specification
type GetEulaRequest struct {
}

// VolumeConfiguration represents the VolumeConfiguration schema from the OpenAPI specification
type VolumeConfiguration struct {
	Size interface{} `json:"size,omitempty"`
	Throughput interface{} `json:"throughput,omitempty"`
	Iops interface{} `json:"iops,omitempty"`
}

// UpdateStudioRequest represents the UpdateStudioRequest schema from the OpenAPI specification
type UpdateStudioRequest struct {
	Displayname interface{} `json:"displayName,omitempty"`
	Userrolearn interface{} `json:"userRoleArn,omitempty"`
	Adminrolearn interface{} `json:"adminRoleArn,omitempty"`
}

// GetStudioRequest represents the GetStudioRequest schema from the OpenAPI specification
type GetStudioRequest struct {
}

// PutStudioMembersRequest represents the PutStudioMembersRequest schema from the OpenAPI specification
type PutStudioMembersRequest struct {
	Identitystoreid interface{} `json:"identityStoreId"`
	Members interface{} `json:"members"`
}

// ListStreamingSessionsRequest represents the ListStreamingSessionsRequest schema from the OpenAPI specification
type ListStreamingSessionsRequest struct {
}

// ListStreamingSessionBackupsResponse represents the ListStreamingSessionBackupsResponse schema from the OpenAPI specification
type ListStreamingSessionBackupsResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Streamingsessionbackups interface{} `json:"streamingSessionBackups,omitempty"`
}

// CreateStudioResponse represents the CreateStudioResponse schema from the OpenAPI specification
type CreateStudioResponse struct {
	Studio interface{} `json:"studio,omitempty"`
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Tags interface{} `json:"tags,omitempty"`
}

// GetLaunchProfileDetailsRequest represents the GetLaunchProfileDetailsRequest schema from the OpenAPI specification
type GetLaunchProfileDetailsRequest struct {
}

// ListLaunchProfilesRequest represents the ListLaunchProfilesRequest schema from the OpenAPI specification
type ListLaunchProfilesRequest struct {
}

// CreateStreamingImageResponse represents the CreateStreamingImageResponse schema from the OpenAPI specification
type CreateStreamingImageResponse struct {
	Streamingimage interface{} `json:"streamingImage,omitempty"`
}

// LaunchProfile represents the LaunchProfile schema from the OpenAPI specification
type LaunchProfile struct {
	Arn interface{} `json:"arn,omitempty"`
	Launchprofileprotocolversions interface{} `json:"launchProfileProtocolVersions,omitempty"`
	Ec2subnetids interface{} `json:"ec2SubnetIds,omitempty"`
	Statuscode interface{} `json:"statusCode,omitempty"`
	Streamconfiguration interface{} `json:"streamConfiguration,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Statusmessage interface{} `json:"statusMessage,omitempty"`
	Launchprofileid interface{} `json:"launchProfileId,omitempty"`
	State interface{} `json:"state,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Studiocomponentids interface{} `json:"studioComponentIds,omitempty"`
	Updatedby interface{} `json:"updatedBy,omitempty"`
	Validationresults interface{} `json:"validationResults,omitempty"`
	Createdby interface{} `json:"createdBy,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Updatedat interface{} `json:"updatedAt,omitempty"`
}

// StopStreamingSessionRequest represents the StopStreamingSessionRequest schema from the OpenAPI specification
type StopStreamingSessionRequest struct {
	Volumeretentionmode interface{} `json:"volumeRetentionMode,omitempty"`
}

// ListEulasRequest represents the ListEulasRequest schema from the OpenAPI specification
type ListEulasRequest struct {
}

// EulaAcceptance represents the EulaAcceptance schema from the OpenAPI specification
type EulaAcceptance struct {
	Eulaid interface{} `json:"eulaId,omitempty"`
	Acceptedat interface{} `json:"acceptedAt,omitempty"`
	Acceptedby interface{} `json:"acceptedBy,omitempty"`
	Accepteeid interface{} `json:"accepteeId,omitempty"`
	Eulaacceptanceid interface{} `json:"eulaAcceptanceId,omitempty"`
}

// PutLaunchProfileMembersRequest represents the PutLaunchProfileMembersRequest schema from the OpenAPI specification
type PutLaunchProfileMembersRequest struct {
	Members interface{} `json:"members"`
	Identitystoreid interface{} `json:"identityStoreId"`
}

// GetStreamingImageRequest represents the GetStreamingImageRequest schema from the OpenAPI specification
type GetStreamingImageRequest struct {
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// AcceptEulasRequest represents the AcceptEulasRequest schema from the OpenAPI specification
type AcceptEulasRequest struct {
	Eulaids interface{} `json:"eulaIds,omitempty"`
}

// CreateLaunchProfileRequest represents the CreateLaunchProfileRequest schema from the OpenAPI specification
type CreateLaunchProfileRequest struct {
	Launchprofileprotocolversions interface{} `json:"launchProfileProtocolVersions"`
	Name interface{} `json:"name"`
	Streamconfiguration interface{} `json:"streamConfiguration"`
	Studiocomponentids interface{} `json:"studioComponentIds"`
	Tags interface{} `json:"tags,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Ec2subnetids interface{} `json:"ec2SubnetIds"`
}

// GetStudioComponentRequest represents the GetStudioComponentRequest schema from the OpenAPI specification
type GetStudioComponentRequest struct {
}

// GetEulaResponse represents the GetEulaResponse schema from the OpenAPI specification
type GetEulaResponse struct {
	Eula interface{} `json:"eula,omitempty"`
}

// DeleteLaunchProfileRequest represents the DeleteLaunchProfileRequest schema from the OpenAPI specification
type DeleteLaunchProfileRequest struct {
}

// ListEulaAcceptancesRequest represents the ListEulaAcceptancesRequest schema from the OpenAPI specification
type ListEulaAcceptancesRequest struct {
}

// StudioComponent represents the StudioComponent schema from the OpenAPI specification
type StudioComponent struct {
	Description interface{} `json:"description,omitempty"`
	Ec2securitygroupids interface{} `json:"ec2SecurityGroupIds,omitempty"`
	Statusmessage interface{} `json:"statusMessage,omitempty"`
	Scriptparameters interface{} `json:"scriptParameters,omitempty"`
	State interface{} `json:"state,omitempty"`
	Subtype interface{} `json:"subtype,omitempty"`
	TypeField interface{} `json:"type,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Configuration interface{} `json:"configuration,omitempty"`
	Statuscode interface{} `json:"statusCode,omitempty"`
	Updatedby interface{} `json:"updatedBy,omitempty"`
	Runtimerolearn interface{} `json:"runtimeRoleArn,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Createdby interface{} `json:"createdBy,omitempty"`
	Initializationscripts interface{} `json:"initializationScripts,omitempty"`
	Studiocomponentid interface{} `json:"studioComponentId,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Secureinitializationrolearn interface{} `json:"secureInitializationRoleArn,omitempty"`
	Updatedat interface{} `json:"updatedAt,omitempty"`
}

// StreamingSession represents the StreamingSession schema from the OpenAPI specification
type StreamingSession struct {
	Statuscode interface{} `json:"statusCode,omitempty"`
	Sessionpersistencemode interface{} `json:"sessionPersistenceMode,omitempty"`
	Startedby interface{} `json:"startedBy,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Ec2instancetype interface{} `json:"ec2InstanceType,omitempty"`
	Streamingimageid interface{} `json:"streamingImageId,omitempty"`
	Volumeconfiguration interface{} `json:"volumeConfiguration,omitempty"`
	Launchprofileid interface{} `json:"launchProfileId,omitempty"`
	Startedat interface{} `json:"startedAt,omitempty"`
	Maxbackupstoretain interface{} `json:"maxBackupsToRetain,omitempty"`
	Stoppedby interface{} `json:"stoppedBy,omitempty"`
	Stoppedat interface{} `json:"stoppedAt,omitempty"`
	Stopat interface{} `json:"stopAt,omitempty"`
	State interface{} `json:"state,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Updatedby interface{} `json:"updatedBy,omitempty"`
	Createdby interface{} `json:"createdBy,omitempty"`
	Ownedby interface{} `json:"ownedBy,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Sessionid interface{} `json:"sessionId,omitempty"`
	Statusmessage interface{} `json:"statusMessage,omitempty"`
	Automaticterminationmode interface{} `json:"automaticTerminationMode,omitempty"`
	Startedfrombackupid interface{} `json:"startedFromBackupId,omitempty"`
	Updatedat interface{} `json:"updatedAt,omitempty"`
	Backupmode interface{} `json:"backupMode,omitempty"`
	Terminateat interface{} `json:"terminateAt,omitempty"`
	Volumeretentionmode interface{} `json:"volumeRetentionMode,omitempty"`
}

// CreateStreamingSessionResponse represents the CreateStreamingSessionResponse schema from the OpenAPI specification
type CreateStreamingSessionResponse struct {
	Session interface{} `json:"session,omitempty"`
}

// ListLaunchProfileMembersRequest represents the ListLaunchProfileMembersRequest schema from the OpenAPI specification
type ListLaunchProfileMembersRequest struct {
}

// DeleteStreamingImageRequest represents the DeleteStreamingImageRequest schema from the OpenAPI specification
type DeleteStreamingImageRequest struct {
}

// GetStreamingImageResponse represents the GetStreamingImageResponse schema from the OpenAPI specification
type GetStreamingImageResponse struct {
	Streamingimage interface{} `json:"streamingImage,omitempty"`
}

// ValidationResult represents the ValidationResult schema from the OpenAPI specification
type ValidationResult struct {
	State interface{} `json:"state"`
	Statuscode interface{} `json:"statusCode"`
	Statusmessage interface{} `json:"statusMessage"`
	TypeField interface{} `json:"type"`
}

// UpdateStreamingImageResponse represents the UpdateStreamingImageResponse schema from the OpenAPI specification
type UpdateStreamingImageResponse struct {
	Streamingimage StreamingImage `json:"streamingImage,omitempty"` // <p>Represents a streaming image resource.</p> <p>Streaming images are used by studio users to select which operating system and software they want to use in a Nimble Studio streaming session.</p> <p>Amazon provides a number of streaming images that include popular 3rd-party software.</p> <p>You can create your own streaming images using an Amazon EC2 machine image that you create for this purpose. You can also include software that your users require.</p>
}

// StreamingSessionBackup represents the StreamingSessionBackup schema from the OpenAPI specification
type StreamingSessionBackup struct {
	State string `json:"state,omitempty"` // The streaming session state.
	Backupid interface{} `json:"backupId,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Ownedby interface{} `json:"ownedBy,omitempty"`
	Statuscode interface{} `json:"statusCode,omitempty"`
	Statusmessage interface{} `json:"statusMessage,omitempty"`
	Launchprofileid interface{} `json:"launchProfileId,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Sessionid interface{} `json:"sessionId,omitempty"`
}

// LicenseServiceConfiguration represents the LicenseServiceConfiguration schema from the OpenAPI specification
type LicenseServiceConfiguration struct {
	Endpoint interface{} `json:"endpoint,omitempty"`
}

// ListLaunchProfilesResponse represents the ListLaunchProfilesResponse schema from the OpenAPI specification
type ListLaunchProfilesResponse struct {
	Launchprofiles interface{} `json:"launchProfiles,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// DeleteStudioComponentRequest represents the DeleteStudioComponentRequest schema from the OpenAPI specification
type DeleteStudioComponentRequest struct {
}

// DeleteStudioMemberResponse represents the DeleteStudioMemberResponse schema from the OpenAPI specification
type DeleteStudioMemberResponse struct {
}

// StartStreamingSessionRequest represents the StartStreamingSessionRequest schema from the OpenAPI specification
type StartStreamingSessionRequest struct {
	Backupid interface{} `json:"backupId,omitempty"`
}

// GetLaunchProfileMemberRequest represents the GetLaunchProfileMemberRequest schema from the OpenAPI specification
type GetLaunchProfileMemberRequest struct {
}

// CreateStudioRequest represents the CreateStudioRequest schema from the OpenAPI specification
type CreateStudioRequest struct {
	Studioencryptionconfiguration interface{} `json:"studioEncryptionConfiguration,omitempty"`
	Studioname interface{} `json:"studioName"`
	Tags interface{} `json:"tags,omitempty"`
	Userrolearn interface{} `json:"userRoleArn"`
	Adminrolearn interface{} `json:"adminRoleArn"`
	Displayname interface{} `json:"displayName"`
}

// GetLaunchProfileInitializationResponse represents the GetLaunchProfileInitializationResponse schema from the OpenAPI specification
type GetLaunchProfileInitializationResponse struct {
	Launchprofileinitialization interface{} `json:"launchProfileInitialization,omitempty"`
}

// LaunchProfileInitializationScript represents the LaunchProfileInitializationScript schema from the OpenAPI specification
type LaunchProfileInitializationScript struct {
	Studiocomponentid interface{} `json:"studioComponentId,omitempty"`
	Studiocomponentname interface{} `json:"studioComponentName,omitempty"`
	Runtimerolearn interface{} `json:"runtimeRoleArn,omitempty"`
	Script interface{} `json:"script,omitempty"`
	Secureinitializationrolearn interface{} `json:"secureInitializationRoleArn,omitempty"`
}

// UpdateStudioComponentResponse represents the UpdateStudioComponentResponse schema from the OpenAPI specification
type UpdateStudioComponentResponse struct {
	Studiocomponent interface{} `json:"studioComponent,omitempty"`
}

// UpdateStreamingImageRequest represents the UpdateStreamingImageRequest schema from the OpenAPI specification
type UpdateStreamingImageRequest struct {
	Description interface{} `json:"description,omitempty"`
	Name interface{} `json:"name,omitempty"`
}

// ListStreamingSessionBackupsRequest represents the ListStreamingSessionBackupsRequest schema from the OpenAPI specification
type ListStreamingSessionBackupsRequest struct {
}

// StudioComponentSummary represents the StudioComponentSummary schema from the OpenAPI specification
type StudioComponentSummary struct {
	TypeField interface{} `json:"type,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Subtype interface{} `json:"subtype,omitempty"`
	Updatedat interface{} `json:"updatedAt,omitempty"`
	Updatedby interface{} `json:"updatedBy,omitempty"`
	Createdby interface{} `json:"createdBy,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Studiocomponentid interface{} `json:"studioComponentId,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
}

// DeleteLaunchProfileResponse represents the DeleteLaunchProfileResponse schema from the OpenAPI specification
type DeleteLaunchProfileResponse struct {
	Launchprofile interface{} `json:"launchProfile,omitempty"`
}

// GetLaunchProfileDetailsResponse represents the GetLaunchProfileDetailsResponse schema from the OpenAPI specification
type GetLaunchProfileDetailsResponse struct {
	Streamingimages interface{} `json:"streamingImages,omitempty"`
	Studiocomponentsummaries interface{} `json:"studioComponentSummaries,omitempty"`
	Launchprofile interface{} `json:"launchProfile,omitempty"`
}

// DeleteStreamingSessionResponse represents the DeleteStreamingSessionResponse schema from the OpenAPI specification
type DeleteStreamingSessionResponse struct {
	Session interface{} `json:"session,omitempty"`
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
}

// UntagResourceResponse represents the UntagResourceResponse schema from the OpenAPI specification
type UntagResourceResponse struct {
}

// GetStudioComponentResponse represents the GetStudioComponentResponse schema from the OpenAPI specification
type GetStudioComponentResponse struct {
	Studiocomponent interface{} `json:"studioComponent,omitempty"`
}

// StartStudioSSOConfigurationRepairRequest represents the StartStudioSSOConfigurationRepairRequest schema from the OpenAPI specification
type StartStudioSSOConfigurationRepairRequest struct {
}

// UpdateLaunchProfileRequest represents the UpdateLaunchProfileRequest schema from the OpenAPI specification
type UpdateLaunchProfileRequest struct {
	Studiocomponentids interface{} `json:"studioComponentIds,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Launchprofileprotocolversions interface{} `json:"launchProfileProtocolVersions,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Streamconfiguration interface{} `json:"streamConfiguration,omitempty"`
}

// UpdateStudioResponse represents the UpdateStudioResponse schema from the OpenAPI specification
type UpdateStudioResponse struct {
	Studio interface{} `json:"studio"`
}

// GetStreamingSessionResponse represents the GetStreamingSessionResponse schema from the OpenAPI specification
type GetStreamingSessionResponse struct {
	Session interface{} `json:"session,omitempty"`
}

// GetStudioResponse represents the GetStudioResponse schema from the OpenAPI specification
type GetStudioResponse struct {
	Studio interface{} `json:"studio"`
}

// GetStreamingSessionStreamRequest represents the GetStreamingSessionStreamRequest schema from the OpenAPI specification
type GetStreamingSessionStreamRequest struct {
}

// UpdateStudioComponentRequest represents the UpdateStudioComponentRequest schema from the OpenAPI specification
type UpdateStudioComponentRequest struct {
	Name interface{} `json:"name,omitempty"`
	Subtype interface{} `json:"subtype,omitempty"`
	Ec2securitygroupids interface{} `json:"ec2SecurityGroupIds,omitempty"`
	Initializationscripts interface{} `json:"initializationScripts,omitempty"`
	Scriptparameters interface{} `json:"scriptParameters,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Runtimerolearn interface{} `json:"runtimeRoleArn,omitempty"`
	Secureinitializationrolearn interface{} `json:"secureInitializationRoleArn,omitempty"`
	TypeField interface{} `json:"type,omitempty"`
	Configuration interface{} `json:"configuration,omitempty"`
}

// StreamingSessionStream represents the StreamingSessionStream schema from the OpenAPI specification
type StreamingSessionStream struct {
	Createdby interface{} `json:"createdBy,omitempty"`
	Expiresat interface{} `json:"expiresAt,omitempty"`
	Ownedby interface{} `json:"ownedBy,omitempty"`
	State interface{} `json:"state,omitempty"`
	Statuscode interface{} `json:"statusCode,omitempty"`
	Streamid interface{} `json:"streamId,omitempty"`
	Url interface{} `json:"url,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
}

// DeleteStudioResponse represents the DeleteStudioResponse schema from the OpenAPI specification
type DeleteStudioResponse struct {
	Studio interface{} `json:"studio"`
}

// Eula represents the Eula schema from the OpenAPI specification
type Eula struct {
	Content interface{} `json:"content,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Eulaid interface{} `json:"eulaId,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Updatedat interface{} `json:"updatedAt,omitempty"`
}

// ListLaunchProfileMembersResponse represents the ListLaunchProfileMembersResponse schema from the OpenAPI specification
type ListLaunchProfileMembersResponse struct {
	Members interface{} `json:"members,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// UpdateLaunchProfileMemberRequest represents the UpdateLaunchProfileMemberRequest schema from the OpenAPI specification
type UpdateLaunchProfileMemberRequest struct {
	Persona interface{} `json:"persona"`
}

// DeleteStreamingSessionRequest represents the DeleteStreamingSessionRequest schema from the OpenAPI specification
type DeleteStreamingSessionRequest struct {
}

// LaunchProfileMembership represents the LaunchProfileMembership schema from the OpenAPI specification
type LaunchProfileMembership struct {
	Identitystoreid interface{} `json:"identityStoreId,omitempty"`
	Persona interface{} `json:"persona,omitempty"`
	Principalid interface{} `json:"principalId,omitempty"`
	Sid interface{} `json:"sid,omitempty"`
}

// DeleteStreamingImageResponse represents the DeleteStreamingImageResponse schema from the OpenAPI specification
type DeleteStreamingImageResponse struct {
	Streamingimage interface{} `json:"streamingImage,omitempty"`
}

// ListStudioMembersResponse represents the ListStudioMembersResponse schema from the OpenAPI specification
type ListStudioMembersResponse struct {
	Members interface{} `json:"members,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// ComputeFarmConfiguration represents the ComputeFarmConfiguration schema from the OpenAPI specification
type ComputeFarmConfiguration struct {
	Activedirectoryuser interface{} `json:"activeDirectoryUser,omitempty"`
	Endpoint interface{} `json:"endpoint,omitempty"`
}

// DeleteLaunchProfileMemberRequest represents the DeleteLaunchProfileMemberRequest schema from the OpenAPI specification
type DeleteLaunchProfileMemberRequest struct {
}

// GetStreamingSessionBackupResponse represents the GetStreamingSessionBackupResponse schema from the OpenAPI specification
type GetStreamingSessionBackupResponse struct {
	Streamingsessionbackup interface{} `json:"streamingSessionBackup,omitempty"`
}

// ListStreamingSessionsResponse represents the ListStreamingSessionsResponse schema from the OpenAPI specification
type ListStreamingSessionsResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Sessions interface{} `json:"sessions,omitempty"`
}

// PutLaunchProfileMembersResponse represents the PutLaunchProfileMembersResponse schema from the OpenAPI specification
type PutLaunchProfileMembersResponse struct {
}

// ActiveDirectoryComputerAttribute represents the ActiveDirectoryComputerAttribute schema from the OpenAPI specification
type ActiveDirectoryComputerAttribute struct {
	Name interface{} `json:"name,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

// StreamingSessionStorageRoot represents the StreamingSessionStorageRoot schema from the OpenAPI specification
type StreamingSessionStorageRoot struct {
	Windows interface{} `json:"windows,omitempty"`
	Linux interface{} `json:"linux,omitempty"`
}

// CreateStreamingSessionStreamResponse represents the CreateStreamingSessionStreamResponse schema from the OpenAPI specification
type CreateStreamingSessionStreamResponse struct {
	Stream interface{} `json:"stream,omitempty"`
}

// StreamConfigurationCreate represents the StreamConfigurationCreate schema from the OpenAPI specification
type StreamConfigurationCreate struct {
	Streamingimageids interface{} `json:"streamingImageIds"`
	Ec2instancetypes interface{} `json:"ec2InstanceTypes"`
	Maxsessionlengthinminutes interface{} `json:"maxSessionLengthInMinutes,omitempty"`
	Sessionbackup interface{} `json:"sessionBackup,omitempty"`
	Sessionstorage interface{} `json:"sessionStorage,omitempty"`
	Volumeconfiguration interface{} `json:"volumeConfiguration,omitempty"`
	Maxstoppedsessionlengthinminutes interface{} `json:"maxStoppedSessionLengthInMinutes,omitempty"`
	Sessionpersistencemode interface{} `json:"sessionPersistenceMode,omitempty"`
	Automaticterminationmode interface{} `json:"automaticTerminationMode,omitempty"`
	Clipboardmode interface{} `json:"clipboardMode"`
}

// StreamConfigurationSessionStorage represents the StreamConfigurationSessionStorage schema from the OpenAPI specification
type StreamConfigurationSessionStorage struct {
	Mode interface{} `json:"mode"`
	Root interface{} `json:"root,omitempty"`
}

// StudioComponentConfiguration represents the StudioComponentConfiguration schema from the OpenAPI specification
type StudioComponentConfiguration struct {
	Activedirectoryconfiguration interface{} `json:"activeDirectoryConfiguration,omitempty"`
	Computefarmconfiguration interface{} `json:"computeFarmConfiguration,omitempty"`
	Licenseserviceconfiguration interface{} `json:"licenseServiceConfiguration,omitempty"`
	Sharedfilesystemconfiguration interface{} `json:"sharedFileSystemConfiguration,omitempty"`
}

// ListStudioMembersRequest represents the ListStudioMembersRequest schema from the OpenAPI specification
type ListStudioMembersRequest struct {
}

// StopStreamingSessionResponse represents the StopStreamingSessionResponse schema from the OpenAPI specification
type StopStreamingSessionResponse struct {
	Session StreamingSession `json:"session,omitempty"` // A streaming session is a virtual workstation created using a particular launch profile.
}

// Studio represents the Studio schema from the OpenAPI specification
type Studio struct {
	Tags interface{} `json:"tags,omitempty"`
	Adminrolearn interface{} `json:"adminRoleArn,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Statuscode interface{} `json:"statusCode,omitempty"`
	Userrolearn interface{} `json:"userRoleArn,omitempty"`
	Studioencryptionconfiguration interface{} `json:"studioEncryptionConfiguration,omitempty"`
	Studioid interface{} `json:"studioId,omitempty"`
	Updatedat interface{} `json:"updatedAt,omitempty"`
	Ssoclientid interface{} `json:"ssoClientId,omitempty"`
	Studiourl interface{} `json:"studioUrl,omitempty"`
	Statusmessage interface{} `json:"statusMessage,omitempty"`
	State interface{} `json:"state,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Displayname interface{} `json:"displayName,omitempty"`
	Homeregion interface{} `json:"homeRegion,omitempty"`
	Studioname interface{} `json:"studioName,omitempty"`
}

// ScriptParameterKeyValue represents the ScriptParameterKeyValue schema from the OpenAPI specification
type ScriptParameterKeyValue struct {
	Value interface{} `json:"value,omitempty"`
	Key interface{} `json:"key,omitempty"`
}

// GetLaunchProfileInitializationRequest represents the GetLaunchProfileInitializationRequest schema from the OpenAPI specification
type GetLaunchProfileInitializationRequest struct {
}

// ListEulaAcceptancesResponse represents the ListEulaAcceptancesResponse schema from the OpenAPI specification
type ListEulaAcceptancesResponse struct {
	Eulaacceptances interface{} `json:"eulaAcceptances,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// StudioMembership represents the StudioMembership schema from the OpenAPI specification
type StudioMembership struct {
	Identitystoreid interface{} `json:"identityStoreId,omitempty"`
	Persona interface{} `json:"persona,omitempty"`
	Principalid interface{} `json:"principalId,omitempty"`
	Sid interface{} `json:"sid,omitempty"`
}

// UpdateLaunchProfileResponse represents the UpdateLaunchProfileResponse schema from the OpenAPI specification
type UpdateLaunchProfileResponse struct {
	Launchprofile interface{} `json:"launchProfile,omitempty"`
}

// GetLaunchProfileMemberResponse represents the GetLaunchProfileMemberResponse schema from the OpenAPI specification
type GetLaunchProfileMemberResponse struct {
	Member interface{} `json:"member,omitempty"`
}

// StudioComponentInitializationScript represents the StudioComponentInitializationScript schema from the OpenAPI specification
type StudioComponentInitializationScript struct {
	Script interface{} `json:"script,omitempty"`
	Launchprofileprotocolversion interface{} `json:"launchProfileProtocolVersion,omitempty"`
	Platform interface{} `json:"platform,omitempty"`
	Runcontext interface{} `json:"runContext,omitempty"`
}

// ActiveDirectoryConfiguration represents the ActiveDirectoryConfiguration schema from the OpenAPI specification
type ActiveDirectoryConfiguration struct {
	Computerattributes interface{} `json:"computerAttributes,omitempty"`
	Directoryid interface{} `json:"directoryId,omitempty"`
	Organizationalunitdistinguishedname interface{} `json:"organizationalUnitDistinguishedName,omitempty"`
}

// DeleteLaunchProfileMemberResponse represents the DeleteLaunchProfileMemberResponse schema from the OpenAPI specification
type DeleteLaunchProfileMemberResponse struct {
}

// ListStreamingImagesRequest represents the ListStreamingImagesRequest schema from the OpenAPI specification
type ListStreamingImagesRequest struct {
}

// GetStudioMemberResponse represents the GetStudioMemberResponse schema from the OpenAPI specification
type GetStudioMemberResponse struct {
	Member interface{} `json:"member,omitempty"`
}

// CreateStudioComponentResponse represents the CreateStudioComponentResponse schema from the OpenAPI specification
type CreateStudioComponentResponse struct {
	Studiocomponent interface{} `json:"studioComponent,omitempty"`
}

// LaunchProfileInitialization represents the LaunchProfileInitialization schema from the OpenAPI specification
type LaunchProfileInitialization struct {
	Systeminitializationscripts interface{} `json:"systemInitializationScripts,omitempty"`
	Activedirectory interface{} `json:"activeDirectory,omitempty"`
	Ec2securitygroupids interface{} `json:"ec2SecurityGroupIds,omitempty"`
	Launchprofileid interface{} `json:"launchProfileId,omitempty"`
	Launchprofileprotocolversion interface{} `json:"launchProfileProtocolVersion,omitempty"`
	Userinitializationscripts interface{} `json:"userInitializationScripts,omitempty"`
	Launchpurpose interface{} `json:"launchPurpose,omitempty"`
	Platform interface{} `json:"platform,omitempty"`
	Name interface{} `json:"name,omitempty"`
}

// NewStudioMember represents the NewStudioMember schema from the OpenAPI specification
type NewStudioMember struct {
	Principalid interface{} `json:"principalId"`
	Persona interface{} `json:"persona"`
}

// CreateStreamingSessionStreamRequest represents the CreateStreamingSessionStreamRequest schema from the OpenAPI specification
type CreateStreamingSessionStreamRequest struct {
	Expirationinseconds interface{} `json:"expirationInSeconds,omitempty"`
}

// DeleteStudioMemberRequest represents the DeleteStudioMemberRequest schema from the OpenAPI specification
type DeleteStudioMemberRequest struct {
}

// ListEulasResponse represents the ListEulasResponse schema from the OpenAPI specification
type ListEulasResponse struct {
	Eulas interface{} `json:"eulas,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// PutStudioMembersResponse represents the PutStudioMembersResponse schema from the OpenAPI specification
type PutStudioMembersResponse struct {
}

// GetLaunchProfileRequest represents the GetLaunchProfileRequest schema from the OpenAPI specification
type GetLaunchProfileRequest struct {
}

// ListStudioComponentsResponse represents the ListStudioComponentsResponse schema from the OpenAPI specification
type ListStudioComponentsResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Studiocomponents interface{} `json:"studioComponents,omitempty"`
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Tags interface{} `json:"tags,omitempty"`
}

// Tags represents the Tags schema from the OpenAPI specification
type Tags struct {
}

// CreateStreamingSessionRequest represents the CreateStreamingSessionRequest schema from the OpenAPI specification
type CreateStreamingSessionRequest struct {
	Ownedby interface{} `json:"ownedBy,omitempty"`
	Streamingimageid interface{} `json:"streamingImageId,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Ec2instancetype interface{} `json:"ec2InstanceType,omitempty"`
	Launchprofileid interface{} `json:"launchProfileId"`
}

// LaunchProfileInitializationActiveDirectory represents the LaunchProfileInitializationActiveDirectory schema from the OpenAPI specification
type LaunchProfileInitializationActiveDirectory struct {
	Directoryname interface{} `json:"directoryName,omitempty"`
	Dnsipaddresses interface{} `json:"dnsIpAddresses,omitempty"`
	Organizationalunitdistinguishedname interface{} `json:"organizationalUnitDistinguishedName,omitempty"`
	Studiocomponentid interface{} `json:"studioComponentId,omitempty"`
	Studiocomponentname interface{} `json:"studioComponentName,omitempty"`
	Computerattributes interface{} `json:"computerAttributes,omitempty"`
	Directoryid interface{} `json:"directoryId,omitempty"`
}

// GetStreamingSessionRequest represents the GetStreamingSessionRequest schema from the OpenAPI specification
type GetStreamingSessionRequest struct {
}

// GetStreamingSessionStreamResponse represents the GetStreamingSessionStreamResponse schema from the OpenAPI specification
type GetStreamingSessionStreamResponse struct {
	Stream interface{} `json:"stream,omitempty"`
}

// SharedFileSystemConfiguration represents the SharedFileSystemConfiguration schema from the OpenAPI specification
type SharedFileSystemConfiguration struct {
	Windowsmountdrive interface{} `json:"windowsMountDrive,omitempty"`
	Endpoint interface{} `json:"endpoint,omitempty"`
	Filesystemid interface{} `json:"fileSystemId,omitempty"`
	Linuxmountpoint interface{} `json:"linuxMountPoint,omitempty"`
	Sharename interface{} `json:"shareName,omitempty"`
}

// StreamConfiguration represents the StreamConfiguration schema from the OpenAPI specification
type StreamConfiguration struct {
	Ec2instancetypes interface{} `json:"ec2InstanceTypes"`
	Sessionpersistencemode interface{} `json:"sessionPersistenceMode,omitempty"`
	Volumeconfiguration interface{} `json:"volumeConfiguration,omitempty"`
	Sessionbackup interface{} `json:"sessionBackup,omitempty"`
	Sessionstorage interface{} `json:"sessionStorage,omitempty"`
	Streamingimageids interface{} `json:"streamingImageIds"`
	Automaticterminationmode interface{} `json:"automaticTerminationMode,omitempty"`
	Maxsessionlengthinminutes interface{} `json:"maxSessionLengthInMinutes,omitempty"`
	Clipboardmode interface{} `json:"clipboardMode"`
	Maxstoppedsessionlengthinminutes interface{} `json:"maxStoppedSessionLengthInMinutes,omitempty"`
}

// GetStudioMemberRequest represents the GetStudioMemberRequest schema from the OpenAPI specification
type GetStudioMemberRequest struct {
}

// StreamingImage represents the StreamingImage schema from the OpenAPI specification
type StreamingImage struct {
	Eulaids interface{} `json:"eulaIds,omitempty"`
	Owner interface{} `json:"owner,omitempty"`
	State interface{} `json:"state,omitempty"`
	Streamingimageid interface{} `json:"streamingImageId,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Statuscode interface{} `json:"statusCode,omitempty"`
	Ec2imageid interface{} `json:"ec2ImageId,omitempty"`
	Encryptionconfiguration interface{} `json:"encryptionConfiguration,omitempty"`
	Platform interface{} `json:"platform,omitempty"`
	Statusmessage interface{} `json:"statusMessage,omitempty"`
	Description interface{} `json:"description,omitempty"`
}

// ListStudioComponentsRequest represents the ListStudioComponentsRequest schema from the OpenAPI specification
type ListStudioComponentsRequest struct {
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
}

// AcceptEulasResponse represents the AcceptEulasResponse schema from the OpenAPI specification
type AcceptEulasResponse struct {
	Eulaacceptances interface{} `json:"eulaAcceptances,omitempty"`
}

// StartStudioSSOConfigurationRepairResponse represents the StartStudioSSOConfigurationRepairResponse schema from the OpenAPI specification
type StartStudioSSOConfigurationRepairResponse struct {
	Studio interface{} `json:"studio"`
}

// CreateStudioComponentRequest represents the CreateStudioComponentRequest schema from the OpenAPI specification
type CreateStudioComponentRequest struct {
	Runtimerolearn interface{} `json:"runtimeRoleArn,omitempty"`
	Subtype interface{} `json:"subtype,omitempty"`
	Name interface{} `json:"name"`
	Scriptparameters interface{} `json:"scriptParameters,omitempty"`
	Secureinitializationrolearn interface{} `json:"secureInitializationRoleArn,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Ec2securitygroupids interface{} `json:"ec2SecurityGroupIds,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Initializationscripts interface{} `json:"initializationScripts,omitempty"`
	TypeField interface{} `json:"type"`
	Configuration interface{} `json:"configuration,omitempty"`
}

// StreamConfigurationSessionBackup represents the StreamConfigurationSessionBackup schema from the OpenAPI specification
type StreamConfigurationSessionBackup struct {
	Mode interface{} `json:"mode,omitempty"`
	Maxbackupstoretain interface{} `json:"maxBackupsToRetain,omitempty"`
}

// ListStudiosResponse represents the ListStudiosResponse schema from the OpenAPI specification
type ListStudiosResponse struct {
	Studios interface{} `json:"studios"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// DeleteStudioComponentResponse represents the DeleteStudioComponentResponse schema from the OpenAPI specification
type DeleteStudioComponentResponse struct {
	Studiocomponent interface{} `json:"studioComponent,omitempty"`
}

// GetStreamingSessionBackupRequest represents the GetStreamingSessionBackupRequest schema from the OpenAPI specification
type GetStreamingSessionBackupRequest struct {
}

// StartStreamingSessionResponse represents the StartStreamingSessionResponse schema from the OpenAPI specification
type StartStreamingSessionResponse struct {
	Session StreamingSession `json:"session,omitempty"` // A streaming session is a virtual workstation created using a particular launch profile.
}

// ListStudiosRequest represents the ListStudiosRequest schema from the OpenAPI specification
type ListStudiosRequest struct {
}

// ListStreamingImagesResponse represents the ListStreamingImagesResponse schema from the OpenAPI specification
type ListStreamingImagesResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Streamingimages interface{} `json:"streamingImages,omitempty"`
}

// StreamingImageEncryptionConfiguration represents the StreamingImageEncryptionConfiguration schema from the OpenAPI specification
type StreamingImageEncryptionConfiguration struct {
	Keyarn interface{} `json:"keyArn,omitempty"`
	Keytype interface{} `json:"keyType"`
}

// DeleteStudioRequest represents the DeleteStudioRequest schema from the OpenAPI specification
type DeleteStudioRequest struct {
}

// CreateStreamingImageRequest represents the CreateStreamingImageRequest schema from the OpenAPI specification
type CreateStreamingImageRequest struct {
	Description interface{} `json:"description,omitempty"`
	Ec2imageid interface{} `json:"ec2ImageId"`
	Name interface{} `json:"name"`
	Tags interface{} `json:"tags,omitempty"`
}

// CreateLaunchProfileResponse represents the CreateLaunchProfileResponse schema from the OpenAPI specification
type CreateLaunchProfileResponse struct {
	Launchprofile interface{} `json:"launchProfile,omitempty"`
}
