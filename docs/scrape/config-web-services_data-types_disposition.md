Contains the attributes of a custom disposition.

                                                    | Name | Type | Description |
| --- | --- | --- |
| agentMustCompleteWorksheet | boolean | Whether the agent needs to complete a worksheet before selecting a disposition. True: Worksheet required. False: Worksheet not required. |
| agentMustConfirm | boolean | Whether the agent is prompted to confirm the selection of the disposition. True: Agent prompted to confirm disposition. False: Agent not prompted to confirm disposition. |
| description | string | Description of the disposition. |
| name | string | Name of the disposition. |
| resetAttemptsCounter | boolean | Whether assigning the disposition resets the number of dialing attempts for this contact. True: Disposition resets the counter to zero. False: Disposition does not reset the counter to zero. |
| sendEmailNotification | boolean | Whether call details are sent as an email notification when the disposition is used by an agent. True: Send email notification. False: Do not send email notification. |
| sendIMNotification | boolean | Whether call details are sent as an instant message in the Five9 system when the disposition is used by an agent. True: Send instant message. False: Do not send instant message. |
| trackAsFirstCallResolution | boolean | Whether the call is included in the first call resolution statistics (customer’s needs addressed in the first call). Used primarily for inbound campaigns. True: Include in statistics. False: Do not include in statistics. |
| type | customDispositionType | Disposition type. |
| typeParameters | dispositionTypeParams | Parameters that apply to the disposition type. |

[]()