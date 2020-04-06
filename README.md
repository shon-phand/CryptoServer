# CryptoServer

This microservice hosts the functionality to provide realtime currency information. 
Endpoints exposed to users are as below

- /currency/:symbol           : to fetch details for specific currency
- /currency/                  : to fetch details for all currencies
- /internal/currency/UpdateDb : to update currencies to latest price and to update the DB with any additional currency

To view complete code please visit : https://github.com/shon-phand/CryptoServer 

Code has separated in 3 different sections.
1) controllers
2) domains
2) services

Controllers: this uses gin framework and interact with endpoints and calls the required services to fetch the data

Domains: domains has models like structs to store input and outout while trasferring the objects and DAO.

Services : This layer separate the bussiness logic from database. This will call DAO to get the response from database.
If we require to change the database, we need to change only DAO logic(interactions with databases), underlying logic will be same in the services.



For currency Storage in-memory, Redis is used. 
Redis is very efficient in-memory  key-value pairs database.

External Dependies used:

gin-gonic/gin  : light weight web framework
redis          : in memory storage
/stretchr/testify/assert : for testing purpose

To handle the dependancies go mod is used. 

#### While updating the DB with all current currency info it took around 164 seconds.
#### With go-routines and go channels implemented in the code, it took only 18 seconds to update DB with latest values.





To run a code please use below docker-compose file, you can find the docker file in code or 
```sh
version: "3"
services:
  redis:
    image: shonphand/redis
    restart: always
  go-app:
    image: shonphand/cryptoserver:1.0.1
    restart: always
    depends_on: 
      - redis
    ports:
      - "8080:8080"
      
```





