## To run the application
- create *config.yml* by *config.example.yml* in /config folder
- type *make startapp* to console


## Migrations
- migrate -path db/migrations -database postgresql://postgres:111111@localhost:5432/ads?sslmode=disable up
- migrate -path db/migrations -database postgresql://postgres:111111@localhost:5432/ads?sslmode=disable down


## Routes
### Get all advertisements
- http://localhost:8080/advertisements?offset=0&limit=5 <br>
  offset=0 - first page <br>
  offset=10 - second page <br>
  offset=20 - third page <br>
  ... <br>
  limit=0 - default <br>
  limit=5 - 5 advertisements on page <br>
  ...
- http://localhost:8080/advertisements?offset=0&dateSort=newest <br />
  dateSort=newest <br />
  dateSort=oldest 
- http://localhost:8080/advertisements?offset=0&priceSort=expensive <br />
  dateSort=expensive <br />
  dateSort=cheap
  
### Get one advertisement
- http://localhost:8080/advertisements/id

### Create advertisement
- Send JSON to http://localhost:8080/advertisements

### Delete advertisement
- http://localhost:8080/advertisements/id
