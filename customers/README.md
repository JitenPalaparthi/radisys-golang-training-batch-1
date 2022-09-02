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