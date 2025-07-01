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

*usage example:*
`api.webstatistics.example.com/put-traffic?sessionId=9069c164-d8f5-4734-bb8c-72d12f6e788e&p=/&site=example.com`
