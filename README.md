# Buss Travel Simple CRUD

## Get Started 
 
1. Put environments on ```.env``` using the ```.env.example``` as guide
2. Install Prisma ```go get github.com/prisma/prisma-client-go```
3. Generate ORM needed files ```go run github.com/prisma/prisma-client-go generate```
4. Run the containers ```docker-compose up```
5. Add ```127.0.0.1 db```  to ```/etc/hosts``` file (step required to up migrations on mysql container)
6. Up Migrations ```go run github.com/prisma/prisma-client-go migrate dev```

## Swagger

To generate swagger docs files, use

1. Install GO swagger docs generator ```go get github.com/swaggo/swag/cmd/swag```
2. Generate updated docs ```go run github.com/swaggo/swag/cmd/swag init -g main.go --output docs```
3. Open ```http://localhost:1324/swagger/index.html```

then use /swagger/ to access

## API Routes

| TYPE  | URL | DESCRIPTION | 
| - | - | - |
| POST | /bus-travel  | Creates BusTravel | 
| DELETE | /bus-travel/:id | Delete BusTravel by ID | 
| GET | /bus-travel/:id | Get BusTravel by ID| 
| GET | /bus-travel | Get all BusTravel | 
| PUT | /bus-travel/:id | Update BusTravel by ID | 


## End2End Test

Start the server locally before run the end to end test.

```go test ./tests/e2e_test.go```

The test perform the following action:

1. Create a BusTravel
2. Update this BusTravel
3. Delete this BusTravel
4. Get the BusTravel to verify if is Deleted

## POSTMAN

Public Collection BusTravel:

https://www.postman.com/altimetry-specialist-94040185/workspace/bustravel