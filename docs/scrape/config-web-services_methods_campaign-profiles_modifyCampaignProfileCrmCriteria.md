## modifyCampaignProfileCrmCriteria

Updates the CRM filters of a campaign profile.

#### modifyCampaignProfileCrmCriteria

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| profileName | string | Name of campaign profile. |
| grouping | crmCriteriaGrouping | Filter relationship between the criteria. |
| addCriteria | campaignFilterCriterion[0..unbounded] | Criteria to add to the profile. |
| removeCriteria | campaignFilterCriterion[0..unbounded] | Criteria to remove from profile. |

#### modifyCampaignProfileCrmCriteriaResponse

Empty.

[]()