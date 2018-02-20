#PostgraphQL Test

Example postgraphql express API. Grahpiql and dynamic JSON enabled.

Watching also enabled, but of course we would only want that for early development.

Uses a simplified, cleaned up version of our current dashboard schema.

#Samples

Create a profile:
```graphql
mutation {
  createProfile(input:{profile:{name:"Foo Bar"}}) {
    profile {
      name
    }
  }
}
```
or, parametrized:
```graphql
mutation Mutation($input: CreateProfileInput!) {
  createProfile(input: $input) {
    profile {
      name
    }
  }
}
```
```json
{
  "input":{
    "profile": {
    	"name": "Foo Baz"
    }
  }
}
```

Get all profiles:
```graphql
query {
  allProfiles {
    nodes {
      name
      id
    }
  }
}
```

Get profile by ID:
```graphql
query {
  profileById(id: 5) {
      name
      id
  }
}
```
