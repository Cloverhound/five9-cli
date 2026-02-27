## isReportRunning

Checks whether a report is running and tracks changes in the status of a report. The method returns the response as soon as the report is completed.

                                                    **Important: **To prevent this method from overloading the system, Five9 recommends that you set the timeout parameter to at least five seconds. If report generation takes longer than the specified time-out, the method returns True, and the client resends the method.

#### isReportRunning

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| identifier | string | Unique identifier returned by runReport. |
| timeout | long | Required time in seconds to wait for changes before returning the current state. If set to 0, this method returns the result immediately. Cannot be empty. Long polling. |

#### isReportRunningResponse

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| return | boolean | Whether the report is running. True: Report still being generated. False: Report generation completed. |

[]()