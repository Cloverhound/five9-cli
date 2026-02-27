`Version 9.5`

Contains the possible dialing rule methods for finding the allowed dialing days and hours.

                                                    | Name | Type | Description |
| --- | --- | --- |
|  | string | Possible values: FOLLOW_STATE_RESTRICTIONS: Follow state restrictions. INHERIT_DOMAIN_SETTINGS: Use domain settings. REGION: Use the dialing rules of the contact’s state. REGION_THEN_PHONE_NUMBER: Initially, use the dialing rules of the contact’s state. Afterward, search by using the phone number (area code and prefix). REGION_THEN_POSTCODE: Initially, use the dialing rules of the contact’s state. Afterward, search by using postal code. |

[]()