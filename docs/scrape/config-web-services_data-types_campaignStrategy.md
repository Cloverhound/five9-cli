Version 4

Defines the dialing strategy configuration for campaigns. This includes start time, schedule, and other attributes that control when and how the strategy runs.

This field must be a numeric value, **not** a time-formatted string.

Type: `xs:long`

**Examples:**

00:00 → 0

01:00 → 60

13:30 → 810

## Important Notes

  -

**Start Time Format**:

Specify `startAfterTimeMins` as a numeric value representing minutes after midnight. Do not use a time-formatted string. Use a valid long integer. Passing a string in HH:MM format (e.g., "00:15") results in a `NumberFormatException` due to type mismatch.

Conversion Formula: `minutes` = (`hours` * 60) + `minutes`

  -

**Schedule Requirement**:

The `<schedule>` element is mandatory when modifying a strategy. Without it, changes will fail validation or be ignored.

  -

**Default Strategy Updates**:

When modifying the Default Strategy, ensure the `<description>` is copied. If omitted, schedule changes will not apply.

                                                    **Note: **The <schedule> must be present for the strategy to be valid.
                                                    | Name | Type | Required | Description |
| --- | --- | --- | --- |
| description | string | Optional* (except when modifying Default Strategy) | Optional description of the strategy. Important: When modifying the Default Strategy, this field must be copied (retained or respecified). If omitted, changes to <schedule> will not apply. |
| enabled | boolean | No | Indicates whether the strategy is active. At least one strategy, named Default, must be active. |
| filter | campaignStrategyFilter | No | Conditions that apply to a strategy. |
| name | string | Yes | Unique name of the strategy. However, you can modify the name at any time. |
| schedule | campaignStrategyPeriod [..unbounded] | Yes* | Defines the dialing pattern for the strategy. Required when modifying a strategy. Omitting this element will result in validation errors or ignored changes. |
| startAfterTimeMins | long (xs:long) | No | startAfterTimeMins is a numeric value (type: xs:long) representing the number of minutes after midnight (00:00). For example: 00:00 → 0 01:00 → 60 13:30 → 810 Note: Do not use time-formatted strings like "00:15" — this will result in unmarshalling errors (NumberFormatException). |

[]()