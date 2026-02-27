## addDNISToCampaign

Adds DNIS (Dialed Number Identification Service) numbers to an inbound campaign. A DNIS is a phone number that can be dialed by a caller. When calls are received from that number, Five9 runs the IVR script associated with the campaign to which the DNIS has been added.

#### []()addDNISToCampaign

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| campaignName | string | Name of campaign. |
| DNISList | string [0..unbounded] | List of numbers to add to the campaign. |

#### addDNISToCampaignResponse

Empty.

[]()