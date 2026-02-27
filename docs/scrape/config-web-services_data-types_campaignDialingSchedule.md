Contains the dialing configurations for a campaign profile.

                                                    | Name | Type | Description |
| --- | --- | --- |
| dialASAPSortOrder | dialSortOrder | Order for dialing numbers in the ASAP queue. |
| dialASAPTimeout | int | Duration before records that are not dialed are removed from the ASAP queue and are treated as normal records. |
| dialASAPTimeoutPeriod | timePeriod | Unit that specifies the dial ASAP time-out. |
| dialingOrder | campaignDialingOrder | Dialing order when contact records have multiple phone numbers. |
| dialingSchedules | campaignNumberSchedule [0..unbounded] | Time ranges used to call each of the three possible number associated with a campaign. |
| includeNumbers | campaignDialNumber [0..unbounded] | Whether to call each of the three numbers in the campaign associated with the profile. |

[]()