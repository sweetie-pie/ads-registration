#!/usr/bin/env bash

# login endpoint
curl "http://localhost:8080/api/login" \
    -i -X POST \
    -H 'Content-Type: application/json' \
    -d '
    {
        "username": "admin",
        "password": "12345",
        "email": "najafizadeh21@gmail.com"
    }
'

# signup endpoint
curl "http://localhost:8080/api/signup" \
    -i -X POST \
    -H 'Content-Type: application/json' \
    -d '
    {
        "username": "amirhossein",
        "password": "12345",
        "email": "najafizadeh21@gmail.com"
    }
'

# create new ad
curl "http://localhost:8080/api/ads" \
    -H 'Content-type: multipart/form-data' \
    -H 'x-token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfbGV2ZWwiOjMsImV4cCI6MTcwNTc1OTYwMSwidXNlcl9pZCI6MSwidXNlcm5hbWUiOiJhZG1pbiJ9.ISOYZ4pFgWDRov3OGqggrM0bhnE_f9mQti86mly1Qzc' \
    -F title="my ad" \
    -F description="testing ad" \
    -F categories="test,test2,test3" \
    -F image='@README.md'
