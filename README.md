# Server_Protection_System

This package is built for rapid Backend development.
It provides you with a set of features that enable rapid development.
This package does not make any changes to the main packages used in this project and only makes some tasks easier. For development, the main package is definitely required.
So in addition to this package, you should also install the main packages used in this package.

Example:

In this package, to connect to the mysql database, just call the Run_Server_Sql function name in Go, which easily performs the start operation for you, the point is that you only connect and to work with the database, you need the main mysql package that has been released for Go, which is placed at the end of this text, which you can easily download:

The names of the packages used along with the licenses for each are as follows:

1-mysql    License  ==  [License](https://github.com/go-sql-driver/mysql?tab=MPL-2.0-1-ov-file "License mysql")
Mozilla Public License Version 2.0

---------------------------------------

2-gin      License  ==  [License](https://github.com/gin-gonic/gin?tab=MIT-1-ov-file "License gin") 
Copyright (c) 2014 Manuel Martínez-Almeida

---------------------------------------

3-uuid      License  ==  [License](https://github.com/google/uuid?tab=License-1-ov-file "License uuid") 
Copyright (c) 2009,2014 Google Inc. All rights reserved.

---------------------------------------

4-Redis client for Go      License  ==  [License](https://github.com/redis/go-redis?tab=BSD-2-Clause-1-ov-file "License Redis client for Go")
Copyright (c) 2013 The github.com/redis/go-redis Authors.
All rights reserved.

---------------------------------------

5-Gomail      License  ==  [License](https://github.com/go-gomail/gomail?tab=MIT-1-ov-file "Gomail")
Copyright (c) 2014 Alexandre Cesaro

---------------------------------------

For more advanced development with more features, using this package is not recommended because it is more suitable for creating quick APIs for review or testing.
Of course, the main goal of this package is to help develop faster.

---------------------------------------

Well, after fully understanding this package, now it's time to learn how to use it:


Package features:

  * [Installation](#Installation)  
  * [Start_sql](#Start_sql)
  * [Send_Email_otp](#Send_Email_otp)
  * [GenerateOTP](#GenerateOTP)
  * [Redis_Server](#Redis_Server)
  * [password_env](#password_env)
  * [Otp_Redis](#Otp_Redis)
  * [Movie](#Movie)
  * [image](#image)
 

  
---------------------------------------
## Installation

To install the package, simply enter the following command:
```bash
go get -u github.com/Mreza2020/Server_Protection_System
```

This system helps you write shorter code.

> [!IMPORTANT]
>Before starting the usage tutorial, you should know that you must place all the required passwords in env, which are:
>   * DB_PASSWORD
>   * DB_Redis
>   * DB_Redis_password
>   * Email
>   * Email_pass
>   * Email_smtp_server
>   * Email_smtp_server_port



## Start_sql

Whenever you want to connect to your MySQL server, just call the package name `Run_Server_Sql`.
How to use it is as follows:

```go
database, err := Run_Server_Sql(ch)
```

## Send_Email_otp

To send a one-time password email, simply use the `Send_Email_otp` package, which is used as follows:

```go
Send_Email_otp(ch chan string, name_file_Email string, otp_cod string, SetHeader_server string, SetHeader_Client string)
```
Note:
It only sends one parameter to the `HTML` or `JavaScript` file, which is `otp_cod`.

To use it in your file, you must use the phrase Code in the part where the one-time `code` is placed.

This code only sends an email and you have to create the one-time code yourself, which of course we have a function in this `package` for this purpose called `GenerateOTP`, which is explained below.

## GenerateOTP

To generate random numbers, simply add the `Generate OTP` `package` to your project.

Structure:

```go
GenerateOTP(length int, Number_length int) string
```

## Redis_Server

To work with Redis, just write `Run_Redis_Server` to easily establish your connection.

Structure:

```go
Run_Redis_Server(ch chan string)
```

## Otp_Redis

For your rapid development, this package includes ready-made code for storing your temporary code in Redis so that you can quickly save and retrieve the code by typing the corresponding function. You only need to run the Redis start function before starting.

Save code structure:

```go
SaveOtpToRedis(email string, otp string, Storage int)
```

Receive code structure:

```go
GetOtpFromRedis(email string, ch chan string)
```

## Movie

If you need to receive a movie from the user in your application, don't worry at all. It is easily provided for you in the `Movie_get_Api` package. The method of use is as follows:

Detailed description:

This function is for capturing video from the user.
Parameters:

- Amount int =  File size value, for example 1 MB
- Unit int = The unit of value, for example 20, has three states:
1- 10 KB
2- 20 MB
3- 30 GB

- NameFile string = The name under which the user uploads their file.
- Folder_name string = The name of the folder where the video is saved, for example, Uploads.
- Directory_name string = Directory path inside that folder, for example video
- Drive string = The drive where the file is saved or the path itself, for example: D:/MyApp

var FileName string
var ContentType string
var Path string
var Emali string
var User string
var User_n string

These are the return parameters of this function.

To use it, you need to write an API with the main Gin package, which is described above for downloading.

For example:

```go
Api.POST("/Movie", Movie_get_Api)
```
Note :

To use it, you need to download FFmpeg on your system.
This FFmpeg is used to increase security.

## image

If you need to get a photo from a user in your application, don't worry at all. It is easily provided for you in the `Image_Get_Api` package. The usage is as follows:

Detailed description:
This function is for capturing Image from the user.

Parameters:

Amount_i int =  File size value, for example 1 MB

-Unit_i int = The unit of value, for example 20, has three states:
1- 10 KB
2- 20 MB
3- 30 GB

NameFile_i string = The name under which the user uploads their file.

Folder_name_i string = The name of the folder where the Image is saved, for example, Uploads.

Directory_name_i string = Directory path inside that folder, for example Image

Drive_i string = The drive where the file is saved or the path itself, for example: D:/MyApp

var FileName_i string
var ContentType_i string
var Path_i string
These are the return parameters of this function.

To use it, you need to write an API with the main Gin package, which is described above for downloading.

For example:

```go
Api.POST("/Image", Image_Get_Api)
```
