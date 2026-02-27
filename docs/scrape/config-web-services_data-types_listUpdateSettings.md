Contains the list update settings. `listUpdateSettings` extends `[basicImportSettings](/bundle/api-config-web/page/config-web-services/data-types/basicImportSettings/_ch-basicImportSettings.htm)`.

                                                    | Name | Type | Description |
| --- | --- | --- |
| callNowColumnNumber | int | Column numbers with a range of [1,256]. If a record should be dialed immediately, the content of the column is 1, T, Y, and Yes. This column is not included in the fieldsMapping parameter and is not imported in the contact database. If callNowMode is also specified, only records that have a true value and apply to callNowMode are called immediately. If you do not want to add a column for each imported record, use callNowMode singly. |
| callNowMode | callNowMode | Whether records are dialed immediately. If callNowColumnNumber is also specified, callNowMode applies to all records with a true value in the specified column. If callNowColumnNumber is omitted, the callNowMode applies to all records imported by the request. |
| callTime | long | When to dial the records (Epoch time in milliseconds); applies to all records in the request, except for those with a value in the timeToCallColumn; does not apply to the addToList method, which is used to process batch record transactions. Note: The call time value is applied only if the campaign exists when the record is added to the list assigned to that campaign. However, if a campaign is created or associated with a list after the record is added to the list, calls may be dialed sooner than the specified value, depending on the size of the list, the position of the record in the list, and the other parameters assigned to the list in the campaign. |
| callTimeColumnNumber | int | Column numbers with a range of [1,256]. Column that contains the times (Epoch time) to call individual records. If a record contains a valid time, this time is used instead of the callTime parameter. Does not apply to the addToList method, which is used to process batch record transactions. |
| cleanListBeforeUpdate | boolean | Whether to remove all records in the list before adding new records. True: Remove all records. False: Do not remove all records. |
| crmAddMode | crmAddMode | Describes how to add new contact records into a dialing list. |
| crmUpdateMode | crmUpdateMode | Describes how to update contact records when adding a record to a dialing list. |
| listAddMode | listAddMode | Describes how to update the list. |

[]()