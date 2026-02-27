## getUsersInfo

Returns information about roles and skills of each user name that matches a pattern.

#### getUsersInfo

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| userNamePattern | string | Regular expression that describes the user name pattern. If omitted or equal to an empty string, all objects are returned. For example, a pattern may be the first characters of the user’s name. |

                                                    **Warning: **Five9 recommends that you specify a string for the userNamePattern parameter as part of the method to reduce the number of objects processed and returned. If you do not specify a string, all user objects are returned. Depending on the number of user objects, the response time may result in performance degradation.

#### getUsersInfoResponse

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| return | userInfo [0..unbounded] | Information about each user name that matches the pattern. |

[]()