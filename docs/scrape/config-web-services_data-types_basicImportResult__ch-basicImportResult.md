Contains information about rejected records.

                                                    | Name | Type | Description |
| --- | --- | --- |
| failureMessage | string | Description of the error sent to your email address or in the newly added callback messages. |
| importIdentifier Version 9.5 | string | Identifier used to check the status and outcome of a data import operation. |
| importTroubles Version 3 | importTrouble | Reason for the rejection. |
| keyFields Version 3 | string | Set of fields marked as keys that define the rejected record. |
| success Version 9.5 | boolean | Whether the request was successful. |
| uploadDuplicatesCount | long | Number of duplicate contact records inserted. |
| uploadErrorsCount | long | Number of errors in the updated contact records. |
| warningsCount | entry [0..unbounded] | Number of warnings associated with the imported data. |

In addition, these data types use `basicImportRresult`:

[crmImportResult](/bundle/api-config-web/page/config-web-services/data-types/basicImportResult/crmImportResult.htm)

[dispositionsImportResult](/bundle/api-config-web/page/config-web-services/data-types/basicImportResult/dispositionsImportResult.htm)

[listImportResult](/bundle/api-config-web/page/config-web-services/data-types/basicImportResult/listImportResult.htm)

This figure shows the relationship between the data types.

![](https://documentation-be.five9.com/bundle/api-config-web/page/config-web-services/data-types/basicImportResult/../images/composites/basicImportResult.gif?_LANG=enus)

[]()