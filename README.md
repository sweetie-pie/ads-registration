# Ads Registration

Creating an Ads registration service using __Golang__ and __MySQL__.

## Setup

Use the following command to set up the project:

```shell
docker compose up -d
```

## Models

Our database models are as below:

### User

- ID
- Username (string)
- Password (string)
- AccessLevel (enum:={"viewer=1", "writer=2", "admin=3"})

### Ad

- ID
- Title (string)
- Description (string)
- Status (enum:={"published=1", "rejected=2", "pending=3"})
- Created At (string)
- Image (string)
- User (ID)

### Category

- ID
- Title (string)
- AdID (ID)
