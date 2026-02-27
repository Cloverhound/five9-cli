```
<?php
```

```
 
```

```
echo "Begin<br/><br/>";
```

```
 
```

```
 
```

```
// Import the WSDL and authenticate the user.-----------------------------
```

```
$wsdl_five9 = "https://api.five9.com/wsadmin/<API version>/AdminWebService?wsdl&user=<Five9username>";
```

```
 
```

```
try
```

```
{
```

```
	$soap_options = array( 'login' => 'Five9username', 'password' => 'Five9password', 'trace' => true );
```

```
	$client_five9 = new SoapClient( $wsdl_five9 , $soap_options );
```

```
}
```

```
catch (Exception $e)
```

```
{
```

```
	$error_message = $e->getMessage();
```

```
        echo $error_message;
```

```
}
```

```
 
```

```
//---------------initiate import (asyncAddRecordsToList)-----------------
```

```
$listUpdateSettings = array ( "fieldsMapping" => array (
```

```
                      array ( "columnNumber" => '1', "fieldName" => "number1", "key" => true ),
```

```
                      array ( "columnNumber" => '2', "fieldName" => "first_name", "key" => false ),
```

```
                      array ( "columnNumber" => '3', "fieldName" => "last_name", "key" => false) ),
```

```
			 "reportEmail" => "email@email.com",
```

```
				"separator" => ',',
```

```
				"skipHeaderLine" => false,
```

```
				"callNowMode" => "ANY",   //optional
```

```
                              "callNowColumnNumber" => 4,   //optional
```

```
				"cleanListBeforeUpdate" => false,
```

```
				"crmAddMode" => "ADD_NEW",
```

```
				"crmUpdateMode" => "UPDATE_SOLE_MATCHES",
```

```
				"listAddMode" => "ADD_IF_SOLE_CRM_MATCH" );
```

```
$data = array ( array ( "5555776754" , "Don" , "Draper", "YES" ),
```

```
                array ( "5551112244" , "Betty" , "Smith", "NO" ));
```

```
$xml_data = array ('listName' => "asdf", 'listUpdateSettings' => $listUpdateSettings, 'importData' => $data);   //request parameters
```

```
 
```

```
 
```

```
 
```

```
$result = $client_five9->asyncAddRecordsToList($xml_data);
```

```
$variables = get_object_vars($result);
```

```
$resp = get_object_vars($variables['return']);
```

```
$identifier = $resp['identifier'];  //the ID for the import
```

```
//echo $identifier;
```

```
 
```

```
//-------check progress of import (isImportRunning)----------------------
```

```
$import_running = true;
```

```
$IIR_p = array('identifier'=>array('identifier'=>$identifier), 'waitTime'=>10);
```

```
 
```

```
while($import_running)
```

```
{
```

```
    try
```

```
    {
```

```
    $IIR_result = $client_five9->isImportRunning($IIR_p);
```

```
prevent multiple calls within a second
```

```
     $variables = get_object_vars($IIR_result);
```

```
  // to prevent multiple calls within a second
```

```
  sleep (1);
```

```
  $import_running = $variables['return'];
```

```
    }
```

```
    catch (Exception $e)
```

```
    {
```

```
            $error_message = $e->getMessage();
```

```
            echo $error_message;
```

```
    }
```

```
}
```

```
 
```

```
//------get result (getListImportResult)---------------------------------
```

```
try
```

```
{
```

```
     $GLIR_p = array('identifier'=>array('identifier'=>$identifier));
```

```
     $GLIR_result = $client_five9->getListImportResult($GLIR_p);
```

```
     print_r($GLIR_result);
```

```
}
```

```
catch (Exception $e)
```

```
{
```

```
        $error_message = $e->getMessage();
```

```
        echo $error_message;
```

```
}
```

```
echo "<br/><br/>";
```

```
echo "END";
```

```
 
```

```
?>
```

[]()