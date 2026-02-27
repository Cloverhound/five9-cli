Your HTTPS requests must provide valid Five9 credentials for basic access authentication, which is used by the Web services API to enforce access control.

The credentials included in the header should be those of an administrator. Five9 recommends that you create several administrator users reserved for the Web Services if you plan to send multiple concurrent requests. Otherwise, if you try to use the same user name and password for multiple concurrent requests, your requests are denied, and you receive an exception. The administrator user must have the VCC Administrator role. Be sure to set the appropriate permissions for each administrator.

With HTTP basic authentication, the user name and password are encoded in base 64. In your client, construct your authorization header as follows:

  -

Concatenate the user name and password, for example:

```
ExampleUsername:ExamplePassword
```

  -

Encode the string in base 64, for example:

```
RXhhbXBsZVVzZXJOYW1lOkV4YW1wbGVQYXNzd29yZA==
```

  -

In your client, include the `Authorization` header with the value `Basic <base64-encoded string`>.

                                                    **Example: **

Web services header with encoded user name and password.

                                                        [Copy](javascript:void(0);)

```
POST https://api.five9.com/wsadmin/AdminWebService HTTP/1.1Accept-Encoding: gzip,deflateContent-Type: text/xml;charset=UTF-8SOAPAction: ""Authorization: Basic RXhhbXBsZVVzZXJOYW1lOkV4YW1wbGVQYXNzd29yZA==
```

[]()