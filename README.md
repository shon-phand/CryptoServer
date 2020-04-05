# CryptoServer

endpoints are as below

/currency/:symbol           : to fetch details for specific currency
/currency/                  : to fetch details for all currencies
/internal/currency/UpdateDb : to update currencies to latest price and to update the DB with any additional currency

To view complete code please visit : https://github.com/shon-phand/CryptoServer 

To run a code please use below docker-compose file

version: "3"
services:
  redis:
    image: shonphand/redis
    restart: always
  go-app:
    image: shonphand/cryptoserver:1.0.0
    restart: always
    ports:
      - "8080:8080"
      
   
# you need to run endpoint /internal/currency/UpdateDb to update DB whenever you started the containers
# it will take 18 seconds to update the complete DB (908 currencies)




