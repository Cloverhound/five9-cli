## resetCampaignDispositions

Resets the dispositions of the campaign list records that match the dispositions. Calls that occurred during the date and time interval are reset so that the contacts can be dialed again if their disposition included in the list of dispositions.

#### resetCampaignDispositions

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| campaignName | string | Name of campaign. |
| dispositions | string [0..unbounded] | List of dispositions to reset. |
| after | dateTime | Start time of the call interval. The start time is not included in the interval. |
| before | dateTime | End time of the call interval. The end time is not included in the interval. |

#### resetCampaignDispositionsResponse

Empty.

[]()