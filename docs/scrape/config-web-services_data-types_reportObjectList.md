Contains the type of data to be included in the report.

                                                    | Name | Type | Description |
| --- | --- | --- |
| objectNames | string [0..unbounded] | Names of the objects. |
| objectType | wsObjectType | Possible filters that you can use in reports. |

### Example Usage

SOAP Request

                                                    [Copy](javascript:void(0);)

```
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.admin.ws.five9.com/">   <soapenv:Header/>   <soapenv:Body>      <ser:runReport>         <folderName>List Reports</folderName>         <reportName>List Details</reportName>         <criteria>            <!-- Zero or more repetitions: -->            <reportObjects>               <!-- Zero or more repetitions: -->               <objectNames>My List name</objectNames>               <objectType>List</objectType>            </reportObjects>            <time>               <start>2025-06-01T00:00:00Z</start>               <end>2025-06-30T12:00:00Z</end>            </time>         </criteria>      </ser:runReport>   </soapenv:Body></soapenv:Envelope>
```

#### Explanation

  -

`folderName`: Specifies the folder where the report is located. In this example, it's "[List Reports](/bundle/reports-dashboards/page/reports-dashboards/description-standard-and-social-reports/list-reports.htm)".

  -

`reportName`: The name of the report to run.  "List Details".

  -

**criteria**: Contains the criteria for the report.

    -

`reportObjects`: Specifies the objects to include in the report.

      -

`objectNames`: The names of the objects. In this example, "My List name".

      -

`ObjectType`: The type of the object, which is "List" in this case. The possible filters that you can use in reports are detailed in the [wsObjectType](/bundle/api-config-web/page/config-web-services/data-types/wsObjectType.htm).

 

  -

**time**: Specifies the time range for the report.

    -

`start`: The start time in ISO 8601 format.

    -

`end`: The end time in ISO 8601 format.

[]()