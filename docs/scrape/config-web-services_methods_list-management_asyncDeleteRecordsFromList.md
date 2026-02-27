## asyncDeleteRecordsFromList

Deletes multiple records from a list. Use one of the deleteFromContacts methods to delete a contact. To check the status of the update, use the response identifier in the request of `[getListImportResult](/bundle/api-config-web/page/config-web-services/methods/list-management/getListImportResult.htm)`.

                                                    **Important: **Because this method affects the performance of the dialer, use this method only during off-peak periods. To delete a large number of records, use `deleteFromList` or `deleteFromListCsv` instead.

#### asyncDeleteRecordsFromList

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| listName | string | Name of list. |
| listDeleteSettings | listDeleteSettings | List update settings. |
| importData | importData | Records to be deleted. |

#### asyncDeleteRecordsFromListResponse

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| return | importIdentifier | Identifier for the imported data. Can be used to check import status and result. |

[]()