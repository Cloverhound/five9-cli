## getUserProfiles

Returns a list of user profiles whose names match a string pattern.

#### getUserProfiles

| Parameter | Type | Description |
| --- | --- | --- |
| userProfileNamePatern | string | Name of the profile or regular expression. For all user profiles, omit the parameter or use this pattern: .* |

#### getUserProfilesResponse

| Parameter | Type | Description |
| --- | --- | --- |
| return | userProfile[0..unbounded] | List of user profiles. |

[]()