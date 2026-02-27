Contains the configuration details of a web connector.

                                                    | Name | Type | Description |
| --- | --- | --- |
| addWorksheet | boolean | Applies only to POST requests. Whether to pass worksheet answers as parameters. True: Pass worksheet answers as parameters. False: Do not pass worksheet answers as parameters. |
| agentApplication | webConnectorAgentAppType | If executeInBrowser=true, this parameter specifies whether to open the URL in an external or an embedded browser. |
| clearTriggerDispositions | boolean | When modifying an existing connector, whether to clear the existing triggers. True: Clear existing triggers. False: Do not clear existing triggers. |
| constants | keyValuePair[0..unbounded] | List of parameters passed with constant values. |
| ctiWebServices | webConnectorCTIWebServicesType | In the Internet Explorer toolbar, whether to open the HTTP request in the current or a new browser window. |
| description | string | Purpose of the connector. |
| executeInBrowser | boolean | When enabling the agent to view or enter data, whether to open the URL in an embedded or external browser window. True: External — open a browser window. False: Embedded — do not open a browser window, which is the desired action when using the On Call disposition triggers. |
| name | string | Name of the connector. |
| postConstants | keyValuePair[0..unbounded] | When using the POST method, constant parameters to pass in the URL. |
| postMethod | boolean | Whether the HTTP request type is POST or GET. True: HTTP POST. False: HTTP GET. |
| postVariables | keyValuePair[0..unbounded] | When using the POST method, variable parameters to pass in the URL. |
| startPageText | string | When using the POST method, enables the administrator to enter text to be displayed in the browser (or agent Browser tab) while waiting for the completion of the connector. |
| trigger | webConnectorTriggerType | Available trigger during a call when the request is sent. |
| triggerDispositions | string [0..unbounded] | When the trigger is OnCallDispositioned, specifies the trigger dispositions. |
| url | string | URL of the external Web site. |
| variables | keyValuePair[0..unbounded] | When using the POST method, connectors can include worksheet data as parameter values. The variable placeholder values are surrounded by @ signs. For example, the parameter ANI has the value @Call.ANI@ |

[]()