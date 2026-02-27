## addToList

Imports batches of list records into a list. Depending on the settings, importing records may affect the contact database records. Records are passed as a two-dimensional array (collection of strings). To check the status of the update, use the response identifier in the request of `[getListImportResult](/bundle/api-config-web/page/config-web-services/methods/list-management/getListImportResult.htm)`.

addToList imports multiple records into a Five9 list by sending structured data (importData) along with mapping and update rules (listUpdateSettings). Choose it when you need to bulk-load or update contacts programmatically from CRMs, data pipelines, or scheduled jobs. Compared to addRecordToList (single record) and addToListCsv (CSV string payload), addToList is best for structured multi-record updates without embedding CSV.

Permissions

  -

Role/Permission required: Admin with *User can use Administrator Services* enabled.

If the permission is not granted, the operation will fail authorization.

                                                    **Important: **Because this batch method affects the performance of the dialer and uses significant database resources, use this method only during off-peak periods. To insert single records while an outbound campaign is running, use a`ddRecordToListSimple` instead. To insert up to 100 records, use `asyncAddRecordsToList`.

#### addToList

                                                    | Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| listName | string | Yes | Target list name. |
| listUpdateSettings | object listUpdateSettings | Yes | Controls dedupe, add/update modes, mapping, and optional call-now behavior; extends basicImportSettings |
| importData | array importData | Yes | Record data aligned to fieldsMapping.columnNumber. |

#### addToListResponse

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| return | importIdentifier | Identifier for the imported data. Can be used to check import status and result. |

### Parameters

Enter the parameters needed.

                                                    **Important: **Optional parameters must be removed and must not be left blank.

  -

listName (string) - Name of list.

  -

listUpdateSettings - This data type contains the list update settings. listUpdateSettings extends basicImportSettings.

  -

allowDataCleanup (boolean) - Whether to remove duplicate entries from a list.

  -

callbackFormat (string) - Optional. Specify the format of the call-back message in the HTTP content-type header

  -

callbackUrl (string) - Optional. URL of the HTTP callback

  -

countryCode (string) - Two-letter codes defined in ISO 3166-1.

  -

failOnFieldParseError (boolean) - Whether to stop the import if incorrect data is found.

  -

fieldsMapping - Mapping of the column number to the field name in the contact database.

  -

columnNumber (int) - Starting with 1, column number in a CSV file or importData array that contains data for the associated contact field.

  -

fieldName (string) - Name of the contact field associated with the column number.

  -

key (boolean) - Whether the key is used to find the record in the contact database.

  -

reportEmail (string) - Notification about import results is sent by email.

  -

separator (string) - Any ASCII character, such as a comma, used to separate entries in a list.

  -

skipHeaderLine (boolean) - Whether to omit the top row that contains the names of the fields.

  -

callNowColumnNumber (int) - Column number, starting with 1. If a record should be dialed immediately, the content of the column is 1, T, Y, and Yes. This column is not included in the fieldsMapping parameter and is not imported into the contact database. If callNowMode is also specified, only records that have a true value and apply to callNowMode are called immediately. If you do not want to add a column for each imported record, use callNowMode singly.

  -

callNowMode - Whether records are dialed immediately.

  -

NONE (string) - Default. No records are dialed immediately.

  -

NEW_CRM_ONLY (string) - Newly created CRM records are dialed immediately.

  -

NEW_LIST_ONLY (string) - New list records are dialed immediately even if the corresponding CRM records existed before the import.

  -

ANY (string) - All imported records are dialed immediately.

  -

callTime (long) - When to dial the records (Epoch time in milliseconds); applies to all records in the request, except for those with a value in the timeToCallColumn; does not apply to the addToList method, which is used to process batch record transactions.

  -

callTimeColumnNumber (int) - Column that contains the times (Epoch time) to call individual records. If a record contains a valid time, this time is used instead of the callTime parameter. Does not apply to the addToList method, which is used to process batch record transactions.

  -

cleanListBeforeUpdate (boolean) - Whether to remove all records in the list before adding new records.

  -

crmAddMode - This data type specifies whether a contact record is added to the contact database when

  -

a new record is added to a dialing list.

  -

ADD_NEW (string) - Contact records are created in the contact database and are added to the dialing list.

  -

DONT_ADD (string) - Records are added to the dialing list but no records are created in the contact database.

  -

crmUpdateMode - Describes how to update contact records when adding a record to a dialing list.

  -

UPDATE_FIRST (string) - Update the first matched record.

  -

UPDATE_ALL (string)  - Update all matched records. Does not apply to single record transactions, such as with the updateCrmRecord method.

  -

UPDATE_SOLE_MATCHES (string) - Update only if one matched record is found.

  -

DONT_UPDATE (string) - Do not update any record.

  -

listAddMode - Describes how to update the list.

  -

ADD_FIRST (string) - Adds the first record when multiple matches exist.

  -

ADD_ALL (string) - Add all records. Does not apply to asynchronous transactions, such as with the addRecordToList and asyncAddRecordsToList methods.

  -

ADD_IF_SOLE_CRM_ MATCH (string) - Add a record if only one match exists in the database.

  -

importData - This data type contains the data to be imported.

  -

values (stringarray) - Array to import. Each item corresponds to the fieldsMapping element specified in listUpdateSettings. Depending on the value of skipHeaderLine, the first record may not be read. If you would like more information, you can see basicImportSettings.

  -

item (string) - Value of a record that corresponds to a field specified in import settings.

                                                    **Note: **International dialing numbers with prefixes, eg. “+1” will only be read by e164 domains. If the value of the field “number1” is with the prefix and the domain is not an e164 domain, the contact will not be added to the list.

Example Request

                                                    [Copy](javascript:void(0);)

```
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.admin.ws.five9.com/">   <soapenv:Header/>   <soapenv:Body>      <ser:addToList>         <listName>{{List}}</listName>        <listUpdateSettings>            <fieldsMapping>                <columnNumber>1</columnNumber>                <fieldName>number1</fieldName>                <key>true</key>            </fieldsMapping>            <fieldsMapping>                <columnNumber>2</columnNumber>                <fieldName>number2</fieldName>                <key>false</key>            </fieldsMapping>            <fieldsMapping>                <columnNumber>3</columnNumber>                <fieldName>first_name</fieldName>                <key>false</key>            </fieldsMapping>            <fieldsMapping>                <columnNumber>4</columnNumber>                <fieldName>last_name</fieldName>                <key>false</key>            </fieldsMapping>            <separator>,</separator>            <skipHeaderLine>false</skipHeaderLine>            <callNowMode>ANY</callNowMode>            <cleanListBeforeUpdate>false</cleanListBeforeUpdate>            <crmAddMode>ADD_NEW</crmAddMode>            <crmUpdateMode>UPDATE_FIRST</crmUpdateMode>            <listAddMode>ADD_FIRST</listAddMode>        </listUpdateSettings>         <importData>            <values>                <item>5551208111</item>                <item>5551208121</item>                <item>John</item>                <item>Smith</item>            </values>            <values>                <item>5551208112</item>                <item>5551208122</item>                <item>John</item>                <item>Smith</item>            </values>            <values>                <item>5551208113</item>                <item>5551208123</item>                <item>John</item>                <item>Smith</item>            </values>         </importData>      </ser:addToList>   </soapenv:Body></soapenv:Envelope>
```

 

[]()