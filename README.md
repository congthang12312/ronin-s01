# Implement local cache

## Content: List of Airports (Code, Name)

- Read Aside with TTL
- Allow to use libs
- Get used to Redis

### Quick Run Project:
    
    Prepare redis in local:
        docker-compose up -d

    Run go in package: /api/cmd/serverd/main with local.env file

    curl: 
        curl --location 'http://127.0.0.1:3000/airport-service/v1/airports?code=VJ1'
    