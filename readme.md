## Migrations
- migrate -path db/migrations -database postgresql://postgres:111111@localhost:5432/ads?sslmode=disable up
- migrate -path db/migrations -database postgresql://postgres:111111@localhost:5432/ads?sslmode=disable down


## Routes
### Get all advertisements
- http://localhost:8080/advertisements?offset=0 <br />
  offset=0 - first page <br />
  offset=10 - second page <br />
  offset=20 - third page <br />
  ...
- http://localhost:8080/advertisements?offset=0&dateSort=newest <br />
  dateSort=newest <br />
  dateSort=oldest 
- http://localhost:8080/advertisements?offset=0&priceSort=expensive <br />
  dateSort=expensive <br />
  dateSort=cheap
  
### Create advertisement
- http://localhost:8080/create?Title=Mercedes S63 AMG Coupe&Description=Отличное состояние, максимальная комплектация.&Price=2827165&Photo_1=https://cdn4.riastatic.com/photosnew/auto/photo/mercedes-benz_s-63-amg__451377334hd.webp&Photo_2=https://cdn3.riastatic.com/photosnew/auto/photo/mercedes-benz_s-63-amg__451377378hd.webp&Photo_3=https://cdn0.riastatic.com/photosnew/auto/photo/mercedes-benz_s-63-amg__451378785hd.webp

### Delete advertisement
- http://localhost:8080/delete/id
