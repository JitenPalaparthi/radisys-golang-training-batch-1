## Implementation details
- Create Customer
- Read Customer
- Update Customer
- Delete Customer
- Read All Customers

## Database

- To create docker instance of postgres

    ```docker run -d --name customersdbpg -p 5432:5432 -e POSTGRES_PASSWORD=admin123 -e POSTGRES_USER=admin -e POSTGRES_DB=customersdb postgres```

- To run application with logger , glog

    ```go run main.go -stderrthreshold=INFO -logtostderr=true```

- to run kafka

    ```docker-compose -f components/kafka-zookeeper.yaml up -d```

- To run a test

    ```go test -timeout 30s -run ^TestValidate$ customers/models```

- To execute all tests from a package 

    ```go test customers/models```

- To test everything

    ```go test -v ./...```

- To check coverage of tests

- First write the coverage using -coverprofile 
    ```go test -coverprofile=coverage.out ./...```
    
- can check in html output
    ```go tool cover -html=coverage.out```


protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    protos/customer.proto

    9618558500
    jitenp@outlook.com