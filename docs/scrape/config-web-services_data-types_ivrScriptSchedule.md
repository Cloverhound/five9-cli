Contains the configuration of an IVR script.

**Note: **When using this in the [createInboundCampaign](/bundle/api-config-web/page/config-web-services/data-types/../methods/campaign-configuration/createInboundCampaign.htm), only the default IVR schedule can be set or modified via the API. Additionally, the name parameter of the IVR script schedule object must be excluded.

                                                    | Name | Type | Description |
| --- | --- | --- |
| name | string | Name of the IVR schedule. |
| scriptName | string | Name of the IVR script. |
| scriptParameters | scriptParameterValue [0..unbounded] | Parameters of a foreign script module used in the IVR script. |

[]()