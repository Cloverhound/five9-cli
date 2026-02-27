## asyncAddRecordsToList

Adds up to 100 records to a list. If the records do not exist in the contact database, they are added. Otherwise, the existing records are updated based on the list update settings. To check the status of the update, use the response identifier in the request of [getListImportResult](/bundle/api-config-web/page/config-web-services/methods/list-management/getListImportResult.htm).

                                                    **Important: **To insert a large number of records, use `addToList` or `addToListCsv`.

#### asyncAddRecordsToList

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| listName | string | Name of list. |
| listUpdateSettings | listUpdateSettings | List update settings. |
| importData | importData | Data to import in XML format. |
| resetDispositionsInCampaignsImportData Version 9.5 | string [0..unbounded] | Optional list of campaign names in which to reset the dispositions. |

#### asyncAddRecordsToListResponse

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| return | importIdentifier | Identifier for the imported data. Can be used to check import status and result. |

[]()