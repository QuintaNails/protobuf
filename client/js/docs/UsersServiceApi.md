# Botsbotsproto.UsersServiceApi

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**usersServiceGetByTelegramID**](UsersServiceApi.md#usersServiceGetByTelegramID) | **POST** /users.UsersService/GetByTelegramID | 

<a name="usersServiceGetByTelegramID"></a>
# **usersServiceGetByTelegramID**
> UsersGetByTelegramIDResponse usersServiceGetByTelegramID(body)



### Example
```javascript
import {Botsbotsproto} from 'botsbotsproto';

let apiInstance = new Botsbotsproto.UsersServiceApi();
let body = new Botsbotsproto.UsersGetByTelegramIDRequest(); // UsersGetByTelegramIDRequest | 

apiInstance.usersServiceGetByTelegramID(body, (error, data, response) => {
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
 **body** | [**UsersGetByTelegramIDRequest**](UsersGetByTelegramIDRequest.md)|  | 

### Return type

[**UsersGetByTelegramIDResponse**](UsersGetByTelegramIDResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

