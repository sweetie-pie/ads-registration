# Ads Registration

Creating an Ads registration service.

## Models

### User

- ID
- Username (string)
- Password (string)
- Email (string)
- Banned (boolean)

### Admin

- ID
- Username (string)
- Password (string)
- Email (string)
- Active (boolean)
- Access Level (enum:={"viewer", "writer", "admin"})

### Ad

- ID
- Title (string)
- Description (string)
- Status (enum:={"published", "rejected", "pending"})
- Created At (string)
- Image (string)
- Creator (ID)

### Category

- ID
- Title (string)
- Description (string)
