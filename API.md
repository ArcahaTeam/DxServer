# API
接口设计中，仅供参考。

## 用户类

### /user/register
Method: `POST`  
Description: `Create a new account.`
```json
{
  "UserID": 123456,
  "Password": 123456,
  "Code": ""
}
```

### /user/login
Method: `POST`  
Description: `Log into DxServer` 
```json
{
  "UserID": 123456,
  "Password": 123456
}
```

## 分数类

### /score/song
Method: `POST`  
Description: `Upload a score with play result` 
```
Not Available
```