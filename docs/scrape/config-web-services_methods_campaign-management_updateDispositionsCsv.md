## updateDispositionsCsv

Updates batches of disposition values in a campaign. Records are passed as a string of comma-separated values. The method returns an importIdentifier that you can use to query import state and import result. To check the status of the update, use the response identifier in the request of [getDispositionsImportResult](/bundle/api-config-web/page/config-web-services/methods/campaign-management/getDispositionsImportResult.htm).

#### updateDispositionsCsv

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| campaignName | string | Name of campaign. |
| DispositionsUpdateSettings Note: This is case sensitive. | dispositionsUpdateSettings | Update settings. |
| csvData | string | List of dispositions to be imported in CSV format. |

#### updateDispositionsCsvResponse

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| return | importIdentifier | Identifier for the requested dispositions import. This identifier can be used to check import status and outcome. |

[]()