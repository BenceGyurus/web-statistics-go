# web-statistics-go


## Documentáció

### Environment variables:
```bash
PORT=3000
GIN_MODE=relase
CASSANDRA_CONTACT_POINT=
CASSANDRA_USERNAME=
CASSANDRA_PASSWORD=
PREFIX=
```

### API
#### `GET` /put-traffic

##### queries:
`sessionId`: the unique id of the user

`p`: the URL path of the visited page

`site`: the domain of the visited website 

##### usage example:
`api.webstatistics.example.com/put-traffic?sessionId=9069c164-d8f5-4734-bb8c-72d12f6e788e&p=/&site=example.com`

##### response:
the response will contain the sent sessionId. If you didn’t send any sessionId, a new one will be generated, which you should use to identify the user from now on.

#### `POST` /traffic
It provides the number of unique visitors

##### queries:
`from`: the starting date for retrieving visitors

`to` : the ending date for retrieving visitors

`page` : the website to be tracked

##### response:
```
{
    "traffic": 2
}
```

#### `POST` /sites
It shows how many unique visitors the different pages have.

##### queries:
`from`: the starting date for retrieving visitors

`to` : the ending date for retrieving visitors

`page` : the website to be tracked

##### response:
```
{
    
}
```


#### `POST` /graph
it shows unique visitors across multiple sections

##### queries:
`from`: the starting date for retrieving visitors

`to` : the ending date for retrieving visitors

`page` : the website to be tracked

`intervals` : number of sections

##### response:
```
[
    {
        "interval": 0,
        "uniqueSessions": 0,
        "totalRequests": 0
    },
    {
        "interval": 1,
        "uniqueSessions": 0,
        "totalRequests": 0
    },
    
    .
    .
    .
    
    {
        "interval": 8,
        "uniqueSessions": 0,
        "totalRequests": 0
    },
    {
        "interval": 9,
        "uniqueSessions": 2,
        "totalRequests": 3
    }
]
```

### `POST` /active
it shows how many active users are on the pages in real time


##### response:
```
{
    "count": 0
}
```


#### `POST` /time
it returns the average time that users spent on the site

##### queries:
`from`: the starting date for retrieving visitors

`to` : the ending date for retrieving visitors

`page` : the website to be tracked

##### response:
```
{
    "count": 0
}
```