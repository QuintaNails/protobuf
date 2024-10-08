# Botsbotsproto.CompaniesServiceApi

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**companiesServiceReserve**](CompaniesServiceApi.md#companiesServiceReserve) | **POST** /companies.CompaniesService/Reserve | 

<a name="companiesServiceReserve"></a>
# **companiesServiceReserve**
> CompaniesReserveResponse companiesServiceReserve(body)



### Example
```javascript
import {Botsbotsproto} from 'botsbotsproto';

let apiInstance = new Botsbotsproto.CompaniesServiceApi();
let body = new Botsbotsproto.CompaniesReserveRequest(); // CompaniesReserveRequest | 

apiInstance.companiesServiceReserve(body, (error, data, response) => {
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
 **body** | [**CompaniesReserveRequest**](CompaniesReserveRequest.md)|  | 

### Return type

[**CompaniesReserveResponse**](CompaniesReserveResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

