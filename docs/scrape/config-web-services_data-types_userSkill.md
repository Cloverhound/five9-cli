Contains the queue assigned to a user.

                                                    | Name | Type | Description |
| --- | --- | --- |
| id | long | Queue ID. |
| userSkill.level level | int | Priority of agent to receive calls sent to this queue. Level 1 is the highest priority; level 10 is the lowest. Higher level receive calls first. When agents are assigned to several queues, each queue may have a different priority. |
| skillName | string | Queue name. |
| userName | string | User name assigned the queue. |

[]()