Data type of autodial campaigns that enables the dialer to dial only when agents are available to take calls. If you do not use this object, numbers are dialed continuously whether or not agents are available. See `[baseOutboundcampaign](/bundle/api-config-web/page/config-web-services/data-types/baseOutboundcampaign/_ch-baseOutboundcampaign.htm).`

                                                    | Name | Type | Description |
| --- | --- | --- |
| ReadyToReceiveCalls | string | One or more agents are available, ready, and not on call. |
| ReadyToReceiveCallsOrBusy | string | One or more agents are logged in and are either ready or busy taking or finishing another call. Agents are not considered busy if they are making a manual call after having been on break. |
| LoggedIn | string | One or more agents are logged in, regardless of their current status. |

[]()