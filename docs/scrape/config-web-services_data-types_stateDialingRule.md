`Version 9.5`

Contains the possible methods for finding the state dialing rules that apply to the contact.

                                                    | Name | Type | Description |
| --- | --- | --- |
| REGION | string | Use the dialing rules of the contact’s state. |
| REGION_THEN_PHONE_NUMBER | string | Initially, use the dialing rules of the contact’s state. Afterward, search by using the phone number (area code and prefix). |
| REGION_THEN_POSTCODE | string | Initially, use the dialing rules of the contact’s state. Afterward, search by using postal code. |

[]()