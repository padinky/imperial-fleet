# IMPERIAL FLEET

This is very simple version for fleet management. Written in Golang and MySQL for database

# Framework & Library Used
1. [Fiber](https://gofiber.io/): Golang http framework
2. [Gorm](https://gorm.io/): ORM library for Golang

# How To Run
Simply clone this repository, then choose either you want to run it using docker-compose or manual (unfortunatelly the docker-compose version is still buggy)

### 1.Using docker-compose
Simply just run this code from the project root directory
```sh
docker-compose up
```

### 2. Manual
1. Prepare your local PostgreSQL database engine
2. Duplicate file `.env-example` and rename it as `.env`, then adjust the configuration following your local setting
3. Run: ```go mod tidy``` to download all external libraries
4. Run: ```go run main```



# Postman Collection
Please find postman collection under folder : `./postman-collection`.
Then we can import those files to postman, adjust the endpoint base url, and play around hitting the endpoints.

