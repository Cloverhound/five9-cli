## getSkillsInfo

Returns a list of users that possess the skill that matches a skill name pattern.

#### getSkillsInfo

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| skillNamePattern | string | Pattern of skill name. The skill name pattern is a regular expression. If omitted or equal to an empty string, all objects are returned. |

#### getSkillsInfoResponse

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| return | skillInfo [0..unbounded] | Skill information for each skill that matches the pattern. |

[]()