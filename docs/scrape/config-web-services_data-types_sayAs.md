Contains the possible text elements of a prompt should be read. This information is located in the TTS Builder.

                                                    | Name | Type | Description |
| --- | --- | --- |
| Default | string |  |
| Words | string | Word strings are spoken as words instead of being pronounced as strings of individual letters and digits. However, the characters of a word may be spoken individually when character sequences are particularly difficult to pronounce. Meant for acronyms to be read as words. |
| Acronym | string | Letters and/or digits, such as NATO and UNESCO in English, that are spoken as words when considered natural in the target language. Otherwise, letters and digits are pronounced individually, for example: API in English. Use detail strict to force spelling mode. In that case, punctuation is also spoken, for example: speaking a comma as comma. Acronym with detail strict is equivalent to letters with detail strict. don’t understand the last 2 sentences about strict something |
| Address | string | Postal addresses. |
| Cardinal | string | Supported if relevant in the target language. Roman cardinals are often supported. |
| Currency | string | Text is a currency amount with or without the currency symbol. Supports currencies commonly specified in the country corresponding to the target language. |
| Date | string |
| Decimal | string | Same as number with format decimal. The separator for the integral part is optional. For example, in U.S. English, 123456.123 and 123,456.123 are pronounced in the same way. |
| Digits | string | Same as number with format digits. Numbers must be read digit by digit, including decimal periods and commas. |
| Duration | string | For example, duration with format hms is read as <h> hour(s), <m> minute(s), and <s> seconds. |
| Fraction | string | Same as number with format fraction. For example, pronounce 1/3 as one third. |
| Letters | string | Strings are pronounced as a sequence of single letters and/or digits. With detail strict punctuation is also spoken, for example: speaking a comma as comma. Letters with detail strict is equivalent to acronym with detail strict. For the true spelling of all readable characters, use the interpret-as value spell. |
| Measure | string | Many units, such as km, hr, dB, lb, or MHz, are supported. Units may appear immediately next to a number, such as 1cm, or be separated by a space, such as 15 ms. For some units, the distinction between singular and plural may not always be made correctly. |
| Name | string | Interpret a string as a proper name if possible. |
| Net | string | Email can be used for email addresses. |
| Telephone | string | Supports phone numbers specified in the country corresponding to the target language. See the language-specific User Guide for a list of the supported formats. Use detail punctuation to speak punctuation, such as speaking a dash as dash. |
| Ordinal | string | Same as number with format ordinal. See the language-specific User’s Guide for a list of the supported formats. |
| Spell | string | Characters in text string are pronounced as individual characters. |
| Time | string | Hour must be less than 24; minutes and seconds must be less than 60; AM/PM is read only if explicitly specified. See the language-specific User’s Guide for a list of the supported formats. |

[]()