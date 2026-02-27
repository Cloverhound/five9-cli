## updateContactsCsv

Updates batches of records in CSV format. Records are passed as a two-dimensional string array. To check the status of the update, use the response identifier in the request of `[getCrmImportResult](/bundle/api-config-web/page/config-web-services/methods/contact-management/getCrmImportResult.htm)`.

                                                    **Important: **Because this batch method affects the performance of the dialer and uses significant database resources, use this method only during off-peak periods. To update single records while an outbound campaign is running, use `updateCrmRecord` instead. To update up to 100 records, use `asyncUpdateCrmRecords`.

#### updateContactsCsv

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| crmUpdateSettings | crmUpdateSettings | Options that determine how contact records are updated. |
| csvData | string | Data in CSV format to be imported to the Contacts database. |

#### updateContactsCsvResponse

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| return | importIdentifier | Identifier for the imported data. This identifier can be used to check import status and result. |

[]()