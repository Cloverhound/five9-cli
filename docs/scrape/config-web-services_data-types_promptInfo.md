Contains information about a prompt.

                                                    | Name | Type | Description |
| --- | --- | --- |
| description | string | Description of the prompt. |
| languages Version 9.5 | string [0..unbounded] | List of languages for getPrompt() or one language for the methods that modify prompts This field is ignored for the methods that add prompts. If you try to create or modify a prompt in more than one language in the same request, an exception occurs. For several languages, use several requests. |
| name | string | Name of the prompt. |
| type | promptType | Type of prompt. |

[]()