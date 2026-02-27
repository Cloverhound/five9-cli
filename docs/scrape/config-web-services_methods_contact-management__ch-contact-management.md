You can create, update, and manage contact records in your database using XML or CSV formats. These operations support both single-record updates and large-scale batch imports, giving you flexibility for real-time changes or scheduled bulk updates.

                                                    | Method | Purpose | Typical Use | Key Inputs | Returns |
| --- | --- | --- | --- | --- |
| asyncUpdateCrmRecords | Asynchronously update up to 100 records | Small near real-time batches | crmUpdateSettings, importData | importIdentifier |
| deleteFromContacts | Batch delete (XML) | Large deletes (off-peak) | crmDeleteSettings, importData | importIdentifier |
| deleteFromContactsCsv | Batch delete (CSV) | Large deletes (off-peak) | crmDeleteSettings, csvData | importIdentifier |
| deleteFromContactsFtp | Batch delete via FTP | Scheduled deletes | crmDeleteSettings, ftpImportSettings | Empty response |
| getContactRecords | Lookup records | Search before update/delete | crmLookupCriteria | contactsLookupResult |
| getCrmImportResult | Import status/result | Poll after async/batch | importIdentifier | crmImportResult |
| updateContacts | Batch upsert (XML) | Large upserts (off-peak) | crmUpdateSettings, importData | importIdentifier |
| updateContactsCsv | Batch upsert (CSV) | Large upserts (off-peak) | crmUpdateSettings, csvData | importIdentifier |
| updateContactsFtp | Batch upsert via FTP | Scheduled imports | crmUpdateSettings, ftpImportSettings | Empty response |
| updateCrmRecord | Single record upsert | Real-time single changes | crmUpdateSettings, record | crmImportResult |

[]()