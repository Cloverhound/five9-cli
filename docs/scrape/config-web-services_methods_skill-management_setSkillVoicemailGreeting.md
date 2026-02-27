## setSkillVoicemailGreeting

Uploads the audio file in WAV format provided by the Web Services user and associates it with a skill to be played when no agents are available to take calls. If a voicemail greeting already exists for the skill, it is replaced. For more information about the supported WAV formats, refer to the [Administrator’s Guide](https://documentation.five9.com/en-us/assets/documentation/vcc-applications/administrator/administrators-guide.pdf).

#### setSkillVoicemailGreeting

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| skillName | string | New prompt attributes. |
| wavFile | base64Binary | Audio file that contains the recording. For more information, refer to the XML definition. |

#### setSkillVoicemailGreetingResponse

Empty.

[]()