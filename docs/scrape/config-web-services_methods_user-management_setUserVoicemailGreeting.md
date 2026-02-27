## setUserVoicemailGreeting

Uploads the audio file in WAV format provided by the Web Services user and to play the file when the user is not available to take calls. If a voicemail greeting already exists for the user, it is replaced. For more information about the WAV formats supported by the VCC, see the [Administrator’s Guide](https://documentation.five9.com/en-us/assets/documentation/vcc-applications/administrator/administrators-guide.pdf).

#### setUserVoicemailGreeting

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| userName | string | Name of user. |
| wavFile | base64Binary | File name. |

#### setUserVoicemailGreetingResponse

Empty.

[]()