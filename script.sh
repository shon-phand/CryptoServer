!#/bin/bash

echo "script started" >> ~/Documents/linux/cron.log

curl -X GET localhost:8080/internal/currency/UpdateDb >> ~/Documents/linux/cron.log

echo "script ended" >> ~/Documents/linux/cron.log

