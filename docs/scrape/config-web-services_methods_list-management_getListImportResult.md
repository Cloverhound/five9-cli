## getListImportResult

Returns the detailed outcome of a list import. The import is identified by the identifier object returned in the list import response.

#### getListImportResult

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| identifier | importIdentifier | Identifier returned in one of the import responses. |

#### getListImportResultResponse

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| return | listImportResult | Detailed description of the import results. |

Example

                                                    [Copy](javascript:void(0);)

```
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.admin.ws.five9.com/">   <soapenv:Header/>   <soapenv:Body>      <ser:getListImportResult>         <identifier>            <identifier>{{identifier}}</identifier>         </identifier>      </ser:getListImportResult>   </soapenv:Body></soapenv:Envelope>
```

[]()