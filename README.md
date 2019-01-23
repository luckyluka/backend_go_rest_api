# BACKEND GO REST API



**INSTRUCTIONS - Starting API**

1. download the api:

>go get github.com/luckyluka/backend_go_rest_api

2. create docker container
>docker run golang go get -v github.com/luckyluka/backend_go_rest_api

3. get container id
>docker ps -lq

4. create an image (using id from previous step)
>docker commit "container_id" backend_go_rest_api

5. run api
>docker run -p 8080:8080 backend_go_rest_api backend_go_rest_api



**API USAGE**



_GET CERTIFICATES_

>curl localhost:12345/users/{user_id}/certificates/ -X GET

example response

{"id":"1","title":"SunScape","createdAt":"Florence","ownerId":"1","year":"1982","transfer":{}}
{"id":"2","title":"EveningGrass","createdAt":"Milan","ownerId":"1","year":"1972","transfer":{}}



_CREATE CERTIFICATE_

request body (json format)

{
    "title":"novaPictura",
    "createdAt":"Mexico",
    "year":"1992",
    "note":"green hills"
}


>curl localhost:12345/users/{user_id}/certificates/{certificate_id} -X POST -d "@file.json" -H "OwnerId:{owner_id}"

example:
>curl localhost:12345/users/1/certificates/1 -X POST -d "@test.json" -H "OwnerId:1"


_UPDATE CERTIFICATE_

request body (json format)

{
    "title":"Vases",
    "createdAt":"Rome",
    "year":"1984",
    "note":"scenery"
}


>curl localhost:12345/users/{user_id}/certificates/{certificate_id}/transfers -X PATCH -d "@file.json" -H "OwnerId:{owner_id}"

example:

>curl localhost:12345/users/1/certificates/1 -X PATCH -d "@test.json" -H "OwnerId:1"


_DELETE CERTIFICATE_


>curl localhost:12345/users/{user_id}/certificates/{certificate_id} -X DELETE -H "OwnerId:{owner_id}"

>curl localhost:12345/users/1/certificates/1 -X DELETE -H "OwnerId:1"


_CREATE CERTIFICATE TRANSFER_

request body (json)
{
    "transfer":{
    "to":"{owner_id}"}
}

curl localhost:12345/users/{user_id}/certificates/{certificate_id}/transfers -X POST -d "@transfer.json" -H "OwnerId:{owner_id}"

curl localhost:12345/users/2/certificates/1/transfers -X POST -d "@transfer.json" -H "OwnerId:1"


_ACCEPT CERTIFICATE TRANSFER_

request body (json)
{
    "transfer":{
    "status":"completed"}
}

curl localhost:12345/users/{user_id}/certificates/{certificate_id}/transfers -X PATCH-d "@transfer.json" -H "OwnerId:{owner_id}"

curl localhost:12345/users/2/certificates/1/transfers -X PATCH-d "@transfer.json" -H "OwnerId:1"





