## addToListCsv

Imports batches of records into a list. Depending on the import settings, contact records may be affected. Records are passed as a string of comma-separated values. To check the status of the update, use the response identifier in the request of `[getListImportResult](/bundle/api-config-web/page/config-web-services/methods/list-management/getListImportResult.htm)`.

                                                    **Important: **Because this batch method affects the performance of the dialer and uses significant database resources, use this method only during off-peak periods. To insert single records while an outbound campaign is running, use `addRecordToListSimple` instead. To insert up to 100 records, use `asyncAddRecordsToList`.

#### addToListCsv

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| listName | string | Name of list. |
| listUpdateSettings | listUpdateSettings | List update settings. |
| csvData | string | Data to import. Fields are separated by commas; records are separated by new lines. |

#### addToListCsvResponse

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| return | importIdentifier | Identifier for the imported data. Can be used to check import status and result. |

[]()