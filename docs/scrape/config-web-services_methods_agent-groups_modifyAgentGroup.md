## modifyAgentGroup

Updates an agent group.

#### modifyAgentGroup

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| group | agentGroup | Group to be modified with the name of an existing group. If included, the new parameters replace those of the existing group. |
| addAgents | string [0..unbounded] | List of agent names to be added to the group. |
| removeAgents | string [0..unbounded] | List of agent names to be removed from the group. |

#### modifyAgentGroupResponse

Empty.

[]()