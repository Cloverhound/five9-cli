## getCampaignProfiles

Returns information about campaign profiles that match a pattern.

#### getCampaignProfiles

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| namePattern | string | Name of campaign profile or regular expression. If omitted, all profiles are returned. |

#### getCampaignProfilesResponse

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| return | campaignProfileInfo[0..unbounded] | Campaign profiles that match the pattern. |

[]()