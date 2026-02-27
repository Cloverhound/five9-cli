## checkDncForNumbers

Checks whether phone numbers are part of a DNC list. The response contains the numbers found in the DNC list.

#### checkDncForNumbers

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| numbers | string [0..unbounded] | List of numbers to search in the DNC list. You may include up to 50000 phone numbers in a request. |

#### checkDncForNumbersResponse

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| return | string [0..unbounded] | Numbers found in the DNC list. |

[]()