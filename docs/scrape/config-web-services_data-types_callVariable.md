Contains the details of a call variable.

                                                    | Name | Type | Description |
| --- | --- | --- |
| applyToAllDispositions | boolean | Whether to use the variable for all dispositions. True: Use the variable for all dispositions. False: Do not use the variable for all dispositions. |
| defaultValue | string | Optional value that may be assigned to a call variable. Some data types (callVariableType) can have a default value. |
| description | string | Description of the variable. |
| dispositions | string [0..unbounded] | If applyToAllDispositions is False, this parameter lists the names of the dispositions for which to set this variable. |
| group | string | Group name of the call variable. |
| name | string | Name of the call variable. |
| reporting | boolean | Whether to add the values to reports: True: Variables are added to reports. False: Default. Variables are not added to reports. |
| restrictions | callVariableRestriction [0..unbounded] | Possible values of the variable. |
| sensitiveData Version 9.5 | boolean | Whether the variable contains personal data that identifies the customer. |
| type | callVariableType | One of the available types of call variables. |

[]()