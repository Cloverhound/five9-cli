## getListsInfo

Returns the names and sizes of all the dialing lists in a domain. The request contains the specified list name starts-with pattern. The response contains all the matching lists in the Five9 domain where the Five9 user ID in the HTTP header is located. For more information about the HTTP header, see Basic Access Authentication.

#### getListsInfo

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| listNamePattern | string | Pattern that is matched to the list names in the user’s domain. |

#### getListsInfoResponse

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| return | listInfo[0..unbounded] | All dialing lists in domain with the size of each list. |

[]()