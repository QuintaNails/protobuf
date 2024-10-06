# Botsbotsproto.TelegramBotsServiceApi

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**telegramBotsServiceAddBot**](TelegramBotsServiceApi.md#telegramBotsServiceAddBot) | **POST** /bots.TelegramBotsService/AddBot | 
[**telegramBotsServiceSetupBot**](TelegramBotsServiceApi.md#telegramBotsServiceSetupBot) | **POST** /bots.TelegramBotsService/SetupBot | 

<a name="telegramBotsServiceAddBot"></a>
# **telegramBotsServiceAddBot**
> BotsAddBotResponse telegramBotsServiceAddBot(body)



### Example
```javascript
import {Botsbotsproto} from 'botsbotsproto';

let apiInstance = new Botsbotsproto.TelegramBotsServiceApi();
let body = new Botsbotsproto.BotsAddBotRequest(); // BotsAddBotRequest | 

apiInstance.telegramBotsServiceAddBot(body, (error, data, response) => {
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
 **body** | [**BotsAddBotRequest**](BotsAddBotRequest.md)|  | 

### Return type

[**BotsAddBotResponse**](BotsAddBotResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="telegramBotsServiceSetupBot"></a>
# **telegramBotsServiceSetupBot**
> BotsSetupBotResponse telegramBotsServiceSetupBot(body)



### Example
```javascript
import {Botsbotsproto} from 'botsbotsproto';

let apiInstance = new Botsbotsproto.TelegramBotsServiceApi();
let body = new Botsbotsproto.BotsSetupBotRequest(); // BotsSetupBotRequest | 

apiInstance.telegramBotsServiceSetupBot(body, (error, data, response) => {
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
 **body** | [**BotsSetupBotRequest**](BotsSetupBotRequest.md)|  | 

### Return type

[**BotsSetupBotResponse**](BotsSetupBotResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

