# golang-training-Theater

## Description

##### Backend Application for storing theater

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

|Path|Method|Description|Body example|
|:---:|:---:|:---:|:---|
|```/account```|```POST```|create new account|```{"FirstName":"ExampleFirstName","LastName":"ExampleLastName","PhoneNumber":"ExamplePhoneNumber","Email":"Example@gmail.com"}```|
|```/genre```|```POST```|create new genre|```{"Name":"ExampleName"}```|
|```/hall```|```POST```|create new hall|```{"AccountId":"1","Name":"ExampleName","Capacity":"9999","LocationId":"1"}```|
|```/location```|```POST```|create new location|```{"AccountId":"1","Address":"ExampleAddress","PhoneNumber":"+777777777777"}```|
|```/performance```|```POST```|create new performance|```{"AccountId":"1","Name":"ExampleName","GenreId":"1","Duration":"1:00"}```|
|```/place```|```POST```|create new place|```{"SectorId":"10","Name":"1"}```|
|```/poster```|```POST```|create new poster|```{"AccountId":"1","ScheduleId":"6","Comment":"ExampleCommit"}```|
|```/price```|```POST```|create new price|```{"AccountId":"1","SectorId":"9","PerformanceId":"1","Price":"99"}```|
|```/role```|```POST```|create new role|```{"Name":"ExampleName"}```|
|```/schedule```|```POST```|create new schedule|```{"AccountId":"1","PerformanceId":"1","Date":"2021-04-25 13:00","HallId":"1"}```|
|```/sector```|```POST```|create new sector|```{"Name":"ExampleName"}```|
|```/ticket```|```POST```|create new ticket|```{"AccountId":"1","ScheduleId":"6","PlaceId":"1","DateOfIssue":"now()","paid":"true","reservation":"true","destroyed":"true"}```|
|```/user```|```POST```|create new user|```{"AccountId":"1","FirstName":"ExampleFirstName","LastName":"ExampleLastName","RoleId":"1","LocationId":"1","PhoneNumber":"+777777777777"}```|
|```/account```|```PUT```|update account|```{"Id":"18","FirstName":"ExampleFirstName","LastName":"ExampleLastName","PhoneNumber":"ExamplePhoneNumber","Email":"Example@gmail.com"}```|
|```/genre```|```PUT```|update genre|```{"Id":"12","Name":"ExampleName"}```|
|```/hall```|```PUT```|update hall|```{"Id":"6","AccountId":"1","Name":"ExampleName","Capacity":"9999","LocationId":"1"}```|
|```/location```|```PUT```|update location|```{"Id":"6","AccountId":"1","Address":"ExampleAddress","PhoneNumber":"+777777777777"}```|
|```/performance```|```PUT```|update performance|```{"Id":"6","AccountId":"1","Name":"ExampleName","GenreId":"1","Duration":"1:00"}```|
|```/place```|```PUT```|update place|```{"Id":"11","SectorId":"10","Name":"2"}```|
|```/poster```|```PUT```|update poster|```{"Id":"7","AccountId":"1","ScheduleId":"6","Comment":"ExampleCommit"}```|
|```/price```|```PUT```|update price|```{"Id":"36",AccountId":"1","SectorId":"9","PerformanceId":"1","Price":"99"}```|
|```/role```|```PUT```|update role|```{"Id":"8","Name":"ExampleName"}```|
|```/schedule```|```PUT```|update schedule|```{"Id":"11","AccountId":"1","PerformanceId":"1","Date":"2021-04-25 13:00","HallId":"1"}```|
|```/sector```|```PUT```|update sector|```{"Id":"19","Name":"ExampleName"}```|
|```/ticket```|```PUT```|update ticket|```{"Id":"26","AccountId":"1","ScheduleId":"6","PlaceId":"1","DateOfIssue":"now()","paid":"true","reservation":"true","destroyed":"true"}```|
|```/user```|```PUT```|update user|```{"Id":"8","AccountId":"1","FirstName":"ExampleFirstName","LastName":"ExampleLastName","RoleId":"1","LocationId":"1","PhoneNumber":"+777777777777"}```|

## Usage

1. Run server on port ```8080```

> ```go run ./theater-gorm/cmd/main.go```

2. Open URL ```http://localhost:8080```