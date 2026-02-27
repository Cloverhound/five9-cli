## setDefaultIVRSchedule

Assigns a default IVR script to a campaign.

#### []()setDefaultIVRSchedule

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| campaignName | string | Name of campaign. |
| scriptName | string | Name of the script. |
| params | scriptParameterValue[0..unbounded] | List of external variables to set for the IVR script before executing the script. |
| isVisualModeEnabled Version 3 | boolean | Whether Visual IVR is enabled in the campaign schedule. |
| isChatEnabled Version 3 | boolean | Whether chat is enabled in the campaign schedule. |

#### setDefaultIVRScheduleResponse

Empty.

[]()