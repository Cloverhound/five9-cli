## modifyUserCannedReports

Updates the list of canned reports associated with a specific user. To get a list of the user’s current canned reports, use `[getUsersInfo](/bundle/api-config-web/page/config-web-services/methods/user-management/getUsersInfo.htm)`.

#### modifyUserCannedReports

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| userName | string | Name of user. |
| cannedReportsToAdd | cannedReport [0..unbounded] | References to existing reports to associate with user. |
| cannedReportsToRemove | string [0..unbounded] | Names of reports to disassociate from user. |

#### modifyUserCannedReportsResponse

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| return | userInfo | Modified user information. |

[]()