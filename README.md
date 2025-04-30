# codegen-service

This service provides an engine of code generation and related api in SmartAPIForge backend system

Core functionality

### Get prepared:

##### Fast init

1) Fill .env file in project root (check .env.xmpl as reference)
2) Change dsn args in Taskfile.yaml if needed
3) Up dependencies + run application via ```task init```

##### or

1) Fill .env file in project root (check .env.xmpl as reference)
2) Change dsn args in Taskfile.yaml if needed
3) Raise redis storage via ```task redis_raise```
4) Run auth-service application ```task run``` (automatically calls ```task build``` before as dependency)
