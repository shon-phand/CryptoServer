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

Response screenshots :
To fetch one currency :
/currency/BTCUSD

``` sh

shon@shon-SVF15212SNW:~$ curl -X GET localhost:8080/currency/BTCUSD

{
  "id": "BTC",
  "ask": "7039.61",
  "bid": "7038.28",
  "last": "7032.91",
  "open": "6772.25",
  "low": "6673.67",
  "high": "7104.64",
  "feeCurrency": "USD"
}
```

To fetch all currency :
/currency/

``` sh

curl -X GET localhost:8080/currency/

[
  {
    "id": "HGT",
    "ask": "0.00000010608",
    "bid": "0.00000009400",
    "last": "0.00000009400",
    "open": "0.00000009508",
    "low": "0.00000009400",
    "high": "0.00000009508",
    "feeCurrency": "BTC"
  },
  {
    "id": "EOS",
    "ask": "2.40973",
    "bid": "2.39028",
    "last": "2.39690",
    "open": "2.29056",
    "low": "2.26215",
    "high": "2.42470",
    "feeCurrency": "DAI"
  },
  {
    "id": "DRG",
    "ask": "0.000004328",
    "bid": "0.000004210",
    "last": "0.000004260",
    "open": "0.000004380",
    "low": "0.000004214",
    "high": "0.000004402",
    "feeCurrency": "BTC"
  },
  {
    "id": "PBTT",
    "ask": "0.0000100",
    "bid": "0.0000099",
    "last": "0.0000099",
    "open": "0.0000099",
    "low": "0.0000099",
    "high": "0.0000100",
    "feeCurrency": "BTC"
  },
  {
    "id": "BERRY",
    "ask": "0.0000010281",
    "bid": "0.0000009668",
    "last": "0.0000009817",
    "open": "0.0000009858",
    "low": "0.0000009817",
    "high": "0.0000009858",
    "feeCurrency": "ETH"
  },
  ...
  }]


```



