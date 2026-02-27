## updateDispositions

Updates batches of disposition values in a campaign. Records are passed as a two-dimensional array (collection of strings).

Sets or resets a disposition for the specified records. Each request can contain up to 50,000 records. To check the status of the update, use the response identifier in the request of [getDispositionsImportResult](/bundle/api-config-web/page/config-web-services/methods/campaign-management/getDispositionsImportResult.htm).

#### updateDispositions

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| campaignName | string | Name of campaign. |
| DispositionsUpdateSettings Note: This is case sensitive. | dispositionsUpdateSettings | Update settings. |
| importData | importData | List of dispositions to be imported. |

#### updateDispositionsResponse

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| return | importIdentifier | Identifier for the requested dispositions import. This identifier can be used to check import status and outcome. |

[]()