
# User

## User Object
```
type User struct {
	Id
	Identifier
	FirstName
	LastName
	Email
	Picture
	DateCreated
}
```

|Name|Path|Request|Response|Descripiton|
|---|---|---|---|---|
|createUser|/api/user/register|User|void|Creates a user|
|loadUser|/api/user/information|User[]|User[]|Load users. Will return all if left blank, else will only return data of specified users|
|updateUser|/api/user/update|User|void|Update user according to fields|
|deleteUser|/api/user/remove|User|void|Delete user|
|loadRoles| /api/user/roles|User|Role[]|Loads all the roles a user has|
|loadTeams| /api/user/teams|User|Team[]|Loads all the teams a user is in|

# Role

## Role Object
```
type Role struct {
	Id
	RoleName
	DateAdded
}
```

|Name|Path|Request|Response|Descripiton|
|---|---|---|---|---|
|createRole|/api/role/create|Role|void|Creates a role|
|loadRole|/api/role/information|Role[]|Role[]|Load roles. Will return all if left blank, else will only return data of specified roles|
|loadUsersWithRole|/api/role/users|Role|User[]|Returns an array of users with role|
|removeRole|/api/role/remove|Role|void|Deletes a role|
|updateRole|/api/role/update|Role|void|Updates role values|

# Permission
 
## Permission Object
```
type Permission struct {
	Id
	PermissionType
	PermissionCategory
	PermissionTenant
	PermissionTenantId
	DateAdded
}
```

|Name|Path|Request|Response|Descripiton|
|---|---|---|---|---|
|createPermission|/api/permission/create|Permission|void|Creates a permission|
|allocatePermission|/api/permission/allocate|Role,Permission|void|Allocates permission to a role|
|removePermission|/api/permission/remove|Role,Permission|void|Removes a roles permission where permission matches|
|loadRolePermissions|/api/permission/list|Role|Permission[]|Returns all permissions a role has|
|updatePermission|/api/permission/update|Permission|void|Updates Permission|

# Team

## Team Object

```
type Team struct {
	Id
	Name
	Description
	Capacity
	Picture
	DateCreated
}
```

|Name|Path|Request|Response|Descripiton|
|---|---|---|---|---|
|createTeam|/api/team/create|Team|void|Creates a team|
|removeTeam|/api/team/remove|Team|void|Removes a team|
|loadTeam|/api/team/information|Team|Team|Returns all Team data given team_id|
|loadTeamMembers|/api/team/list|Team|Userp[|Returns all team members in team|
|updateTeamMembers|/api/team/members|Team,User[]|void|Update members in team|
|UpdateTeam|/api/team/update|Team|void|Updates a team|

# Resources

## Building

## Room


## createUser

This is used to create a new user 

**PUT** api/user/create <br>

request (application/json)
```    
Body
{
  <user_object>   
}
```
response (application/json)
```   
{
  status,
  result: {
    message
  }
}
```

## validateUsername

**GET** api/user/validate <br>

request (application/json)
```
Parameters{
  user_id
}  
```  
response (application/json)  
```
{
  status
}
```

## loadUser
**GET** api/user/profile <br>

request (application/json)
```
Parameters
{
  users: [] 
}
```
response (application/json)
```
{
  status,
  result:{
    users: []
  }
}
```

## updateUser

**POST** api/user/profile <br>

request (application/json)
```
Parameters
{
  user
}
Body
{
  <user_object>
}
```
response (application/json)
```
{
  status,
  result: {
      message,
      fields: []                
  }
}
```

## deleteUser
**DELETE** api/user/profile <br>

request (application/json)
```
Parameters
{
  user
}
```
response (application/json)
```
{
  status,
  result: {
    message
  }
}
```

## loginUser
**POST** api/user/login <br>

request (application/json)
```
Body
{
  user_id,
  username,
  password
}
```
---
# Team

## createTeam
**PUT** api/team/create <br>

request (application/json)
```
Body
{
  *name: string,
  description: string,
  capacity: integer,
  picture: string    
}
```
response (application/json)
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

request (application/json)
```
Parameters
{
  credential,
  teams:[ //if blank will return all teams
    team_id: string
  ]
}
```
response (application/json)
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

request (application/json)
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
response (application/json)
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

request (application/json)
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
response (application/json)
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

request (application/json)
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
response (application/json)
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

request (application/json)
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
response (application/json)
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

request (application/json)
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
response (application/json)
```
{
    status: integer
}
```

## updateBooking
**POST** api/booking/update <br>

request (application/json)
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
response (application/json)
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

request (application/json)
```
Parameters
{
  credential
  bookings: [
    booking_id: string
  ]
  
}
```
response (application/json)
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

request (application/json)
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
response (application/json)
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

request (application/json)
```
Parameters
{
  credential,
  buildings: [
    building_id: string //if blank will return all buildings
  ]  
}

```
response (application/json)
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

request (application/json)
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
response (application/json)
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

request (application/json)
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
response (application/json)
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

request (application/json)
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
response (application/json)
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

request (application/json)
```
Parameters
{
  credential,
  rooms: [
    room_id: string
  ]  
}
```
response (application/json)
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

request (application/json)
```
Parameters
{
  credential,
  room_id: string
}
```
response (application/json)
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

request (application/json)
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
response (application/json)
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

request (application/json)
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
response (application/json)
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

request (application/json)
```
Parameters
{
  credential,
  workspaces: [
    workspace_id: string
  ]
}
```
response (application/json)
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

request (application/json)
```
Parameters
{
  credential,
  room_id: string
}
```
response (application/json)
```
{
  status: integer,
  result: {
    message: string
  }
}
```