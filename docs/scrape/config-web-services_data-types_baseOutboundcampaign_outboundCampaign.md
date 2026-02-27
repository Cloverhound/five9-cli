## outboundCampaign

Contains information about outbound campaigns. All data types used by `createOutboundCampaign` are listed below. Depending on the campaign mode and your goal, you can use these flags as follows:

                                                    | Goal | limitPreviewTime | dialNumberOnTimeout | previewDialImmediately |
| --- | --- | --- | --- |
| Unlimited preview time | false |  | false |
| Limited preview time | true |  | false |
| Dial number | true | true | false |
| Set agent to not ready | true | false | false |
| Dial immediately |  |  | true |

This figure shows the relationship between the groups of elements.

![](https://documentation-be.five9.com/bundle/api-config-web/page/config-web-services/data-types/baseOutboundcampaign/../images/composites/baseOutboundCampaign.gif?_LANG=enus)

                                                    | Name | Type | Description |
| --- | --- | --- |
| actionOnAnswerMachine | campaignDialingAction | Action to take when the answering machine is detected. |
| actionOnQueueExpiration | campaignDialingAction | Action to take when the maximum queue time expires, which occurs when no agent is available to take a call. |
| callAnalysisMode | callAnalysisMode | Types of attempts when a call is answered. |
| callsAgentRatio | double | For campaigns in the power mode, count of phone numbers dialed for an agent ready for a new call. |
| dialNumberOnTimeout | boolean | For campaigns in the preview mode, use as follows: True: Dial number when preview time expires. False: Set agent to Not Ready state. |
| dialingMode | campaignDialingMode | Types of dialing modes. |
| dialingPriority Version 9.5 | int | Method to set the dialing priority of the running outbound campaign by assigning a priority to each campaign. The default for all campaigns is 3; the range is 1 to 99. To enable this feature, see campaignsSettings. |
| dialingRatio Version 9.5 | int | Method to set the dialing priority of the running outbound campaign by using a ratio (1 to 99). The default is 50. To enable this feature, see campaignsSettings. |
| distributionAlgorithm | distributionAlgorithm | Method used by the ACD to transfer calls to agents. |
| distributionTimeFrame | distributionTimeFrame | Time intervals used by distributionAlgorithm. |
| limitPreviewTime | boolean | For campaigns in the preview mode, use as follows: True: Dial contact number after maxPreviewTime. False: Allow agents to preview the contact number for an unlimited time. |
| maxDroppedCallsPercentage | float | Maximum allowed percentage of dropped calls. Use when monitorDroppedCalls is set to true. |
| maxPreviewTime | timer | Duration until expiration of the preview time. |
| maxQueueTime | timer | Maximum time allowed for calls in a queue. |
| monitorDroppedCalls | boolean | Whether to keep track of the dropped call percentage of the campaign. True: Keep track of the dropped calls for the last 30 days or since the last reset. False: Do not keep track of the dropped calls. |
| previewDialImmediately | boolean | For outbound campaigns in preview mode, use as follows: True: Automatically dial the number without waiting for an action from the agent. False: Do not dial the number automatically. |
| useTelemarketingMaxQueTimeEq1 Version 9.5 | boolean | Whether to enable maximum queue time for telemarketing in campaigns to one second. |
| generalCampaign |  |  |
| analyzeLevel | int | Voice detection level for an answering machine. The values range from 0 (fast detection) to 11 (accurate detection). |
| CRMRedialTimeout | timer | Minimum time before redialing a contact record after all numbers for the contact record have been dialed or skipped. The default is 10 minutes. |
| dnisAsAni | boolean | When transferring calls to third parties, whether to override the default DNIS of the domain by using the contact’s phone number (ANI) as the DNIS (caller ID). True: Override the default DNIS. False: Do not override the default DNIS. |
| enableListDialingRatios | boolean | Whether to use list dialing ratios, which enable multiple lists to be dialed at specified frequencies. True: Enable dialing ratios. False: Do not enable dialing ratios. |
| listDialingMode | listDialingMode | Describes the list dialing mode. |
| noOutOfNumbersAlert | boolean | When an outbound campaign runs out of numbers to dial, whether to turn off notification messages to administrators and supervisors that the campaign is no longer dialing because the lists are complete. True: Turn off notification messages. False: Do not turn off notification messages. |
| stateDialingRule Version 9.5 | campaignStateDialingRule | How dialing rule options are used in the campaign. |
| timeZoneAssignment Version 9.5 | campaignTimeZoneAssignment | How time zone are assigned the campaign. |
| campaign |  |  |
| autoRecord | boolean | Whether to record all calls of the campaign. True: Record all calls. False: Do not record all calls. |
| callWrapup | campaignCallWrapup | Details for the work time after the call. |
| ftpHost | string | Host name of the FTP server. |
| ftpPassword | string | Password of the FTP server. |
| ftpUser | string | User name for the FTP server. |
| recordingNameAsSid | boolean | For FTP transfer, whether to use the session ID as the recording name. True: Use the session ID as recording name. False: Do not use the session ID as recording name. |
| useFtp | boolean | Whether to use FTP to transfer recordings. True: Use FTP to transfer recordings False: Do not use FTP to transfer recordings. |

[]()