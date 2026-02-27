Your client sends Web service requests that are acknowledged by Five9 with a Web services response. Your client needs to check periodically for the status and outcome of the operations requested. The Web services API contains the XML-encoded SOAP methods used to communicate with your client application.

Five9 recommends that you use the latest version. Previous versions are still supported but are not recommended because they do not contain all the features. `Five9username` is the user name of the administrator:

                                                    **Note: **

In the rest of this guide, the features that apply to these versions are highlighted. Release 13 supports SOAP web services but adds no new features.

Use the appropriate data center for your location:

US: `api.five9.com`

UK: `api.five9.eu`

Canada: `api.five9.ca`

Frankfurt: `api.eu.five9.com`

  -

Version 13: `https://api.five9.com/wsadmin/v13/
AdminWebService?wsdl&user=***<Five9username>***`

  -

Version 12: `https://api.five9.com/wsadmin/v12/
AdminWebService?wsdl&user=***<Five9username>***`

  -

Version 11: `https://api.five9.com/wsadmin/v11/
AdminWebService?wsdl&user=***<Five9username>***`

  -

Version 10.2: `https://api.five9.com/wsadmin/v10_2/
AdminWebService?wsdl&user=***<Five9username>***`

  -

Version 10: `https://api.five9.com/wsadmin/v10/
AdminWebService?wsdl&user=***<Five9username>***`

  -

`Version 9.5`: `https://api.five9.com/wsadmin/v9_5/
AdminWebService?wsdl&user=***<Five9username>***`

  -

Version 9.3: `https://api.five9.com/wsadmin/v9_3/
AdminWebService?wsdl&user=***<Five9username>***`

  -

Version 4: `https://api.five9.com/wsadmin/v4/
AdminWebService?wsdl&user=***<Five9username>***`

  -

Version 3: `https://api.five9.com/wsadmin/v3/
AdminWebService?wsdl&user=***<Five9username>***`

  -

Version 2: `https://api.five9.com/wsadmin/v2/
AdminWebService?wsdl&user=***<Five9username>***`

  -

Default version (common to all versions): `https://api.five9.com/wsadmin/AdminWebService?wsdl&user=***<Five9username>***`

To ensure that connections are secure, send all requests by Transport Layer Security protocol (HTTPS) or VPN (IPSec or SSH) to the URL for your version, for example:

`https://api.five9.com/wsadmin[/***<version>***]/AdminWebService`

[]()