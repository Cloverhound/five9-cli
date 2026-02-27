## deleteLanguagePrompt

`Version 9.5`

Deletes a prompt in the specified language.

You cannot overwrite an existing prompt. If you try to modify a prompt in any language, an exception occurs.

Each prompt can exist in only one version in each language. If you try to modify an existing prompt, you get an exception. Therefore, use this method to delete the current language version. Afterward, use one of the modification methods to upload the new version.

#### deleteLanguagePrompt

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| promptName | string | Name of the prompt. |
| language | string | Language of the prompt. |

#### deleteLanguagePromptResponse

Empty.

[]()