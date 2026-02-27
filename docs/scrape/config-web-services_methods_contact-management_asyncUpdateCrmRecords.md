## asyncUpdateCrmRecords

Updates up to 100 records in the contact database.

To check the status of the update, use the response identifier in the request of `[getCrmImportResult](/bundle/api-config-web/page/config-web-services/methods/contact-management/getCrmImportResult.htm)`.

#### asyncUpdateCrmRecords

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| crmUpdateSettings | crmUpdateSettings | Update options. |
| importData | importData | List of records to update. |

#### asyncUpdateCrmRecordsResponse

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| return | importIdentifier | Identifier for the request. This identifier can be used to check status and result. |

[]()