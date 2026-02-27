This section contains the settings used for updating CRM records. These settings specify how records are added and updated in the system.

                                                    | Name | Type | Description |
| --- | --- | --- |
| basicImportSettings | crmUpdateSettings | Detailed information about the settings. See section on basicImportSettings. |
| crmAddMode | crmAddMode | Describes how to add a contact record. |
| crmUpdateMode | crmUpdateMode | Specifies how to update an existing contact record.Important: The UPDATE_ALL value of the crmUpdateMode parameter is not supported for the asyncAddrecordsToList and asyncUpdateCrmrecords methods. You cannot use the UPDATE_ALL mode to update all the fields of a contact record at once when using these methods. |

[]()