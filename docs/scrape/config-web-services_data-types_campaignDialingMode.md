Contains the possible dialing modes.

                                                    | Name | Type | Description |
| --- | --- | --- |
| PREDICTIVE | string | Depending on campaign statistics, dials at a variable calls-to-agent ratio. For maximum agent use, predicts agent availability to begin dialing calls before an agent becomes ready for calls. |
| PROGRESSIVE | string | Depending on campaign statistics, dials at a variable calls-to-agent ratio when an agent becomes available. |
| TCPA PREVIEW | string | Enables the agent to review the contact details before dialing or skipping the record. To use the Preview-Only Manual-Dialing mode (for outbound campaigns only), you must set limitPreviewTime=True and dialNumberOnTimeout=False. For more information on setting related flags, see outboundCampaign. Note: For domains enabled for TCPA, only this dialing mode is allowed. |
| POWER | string | Dials at a fixed calls-to-agent ratio (1-to-1 or higher) when an agent becomes available. |

[]()