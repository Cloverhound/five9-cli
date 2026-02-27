Maps the column number to the contact database field name for each field imported into the dialing list and contact database.

                                                    | Name | Type | Description |
| --- | --- | --- |
| columnNumber | int | Column numbers with a range of [1,256]. Column number in a CSV file or importData array that contains data for the associated contact field. |
| fieldName | string | Name of the contact field associated with the column number. |
| key | boolean | Whether the key is used to find the record in the contact database. When a record needs to be updated, the key is used to find the record to update in the contact database. For example, the key can be first_name, first_name+last_name, Number1, or a combination. When a record is added, the value of the key determines if the record already exists. If so, the values in crmAddMode, crmUpdateMode, and listAddMode determine how to handle matching records. True: Use the key. False: Do not use the key. |

[]()