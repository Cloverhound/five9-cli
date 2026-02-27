## addToListCsv

#### Request

`listUpdateSettings`, highlighted in blue, extends `basicImportSettings`, highlighted in red.

```
<env:Envelope xmlns:env="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.admin.ws.five9.com/">
```

```
<env:Header/>
```

```
<env:Body>
```

```
	<ser:addToListCsv>
```

```
		<listName>hotleadslist</listName>
```

```
		<listUpdateSettings>
```

```
<fieldsMapping>
```

```
			   <columnNumber>1</columnNumber>
```

```
<fieldName>number1</fieldName>
```

```
<key>true</key>
```

```
			</fieldsMapping>
```

```
			<fieldsMapping>
```

```
			   <columnNumber>2</columnNumber>
```

```
<fieldName>first_name</fieldName>
```

```
<key>false</key>
```

```
			</fieldsMapping>
```

```
			<fieldsMapping>
```

```
			   <columnNumber>3</columnNumber>
```

```
<fieldName>last_name</fieldName>
```

```
<key>false</key>
```

```
			</fieldsMapping>
```

```
<reportEmail>admin@yourcompany.com</reportEmail>
```

```
			<separator>,</separator>
```

```
			<skipHeaderLine>true</skipHeaderLine>
```

```
<cleanListBeforeUpdate>false</cleanListBeforeUpdate>
```

```
			<crmAddMode>ADD_NEW</crmAddMode>
```

```
			<crmUpdateMode>UPDATE_ALL</crmUpdateMode>
```

```
			<listAddMode>ADD_ALL</listAddMode>
```

```
		</listUpdateSettings>
```

```
		<csvData>number1,first_name,last_name</csvData>
```

```
	</ser:addToListCsv>
```

```
</env:Body>
```

```
</env:Envelope>
```

#### Response

```
<env:Envelope xmlns:env="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.admin.ws.five9.com/">
```

```
<env:Header/>
```

```
<env:Body>
```

```
	<ser:addToListCsvResponse>
```

```
		<return>
```

```
			<identifier>0a2c9316-1a68-4be1-b817-c885326018c6</identifier>
```

```
		</return>
```

```
	</ser:addToListCsvResponse>
```

```
</env:Body>
```

```
</env:Envelope>
```

## []()addRecordToList

#### Request

```
<env:Envelope xmlns:env="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.admin.ws.five9.com/">
```

```
<env:Header/>
```

```
<env:Body>
```

```
	<ser:addRecordToList>
```

```
      <listName>some_list_name</listName>
```

```
      <listUpdateSettings>
```

```
        <fieldsMapping>
```

```
          <columnNumber>1</columnNumber>
```

```
          <fieldName>number1</fieldName>
```

```
          <key>true</key>
```

```
        </fieldsMapping>
```

```
        <fieldsMapping>
```

```
          <columnNumber>2</columnNumber>
```

```
          <fieldName>first_name</fieldName>
```

```
          <key>false</key>
```

```
        </fieldsMapping>
```

```
        <fieldsMapping>
```

```
          <columnNumber>3</columnNumber>
```

```
          <fieldName>last_name</fieldName>
```

```
          <key>false</key>
```

```
        </fieldsMapping>
```

```
        <separator>,</separator>
```

```
        <skipHeaderLine>false</skipHeaderLine>
```

```
        <callNowMode>ANY</callNowMode>
```

```
        <cleanListBeforeUpdate>false</cleanListBeforeUpdate>
```

```
        <crmAddMode>ADD_NEW</crmAddMode>
```

```
        <crmUpdateMode>UPDATE_FIRST</crmUpdateMode>
```

```
        <listAddMode>ADD_FIRST</listAddMode>
```

```
      </listUpdateSettings>
```

```
      <record>
```

```
        <fields>5551208111</fields>
```

```
        <fields>John</fields>
```

```
        <fields>Smith</fields>
```

```
      </record>
```

```
	</ser:addRecordToList>
```

```
</env:Body>
```

```
</env:Envelope>
```

#### Response

`listImportResult`, highlighted in blue, extends `basicImportResults`, highlighted in red.

```
<env:Envelope xmlns:env="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.admin.ws.five9.com/">
```

```
<env:Header/>
```

```
<env:Body>
```

```
	<ser:addRecordToListResponse>
```

```
		<return>
```

```
<uploadDuplicatesCount>0</uploadDuplicatesCount>
```

```
			<uploadErrorsCount>0</uploadErrorsCount>
```

```
<warningsCount/>
```

```
<crmRecordsInserted>0</crmRecordsInserted>
```

```
			<crmRecordsUpdated>1</crmRecordsUpdated>
```

```
			<listName>some_list_name</listName>
```

```
			<listRecordsDeleted>0</listRecordsDeleted>
```

```
			<listRecordsInserted>0</listRecordsInserted>
```

```
		</return>
```

```
	</ser:addRecordToListResponse>
```

```
</env:Body>
```

```
</env:Envelope>
```

## asyncAddRecordsToList

#### Request

```
<env:Envelope xmlns:env="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.admin.ws.five9.com/">
```

```
<env:Header/>
```

```
<env:Body>
```

```
	<ser:asyncAddRecordsToList>
```

```
      <listName>asdf</listName>
```

