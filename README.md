
# 🚲 Bike Rental GraphQL API

This API allows management of bikes and rentees in a bike rental system. It is built using `gqlgen` in Go and communicates with bike and rentee microservices via gRPC.

## 📦 Models

### `Bike`
```graphql
type Bike {
  id: ID!
  ownerName: String!
  type: String!
  make: String!
  serial: String!
  renteeId: String
}
```

### `Rentee`
```graphql
type Rentee {
  id: ID!
  firstName: String!
  lastName: String!
  nationalIdNumber: String!
  phone: String!
  email: String!
  heldBikes: [String!]!
}
```

## ✍️ Inputs

### `NewBike`
```graphql
input NewBike {
  ownerName: String!
  type: String
  make: String
  serial: String
}
```

### `UpdateBikeInput`
```graphql
input UpdateBikeInput {
  id: ID!
  type: String
  make: String
  serial: String
  ownerName: String
}
```

### `NewRenteeInput`
```graphql
input NewRenteeInput {
  firstName: String!
  lastName: String!
  nationalIdNumber: String!
  phone: String!
  email: String!
  heldBikes: [String!]!
}
```

### `RenteeInput`
```graphql
input RenteeInput {
  id: ID!
  firstName: String!
  lastName: String!
  nationalIdNumber: String!
  phone: String!
  email: String!
  heldBikes: [String!]!
}
```

## 🔍 Queries

### Get a list of all bikes
```graphql
query {
  listBikes {
    id
    type
    make
    ownerName
    serial
  }
}
```

### Get bike by ID
```graphql
query {
  getBike(id: "bike-id") {
    id
    type
    make
    serial
    ownerName
  }
}
```

### Get bikes by type
```graphql
query {
  getBikesByTYPE(typeArg: "mountain") {
    id
    make
    ownerName
  }
}
```

### Get bikes by make
```graphql
query {
  getBikesByMAKE(make: "Yamaha") {
    id
    type
    ownerName
  }
}
```

### List all rentees
```graphql
query {
  listRentees {
    id
    firstName
    lastName
    phone
    heldBikes
  }
}
```

## 🔧 Mutations

### Add a new bike
```graphql
mutation {
  addBike(bike: {
    ownerName: "John Doe",
    type: "road",
    make: "Trek",
    serial: "ABC123"
  }) {
    id
    type
    ownerName
  }
}
```

### Delete a bike
```graphql
mutation {
  deleteBike(id: "bike-id")
}
```

### Add a new rentee
```graphql
mutation {
  addRentee(rentee: {
    firstName: "Jane",
    lastName: "Smith",
    nationalIdNumber: "123456789",
    phone: "123-456-7890",
    email: "jane@example.com",
    heldBikes: []
  }) {
    id
    email
  }
}
```

### Update a rentee
```graphql
mutation {
  updateRentee(rentee: {
    id: "rentee-id",
    firstName: "Jane",
    lastName: "Doe",
    nationalIdNumber: "123456789",
    phone: "987-654-3210",
    email: "jane.doe@example.com",
    heldBikes: ["bike-id"]
  }) {
    id
    heldBikes
  }
}
```

---

## 📘 Notes

- Backend uses gRPC clients (`BikeClient`, `RenteeClient`) to perform actual data operations.
- Some queries like `getBike` are not yet implemented and may panic when called.
- `HeldBikes` refers to the IDs of bikes currently held by the rentee.
