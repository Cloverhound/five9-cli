Contains the options for email notifications.

                                                    | Name | Type | Description |
| --- | --- | --- |
| emailAddress | string | Email address to receive error messages. This address is used for bounced email messages and as the Reply-To address. |
| maxAttachmentSize | int | Maximum MB for a voicemail attachment. If a voicemail attachment exceeds the specified size, the email notification is sent without the attachment. |
| newUserNotification | boolean | Whether to send the login credentials to the newly created email address of a user. True: Send the login credentials. False: Do not send the login credentials. |

[]()