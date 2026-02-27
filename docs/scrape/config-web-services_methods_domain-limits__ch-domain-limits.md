For security reasons, Five9 has implemented default limits for each type of Web Services request. If you require higher limits for some types of requests, contact Five9 Customer Support.

                                                    **Important: **API request limits may change over time. To ensure that you always know the correct limits for your domain, use `[getCallCountersState](/bundle/api-config-web/page/config-web-services/methods/domain-limits/getCallCountersState.htm)`. The response contain the current limits for each type of request in the table below.

The minute, hour, or day count starts at the beginning of the time interval at the location of the data center, including its time zone.

If you start to send requests at 9:30 PM in California, the hourly interval ends at 10:00 PM. At that time, the limit is reset for the next hour.

If your production domain has 4000 agents, calculate as follows the number of records that you can update or insert each hour with `asyncUpdateCrmRecords`. Your domain may have different limits from those in the example.

Maximum number of requests each hour: 1000

   X

Maximum number of records in each request: 100

Result: 100,000 records

                                                    | Types of API Requests | Minute | Hour | Day |
| --- | --- | --- | --- |
| Query getContactRecords returns a maximum of 1000 records. checkDncForNumbers, getAgentGroup, getAgentGroups, getAutodialCampaign, getCallVariableGroups, getCallVariables, getCampaignDNISList, getCampaignProfileDispositions, getCampaignProfileFilter, getCampaignProfiles, getCampaigns, getCampaignState, getContactFields, getContactRecords, getCrmImportResult, getDialingRules, getDisposition, getDispositionsImportResult, getDispositions, getDNISList, getInboundCampaign, getIVRScripts, getListImportResult, getListsForCampaign, getListsInfo, getOutboundCampaign, getPrompt, getReasonCodeByType, getReasonCode, getSkillInfo, getSkill, getSkillsInfo, getSkills, getSkillVoicemailGreeting, getUserGeneralInfo, getUserInfo, getUserProfile, getUserProfiles, getUsersGeneralInfo, getUserInfo, getUserVoicemailGreeting, getVCCConfiguration, getWebConnectors, isImportRunning, and isReportRunning | 450 | 15000 | 172800 |
| Modify (creating, modifying, or deleting values or objects) addDispositionsToCampaign, addDNISToCampaign, addListsToCampaign, addNumbersToDnc, addPromptTTS, addPromptWavInline, addPromptWav, addSkillsToCampaign, createAgentGroup, createAutodialCampaign, createCallVariable, createCallVariablesGroup, createCampaignProfile, createContactField, createDisposition, createInboundCampaign, createIVRScript, createList, createOutboundCampaign, createReasonCode, createSkill, createUser, createUserProfile, createWebConnector, deleteAgentGroup, deleteCallVariable, deleteCallVariablesGroup, deleteCampaign, deleteCampaignProfile, deleteContactField, deleteIVRScript, deleteList, deletePrompt, deleteReasonCodeByType, deleteReasonCode, deleteSkill, deleteUser, deleteUserProfile, deleteWebConnector, forceStopCampaign, modifyAgentGroup, modifyAutodialCampaign, modifyCallVariable, modifyCallVariablesGroup, modifyCampaignLists, modifyCampaignProfileCrmCriteria, modifyCampaignProfileDispositions, modifyCampaignProfileFilterOrder, modifyCampaignProfile, modifyContactField, modifyDisposition, modifyInboundCampaign, modifyIVRScript, modifyOutboundCampaign, modifyPromptTTS, modifyPromptWavInline, modifyPromptWav, modifyReasonCode, modifySkill, modifyUserCannedReports, modifyUser, modifyUserProfile, modifyUserProfileSkills, modifyUserProfileUserList, modifyVCCConfiguration, modifyWebConnector, removeDisposition, removeDispositionsFromCampaign, removeDNISFromCampaign, removeListsFromCampaign, removeNumbersFromDnc, removeSkillsFromCampaign, renameCampaign, renameDisposition, resetCampaignDispositions, resetCampaign, resetListPosition, setDefaultIVRSchedule, setDialingRules, setSkillVoicemailGreeting, setUserVoicemailGreeting, startCampaign, stopCampaign, userSkillAdd, userSkillModify, userSkillRemove | 140 | 7200 | 172800 |
| Importing single records addRecordToList, deleteRecordFromList, updateCrmRecord | 160 | 7200 | 172800 |
| Importing multiple records asynchronously You can upload up to 100 records in each request. asyncAddRecordsToList, asyncDeleteRecordsFromList, asyncUpdateCampaignDispositions, asyncUpdateCrmRecords | 20 | 1000 | 2000 |
| Generating reports runReport | 16 | 120 | 800 |
| Retrieving reports You can retrieve up to 50,000 records in each report. For other reporting limits, refer to the Dashboard and Reporting User’s Guide. getReportResultCsv, getReportResult | 20 | 200 | 1000 |
| Uploading You can upload up to 50,000 records in each request. addNumbersToDnc, addToList, addToListCsv, addToListFtp, deleteAllFromList, deleteFromContactsCsv, deleteFromContactsFtp, deleteFromContacts, deleteFromList, deleteFromListCsv, deleteFromListFtp, updateContacts, updateContactsCsv, updateContactsFtp, updateDispositions, updateDispositionsCsv, and updateDispositionsFtp | 20 | 400 | 2000 |

[]()