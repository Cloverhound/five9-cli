Contains the attributes and permissions that can be assigned to an agent.

                                                    | Name | Type | Description |
| --- | --- | --- |
| alwaysRecorded | boolean | Whether the agent’s calls are recorded. True: Agent’s calls are recorded. False: Agent’s calls are not recorded. |
| attachVmToEmail | boolean | Whether the agent is allowed to attach a voicemail message to an email message. True: Agent is allowed. False: Agent is not allowed. |
| permissions | agentPermission [0..unbounded] | List of permissions that can be assigned to an agent. |
| sendEmailOnVm | boolean | Whether the agent is allowed to send an email message as a follow-up to a voicemail message. True: Agent is allowed. False: Agent is not allowed. |

[]()