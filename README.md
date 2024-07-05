# Mock API for time-tracker

External API imitator for https://github.com/nquidox/time-tracker/

## What it does
Generates sqlite DB with some sample data on startup.

App recieves GET request /info?passportSerie=1234&passportNumber=567890 and responses with JSON:
```
{"name":"Name","patronymic":"Patronymic","surname":"Surname","Address":"City name","passportSerie":1234,"passportNumber":567891}
```
Generates random person credentials for every unique serie + number and saves it to DB.


## Setup
Edit .env file params if needed.
```sh
#Server address
FAKEAPI_HOST=127.0.0.1
FAKEAPI_PORT=9001

#SQLite DB
DB_NAME=database.db

#Loggers
API_LOG_LVL=Info
DB_LOG_LVL=Silent
```

## Licence
```
MIT
```

