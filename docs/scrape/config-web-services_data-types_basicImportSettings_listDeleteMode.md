## listDeleteMode

Contains the options for deleting records from a list.

                                                    | Name | Type | Description |
| --- | --- | --- |
| DELETE_ALL | string | Delete all records. Does not apply to single record transactions, such as with the deleteRecordFromList method. |
| DELETE_IF_SOLE_CRM_MATCH | string | Delete only if a single match is found in the database. |
| DELETE_EXCEPT_FIRST | string | Delete all records except the first matched record. |

[]()