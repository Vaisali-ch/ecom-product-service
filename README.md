# ecom-product-service
This microservice featuring product service for a ecommerce application, This has been done as part of college project

## folder structure
ecom-product-service
   1. config/  // this folder contains db configurations details needed for the product services
   2. controllers/ // this folder contains business logic to perform crud operationd needed for the service
   3. models/ // models contains request and response type/model structure of the product service
   4. routes/ // routes contains logic for the URL mapping for the api.
   main.go file used to call handlers to with start of services.
   Docker file used to create docker docker image to deploy service.
   .gitignore file used to add those files which needs to avoid adding in git repo.
   go.mod & go.sum having all the dependency package needed for this service.

### Rest APIs:
     1. GET	/products	                ///  List all products
     2. GET	/products/:id	            ///  Get a product
     3. POST	/products	            ///  Add a product
     4. PUT	/products/:id	            ///  Update a product
     5. DELETE	/products/:id	        ///  Delete a product
     6. GET	/products/:id/availability	///  Check availability