Contains the dates for which a dialing rule applies. You must specify either the days of the week or a date range. If both are specified, `daysOfWeek` is used, and the date range is ignored.

                                                    | Name | Type | Description |
| --- | --- | --- |
| daysOfWeek | dayOfWeek [0..unbounded] | Array of the days of the week. |
| endDate | dateTime | Last day of the date range. |
| startDate | dateTime | First day of the date range. |

[]()