## modifyUser

Modifies the user’s attributes.

An exception is thrown if the user already exists, if the limit number of users is reached, or if user attributes are invalid.

#### modifyUser

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| userGeneralInfo | userGeneralInfo | User information to modify. |
| rolesToSet | userRoles | User roles to modify. |
| rolesToRemove | userRoleType [0..unbounded] | User roles to remove. |

#### modifyUserResponse

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| return | userInfo | Modified user information. |

[]()