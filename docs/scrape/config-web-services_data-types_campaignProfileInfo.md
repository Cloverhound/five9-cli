Contains the details of a campaign profile.

                                                    | Name | Type | Description |
| --- | --- | --- |
| ANI | string | ANI to send with outbound call. |
| description | string | Description of the profile. |
| dialingSchedule | campaignDialingSchedule | Order and time for dialing the numbers in a record. |
| dialingTimeout | int | Time to wait before disconnecting an unanswered call and logging it as No Answer. The default is 17 seconds. |
| initialCallPriority | int | Priority initially assigned to inbound and outbound calls on a scale of 1 to 100. Inbound calls have a default priority of 60. Calls with a higher priority are answered first, regardless of their time in a queue. To force calls from a campaign to be answered before those from other campaigns, increase the priority by 1. |
| maxCharges | int | Applies to inbound and outbound calls. Maximum dollar amount for long distance charges. The campaign stops automatically when this amount is reached. Zero means no limit. |
| name | string | Name of campaign profile. |
| numberOfAttempts | int | For outbound campaigns, number of dialing attempts for phone numbers in a list record, including redials due to disposition settings. |

[]()