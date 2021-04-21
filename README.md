# golang-training-Theater

## Description

##### Full Stack Application for storing theater

|Path|Method|Description|
|:---:|:---:|:---:|
|```/tickets```|```GET```|get all tickets|
|```/posters```|```GET```|get all posters|
|```/users?idAccount={id}```|```GET```|get all users by ```account```|
|```/account?id={id}```|```GET```|get account by ```id```|
|```/genre?id={id}```|```GET```|get genre by ```id```|
|```/hall?id={id}```|```GET```|get hall by ```id```|
|```/location?id={id}```|```GET```|get location by ```id```|
|```/performance?id={id}```|```GET```|get performance by ```id```|
|```/place?id={id}```|```GET```|get place by ```id```|
|```/poster?id={id}```|```GET```|get poster by ```id```|
|```/price?id={id}```|```GET```|get price by ```id```|
|```/role?id={id}```|```GET```|get role by ```id```|
|```/schedule?id={id}```|```GET```|get schedule by ```id```|
|```/sector?id={id}```|```GET```|get sector by ```id```|
|```/ticket?id={id}```|```GET```|get ticket by ```id```|
|```/user?id={id}```|```GET```|get user by ```id```|
|```/account?id={id}```|```DELETE```|delete account by ```id```|
|```/genre?id={id}```|```DELETE```|delete genre by ```id```|
|```/hall?id={id}```|```DELETE```|delete hall by ```id```|
|```/location?id={id}```|```DELETE```|delete location by ```id```|
|```/performance?id={id}```|```DELETE```|delete performance by ```id```|
|```/place?id={id}```|```DELETE```|delete place by ```id```|
|```/poster?id={id}```|```DELETE```|delete poster by ```id```|
|```/price?id={id}```|```DELETE```|delete price by ```id```|
|```/role?id={id}```|```DELETE```|delete role by ```id```|
|```/schedule?id={id}```|```DELETE```|delete schedule by ```id```|
|```/sector?id={id}```|```DELETE```|delete sector by ```id```|
|```/ticket?id={id}```|```DELETE```|delete ticket by ```id```|
|```/user?id={id}```|```DELETE```|delete user by ```id```|
|```/account```|```POST```|create new account|
|```/genre```|```POST```|create new genre|
|```/hall```|```POST```|create new hall|
|```/location```|```POST```|create new location|
|```/performance```|```POST```|create new performance|
|```/place```|```POST```|create new place|
|```/poster```|```POST```|create new poster|
|```/price```|```POST```|create new price|
|```/role```|```POST```|create new role|
|```/schedule```|```POST```|create new schedule|
|```/sector```|```POST```|create new sector|
|```/ticket```|```POST```|create new ticket|
|```/user```|```POST```|create new user|
|```/account```|```PUT```|update account|
|```/genre```|```PUT```|update genre|
|```/hall```|```PUT```|update hall|
|```/location```|```PUT```|update location|
|```/performance```|```PUT```|update performance|
|```/place```|```PUT```|update place|
|```/poster```|```PUT```|update poster|
|```/price```|```PUT```|update price|
|```/role```|```PUT```|update role|
|```/schedule```|```PUT```|update schedule|
|```/sector```|```PUT```|update sector|
|```/ticket```|```PUT```|update ticket|
|```/user```|```PUT```|update user|

## Usage

1. Run server on port ```8080```

> ```go run ./theater-gorm/cmd/main.go```

2. Open URL ```http://localhost:8080```