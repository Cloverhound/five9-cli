Contains the configuration of a contact field.

                                                    | Name | Type | Description |
| --- | --- | --- |
| displayAs | contactFieldDisplay | Display options for the data in the Agent desktop. |
| mapTo | contactFieldMapping | Map of the system information into the field. The field is updated when a disposition is set. |
| name | string | Name of the contact field. |
| restrictions | contactFieldRestriction [0..unbounded] | Restrictions imposed on the data that can be stored in this field. |
| system | boolean | Whether this field is set by the system or an agent. True: Field set by system. False: Field set by agent. |
| type | contactFieldType | Type of data stored in this field. |

[]()