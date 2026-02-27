## addRecordToList

Adds a record to a list. If a corresponding record does not already exist in the contact database, a new record is added. Otherwise, the existing record is updated based on the options settings.

                                                    **Important: **To import large numbers of records, use `addToList` or `addToListCsv`.

#### addRecordToList

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| listName | string | Name of list. |
| listUpdateSettings | listUpdateSettings | List update settings. |
| record | recordData | Data to import. |

#### addRecordToListResponse

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| return | listImportResult | Result of the addition. |

[]()