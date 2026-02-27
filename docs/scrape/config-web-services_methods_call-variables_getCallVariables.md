## getCallVariables

Returns information about a group of call variables.

#### getCallVariables

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| namePattern | string | Name of call variables group or regular expression. If omitted, all call variables are returned. |
| groupName | string | Group name of call variables. |

#### getCallVariablesResponse

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| return | callVariable[0..unbounded] | Call variables that match the pattern. |

[]()