# System Design Choices 

## Goals
- Service should generate a shortened link when given a URL 
- User should be redirected to original link when given a shortened link 
- User can pick a custom shortened url
- Links will expire after a certain time period.  
- Minimal latency when happening in real time 
- Caching a percentage of the links for faster access times 

## Assumptions
- Assume a lot more read access than writing to database 
- Not going to be that many people using this app. 

## Database Design 
- Choosing to use Redis Database for fast access and caching. 

## Algorithm Choices 
- Base 62 Encoding for a randomized number generator (can lead to replicating issues)

# To Test Program: 
```
docker-compose up --build
```

Example command of POST to redis db
```
curl -L -X POST 'localhost:8080/encode' \         
-H 'Content-Type: application/json' \
--data-raw '{
    "url": "https://www.example.com",
    "expires": "2022-10-04 17:18:00"
}'
```

Example command of GET from redis db
```
curl -L -X GET "http://localhost:8080/ZNkw23qpiss"
```