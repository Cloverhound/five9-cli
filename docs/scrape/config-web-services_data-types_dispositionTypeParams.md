Contains the information returned by a dispositions update transaction.

                                                    | Name | Type | Description |
| --- | --- | --- |
| allowChangeTimer | boolean | Whether the agent can change the redial timer for this disposition. True: Agent can change redial timer. False: Agent cannot change redial timer. |
| attempts | byte | Number of redial attempts. |
| timer | timer | Redial timer. |
| useTimer | boolean | Whether this disposition uses a redial timer. True: Use a redial timer. False: Do not use a redial timer. |

[]()