```
      <listUpdateSettings>
```

```
        	<fieldsMapping>
```

```
          <columnNumber>1</columnNumber>
```

```
          <fieldName>number1</fieldName>
```

```
          <key>true</key>
```

```
        	</fieldsMapping>
```

```
        	<fieldsMapping>
```

```
          <columnNumber>2</columnNumber>
```

```
          <fieldName>first_name</fieldName>
```

```
          <key>false</key>
```

```
        	</fieldsMapping>
```

```
        	<callTimeColumnNumber>3</callTimeColumnNumber>
```

```
        	<crmAddMode>ADD_NEW</crmAddMode>
```

```
        	<callNowMode>ANY</callNowMode>
```

```
        	<crmUpdateMode>UPDATE_FIRST</crmUpdateMode>
```

```
        	<listAddMode>ADD_IF_SOLE_CRM_MATCH</listAddMode>
```

```
      </listUpdateSettings>
```

```
      <importData>
```

```
        	<values>
```

```
          <item>6665554499</item>
```

```
          <item>George</item>
```

```
          <item>1341957101000</item>
```

```
        	</values>
```

```
        	<values>
```

```
          <item>9995554499</item>
```

```
          <item>Ringo</item>
```

```
          <item>1341957500000</item>
```

```
        	</values>
```

```
      </importData>
```

```
	</ser:asyncAddRecordsToList>
```

```
</env:Body>
```

```
</env:Envelope>
```

#### Response

```
<env:Envelope xmlns:env="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.admin.ws.five9.com/">
```

```
<env:Header/>
```

```
<env:Body>
```

```
    <ser:asyncAddRecordsToListResponse>
```

```
		<return>
```

```
			<identifier>some_string</identifier>
```

```
		</return>
```

```
	</ser:addRecordToListResponse>
```

```
</env:Body>
```

```
</env:Envelope>
```

## deleteFromContacts

#### Request

```
<env:Envelope xmlns:env="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.admin.ws.five9.com/">
```

```
<env:Header/>
```

```
<env:Body>
```

```
	<ser:deleteFromContacts>
```

```
		<crmDeleteSettings>
```

```
			<fieldsMapping>
```

```
			   <columnNumber>1</columnNumber>
```

```
			   <fieldName>number2</fieldName>
```

```
			   <key>true</key>
```

```
			</fieldsMapping>
```

```
			<reportEmail>admin@yourcompany.com</reportEmail>
```

```
			<separator>,</separator>
```

```
			<skipHeaderLine>false</skipHeaderLine>
```

```
			<crmDeleteMode>DELETE_ALL</crmDeleteMode>
```

```
		</crmDeleteSettings>
```

```
		<importData>
```

```
			<values>
```

```
			   <item>4155551234</item>
```

```
			</values>
```

```
			<values>
```

```
			   <item>5552654455</item>
```

```
			</values>
```

```
		</importData>
```

```
	</ser:deleteFromContacts>
```

```
</env:Body>
```

```
</env:Envelope>
```

#### Response

```
<env:Envelope xmlns:env="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.admin.ws.five9.com/">
```

```
<env:Header/>
```

```
   <env:Body>
```

```
      <ser:deleteFromContactsResponse>
```

```
         <return>
```

```
			  <identifier>91d4fb84-223f-49b9-8a12-c91484bc5c78</identifier>
```

```
         </return>
```

```
      </ser:deleteFromContactsResponse>
```

```
   </env:Body>
```

```
</env:Envelope>
```

## []()runReport

This example contains multiple `<criteria>` objects.

#### Request

```
<env:Envelope xmlns:xsd="http://www.w3.org/2001/XMLSchema"
```

```
			  xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
```

```
			  xmlns:tns="http://service.admin.ws.five9.com/"
```

```
		  	  xmlns:env="http://schemas.xmlsoap.org/soap/envelope/"
```

```
			  xmlns:ins0="http://jaxb.dev.java.net/array">
```

```
  <env:Body>
```

```
    <tns:runReport>
```

```
      <folderName>Shared Reports</folderName>
```

```
      <reportName>Test Report</reportName>
```

```
      <criteria>
```

```
        <time>
```

```
			<start>2019-04-16T20:22:17.173</start>
```

```
			<end>2019-05-08T02:37:42.4786435-05:00</end>
```

```
			</time>
```

```
        <reportObjects>
```

```
			<objectNames>Test Campaign1</objectNames>
```

```
			<objectNames>Test Campaign2</objectNames>
```

```
			<objectType>Campaign</objectType>
```

```
			</reportObjects>
```

```
		</criteria>
```

```
    </tns:runReport>
```

```
  </env:Body>
```

```
</env:Envelope>
```

#### Response

```
<env:Envelope xmlns:env="http://schemas.xmlsoap.org/soap/envelope/">
```

```
	<env:Header/>
```

```
		<env:Body>
```

```
			<ns2:runReportResponse
```

```
			     xmlns:ns2="http://service.admin.ws.five9.com/">
```

```
				<return>E0F0BC9A5544767BBrt0.c.ie.oA6DAA@7p90E4Bs7lCf4v49BcAm			</return>
```

```
			</ns2:runReportResponse>
```

```
	</env:Body>
```

```
</env:Envelope>
```

[]()