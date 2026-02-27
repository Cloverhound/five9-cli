Version 3

Contains list update settings for the `[addRecordToListSimple](/bundle/api-config-web/page/config-web-services/data-types/../methods/list-management/addRecordToListSimple.htm)` method.

                                                    | Name | Type | Description |
| --- | --- | --- |
| callAsap | boolean | Whether to call the contact as soon as possible. |
| countryCode countryCode Version 4 | string | Two-letter codes defined in ISO 3166-1. |
| fieldsMapping | fieldEntry [0..unbounded] | Mapping of the column number to the field name in the contact database. Column numbers start at 0 whereas in listUpdateSettings, column numbers start at 1. |
| timeToCall | long | When to dial the records (Epoch time in milliseconds). |
| updateCRM | boolean | Whether to update the contact field data of an existing record: True: Changes to the value of a contact field are saved. False: Changes to the value of a contact field are not saved, but new records are inserted. |

[]()