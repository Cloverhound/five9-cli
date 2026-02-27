Contains the details for the work time after the end of a call.

                                                    | Name | Type | Description |
| --- | --- | --- |
| agentNotReady | boolean | Whether to automatically place agents who reach a call timeout in a Not Ready state. True: Set agents to Not Ready state. False: Do not set agents to Not Ready state. |
| dispostionName Note spelling. | string | Name of disposition automatically set for the call if the timeout is reached. |
| enabled | boolean | Whether to limit the wrap-up time of agents. True: Limit the wrap-up time. False: Do not limit the wrap-up time. |
| reasonCodeName | string | Not Ready reason code for agents who are automatically placed in Not Ready state after reaching the timeout. |
| timeout | timer | Time limit for agents in wrap-up mode. |

[]()