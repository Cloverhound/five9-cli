Contains global configuration options that are not found in other categories.

                                                    | Name | Type | Description |
| --- | --- | --- |
| defaultCampaign | string | Name of the campaign selected by default when agents start manual calls to external numbers. |
| enableReasonCodes | boolean | Whether agents can choose reason codes when selecting Not Ready and Logout. True: Agents can choose reason codes. False: Agents cannot choose reason codes. |
| internalCallTimeout | int | Number of seconds to wait for a call to be answered by another agent. The default value is 20 seconds. |
| maySelectCampaign | boolean | In the softphone menu, whether agents can select a campaign other than the default. True: Agents can select another campaign. False: Agents cannot select another campaign. |
| maySelectNone | boolean | Whether agents can make manual calls not associated with a campaign. True: Agents can make manual calls. False: Agents cannot make manual calls. |
| showDialAttempts | boolean | Whether agents can see call attempts automatically assigned a disposition by the dialer in the Contact Sessions panel. True: Agents can see call attempts. False: Agents cannot see call attempts. |
| voicemailTimeout | int | Number of seconds for an agent to wait before accepting a transferred skill group voicemail. If the agent does not accept the voicemail message within the set time, the voicemail message is transferred to the next agent in the skill group. |

[]()