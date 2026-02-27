## modifyCampaignProfileDispositions

Updates the disposition count limits for a campaign. If a campaign with this profile reaches the maximum count for a disposition, the campaign stops automatically.

#### modifyCampaignProfileDispositions

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| profileName | string | Name of campaign profile. |
| addDispositionCounts | dispositionCount[0..unbounded] | Limits for the number of dispositions. |
| removeDispositionCounts | dispositionCount[0..unbounded] | Disposition limits to remove from the profile. |

#### modifyCampaignProfileDispositionsResponse

Empty.

[]()