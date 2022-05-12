
# User

## createUser

**PUT** api/user/create <br>

send (application/json)
```    
Body
{
  *username: string,
  first: string,
  surname: string,
  email: string,
  picture: string
  *password: string,
  *password_key: string     
}
```
recieve (application/json)
```   
{
  status: integer,
  result: {
    message: string
    credential
  }
}
```

## validateUsername

**GET** api/user/validate <br>

send (application/json)
```
Parameters{
  username: string
}  
```  
recieve (application/json)  
```
{
    status: integer,
    result:{
      valid: boolean
    }
}
```

## loadUser
**GET** api/user/profile <br>

send (application/json)
```
Parameters
{
  credential,
  usernames: [ //if left blank will return all users
    username: string
  ] 
}
```
recieve (application/json)
```
{
  status: integer,
  result:{
    users: [
      username: string,
      first: string,
      surname: string,
      email: string,
      picture: string,
      teams: {
        team_name: string,
        team_id: string
      }
    ]
  }
}
```

## updateUser

**POST** api/user/profile <br>

send (application/json)
```
Parameters
{
  credential,
  username: string
}
Body
{
  username: string,
  first: string,
  surname: string,
  email: string,
  picture: string
  password: string,
  password_key: string
}
```
recieve (application/json)
```
{
  status: integer,
  result: {
      message: string,
      fields: [ //list of fields updated
        field: value
      ]                
  }
}
```

## deleteUser
**DELETE** api/user/profile <br>

send (application/json)
```
Parameters
{
  credential,
  username: string
}
```
recieve (application/json)
```
{
  status: integer,
  result: {
    message: string
  }
}
```

## loginUser
**POST** api/user/login <br>

send (application/json)
```
Body
{
  user_id: string,
  username: string,
  password: string
}
```
---
# Team

## createTeam
**PUT** api/team/create <br>

send (application/json)
```
Body
{
  *name: string,
  description: string,
  capacity: integer,
  picture: string    
}
```
recieve (application/json)
```
{
  status: integer,
  result: {
      message: string
  }
}
```

## loadTeam
**GET** api/team/profile <br>

send (application/json)
```
Parameters
{
  credential,
  teams:[ //if blank will return all teams
    team_id: string
  ]
}
```
recieve (application/json)
```
{
  status: integer,
  result: {
    teams: [
      team_id: {
        name: string,
        description: string,
        capacity: integer,
        picture: string,
        date_created: date,
        users: [
          username: string
        ]
      }
    ]
  }
}
```

## updateTeam
**POST** api/team/profile <br>

send (application/json)
```
Parameters
{
  credentials,
  team_id: string
}
Body
{
  name: string,
  description: string,
  picture: string,
  capacity: string 
}
```
recieve (application/json)
```
{
  status: integer,
  result: {
    message: string
  }
}
```

## addTeamMember
**PUT** api/team/members <br>

send (application/json)
```
Parameters
{
  credential
}
Body
{
  team_id: string,
  users: [(list of usernames to add)]
}
```
recieve (application/json)
```
{
  status: integer,
  result: {
    message: string
  }
}
```

## removeTeamMember
**DELETE** api/team/members <br>

send (application/json)
```
Parameters
{
  credential
}
Body
{
  team_id: string,
  users: [(list of usernames to add)]
}
```
recieve (application/json)
```
{
  status: integer,
  result: {
    message: string
  }
}
```
---

# Role (WIP)

## createRole
**PUT** api/role/create <br>

send (application/json)
```
Parameters
{
  credential
}
Body
{
  name: string
}
```
recieve (application/json)
```
{
  status: integer,
  result: {
      role_id: string
  }
}  
```
---

# Bookings

## createBooking
**PUT** api/booking/reserve <br>

send (application/json)
```
Body        
{
  users: [(list of usernames)
    username: string
  ],
  resource_type: string,
  resource_preference_id: string,
  start: Date,
  end: Date   
}
```
recieve (application/json)
```
{
    status: integer
}
```

## updateBooking
**POST** api/booking/update <br>

