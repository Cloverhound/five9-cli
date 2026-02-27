## isImportRunning

Checks whether data is being imported by using the importIdentifier object returned in the list import response. To track changes in the import process, use the `waitTime` parameter (long-polling technique). The method returns the new state when it is changed or the current state after the specified `waitTime`.

#### isImportRunning

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| identifier | importIdentifier | Identifier for the import session. |
| waitTime | long | Optional number of seconds to wait for changes. If not specified, This method contains the result immediately. |

#### isImportRunningResponse

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| return | boolean | Whether import is running. |

[]()