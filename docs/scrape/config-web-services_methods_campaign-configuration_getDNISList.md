## getDNISList

Returns the list of DNIS for the domain.

#### []()getDNISList

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| selectUnassigned | boolean | Whether to return the list of DNIS unassigned to a campaign: True: the list is returned. False: all DNIS provisioned for the domain are returned. |

#### getDNISListResponse

                                                    | Parameter | Type | Description |
| --- | --- | --- |
| return | string [0..unbounded] | List of unassigned DNIS associated with the domain. |

[]()