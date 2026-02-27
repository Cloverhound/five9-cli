## autodialCampaign

Contains information about autodial campaigns. It is used by `createAutodialResponse`, `getAutodialCampaignResponse`, and `modifyAutodialCampaign`

                                                    | Name | Type | Examples |
| --- | --- | --- |
| agentAvailability | agentAvailability | Types of agent states considered available. Autodial campaigns dial only agents that are available to process calls. If omitted, the list is dialed continuously regardless of agent availability. |
| agentSkillName | string | Numbers dialed only if agents with the specified skill are available. If empty when dialIfAgentsAvailable=True, agent availability is for any skill. When modifyAutodialCampaign specifies agentSkillName=null, the value remains. |
| defaultIvrSchedule | ivrScriptSchedule | Schedule of the IVR script that processes call flow. Required for inbound and autodial calls. |
| dialIfAgentsAvailable | boolean | Whether to dial numbers only if agents with agentSkillName are available. True: Dial only if agents are available in the skill group. False: Dial regardless of agent availability in the skill group. |
| maxNumOfLines | int | Maximum number of outbound phone lines dedicated to the campaign. |

[]()