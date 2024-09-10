# About
Boilerplate of Fairplay KSM implementation with Golang Fiber following https://github.com/payt0nc/fairplay-ksm and https://github.com/easonlin404/ksm.
Please check respective github repos for further details.


# Development
### In order to have an optimal development experience you need to have Docker installed.
Create a copy of `.env.example`, rename it to `.env` and update the following keys accordingly:
- FAIRPLAY_CERTIFICATION
- FAIRPLAY_PRIVATE_KEY
- FAIRPLAY_APPLICATION_SERVICE_KEY

If you follow along, other environment variables won't be required to modify.

After pulling, go to the root of the project directory, make sure `docker`
is running and then run the commands:
```
docker compose up --build -d
```
This will spin everything up with docker-compose, ie, start the service, database and redis.

For consecutive times, you can run:
`docker compose up -d`


## Health Check

`curl --location 'localhost:8080/v1/fps/health'`

### Sample Response
`{
"message": "Service is healthy",
"status": "ok"
}`

## Get License (Couldn't make it work till now)

```
curl --location 'localhost:8080/v1/fps/license' \
--header 'Content-Type: application/json' \
--data '{
		"spc": "<SPC>",
		"assetId": "987654321"
	}' 
```

### Expected Response

```
HTTP/1.1 200 OK
Content-Length: 902
Content-Type: application/json; charset=utf-8
Date: Thu, 08 Mar 2018 05:51:02 GMT

{
    "ckc": "<CKC>"
}



```


## Stop the Service

`docker compose stop`



