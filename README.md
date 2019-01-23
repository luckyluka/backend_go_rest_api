# BACKEND GO REST API

## INSTRUCTIONS - Starting API
<br/>

##### 1. DOWNLOAD THE API

>go get github.com/luckyluka/backend_go_rest_api
<br/>

##### 2. CREATE DOCKER CONTAINER

>docker run golang go get -v github.com/luckyluka/backend_go_rest_api
<br/>

##### 3. GET CONTAINER ID

>docker ps -lq
<br/>

##### 4. CREATE A DOCKER IMAGE (using id from previous step)

>docker commit "container_id" backend_go_rest_api
<br/>

##### 5. RUN API

>docker run -p 8080:8080 backend_go_rest_api backend_go_rest_api
<br/>


## API USAGE

<br/>

To interface with the API you may use curl commands, as is done in the examples below.
You may also use provided json files, where there is a need for request body.

<br/>

### GET CERTIFICATES
<br/>

##### API REQUEST

###### /users/{user_id}/certificates/


##### RESPONSE BODY

>{"id":"1","title":"SunScape","createdAt":"Florence","ownerId":"1","year":"1982",Note: "River"}
>{"id":"2","title":"EveningGrass","createdAt":"Milan","ownerId":"1","year":"1972",Note: "Moon"}


##### EXAMPLE
>curl localhost:8080/users/1/certificates -X GET
<br/>

### CREATE CERTIFICATE
<br/>

##### API REQUEST

###### /users/{user_id}/certificates/{certificate_id}

##### REQUEST HEADER
###### POST /users/1/certificates/1/ HTTP/1.1 / Host: localhost:12345 /OwnerId: {owner_id} / Content-Type: application/x-www-form-urlencoded

##### REQUEST BODY
{
    
    "title":"novaPictura",
    
    "createdAt":"Mexico",
   
    "year":"1992",
   
    "note":"green hills"

}

##### RESPONSE BODY
{
    
    "title":"novaPictura",
    
    "createdAt":"Mexico",
   
    "year":"1992",
   
    "note":"green hills"

}


##### EXAMPLE
>curl localhost:8080/users/1/certificates/3 -X POST -d "@test.json" -H "OwnerId:1"
<br/>

### UPDATE CERTIFICATE
<br/>

##### API REQUEST

###### /users/{user_id}/certificates/{certificate_id}/transfers

##### REQUEST HEADER
###### POST /users/1/certificates/1/transfers HTTP/1.1 / Host: localhost:12345 /OwnerId: {owner_id} / Content-Type: application/x-www-form-urlencoded

##### REQUEST BODY
{
    
    "title":"novaPictura",
    
    "createdAt":"Mexico",
   
    "year":"1992",
   
    "note":"green hills"

}


##### RESPONSE BODY
{
    
    "title":"novaPictura",
    
    "createdAt":"Mexico",
   
    "year":"1992",
   
    "note":"green hills"

}


##### EXAMPLE
>curl localhost:8080/users/1/certificates/1 -X PATCH -d "@test.json" -H "OwnerId:1"
<br/>

### DELETE CERTIFICATE
<br/>

##### API REQUEST

###### /users/{user_id}/certificates/{certificate_id}

##### REQUEST HEADER
###### POST /users/{user_id}/certificates/{certificate_id}/ HTTP/1.1 / Host: localhost:12345 /OwnerId: {owner_id} / Content-Type: application/x-www-form-urlencoded

##### EXAMPLE
>curl localhost:8080/users/1/certificates/1 -X DELETE -H "OwnerId:1"
<br/>

### CREATE CERTIFICATE TRANSFER
<br/>

##### API REQUEST

###### /users/{user_id}/certificates/{certificate_id}/transfers

##### REQUEST HEADER
###### POST /users/{user_id}/certificates/{certificate_id}/transfers HTTP/1.1 / Host: localhost:12345 /OwnerId: {owner_id} / Content-Type: application/x-www-form-urlencoded

##### REQUEST BODY

{

    "transfer":{
    
    "to":"{owner_id}"}

}

##### RESPONSE BODY
{
    
    "title":"novaPictura",
    
    "createdAt":"Mexico",
   
    "year":"1992",
   
    "note":"green hills"
   
    "transfer":{
    
    "to":"{owner_id}"}

}


##### EXAMPLE
>curl localhost:8080/users/1/certificates/1/transfers -X POST -d "@transfer.json" -H "OwnerId:1"
<br/>

### ACCEPT CERTIFICATE TRANSFER
<br/>

##### API REQUEST

###### /users/{user_id}/certificates/{certificate_id}/transfers

##### REQUEST HEADER
###### POST /users/{user_id}/certificates/{certificate_id}/transfers HTTP/1.1 / Host: localhost:12345 /OwnerId: {owner_id} / Content-Type: application/x-www-form-urlencoded


##### REQUEST BODY
{
    
    "transfer":{
    
    "status":"completed"}
}

##### RESPONSE BODY
{
    
    "title":"novaPictura",
    
    "createdAt":"Mexico",
   
    "year":"1992",
   
    "note":"green hills"

}

##### EXAMPLE
>curl localhost:8080/users/2/certificates/1/transfers -X PATCH -d "@transfer_accept.json" -H "OwnerId:2"


<br/>

## UNIT TESTING
<br/>

There are basic unit tests provided in the backend_test.go


#### RUN TESTS

>go test
