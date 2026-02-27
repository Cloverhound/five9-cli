Contains the settings for updating all lists and contacts.

                                                    | Name | Type | Description |
| --- | --- | --- |
| allowDataCleanup | boolean | Whether to remove duplicate entries from a list. True: Remove duplicate entries. False: Ignore duplicate entries. |
| callbackFormat Version 9.5 | callbackFormat | File format of the HTTP callback message. |
| callbackUrl Version 9.5 | string | URL of the callback passed in the batch method. |
| countryCode countryCodeVersion 4 | string | Two-letter codes defined in ISO 3166-1. |
| failOnFieldParseError Version 3 | boolean | Whether to stop the import if incorrect data is found: True: The record is rejected when at least one field fails validation. False: Default. The record is accepted. However, changes to the fields that fail validation are rejected. |
| fieldsMapping | fieldEntry [1..unbounded] | Mapping of the column number to the field name in the contact database. For listUpdateSettings, the column number starts at 1, whereas for listUpdateSimpleSettings, the column number starts at 0. |
| reportEmail | string | Notification about import results is sent to the email addresses that you set for your application. See also basicImportResult. The following methods ignore the value of reportEmail: addRecordToList addRecordToListSimple deleteRecordFromList updateCrmRecord asyncAddRecordsToList asyncDeleteRecordsFromList asyncUpdateCrmRecords asyncUpdateCampaignDispositions |
| separator | string | Any ASCII character, such as a comma, used to separate entries in a list. |
| skipHeaderLine | boolean | Whether to omit the top row that contains the names of the fields. True: Omit the top row. False: Include the top row. |

These data types use `basicImportSettings`:

[crmDeleteSettings](/bundle/api-config-web/page/config-web-services/data-types/basicImportSettings/crmDeleteSettings.htm)

[crmUpdateSettings](/bundle/api-config-web/page/config-web-services/data-types/basicImportSettings/crmUpdateSettings.htm)

[dispositionsUpdateSettings](/bundle/api-config-web/page/config-web-services/data-types/basicImportSettings/dispositionsUpdateSettings.htm)

[listDeleteMode](/bundle/api-config-web/page/config-web-services/data-types/basicImportSettings/listDeleteMode.htm)

[listUpdateSettings](/bundle/api-config-web/page/config-web-services/data-types/basicImportSettings/listUpdateSettings.htm)

This figure shows the relationship between the data types.

![](https://documentation-be.five9.com/bundle/api-config-web/page/config-web-services/data-types/basicImportSettings/../images/composites/basicImportSettings.gif?_LANG=enus)

[]()