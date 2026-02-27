Version 3

Contains detailed information about import problems.

                                                    | Name | Type | Description |
| --- | --- | --- |
| ParseError | string | Error message. |
| DuplicateKey | string | Record with the key field that already exists. |
| AllKeyFieldsAreEmpty | string | At least one key field should not be empty. |
| NoMatchesInContacts | string | No corresponding contact in the database. |
| OneMatchInContacts | string | Record being imported already exists in CRM. Occurs when CrmUpdateMode=DONT_UPDATE. |
| MultipleMatchesInContacts | string | Several contacts in the list have the same key value. Not allowed by the settings of another request. |
| InternalImportError | string | Undefined error. |

[]()