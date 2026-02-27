## runReport

Use this method to generate a report. Because the time to generate a report varies, you need to follow up with `[isReportRunning](/bundle/api-config-web/page/config-web-services/methods/reports/isReportRunning.htm)` until it returns false. If you omit this step, the report that you retrieve may not contain all the data that you expect. When the report is completely generated, retrieve the data with `[getReportResult](/bundle/api-config-web/page/config-web-services/methods/reports/getReportResult.htm)` or `[getReportResultCsv](/bundle/api-config-web/page/config-web-services/methods/reports/getReportResultCsv.htm)`.

To reduce network traffic when calling runReport, Five9 recommends that you limit the criteria to a smaller time period to reduce the data returned. For example, to obtain data for a year, split the time period into months or weeks to return smaller data amounts instead of requesting one large report for an entire year.

#### runReport

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| folderName | string | Report category, such as Call Log Reports. For more information, see the Dashboards and Reports User’s Guide. |
| reportName | string | Name of the report, such as Call Log. |
| criteria | customReportCriteria | Filters used to generate the report. |

#### runReportResponse

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| return | string | Identifier used to check the status of the report (isReportRunning) and to retrieve the results (getReportResult). |

[]()