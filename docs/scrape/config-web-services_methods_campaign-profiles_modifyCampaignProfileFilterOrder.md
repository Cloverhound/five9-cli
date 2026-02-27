## modifyCampaignProfileFilterOrder

Updates the criteria by which to dial the records of a campaign profile.

#### modifyCampaignProfileFilterOrder

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| campaignProfile | string | Name of campaign profile. |
| addOrderByField | orderByField[0..unbounded] | List, priority, and order of columns to sort. |
| removeOrderByField | string [0..unbounded] | Name of the contact field to remove from the filter order. |

#### modifyCampaignProfileFilterOrderResponse

Empty.

[]()