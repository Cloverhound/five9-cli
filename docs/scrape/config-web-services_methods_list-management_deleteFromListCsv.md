## deleteFromListCsv

Deletes batches of records from a list. Records are passed as a string of comma-separated values.

To check the status of the update, use the response identifier in the request of `[getListImportResult](/bundle/api-config-web/page/config-web-services/methods/list-management/getListImportResult.htm)`.

                                                    **Important: **Because this batch method affects the performance of the dialer and uses significant database resources, use this method only during off-peak periods. To delete single records while an outbound campaign is running, use `deleteRecordFromList` instead. To delete up to 100 records, use `asyncDeleteRecordsFromList`. If you require a larger batch, contact your Five9 representative.

#### deleteFromListCsv

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| listName | string | Name of list. |
| listDeleteSettings | listDeleteSettings | List deletion settings. |
| csvData | string | Records to remove from the list in CSV format. |

#### deleteFromListCsvResponse

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| return | importIdentifier | Identifier for the deleted data. Can be used to check import status and result. |

[]()