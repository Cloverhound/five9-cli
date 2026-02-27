## setDialingRules

Creates, modifies, or deletes dialing rules.

#### setDialingRules

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| setRules | dialingRule[0..unbounded] | Dialing rules to create or modify. If a rule exists with the same name, it is replaced with the new configuration. |
| removeRules | string [0..unbounded] | Names of the dialing rules to delete. |

#### setDialingRulesResponse

Empty.

[]()