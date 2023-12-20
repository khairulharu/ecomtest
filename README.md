# ecomtest

|HOW TO RUN|

clone repository and run the commands below to downloading other library
```go
go get
```
first change the .env file with database configuration
```
DATABASE_HOST=host
DATABASE_PORT=port
DATABASE_USER=user 
DATABASE_PASS=password
DATABASE_NAME=dbname

SERVER_HOST=localhost
SERVER_PORT=5656
```
the schema database internal/db/migration folder, or in this app, i set automaticly migrate models to database using golang migrator
```go
if err := db.AutoMigrate(&domain.Product{}); err != nil {
		log.Fatalf("migrate error:%s", err.Error())
	}
```  

this code in internal/component/connection.go
the product will auto migration using gorm migrator, if you want some using golang migration there a schema and use this code
i use golang migration cli, run this code in cli but you must have download the golang migrator first
```
migrate -path internal/db/migration -database "mysql://user:pass@tcp(host:port)/dbname" -verbose up
```

run the app with command noted you have 
```go
go run main.go
```


link drive of the dump and collection json

```
https://drive.google.com/drive/folders/1U7jxs3Ow-TUKOKrlj0SZivm1DmuBJxOK?usp=sharing
```