## updateConfigurationTranslations

`Version 9.5`

Updates the configuration translations for the locale.

#### updateConfigurationTranslations

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| localeName | string | Name of the locale. |
| diffTranslations | adminConfigTranslation[0..unbounded] | List of configuration translations for the locale. |

#### updateConfigurationTranslationsResponse

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| translationsWithNonExistingKey | adminConfigTranslation[0..unbounded] | Updated list of configuration translations for the locale. |

[]()