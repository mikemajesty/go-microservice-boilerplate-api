db.createUser(
  {
      user: "admin",
      pwd: "admin",
      roles: [
          {
              role: "readWrite",
              db: "go-microservice-boilerplate-api"
          }
      ]
  }
);