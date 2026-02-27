## getSkills

Returns information about each skill name that matches a pattern.

#### getSkills

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| skillNamePattern | string | Pattern of skill name. The skill name pattern is a regular expression. If omitted or equal to an empty string, all objects are returned. |

#### getSkillsResponse

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| return | skill [0..unbounded] | Skill information for each skill that matches the pattern. |

[]()