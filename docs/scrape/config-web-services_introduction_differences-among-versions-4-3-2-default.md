The differences between API versions 4, 3, 2, and the default are listed in this table.

                                                    |  | Version 4 | Version 3 | Version 2 | Default |
| --- | --- | --- | --- | --- |
| Data Types |
| adminPermissionType | Added | Added |  |  |
| apiOperationType (VIVRCreateSession Version 3) |  | Added |  |  |
| basicImportResult (importTroubles , keyFieldsVersion 3 ) |  | Added |  |  |
| basicImportSettings (countryCode, failOnFieldParseError ) | Added | Added |  |  |
| campaignStrategies | New |  |  |  |
| campaignStrategy | New |  |  |  |
| campaignStrategyFilter | New |  |  |  |
| campaignStrategyPeriod | New |  |  |  |
| contactFieldRestriction (isEnabled) |  | Deleted | Added |  |
| importTrouble |  | New |  |  |
| importTroubleKind |  | New |  |  |
| listDialingMode(chatEnabled, visualModeEnabled) |  | Added |  |  |
| listDialingMode(EXTENDED_STRATEGY) | Added |  |  |  |
| listUpdateSimpleSettings (countryCode) | Added | New |  |  |
| passwordPolicies (entryValues) |  | Modified | Added | Modified |
| passwordPolicyEntries |  | Removed | Added | Removed |
| passwordPolicyEntryValue |  | Removed | Added | Removed |
| supervisorPermissionType (CanViewTextDetailsTab Version 3, CanAccessShowFields Version 3 |  | Added |  |  |
| Methods |
| addRecordToListSimple |  | New |  |  |
| createContactField |  | Modified | Modified | Modified |
| getCampaignStrategies | New |  |  |  |
| getContactFields |  | Modified | Modified | Modified |
| modifyContactField |  | Modified | Modified | Modified |
| setCampaignStrategies | New |  |  |  |
| setDefaultIVRSchedule (isVisualModeEnabled Version 3 Version 3, isChatEnabled ) |  | Added |  |  |
| Exceptions |
| AddRecordToListFault |  | New |  |  |
| CommonCampaignFault Version 3 |  | New |  |  |
| WrongListDialingModeFault | New |  |  |  |

[]()