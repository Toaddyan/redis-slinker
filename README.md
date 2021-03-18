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
- Base 62 Encoding for a SHA256 Hash (can lead to replicating issues)
