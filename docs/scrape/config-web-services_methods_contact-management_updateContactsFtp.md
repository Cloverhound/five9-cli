## updateContactsFtp

Imports batches of records by using FTP. For information about the format of the file obtained from the FTP server, see the [Administrator’s Guide](http://webapps.five9.com/assets/files/for_customers/documentation/vcc-applications/administrator/administrators-guide.pdf).

                                                    **Important: **Because this batch method affects the performance of the dialer and uses significant database resources, use this method only during off-peak periods. To update single records while an outbound campaign is running, use `updateCrmRecord` instead. To update up to 100 records, use `asyncUpdateCrmRecords`.

#### updateContactsFtp

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| updateSettings | crmUpdateSettings | Options that determine how contact records are updated. |
| ftpSettings | ftpImportSettings | FTP settings. |

#### updateContactsFtpResponse

This method contains no parameters.

[]()