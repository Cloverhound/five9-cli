## deleteFromListFtp

Removes batches of records from a list by FTP. For more information about the format of the file obtained from the FTP server, refer to the [Administrator’s Guide](https://documentation.five9.com/en-us/assets/documentation/vcc-applications/administrator/administrators-guide.pdf).

                                                    **Important: **Because this batch method affects the performance of the dialer and uses significant database resources, use this method only during off-peak periods. To delete single records while an outbound campaign is running, use `deleteRecordFromList` instead. To delete up to 100 records, use `asyncDeleteRecordsFromList`. If you require a larger batch, contact your Five9 representative.

#### deleteFromListFtp

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| listName | string | Name of list. |
| deleteSettings | listDeleteSettings | List deletion settings. |
| ftpSettings | ftpImportSettings | FTP settings. |

#### deleteFromListFtpResponse

Empty.

[]()