Contains CRM update settings.

                                                    | Name | Type | Description |
| --- | --- | --- |
| FinalDisp | string | Any contact number of the contact is not dialed again by the current campaign. |
| FinalApplyToCampaigns | string | Contact is not dialed again by any campaign that contains the disposition. |
| AddActiveNumber | string | Adds the number dialed to the DNC list. |
| AddAndFinalize | string | Adds the call results to the campaign history. This record is no longer dialing in this campaign. Does not add the contact’s other phone numbers to the DNC list. |
| AddAllNumbers | string | Adds all the contact’s phone numbers to the DNC list. |
| DoNotDial | string | Number is not dialed in the campaign, but other numbers from the CRM record can be dialed. |
| RedialNumber | string | Number is dialed again when the list to dial is completed, and the dialer starts again from the beginning. |

[]()