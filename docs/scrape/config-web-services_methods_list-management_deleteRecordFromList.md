## deleteRecordFromList

Deletes a record from a list. Use one of the deleteFromContacts methods to delete a contact.

                                                    **Important: **Because this method affects the performance of the dialer, use this method only during off-peak periods. To delete a large number of records, use `deleteFromList` or `deleteFromListCsv` instead.

#### deleteRecordFromList

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| listName | string | Name of list. |
| listDeleteSettings | listDeleteSettings | List deletion settings. |
| record | recordData | Records to delete from the list. |

#### deleteRecordFromListResponse

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| return | listImportResult | Result of the deletion. |

[]()