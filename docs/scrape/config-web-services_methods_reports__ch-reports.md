**Important: **When using the SOAP API reporting methods (e.g., `runReport`, `getReportResult`, `getReportResultCsv`, `isReportRunning`), avoid running reports frequently to retrieve near real-time data. Reports are not designed for real-time data extraction. Consider alternative approaches for real-time data needs.

Reports obtained with the API use the Pacific time zone, adjusted for daylight savings time if needed. For example, for users in the Central time zone, the report time is minus (-) 2 hours.

The methods `runReport` and `get*Report*` use the user’s configured locale. If no locale is configured, the system defaults to the domain’s locale. The system follows this process when searching for report folders or report names.

To generate and retrieve reports, use the methods in this order:

  -

Authenticate the reporting user.

  -

Send `runReport`.

  -

Send `isReportRunning` repeatedly until the response is false, which indicates that the report is completed.

  -

To obtain the results, send `getReportResults`.

[getReportResult](/bundle/api-config-web/page/config-web-services/methods/reports/getReportResult.htm)

[getReportResultCsv](/bundle/api-config-web/page/config-web-services/methods/reports/getReportResultCsv.htm)

[isReportRunning](/bundle/api-config-web/page/config-web-services/methods/reports/isReportRunning.htm)

[runReport](/bundle/api-config-web/page/config-web-services/methods/reports/runReport.htm)

[]()