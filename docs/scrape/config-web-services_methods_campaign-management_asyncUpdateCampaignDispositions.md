## asyncUpdateCampaignDispositions

`Version 9.5`

Updates campaign dispositions asynchronously.

Use this method to update the disposition of a small number of records. The method returns an `importIdentifier` object that you can use to query the import status and result. To check the status of the update, use the response identifier in the request of `[getDispositionsImportResult](/bundle/api-config-web/page/config-web-services/methods/campaign-management/getDispositionsImportResult.htm)`.

#### asyncUpdateCampaignDispositions

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| campaignName | string | Name of campaign. |
| dispositionsUpdateSettings | dispositionsUpdateSettings | Update settings. |
| importData | importData | List of dispositions to be imported. |

#### asyncUpdateCampaignDispositionsResponse

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| return | importIdentifier | Identifier for the imported data. Use this identifier to check the import status and result. |

[]()