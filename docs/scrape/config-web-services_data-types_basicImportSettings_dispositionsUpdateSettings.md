## dispositionsUpdateSettings

Contains the disposition update settings.

                                                    | Name | Type | Description |
| --- | --- | --- |
| commonDispositionValue | string | Disposition value when the same disposition is assigned to all records in the list. |
| dispositionColumnNumber | int | Column number for the disposition value of a record. The column is created when not assigning the same disposition for all records in the list. Column numbers with a range of [1, 256] for the disposition value of a record. The column is created when not assigning the same disposition for all records in the list. |
| dispositionsUpdateMode | dispositionsUpdateMode | Describes how dispositions are updated. |
| updateToCommonDisposition | boolean | For all records, whether to use the disposition value specified in commonDispositionValue. True: Use commonDispositionValue. False: Do not use commonDispositionValue. |
| warnIfNoCrmMatchFound | boolean | Whether to add a warning in the import transaction report when records do not match a CRM record. True: Warn when no CRM match exists. False: Do not warn when no CRM match exists. |

[]()