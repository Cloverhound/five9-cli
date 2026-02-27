## listUpdateSettings

Contains the list update settings.

                                                    | Name | Type | Description |
| --- | --- | --- |
| callNowColumnNumber | int | Column number, starting with 1. If a record should be dialed immediately, the content of the column is 1, T, Y, and Yes. This column is not included in the fieldsMapping parameter and is not imported in the contact database. If callNowMode is also specified, only records that have a true value and apply to callNowMode are called immediately. If you do not want to add a column for each imported record, use callNowMode singly. |
| callNowMode | callNowMode | Whether records are dialed immediately. If callNowColumnNumber is also specified, callNowMode applies to all records with a true value in the specified column. If callNowColumnNumber is omitted, the callNowMode applies to all records imported by the request. |
| callTime | long | When to dial the records (Epoch time in milliseconds); applies to all records in the request, except for those with a value in the timeToCallColumn; does not apply to the addToList method, which is used to process batch record transactions. The call time value is applied only if the campaign exists when the record is added to the list assigned to that campaign. However, if a campaign is created or associated with a list after the record is added to the list, calls may be dialed sooner than the specified value, depending on the size of the list, the position of the record in the list, and the other parameters assigned to the list in the campaign. |
| callTimeColumnNumber | int | Column that contains the times (Epoch time) to call individual records. If a record contains a valid time, this time is used instead of the callTime parameter. Does not apply to the addToList method, which is used to process batch record transactions. |
| cleanListBeforeUpdate | boolean | Whether all records in the list should be removed before adding new records. True: Remove records before adding new ones. False: Do not remove records before adding new ones. |
| crmAddMode | crmAddMode | Whether contact records should be added when a new record is inserted into a dialing list. |
| crmUpdateMode | crmUpdateMode | Whether contact records should be updated when a record is added to a dialing list. |
| listAddMode | listAddMode | Describes how to update the list. |

[]()