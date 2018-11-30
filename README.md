# eve-navi

## Architectures

```
Web Client <--+-- [GraphQL] ---> Gateway <--+-- [gRPC] ---> Shops <---> MySQL
              |                             |
              +-- [HTTP/2] ---> Server      +-- [gRPC] ---> Sourvenirs <---> MySQL
                                            |
                                            +-- [gRPC] ---> Foods <---> MySQL
```
