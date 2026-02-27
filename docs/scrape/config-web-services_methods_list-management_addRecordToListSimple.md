## addRecordToListSimple

                                                    | Version 3 |
| --- |
| Adds a record to a list instead of addRecordToList if you need only a few parameters. Because this method is simpler, it is faster than addRecordToList. If a corresponding record does not exist in the contact database, a new record is added. Otherwise, the existing record is updated according to the parameters. |

                                                    **Important: **To import large numbers of records, use `addToList` or `addToListCsv`.

Because `addRecordToListSimple` does not have a duplicate records flag, if the list record has more than two matches, the behavior is as follows:

  -

No contact field data is updated.

  -

No record is added to the list.

  -

Users receive this fault: More than one record matches specified criteria.

This behavior is equivalent to this:

  -

`crmUpdateMode = UPDATE_SOLE_MATCHES`

  -

`listAddMode = ADD_IF_SOLE_CRM_MATCH`

#### addRecordToListSimple

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| listName | string | Name of list. |
| listUpdateSimpleSettings | listUpdateSimpleSettings | List update settings. |
| record | recordData | Data to import. |

#### addRecordToListSimpleResponse

Empty.

[]()