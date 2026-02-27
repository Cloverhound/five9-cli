## getCampaignState

Returns the state of a campaign and track changes in the state of a campaign by using long polling. The method returns the updated state or the current state after the time-out.

#### getCampaignState

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| campaignName | string | Name of campaign. |
| waitUntilChange | long | Optional duration in seconds to wait for changes. If omitted, the response is returned immediately. |

#### getCampaignStateResponse

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| return | campaignState | Current state of the campaign, which may be running, not running, starting, or stopping. |

[]()