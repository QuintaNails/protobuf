# Botsbotsproto.CompaniesServiceApi

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**companiesServiceReserveFromBot**](CompaniesServiceApi.md#companiesServiceReserveFromBot) | **POST** /companies.CompaniesService/ReserveFromBot | 

<a name="companiesServiceReserveFromBot"></a>
# **companiesServiceReserveFromBot**
> CompaniesReserveFromBotResponse companiesServiceReserveFromBot(body)



### Example
```javascript
import {Botsbotsproto} from 'botsbotsproto';

let apiInstance = new Botsbotsproto.CompaniesServiceApi();
let body = new Botsbotsproto.CompaniesReserveFromBotRequest(); // CompaniesReserveFromBotRequest | 

apiInstance.companiesServiceReserveFromBot(body, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CompaniesReserveFromBotRequest**](CompaniesReserveFromBotRequest.md)|  | 

### Return type

[**CompaniesReserveFromBotResponse**](CompaniesReserveFromBotResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

