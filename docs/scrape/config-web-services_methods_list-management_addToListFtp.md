## addToListFtp

Imports by FTP a batch of list records. Depending on the settings, importing records may affect the contact database records. Records are passed as a string of comma-separated values. The method returns an identifier object that you can use to query the import status and result. For more information about the format of the file obtained from the FTP server, refer to the [Administrator’s Guide](https://documentation.five9.com/en-us/assets/documentation/vcc-applications/administrator/administrators-guide.pdf).

                                                    **Important: **Because this batch method affects the performance of the dialer and uses significant database resources, use this method only during off-peak periods. To insert single records while an outbound campaign is running, use `addRecordToListSimple` instead. To insert up to 100 records, use `asyncAddRecordsToList`.

#### addToListFtp

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| listName | string | Name of list. |
| updateSettings | listUpdateSettings | List update settings. |
| ftpSettings | ftpImportSettings | FTP settings. |

#### addToListFtpResponse

Empty.

[]()