
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

## Building Object
```
type Building struct {
	Id
	Location
	Dimension
}

```
## Room Object
```
type Room struct {
	Id
	BuildingId
	Location
	Dimension
	RoomAssociates
}
```

|Name|Path|Request|Response|Descripiton|
|---|---|---|---|---|
|createResource|/api/resource/create|Resource|void|Creates a resource|
|updateResource|/api/resource/update|Resource,(Room/Building)|void|Updates a resource|
|removeResource|/api/resource/remove|Resource|void|Remove a resource|
|createRoom|/api/resource/room/create|Room|void|Creates a room|
|updateRoom|/api/resource/room/update|Room|void|Updates a room|
|removeRoom|/api/resource/room/remove|Room|void|Remove a room|
|listRoomResources|/api/resource/room/list|Room|Resource[]|List a rooms resources|
|createBuilding|/api/resource/building/create|Building|void|Creates a room|
|updateBuilding|/api/resource/building/update|Building|void|Updates a room|
|removeBuilding|/api/resource/building/remove|Building|void|Remove a room|
|listBuildingRooms|/api/resource/building/list|Building|Rooms[]|List a buildings rooms|
