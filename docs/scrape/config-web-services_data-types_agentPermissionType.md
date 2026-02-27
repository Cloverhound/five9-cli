Contains all the permissions that can be assigned to an agent. This variable is a string.

                                                    **Important: **Agents who log into a PSTN station or forward calls between sessions may incur high long-distance fees. Only give permission to the appropriate users.

                                                    | Name | Type | Description |
| --- | --- | --- |
| ReceiveTransfer |  | Can receive transfers from other agents. |
| ProcessVoiceMail |  | Can process skill voicemail messages. |
| DeleteVoiceMail |  | Can delete voicemail messages. |
| TransferVoiceMail |  | Can transfer voicemail messages to other users and queues. |
| MakeRecordings |  | Can record calls or a portion of a call. |
| SendMessages |  | Can send messages to agents, administrators, and supervisors. |
| CreateChatSessions |  | Can start chat sessions with agents, administrators, and supervisors. |
| TrainingMode |  | Can initiate and receive a test call in the agent desktop. |
| CannotRemoveCRM |  | Cannot delete contact records. Five9 recommends that you assign this permission to all agents. |
| CannotEditSession |  | Cannot edit CRM session notes. |
| CallForwarding |  | Can enable call forwarding. |
| AddingToDNC |  | Can add numbers to your DNC list. Without this permission, the Add to DNC option in the agent desktop is unavailable, but agents can still use custom dispositions that add numbers to the DNC list. |
| DialManuallyDNC |  | Can manually dial the numbers listed in the DNC List. Without this permission, agents see an error message when they dial a number in the DNC list. |
| CreateCallbacks |  | Can create callback reminders. Without this permission, the Add Callback button in the agent desktop is disabled. |
| PlayAudioFiles |  | Can play prerecorded messages while on call or leave standard recorded messages on answering machines (Play Audio File button). Without this permission, the audio files in the agent desktop are hidden. |
| SkipCrmInPreviewDialMode |  | Can skip records In Preview Dialing Mode. Disabled by default. |
| ManageAvailabilityBySkill |  | Can select the skill groups or ACD queues in which to participate during the active session. |
| BrowseWebInEmbeddedBrowser |  | Can use the browser to open any URL from the agent desktop. |
| ChangePreviewPreferences |  | Can control when and where the preview window is displayed. |
| CanWrapCall |  | Can end a call and assign a disposition to it after spending time in the Wrap-Up state. Without this permission, the agent can end a call only by setting a disposition. |
| CanPlaceCallOnHold |  | Can place calls on hold. |
| CanParkCall |  | Can park a call. |
| CanRejectCalls |  | When auto-answer is disabled, the agent can decline incoming calls (typically inbound calls). To use this permission, be sure to enable CanConfigureAutoAnswer. |
| CanConfigureAutoAnswer |  | Can choose whether to auto-answer calls or be prompted before accepting a call. Enabled by default. |
| ScreenRecording | string | Do not use. |
| RecycleDispositionAllowed Version 3 | string | Enables agents to view and select the Recycle system disposition. |
| MakeTransferToAgents Version 3 | string | Can transfer calls to other agents. |
| MakeTransferToSkills Version 3 | string | Can transfer calls to skill groups. |
| MakeTransferToInboundCampaigns Version 3 | string | Can transfer calls to inbound campaigns. |
| MakeTransferToExternalCalls Version 3 |  | Can transfer calls to external phone numbers. |
| CreateConferenceWithAgents Version 3 |  | Can create a conference with other agents. |
| CreateConferenceWithSkills Version 3 |  | Can create a conference by using skill groups. |
| CreateConferenceWithInboundCampaigns Version 3 |  | Can create a conference with inbound campaigns. |
| CreateConferenceWithExternalCalls Version 3 |  | Can create a conference with external call participants. |
| MakeCallToAgents Version 3 |  | Can call other agents. |
| MakeCallToSkills Version 3 |  | Can call skill groups. |
| MakeCallToExternalCalls Version 3 |  | Can make external calls. |
| CanRunJavaClient Version 4 |  | Can use the Java client of the Agent Desktop Plus version. |
| CanRunWebClient Version 4 |  | Can use the Web client of the Agent Desktop Plus version. |
| CanViewMissedCalls Version 9.3 |  | Can view and return missed personal calls. |
| MakeCallToSpeedDialNumber Version 9.5 |  | Can call speed-dial numbers, including external numbers. |
| CreateConferenceWithSpeedDialNumber Version 9.5 |  | Can add speed-dial numbers to conference calls, including external numbers. |
| MakeTransferToSpeedDialNumber Version 9.5 |  | Can transfer calls to speed-dial numbers, including external numbers. |
| CanSelectDisplayLanguage Version 9.5 |  | Can select a language in the softphone settings. |
| CanViewWebAnalytics Version 9.3 |  | Can view and use web analytics. |
| CanTransferChatsToAgents Version 10 |  | Can transfer chats to agents. |
| CanTransferChatsToSkills Version 10 |  | Can transfer chats to queues. |
| CanTransferEmailsToAgents Version 10 |  | Can transfer email messages to agents. |
| CanTransferEmailsToSkills Version 10 |  | Can transfer emails to queues. |
| CanCreateChatConferenceWithAgents Version 10 |  | Can create chat conferences with agents. |
| CanCreateChatConferenceWithSkills Version 10 |  | Can create chat conferences with queues. |
| CanTransferSocialsToAgents Version 10 |  | Can transfer social interactions to agents. |
| CanTransferSocialsToSkills Version 10 |  | Can transfer social interactions to queues. |

[]()