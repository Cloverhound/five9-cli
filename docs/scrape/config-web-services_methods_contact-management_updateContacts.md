## updateContacts

Updates batches of records in the contact database. Records are passed as a two-dimensional string array. To check the status of the update, use the response identifier in the request of `[getCrmImportResult](/bundle/api-config-web/page/config-web-services/methods/contact-management/getCrmImportResult.htm)`.

                                                    **Note: **In the process of creating and updating records, with multiple files/loads at once, if the same record is being deleted before the list is completed, there is a possibility that the deletion will occur before the update. This can result in the record being recreated with the new updates.

                                                    **Important: **Because this batch method affects the performance of the dialer and uses significant database resources, use this method only during off-peak periods. To update single records while an outbound campaign is running, use `updateCrmRecord` instead. To update up to 100 records, use `asyncUpdateCrmRecords`.

#### updateContacts

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| crmUpdateSettings | crmUpdateSettings | Options that determine how contact records are updated. |
| importData | importData | XML-formatted data to import. |

#### updateContactsResponse

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| return | importIdentifier | Identifier for the imported data. This identifier can be used to check import status and result. |

[]()