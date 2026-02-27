## deleteAllFromList

Deletes all records from a list.

To check the status of the update, use the response identifier in the request of `[getListImportResult](/bundle/api-config-web/page/config-web-services/methods/list-management/getListImportResult.htm)`.

                                                    **Important: **Because this batch method affects the performance of the dialer and uses significant database resources, use this method only during off-peak periods. To delete single records while an outbound campaign is running, use `deleteRecordFromList` instead. To delete up to 100 records, use `asyncDeleteRecordsFromList`. If you require a larger batch, contact your Five9 representative.

#### deleteAllFromList

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| listName | string | Name of list. |
| reportEmail | string | Email address for the deletion report. |
| callbackUrl Version 9.5 | string | URL for the HTTP callback. |
| callbackFormat Version 9.5 | callbackFormat | File format returned by the client. |

#### deleteAllFromListResponse

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| return | importIdentifier | Identifier for the deleted data. Can be used to check import status and result. |

[]()