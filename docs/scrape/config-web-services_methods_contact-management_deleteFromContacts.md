## deleteFromContacts

Deletes batches from the contact database based on the specified import settings. Records are passed as a two-dimensional string array.

To check the status of the update, use the response identifier in the request of `[getCrmImportResult](/bundle/api-config-web/page/config-web-services/methods/contact-management/getCrmImportResult.htm)`.

                                                    **Important: **Because this batch method affects the performance of the dialer and uses significant database resources, use this method only during off-peak periods. To delete single records while an outbound campaign is running, use `deleteRecordFromList` instead. To delete up to 100 records, use `asyncDeleteRecordsFromList`. If you require a larger batch, contact your Five9 representative.

#### deleteFromContacts

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| crmDeleteSettings | crmDeleteSettings | Deletion options. |
| importData | importData | List of XML-formatted records to delete. |

#### deleteFromContactsResponse

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| return | importIdentifier | Identifier for the request. This identifier can be used to check status and result. |

[]()