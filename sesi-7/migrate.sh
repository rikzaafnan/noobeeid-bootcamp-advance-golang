curl \
  -X POST 'http://localhost:7700/indexes/movies/documents?primaryKey=id' \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer ThisIsMasterKey' \
  --data-binary @movies.json