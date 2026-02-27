The Five9 Configuration Web Services enable you to build secure applications to access, configure, manage, and synchronize call center information with work flow-management systems, such as lead management, CRM, and ERP systems.

This guide is intended for developers who understand these technologies and protocols:

  -

Client-server architecture and Web services

  -

SOAP, HTTP, and XML

  -

JSP, ASP, CGI, or Perl

  -

Computer-telephony integration concepts, processes, events, and call routing

  -

Overall call center integration and configuration

                                                    **Note: **
Five9 Configuration Web Services uses the swaRef.xsd to include binary content for types that are used to `upload` greetings in WAV format.

If you use Apache CXF, be aware that it does not support swaRef.xsd types. Before the SOAP envelope, it inserts text for which Five9 returns this exception:

```
Exception=javax.xml.ws.soap.SOAPFaultException: org.xml.sax.SAXParseException: Content is not allowed in prolog.
```

```
 
```

To prevent this issue, remove the SwaOutputInterceptor Apache CXF proxy class. However, with this fix, you cannot use Five9 methods that uses the swaRef types.

[]()