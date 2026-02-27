Contains the configuration for a dialing rule. Dialing rules are used to ensure that numbers are not dialed during certain times or days.

                                                    | Name | Type | Description |
| --- | --- | --- |
| applyToManualCalls | boolean | Whether to restrict manual calls to the dialing rule. True: Restrict manual calls. False: Do not restrict manual calls. |
| contactText | string | Comma-separated list of the possible entries for a state in the calling list. For example, "Alabama, AL" indicates that the state dialing rule applies to contact records that have either Alabama or AL in the state contact field. If omitted, the name and abbreviation of the state are used by default. |
| dateRange | dateRange | Date range when dialers do not dial numbers. |
| fixedTimeZone | string | Time zone used by the dialer to check whether a number can be called. If omitted, the time zone of the dialed number is used by default. For example, US/Pacific is used for PDT time. The names of the time zones are located in the Dialing Rules tab of Administrator Configuration. |
| name | string | Name of the dialing rule. |
| state | stateProvince | State for which to apply this rule. If omitted, the rule applies to numbers of any state. If specified, the rule applies to the value of the state contact field. |
| timeRange | timeRange | Part of the day that applies to the rule. If omitted, the assumption is all day long. |

[]()