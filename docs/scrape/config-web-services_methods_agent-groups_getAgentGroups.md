## getAgentGroups

Returns a list of agent groups whose names match a string pattern.

#### getAgentGroups

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| groupNamePattern | string | Name of the group or regular expression. For all agent groups, use this pattern: .*. |

#### getAgentGroupsResponse

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| return | agentGroup[0..unbounded] | Groups that match the pattern. |

[]()