Contains the list of supervisor permissions.

                                                    **Important: **Supervisors who log into a PSTN station or forward calls between sessions may incur high long-distance fees. Therefore, be sure to give permission only to the appropriate users.

   

                                                    | Name | Type | Description |
| --- | --- | --- |
| CanUseSupervisorSoapApi Version 12 | string | Can submit requests to the Supervisor API if the following permissions are enabled on the VCC Administration console: User Can Execute Web API Requests. Permission to execute the required request. For example, to edit dispositions with the Supervisor API, enable the corresponding permission: User Can Edit Dispositions. |
| Users | string | Can view the Users tab in the Supervisor desktop. |
| Agents | string | Can monitor the status of agents (logged out, on break, on a call) and view agent statistics and ACD queues. |
| CallMonitoring | string | Can listen to agents’ calls (silent monitoring). |
| Stations | string | Can view station IDs assigned to their domain. |
| ChatSessions | string | Can view active chat sessions. |
| Campaigns | string | Can view the Campaigns tab, including campaign status information and statistics, and other information and abilities. |
| CampaignManagement | string | Can be enabled for any of the single campaign management permissions at the bottom of this table: CampaignManagementStart, CampaignManagementStop, CampaignManagementReset, CampaignManagementResetDispositions, CampaignManagement ResetListPositions, and CampaignManagementResetDialerDCP. |
| AllSkills | string | If false, only the skill data assigned to the user is visible. It disables the Filter Statistics by Skill option in the Supervisor View menu. verify |
| BillingInfo | string | Whether the user can view the billing information: true or false. |
| BargeInMonitor | string | Can speak with the customer. Use this value when the supervisor needs to help but not take over the call from the agent. |
| WhisperMonitor | string | Can speak to the agent without being heard by the customer. Use this value when coaching agents or assisting with difficult calls. |
| ViewDataForAllAgentGroups | string | Can access all agent groups. When the permission is disabled, the supervisor can access only agent groups of which the supervisor is a member. |
| ReviewVoiceRecordings | string | Can access voicemail messages and recordings associated with each agent that the supervisor can access. |
| EditAgentSkills | string | Can add and remove skills and change skill levels for agents that the supervisor can view. |
| CanAccessDashboardMenu | string | Can access the Dashboard menu in the supervisor desktop. |
| CampaignManagementStart | string | Can start a campaign. |
| CampaignManagementStop | string | Can stop a campaign. |
| CampaignManagementReset | string | Can reset a campaign. |
| CampaignManagementResetDispositions | string | Can reset the dispositions of a campaign. |
| CanUseSupervisorSoapApi Version 12 | string | Can submit requests to the Supervisor SOAP API if the following permissions are enabled on the Administration application: User Can Execute Web API Requests. User can edit agent skills on the Administration application. |
| CampaignManagementResetListPositions | string | Can reset the dialing list position for outbound and autodial campaigns. |
| CampaignManagementResetAbandonCallRate | string | Can reset the dialer’s dropped call percentage for outbound and autodial campaigns. |
| CanViewTextDetailsTab Version 3 | string | Can view and log into the Text Details tab of the Supervisor desktop to access social media, email, and chat. |
| CanAccessShowFields Version 3 | string | Can use the View > Show Fields menu to set the layout of the application. |
| CanRunJavaClient Version 10 | string | Can run the Java client applications. |
| CanRunWebClient Version 10 | string | Can run the web client applications. |
| CanChangeDisplayLanguage Version 10.1 | string | Can change the display language. |
| CanMonitorIdleAgents Version 11 | string | Supervisors can monitor agents when agents are not on a call. |

[]()