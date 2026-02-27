Contains the algorithm used by the ACD to determine to which agent to transfer a call in a specific queue.

                                                    | Name | Type | Description |
| --- | --- | --- |
| LongestReadyTime | string | Agent with the longest idle time since the last call. |
| LongestReadyTimeExcludeMC | string | Agent with the longest idle time, excluding manual calls, such as callbacks. |
| RoundRobin | string | Agent with the highest idle time among those logged into the queue. When agents log into the queue, they have the lowest idle time. The first queued call is delivered to the agent with the highest idle time. This agent is removed from the list. The process continues with the next agent with the highest idle time and the next queued call. |
| MinCallsHandled | string | Agent who has handled the fewest calls during the interval specified in distributionTimeFrame. |
| MinHandleTime | string | Agent who has the lowest total call handle time during the interval specified in distributionTimeFrame. |

[]()