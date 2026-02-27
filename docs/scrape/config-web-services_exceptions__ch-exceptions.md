All error messages contain at least the `message` parameter, which is a string that describes the exception. Some exceptions contain additional parameters, which are described in the table.

                                                    | Name | Description |
| --- | --- |
| AccessDisallowedFault | Your domain does not have access to the Five9 Configuration Web Services. To request access, contact Five9 Support. |
| AddRecordToListFault | Version 3 Error with addRecordToListSimple due to one of these reasons: The following field(s) do not exist in the CRM table: {0} Sending email not supported in this release. Mandatory field {0} missing. There can be maximum of 64 CRM fields. Number {0} must either be 10 digits for dialing within North America or begin with 011 for International number. Number {0} has to be less than or equal to 16 digit and greater than or equal to 10 digit. List was not found. Specified record already exist in the list. At least one number should be specified for new record. No value provided for key {0}. More than one record matches specified criteria. At least one key must be provided. Time format must be provided along with Time to Dial. Incorrect Time to Dial. Incorrect key {0}. No such field. Value of field {0} is too long. Value of field {0} has incorrect format{1}. Number of requests exceed maximum allowed value: {0}. The requested domain is disabled. There is no resources for processing the request at the moment. Please try to re-send the request. Domain wasn't found. Error while executing request. Unknown error {0}. Error while executing SQL statement {0}. Error while executing SQL statement {0} with values {1}. Error while processing Call ASAP and Time to call parameters - service temporary unavailable. Error while loading external params. |
| AdminSessionClosedFault | System error. |
| AllowedOnlyForPreviewCampaignFault | Incorrect dialing mode. Returned for the preview-only manual dialing mode. WrongDialingModeException: Name Type Description campaignName string Name of campaign. | Name | Type | Description | campaignName | string | Name of campaign. |
| Name | Type | Description |
| campaignName | string | Name of campaign. |
| CampaignAlreadyExistsFault | Attempting to create a campaign that already exists or to rename a campaign to the name of a campaign that already exists. Name Type Description campaignName string Name of campaign. | Name | Type | Description | campaignName | string | Name of campaign. |
| Name | Type | Description |
| campaignName | string | Name of campaign. |
| CampaignNotFoundFault | Campaign name does not exist in the Five9 configuration. Name Type Description campaignName string Name of campaign. | Name | Type | Description | campaignName | string | Name of campaign. |
| Name | Type | Description |
| campaignName | string | Name of campaign. |
| CampaignStateUpdateFault | Campaign state cannot be updated, for example: you cannot stop a campaign that is not running. Name Type Description campaignName string Name of campaign. details string Reason for the failure. | Name | Type | Description | campaignName | string | Name of campaign. | details | string | Reason for the failure. |
| Name | Type | Description |
| campaignName | string | Name of campaign. |
| details | string | Reason for the failure. |
| CantModifyObjectFault | System error. |
| CommonCampaignFault Version 3 | Common part of the campaign error. Name Type Description campaignName string Name of campaign. | Name | Type | Description | campaignName | string | Name of campaign. |
| Name | Type | Description |
| campaignName | string | Name of campaign. |
| ConcurrentModificationFault | Specified object is already being modified by another client. |
| CrmFieldNotFoundFault | Specified contact database field cannot be found in this domain. Name Type Description fieldName string Name of field. | Name | Type | Description | fieldName | string | Name of field. |
| Name | Type | Description |
| fieldName | string | Name of field. |
| DialProfileNotFoundFault | Undefined profile is specified when creating or modifying an outbound campaign. Name Type Description profileName string Name of profile. | Name | Type | Description | profileName | string | Name of profile. |
| Name | Type | Description |
| profileName | string | Name of profile. |
| DispositionAlreadyExistsFault | Attempting to create a disposition that already exists or to rename a disposition to the name of a disposition that already exists. Name Type Description dispositionName string Name of disposition. | Name | Type | Description | dispositionName | string | Name of disposition. |
| Name | Type | Description |
| dispositionName | string | Name of disposition. |
| DispositionIsAlreadyAssignedFault DispositionAlreadyAssignedException Note: The Fault name is different from the Exception name. | Attempting to assign to a campaign a disposition that has already been assigned to the campaign. Name Type Description campaignName string Name of campaign. dispositionName string Information about current and requested disposition. | Name | Type | Description | campaignName | string | Name of campaign. | dispositionName | string | Information about current and requested disposition. |
| Name | Type | Description |
| campaignName | string | Name of campaign. |
| dispositionName | string | Information about current and requested disposition. |
| DispositionIsNotAssisgnedFault DispositionIsNotAssignedException Note: The spelling of the Fault is different from that of the Exception. | Attempting to remove from a campaign a disposition that is not assigned to that campaign. Name Type Description campaignName string Name of campaign. dispositionName string Name of disposition. | Name | Type | Description | campaignName | string | Name of campaign. | dispositionName | string | Name of disposition. |
| Name | Type | Description |
| campaignName | string | Name of campaign. |
| dispositionName | string | Name of disposition. |
| DispositionNotAllowedFault | Attempting to add a disposition type that is not allowed for a campaign. Name Type Description dispositionName string Name of disposition. | Name | Type | Description | dispositionName | string | Name of disposition. |
| Name | Type | Description |
| dispositionName | string | Name of disposition. |
| DispositionNotFoundFault | Attempting to get, remove, or rename a disposition that does not exist in the domain. Name Type Description dispositionName string Name of disposition. | Name | Type | Description | dispositionName | string | Name of disposition. |
| Name | Type | Description |
| dispositionName | string | Name of disposition. |
| DNISAlreadyAssignedFault | If a DNIS number is already assigned to a user, the DNISAlreadyAssignedFault saves the users name in the campaignName field. To ensure compatibility with existing API versions, we've opted to keep the users name stored in this location without altering the API. Campaign and user name may be distinguished via the fault message: 'DNIS "DDD" already assigned to campaign "CCC"' or 'DNIS "DDD" already assigned to user "UUU"'. Name Type Description campaignName string Name of campaign. dnis string DNIS number. | Name | Type | Description | campaignName | string | Name of campaign. | dnis | string | DNIS number. |
| Name | Type | Description |
| campaignName | string | Name of campaign. |
| dnis | string | DNIS number. |
| DNISIsNotAssignedFault | DNIS is not assigned to a campaign. Name Type Description campaignName string Name of campaign. dnis string DNIS number. | Name | Type | Description | campaignName | string | Name of campaign. | dnis | string | DNIS number. |
| Name | Type | Description |
| campaignName | string | Name of campaign. |
| dnis | string | DNIS number. |
| DNISNotFoundFault | DNIS cannot be found in the domain. Name Type Description dnis string DNIS number. | Name | Type | Description | dnis | string | DNIS number. |
| Name | Type | Description |
| dnis | string | DNIS number. |
| ExecutionRestrictionFault | Processing error. |
| ExtensionAlreadyInUseFault | Extension already assigned to another user. Name Type Description ext string 4-digit extension. | Name | Type | Description | ext | string | 4-digit extension. |
| Name | Type | Description |
| ext | string | 4-digit extension. |
| ExtensionsOutOfRangeFault Version 9.5 | Number of digits in the users’ extension is out of the allowed range. Name Type Description campaignName string [0..unbounded] Name of campaign. userNames string [0..unbounded] List of user names. | Name | Type | Description | campaignName | string [0..unbounded] | Name of campaign. | userNames | string [0..unbounded] | List of user names. |
| Name | Type | Description |
| campaignName | string [0..unbounded] | Name of campaign. |
| userNames | string [0..unbounded] | List of user names. |
| FinderException | Object cannot be found. |
| ImportCancelledFault | Import operation was canceled. |
| ImportInProgressFault | Campaign cannot be deleted because data is being imported into this campaign or into a list associated with this campaign |
| ImportSizeLimitExceededFault ImportRecordsCountLimitExceededException Note: The name of the fault is different from that of the exception. | Request exceeds maximum number of records allowed. Name Type Description limit int Maximum number of records allowed. recordsCount int Number of records in the request. | Name | Type | Description | limit | int | Maximum number of records allowed. | recordsCount | int | Number of records in the request. |
| Name | Type | Description |
| limit | int | Maximum number of records allowed. |
| recordsCount | int | Number of records in the request. |
| IncorrectArgumentFault | Request contains incorrect argument name or value, or non-parameterized scripts contain parameters. |
| InternalImportFault | All requests that trigger a data import process may experience an internal import fault. Resubmit the request. |
| InvalidAccountFault | Invalid account regardless of the reason, including password and security questions issues. May be returned by the createUser and modifyUser methods. |
| InvalidDateRangeFault | Invalid date range. Start date and time must precede end date and time. Name Type Description end dateTime End time of range. start dateTime Start time of range. | Name | Type | Description | end | dateTime | End time of range. | start | dateTime | Start time of range. |
| Name | Type | Description |
| end | dateTime | End time of range. |
| start | dateTime | Start time of range. |
| InvalidImportDataFault | Incorrectly formatted source data for import operations. |
| InvalidRegexpPatternFault | Incorrectly formed regular expression used for data lookup. Name Type Description pattern string Invalid regular expression. | Name | Type | Description | pattern | string | Invalid regular expression. |
| Name | Type | Description |
| pattern | string | Invalid regular expression. |
| InvalidUserDataFault | Error in data provided to create or update user. |
| IvrScriptOwnershipNotFoundFault Version 10 | Attempting to create or modify an IVR script with invalid ownership information. |
| IvrScriptNotFoundFault | Attempting to create or modify a campaign with an IVR script that does not exist. Name Type Description ivrScriptName string Name of IVR script. | Name | Type | Description | ivrScriptName | string | Name of IVR script. |
| Name | Type | Description |
| ivrScriptName | string | Name of IVR script. |
| ListAlreadyAssignedFault | List already assigned to this campaign. Name Type Description campaignName string Name of campaign. listName string Name of list. | Name | Type | Description | campaignName | string | Name of campaign. | listName | string | Name of list. |
| Name | Type | Description |
| campaignName | string | Name of campaign. |
| listName | string | Name of list. |
| ListAlreadyExistsFault | List name already in use. Name Type Description listName string Name of list. | Name | Type | Description | listName | string | Name of list. |
| Name | Type | Description |
| listName | string | Name of list. |
| ListCantBeRemovedFault | Attempting to delete a list that is in use. Name Type Description listName string Name of list. | Name | Type | Description | listName | string | Name of list. |
| Name | Type | Description |
| listName | string | Name of list. |
| ListIsNotAssignedFault | Attempting to modify or remove a list that is not assigned to the campaign. Name Type Description campaignName string Name of campaign. listName string Name of list. | Name | Type | Description | campaignName | string | Name of campaign. | listName | string | Name of list. |
| Name | Type | Description |
| campaignName | string | Name of campaign. |
| listName | string | Name of list. |
| ListNotFoundFault | Attempting to modify a list that does not exist. Name Type Description listName string Name of list. | Name | Type | Description | listName | string | Name of list. |
| Name | Type | Description |
| listName | string | Name of list. |
| LocaleNotFoundFault Version 9.5 | Unable to find the contact’s locale. Name Type Description localeName string Name of locale. | Name | Type | Description | localeName | string | Name of locale. |
| Name | Type | Description |
| localeName | string | Name of locale. |
| LogoutReasonCodeNotFoundFault Removed from Version 10 | No reason codes of type logout exist with the name provided. Name Type Description reasonCodeName string Name of reason code. | Name | Type | Description | reasonCodeName | string | Name of reason code. |
| Name | Type | Description |
| reasonCodeName | string | Name of reason code. |
| MaxPlayFileCountForSkillExceededFault Version 9.5 | Number of audio files for the skill has been exceeded. Name Type Description skillName string Name of skill. | Name | Type | Description | skillName | string | Name of skill. |
| Name | Type | Description |
| skillName | string | Name of skill. |
| MissedArgumentFault | Required argument is missing from the request. |
| MissedOsLoginFault | Do not use. |
| NotReadyReasonCodeNotFoundFault Removed from Version 10 | No Not Ready reason code with that name exists. Name Type Description reasonCodeName string Name of reason code. | Name | Type | Description | reasonCodeName | string | Name of reason code. |
| Name | Type | Description |
| reasonCodeName | string | Name of reason code. |
| ObjectAlreadyExistsFault | Object with the same name already exists. Name Type Description id long Object ID. objectName string Name of object. objectType wsObjectType Type of data. | Name | Type | Description | id | long | Object ID. | objectName | string | Name of object. | objectType | wsObjectType | Type of data. |
| Name | Type | Description |
| id | long | Object ID. |
| objectName | string | Name of object. |
| objectType | wsObjectType | Type of data. |
| ObjectInUseFault | Requested object is currently being used. Name Type Description id long Object ID. objectName string Name of object. objectType wsObjectType Type of data. | Name | Type | Description | id | long | Object ID. | objectName | string | Name of object. | objectType | wsObjectType | Type of data. |
| Name | Type | Description |
| id | long | Object ID. |
| objectName | string | Name of object. |
| objectType | wsObjectType | Type of data. |
| ObjectNotFoundFault | Requested object does not exist. Name Type Description id long Object ID. objectName string Name of object. objectType wsObjectType Type of data. | Name | Type | Description | id | long | Object ID. | objectName | string | Name of object. | objectType | wsObjectType | Type of data. |
| Name | Type | Description |
| id | long | Object ID. |
| objectName | string | Name of object. |
| objectType | wsObjectType | Type of data. |
| ObjectsCountLimitExceededFault | Maximum count for this type of object for your domain has been reached. Name Type Description id long Object ID. limit int Limit that has been reached. objectName string Name of object. objectType wsObjectType Type of data. | Name | Type | Description | id | long | Object ID. | limit | int | Limit that has been reached. | objectName | string | Name of object. | objectType | wsObjectType | Type of data. |
| Name | Type | Description |
| id | long | Object ID. |
| limit | int | Limit that has been reached. |
| objectName | string | Name of object. |
| objectType | wsObjectType | Type of data. |
| OperationsLimitExceededFault | Maximum number of Web Services requests for the 24-hour period has been exceeded. Name Type Description limit long Maximum number of requests. operationType string Type of operation. | Name | Type | Description | limit | long | Maximum number of requests. | operationType | string | Type of operation. |
| Name | Type | Description |
| limit | long | Maximum number of requests. |
| operationType | string | Type of operation. |
| ParseException | Error found during parsing. Name Type Description errorOffset int Location of the error. | Name | Type | Description | errorOffset | int | Location of the error. |
| Name | Type | Description |
| errorOffset | int | Location of the error. |
| PromptAlreadyExistsFault | Attempting to create a prompt that already exists. Name Type Description promptName string Name of prompt. | Name | Type | Description | promptName | string | Name of prompt. |
| Name | Type | Description |
| promptName | string | Name of prompt. |
| PromptCantBeDeletedFault | Attempting to delete a prompt that is currently in use. Name Type Description promptName string Name of prompt. | Name | Type | Description | promptName | string | Name of prompt. |
| Name | Type | Description |
| promptName | string | Name of prompt. |
| PromptNotFoundFault | Attempting to create or modify a campaign that requires a prompt. Name Type Description promptName string Name of prompt. | Name | Type | Description | promptName | string | Name of prompt. |
| Name | Type | Description |
| promptName | string | Name of prompt. |
| ReasonCodeCountLimitExceededFault | Limit for the number of reason codes in a domain has been reached. Name Type Description reasonCodeName string Name of reason code. | Name | Type | Description | reasonCodeName | string | Name of reason code. |
| Name | Type | Description |
| reasonCodeName | string | Name of reason code. |
| ReasonCodeNotFoundFault | No reason code of that name exists. Name Type Description reasonCodeName string Name of reason code. | Name | Type | Description | reasonCodeName | string | Name of reason code. |
| Name | Type | Description |
| reasonCodeName | string | Name of reason code. |
| ReportNotFoundFault | Report with that name and category does not exist. Name Type Description folder string Category of report. reportName string Name of report. | Name | Type | Description | folder | string | Category of report. | reportName | string | Name of report. |
| Name | Type | Description |
| folder | string | Category of report. |
| reportName | string | Name of report. |
| ResultIsNotReadyFault | Operation in progress has not been completed. To check status, use isImportRunning or isReportRunning. |
| ScheduleNotFoundFault | FTP schedule was not found. Name Type Description scheduleName string Name of schedule. | Name | Type | Description | scheduleName | string | Name of schedule. |
| Name | Type | Description |
| scheduleName | string | Name of schedule. |
| ScheduleOperationFailedFault | Attempt to schedule FTP event failed. Name Type Description operation operationType Enumeration of type string. scheduleName string Name of schedule. | Name | Type | Description | operation | operationType | Enumeration of type string. | scheduleName | string | Name of schedule. |
| Name | Type | Description |
| operation | operationType | Enumeration of type string. |
| scheduleName | string | Name of schedule. |
| ServerFault | Web Services server error. |
| ServiceUnavailableFault | Web Services are not available. |
| SessionClosedFault | Session closed while the request is executed because another session is started concurrently with the same credentials. |
| SkillAlreadyAssignedFault | Attempting to assign to a campaign a skill already assigned to the campaign. Name Type Description campaignName string Name of campaign. | Name | Type | Description | campaignName | string | Name of campaign. |
| Name | Type | Description |
| campaignName | string | Name of campaign. |
| SkillAlreadyExistsFault | Attempting to create a skill with a name that already exists. Name Type Description skillName string Name of skill. | Name | Type | Description | skillName | string | Name of skill. |
| Name | Type | Description |
| skillName | string | Name of skill. |
| SkillCantBeDeletedFault | Attempting to delete a skill used in other objects. Name Type Description skillName string Name of skill. | Name | Type | Description | skillName | string | Name of skill. |
| Name | Type | Description |
| skillName | string | Name of skill. |
| SkillIsNotAssignedFault | Attempting to remove or modify a skill that is not assigned to that campaign. Name Type Description campaignName string Name of campaign. skillName string Name of skill. | Name | Type | Description | campaignName | string | Name of campaign. | skillName | string | Name of skill. |
| Name | Type | Description |
| campaignName | string | Name of campaign. |
| skillName | string | Name of skill. |
| SkillNotFoundFault | Requested skill cannot be found. Name Type Description skillName string Name of skill. | Name | Type | Description | skillName | string | Name of skill. |
| Name | Type | Description |
| skillName | string | Name of skill. |
| TooManyExtensionsFault | Number of extensions has exceeded the maximum allowed. |
| TooManyItemsFault | Number of items has exceeded the maximum allowed. |
| TooManyUsersFault | Number of users has exceeded the maximum allowed. Returned by createUser. |
| TtsGenerationFailed Note: Fault is absent from the name of the Exception. | Audio file for the TTS prompt cannot be created. Name Type Description promptName string Name of prompt. | Name | Type | Description | promptName | string | Name of prompt. |
| Name | Type | Description |
| promptName | string | Name of prompt. |
| UnknownIdentifierFault | Unknown identifier of import operation is requested. |
| UserAlreadyExistsFault | Attempting to create a user with a name that already exists. Name Type Description userName string Name of user. | Name | Type | Description | userName | string | Name of user. |
| Name | Type | Description |
| userName | string | Name of user. |
| UserAlreadyHasSkillFault | User already has the specified skill. Name Type Description skillName string Name of skill. userName string Name of user. | Name | Type | Description | skillName | string | Name of skill. | userName | string | Name of user. |
| Name | Type | Description |
| skillName | string | Name of skill. |
| userName | string | Name of user. |
| UserAlreadyLoggedInFault | User ID is already logged in. |
| UserCantBeDeletedFault | User name cannot be deleted because it is used in other objects. Name Type Description userName string Name of user. | Name | Type | Description | userName | string | Name of user. |
| Name | Type | Description |
| userName | string | Name of user. |
| UserDoesntHaveSkillFault | User does not have the specified skill. Name Type Description skillName string Name of skill. userName string Name of user. | Name | Type | Description | skillName | string | Name of skill. | userName | string | Name of user. |
| Name | Type | Description |
| skillName | string | Name of skill. |
| userName | string | Name of user. |
| UserHasNoRequiredRoleFault UserHasNoRequiredRolesException Note: The spelling of the Fault is different from that of the Exception. | User being added to agent group does not have the required agent or supervisor role. Name Type Description roles userRoleType [0..unbounded] Types of roles. userName string Name of user. | Name | Type | Description | roles | userRoleType [0..unbounded] | Types of roles. | userName | string | Name of user. |
| Name | Type | Description |
| roles | userRoleType [0..unbounded] | Types of roles. |
| userName | string | Name of user. |
| UserNotFoundFault | User not found. Name Type Description userName string Name of user. | Name | Type | Description | userName | string | Name of user. |
| Name | Type | Description |
| userName | string | Name of user. |
| WavFileUploadFailedFault | WAV file upload failed when attempting to create or modify prompt. Name Type Description promptName string Name of prompt. | Name | Type | Description | promptName | string | Name of prompt. |
| Name | Type | Description |
| promptName | string | Name of prompt. |
| WrongCampaignStateFault | Requesting a campaign that is not in the correct state. For example, attempting to delete, reset, or rename a campaign or dispositions while the campaign is running. Name Type Description actualState campaignState Specified state. campaignName string Name of campaign. desiredState campaignState Campaign state that should be specified. | Name | Type | Description | actualState | campaignState | Specified state. | campaignName | string | Name of campaign. | desiredState | campaignState | Campaign state that should be specified. |
| Name | Type | Description |
| actualState | campaignState | Specified state. |
| campaignName | string | Name of campaign. |
| desiredState | campaignState | Campaign state that should be specified. |
| WrongCampaignTypeFault | Incorrect campaign type. For example, outbound campaign requests should contain types that apply to outbound campaigns. Name Type Description actualType campaignType Specified campaign type. campaignName string Name of campaign. desiredType campaignType Campaign types to specify. | Name | Type | Description | actualType | campaignType | Specified campaign type. | campaignName | string | Name of campaign. | desiredType | campaignType | Campaign types to specify. |
| Name | Type | Description |
| actualType | campaignType | Specified campaign type. |
| campaignName | string | Name of campaign. |
| desiredType | campaignType | Campaign types to specify. |
| WrongListDialingModeFault WrongListDialingModeFault Version 4 | Incorrect dialing mode. Name Type Description campaignName string Name of campaign. | Name | Type | Description | campaignName | string | Name of campaign. |
| Name | Type | Description |
| campaignName | string | Name of campaign. |
| WrongPromptTypeFault | Specified prompt type is incorrect. Name Type Description actualType promptType Specified prompt type. desiredType promptType Prompt type to specify. promptName string Name of prompt. | Name | Type | Description | actualType | promptType | Specified prompt type. | desiredType | promptType | Prompt type to specify. | promptName | string | Name of prompt. |
| Name | Type | Description |
| actualType | promptType | Specified prompt type. |
| desiredType | promptType | Prompt type to specify. |
| promptName | string | Name of prompt. |

[]()