`Version 9.5`

Contains the possible values of the `X-Frame-Options` header for your Visual IVR pages.

                                                    | Name | Type | Description |
| --- | --- | --- |
|  | string | Possible values: DENY: No display even if the frame is located in the same domain as the page. Cannot start a Visual IVR script with this value. Use this value if you do not need to place a Visual IVR link in an iframe. Most secure option. SAME_ORIGIN: Display only if the frame is located in the same domain as the page. Mostly for internal use because customers cannot access internal Web pages. Example: The HTTP header of http://shop.example.com/confirm.asp contains X-FRAME-OPTIONS: SAME_ORIGIN. Any frame in the http://shop.example.com domain can be displayed. ALLOW_FROM: Display only if the frame is located in the domain that you specify in the field. Example: The HTTP header of http://shop.example.com/confirm.asp contains X-FRAME-OPTIONS: ALLOW_FROM https://partner.affiliate.com. The page may be framed only by pages in the https://partner.affiliate.com domain. |

[]()