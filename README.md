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

##### 5. CREATE A DOCKER IMAGE (using id from previous step)

>docker commit "container_id" backend_go_rest_api
<br/>

##### 6. RUN API

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

##### /users/{user_id}/certificates/


##### RESPONSE BODY

>{"id":"1","title":"SunScape","createdAt":"Florence","ownerId":"1","year":"1982",Note: "River"}
>{"id":"2","title":"EveningGrass","createdAt":"Milan","ownerId":"1","year":"1972",Note: "Moon"}


##### EXAMPLE
>curl localhost:8080/users/1/certificates/ -X GET

### CREATE CERTIFICATE
<br/>

##### API REQUEST

##### /users/{user_id}/certificates/{certificate_id}

##### REQUEST HEADER
##### POST /users/1/certificates/1/transfers HTTP/1.1 / Host: localhost:12345 /OwnerId: {owner_id} / Content-Type: application/x-www-form-urlencoded

##### REQUEST BODY
{
    
    "title":"novaPictura",
    
    "createdAt":"Mexico",
   
    "year":"1992",
   
    "note":"green hills"

}


##### EXAMPLE
>curl localhost:8080/users/1/certificates/1 -X POST -d "@test.json" -H "OwnerId:1"

#### UPDATE CERTIFICATE
<br/>

##### API REQUEST

##### /users/{user_id}/certificates/{certificate_id}/transfers

##### REQUEST HEADER
##### POST /users/1/certificates/1/transfers HTTP/1.1 / Host: localhost:12345 /OwnerId: {owner_id} / Content-Type: application/x-www-form-urlencoded

##### REQUEST BODY
{
    
    "title":"novaPictura",
    
    "createdAt":"Mexico",
   
    "year":"1992",
   
    "note":"green hills"

}


##### EXAMPLE
>curl localhost:8080/users/1/certificates/1 -X PATCH -d "@test.json" -H "OwnerId:1"
<br/>

#### DELETE CERTIFICATE
<br/>


>curl localhost:8080/users/{user_id}/certificates/{certificate_id} -X DELETE -H "OwnerId:{owner_id}"

>curl localhost:8080/users/1/certificates/1 -X DELETE -H "OwnerId:1"
<br/>


#### CREATE CERTIFICATE TRANSFER

<br/>

>curl localhost:8080/users/{user_id}/certificates/{certificate_id}/transfers -X POST -d "@transfer.json" -H "OwnerId:{owner_id}"

example:
<br/>

>curl localhost:8080/users/2/certificates/1/transfers -X POST -d "@transfer.json" -H "OwnerId:1"
<br/>

_request body (example)_

{

    "transfer":{
    
    "to":"{owner_id}"}

}

<br/>

#### ACCEPT CERTIFICATE TRANSFER
<br/>

>curl localhost:8080/users/{user_id}/certificates/{certificate_id}/transfers -X PATCH-d "@transfer.json" -H "OwnerId:{owner_id}"

>curl localhost:8080/users/2/certificates/1/transfers -X PATCH-d "@transfer.json" -H "OwnerId:1"

<br/>

_request body (example)_
{
    
    "transfer":{
    
    "status":"completed"}
}