send (application/json)
```
Parameters
{
  credentials,
  booking_id: string
}
Body
{
  resource_type: string,
  resource_preference_id: string,
  start: date,
  end: date,
  booked: boolean
}
```
recieve (application/json)
```
{
  status: integer,
  result: {
    message: string
  }
}
```

## deleteBooking
**DELETE** api/booking/update <br>

send (application/json)
```
Parameters
{
  credential
  bookings: [
    booking_id: string
  ]
  
}
```
recieve (application/json)
```
{
  status: integer,
  result: {
    message: string
  }
}
```

---

# Workspace

## createBuilding
**PUT** api/resources/building/create <br>

send (application/json)
```
Parameters
{
  credential
}
Bodyload
{
  location: {},
  dimensions: {},
  name: string
}
```
recieve (application/json)
```
{
  status: integer,
  result: {
    message: string
  }
}
```

## loadBuidling
**GET** api/resources/building <br>

send (application/json)
```
Parameters
{
  credential,
  buildings: [
    building_id: string //if blank will return all buildings
  ]  
}

```
recieve (application/json)
```
{
  status: integer,
  result: {
  buildings: [
    building_id: {
      location: {},
      dimensions: {},
      name: string,
      rooms: [
        room_id: string
      ]
    }
  ]


  }
}
```

## updateBuilding
**POST** api/resources/building/update <br>

send (application/json)
```
Parameters
{
  credential
}
Body
{
  building_id: string,
  fields: {
    (fields to update with new value)
    field: value
  }
}
```
recieve (application/json)
```
{
  status: integer,
  result: {
    message: string
  }
}
```

## createRoom
**PUT** api/resources/room/create <br>

send (application/json)
```
Parameters
{
  credential
}
Body
{
  location: {},
  dimensions: {},
  name: string,
  building_id: string
}
```
recieve (application/json)
```
{
  status: integer,
  result: {
    message: string
  }
}
```
## updateRoom
**POST** api/resources/room/update <br>

send (application/json)
```
Parameters
{
  credential
}
Body
{
  room_id: string,
  fields: {
    (fields to update with new value)
    field: value
  }
}
```
recieve (application/json)
```
{
  status: integer,
  result: {
    message: string
  }
}
```

## loadRoom
**GET** api/resources/room <br>

send (application/json)
```
Parameters
{
  credential,
  rooms: [
    room_id: string
  ]  
}
```
recieve (application/json)
```
{
  status: integer,
  result: {
    rooms: [
      room_id: {
        features: [],
        building_id: string,
        location: {},
        dimension: {},
        associations: [
          room_id: string
        ]
      }
    ]
  }
}
```

## deleteRoom
**DELETE** api/resources/room/remove <br>

send (application/json)
```
Parameters
{
  credential,
  room_id: string
}
```
recieve (application/json)
```
{
  status: integer,
  result: {
    message: string
  }
}
```
=========
# vvvvv WIP vvvvv
=========

## createWorkspace
**PUT** api/resources/workspace/create <br>

send (application/json)
```
Parameters
{
  credential
}
Body
{
  location: {},
  dimensions: {},
  name: string,
  building_id: string
}
```
recieve (application/json)
```
{
  status: integer,
  result: {
    message: string
  }
}
```
## updateWorkspace
**POST** api/resources/workspace/update <br>

send (application/json)
```
Parameters
{
  credential
}
Body
{
  room_id: string,
  fields: {
    (fields to update with new value)
    field: value
  }
}
```
recieve (application/json)
```
{
  status: integer,
  result: {
    message: string
  }
}
```

## loadWorkspace
**GET** api/resources/workspace <br>

send (application/json)
```
Parameters
{
  credential,
  workspaces: [
    workspace_id: string
  ]
}
```
recieve (application/json)
```
{
  status: integer,
  result: {
    workspaces: [
      workspace_id: {
        rooms: [
          room_id: 
        ]
      }
    ]
  }
}
```

## deleteWorkspace
**DELETE** api/resources/room/remove <br>

send (application/json)
```
Parameters
{
  credential,
  room_id: string
}
```
recieve (application/json)
```
{
  status: integer,
  result: {
    message: string
  }
}
```