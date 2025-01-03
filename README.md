
# Golang boilerplate Microservice API

##### In this Golang boilerplate I used the main concepts of DDD, Clean Architecture, Ports and Adapters (Hexagonal Architecture).

- infra
  ```
   docker-compose up --build
  ```

- run

  ```
  CompileDaemon -command="./go-microservice-boilerplate-api"
  ```

### Test

- run
  ```
  $ go test ./core/dog/use-case/__tests__/ ./core/cat/use-case/__test__/
  ```

#### CRUD features

- List
  - mongo
    - search
    - pagination
    - sort
    - entity validation
  - postgres
    - search
    - pagination
    - sort
    - entity validation
- Delete
  - mongo
    - Logical deletion
    - entity validation
  - postgres
    - Logical deletion
    - entity validation
- Update
  - mongo
    - Update Partial entity
    - entity validation
  - postgres
    - Update Partial entity
    - entity validation
- Create
  - mongo
    - entity validation
    - Not allow creating duplicates
  - postgres
    - entity validation

### Architecture diagram

## ![alt text](OnionGraph.jpg)

[Architecture documentation](https://jeffreypalermo.com/2008/07/the-onion-architecture-part-1/)

### User diagram

![alt text](diagram.png)

### Microservice architecture.


##### App Skeleton

```
.
├── core
│   ├── cat
│   │   ├── entity
│   │   │   └── cat.go
│   │   ├── repository
│   │   │   └── cat.go
│   │   └── use-case
│   │       ├── cat-create.go
│   │       ├── cat-delete.go
│   │       ├── cat-get-by-id.go
│   │       ├── cat-list.go
│   │       ├── cat-update.go
│   │       └── __test__
│   │           ├── cat-create_test.go
│   │           ├── cat-delete_test.go
│   │           ├── cat-get-by-id_test.go
│   │           ├── cat-list_test.go
│   │           └── cat-update_test.go
│   └── dog
│       ├── entity
│       │   └── dog.go
│       ├── repository
│       │   └── dog.go
│       └── use-case
│           ├── dog-create.go
│           ├── dog-delete.go
│           ├── dog-get-by-id.go
│           ├── dog-list.go
│           ├── dog-update.go
│           └── __tests__
│               ├── dog-create_test.go
│               ├── dog-delete_test.go
│               ├── dog-get-by-id_test.go
│               ├── dog-list_test.go
│               └── dog-update_test.go
├── docker-compose.yml
├── go.mod
├── go.sum
├── infra
│   ├── cache
│   │   ├── adapter.go
│   │   ├── memory
│   │   │   └── service.go
│   │   └── redis
│   │       └── service.go
│   ├── database
│   │   ├── adapter.go
│   │   ├── mongo
│   │   │   └── service.go
│   │   └── postgres
│   │       ├── migrations
│   │       │   └── create-cat-table.go
│   │       └── service.go
│   ├── logger
│   │   ├── adapter.go
│   │   └── service.go
│   ├── repository
│   │   ├── adapter.go
│   │   ├── mongo
│   │   │   └── repository.go
│   │   └── postgres
│   │       └── repository.go
│   └── secret
│       ├── adapter.go
│       └── service.go
├── main.go
├── modules
│   ├── cat
│   │   ├── adapter.go
│   │   ├── controller.go
│   │   ├── repository.go
│   │   ├── routes.go
│   │   └── validator.go
│   └── dog
│       ├── adapter.go
│       ├── controller.go
│       ├── repository.go
│       ├── routes.go
│       └── validator.go
├── observables
│   └── log-middlaware.go
├── README.md
├── textinho
└── utils
    ├── context.go
    ├── entity.go
    ├── exception.go
    ├── pagination.go
    ├── route.go
    ├── search.go
    ├── sort.go
    ├── type.go
    └── validator.go
```

---

The following is a list of all the people that have contributed Nestjs monorepo boilerplate. Thanks for your contributions!

[<img alt="mikemajesty" src="https://avatars1.githubusercontent.com/u/11630212?s=460&v=4&s=117" width="117">](https://github.com/mikemajesty)

## License

It is available under the MIT license.
[License](https://opensource.org/licenses/mit-license.php)