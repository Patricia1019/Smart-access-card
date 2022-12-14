# Backend Interface

## GET getByID (Done)
### Description
通过CardID获取权限。输入cardID，输出包含cardID(string), username(string)和permission(bool)字段的Json。

### Example
GET http://8.134.190.122/api/getByID?cardID=100
如果没有权限或查表失败返回
```
{
    "cardID": "100",
    "permission": "no",
    "username": "unknown"
}
```
否则返回
```
{
    "cardID": "100",
    "permission": "yes",
    "username": "user"
}
```

## POST addUserID (Done)
### Description
提交用户权限。输入cardID, userame，输出包含cardID(string), username(string)和permission(bool)字段的Json。注意返回的permission表示操作后是否有权限，不表示操作本身是否成功。

### Example
POST http://8.134.190.122/api/addUserID?cardID=100&username=user
操作后（不论是否成功）没有权限则返回
```
{
    "cardID": "100",
    "permission": "no",
    "username": "unknown"
}
```
成功添加后有权限或重复添加则返回
```
{
    "cardID": "100",
    "permission": "yes",
    "username": "user"
}
```

## POST deleteByID (Done)
### Description
删除用户权限。输入cardID，输出包含cardID(string), username(string)和permission(bool)字段的Json。注意返回的permission表示操作后是否有权限，不表示操作本身是否成功。

### Example
POST http://8.134.190.122/api/deleteByID?cardID=100
操作后（不论是否成功）没有权限则返回
```
{
    "cardID": "100",
    "permission": "no",
    "username": "unknown"
}
```
否则返回
```
{
    "cardID": "100",
    "permission": "yes",
    "username": "user"
}
```
## getByUser (Done)
[GIN-debug] GET    /api/getByUser            --> iotproject/service.GetByUser (4 handlers)

## deleteByUser (Done)
[GIN-debug] POST   /api/deleteByUser         --> iotproject/service.DeleteByUser (4 handlers)

## Login (Done)
[GIN-debug] GET    /login                    --> iotproject/service.Login (3 handlers)
注意：添加Login接口后，除了getByID的api接口都需要cookie才能使用
