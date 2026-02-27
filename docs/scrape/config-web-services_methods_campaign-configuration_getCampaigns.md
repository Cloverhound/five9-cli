## getCampaigns

Returns a list of campaigns whose names match a string pattern. If no name is specified, all the existing campaigns in the domain are returned.

#### []()getCampaigns

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| campaignNamePattern | string | Name of the campaign or regular expression that matches several campaign names. For example, for all campaigns, use this pattern: .*. |
| campaignType | campaignType | Type of campaign: inbound, outbound, or autodial. |

#### getCampaignsResponse

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| return | campaign[0..unbounded] | Name and basic attributes of the campaign. |

[]()