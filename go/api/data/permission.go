package data

import (
	"api/db"
	"database/sql"
	"encoding/json"
	"time"
)

//////////////////////////////////////////////////
// Structures and Variables

// Permission identifies a permission via common attributes
type Permission struct {
	Id                 *string    `json:"id,omitempty"`
	PermissionId       *string    `json:"permission_id,omitempty"`
	PermissionIdType   *string    `json:"permission_id_type,omitempty"`
	PermissionType     *string    `json:"permission_type,omitempty"`
	PermissionCategory *string    `json:"permission_category,omitempty"`
	PermissionTenant   *string    `json:"permission_tenant,omitempty"`
	PermissionTenantId *string    `json:"permission_tenant_id,omitempty"`
	DateAdded          *time.Time `json:"date_added,omitempty"`
}

// Permissions represent a splice of Permission
type Permissions []*Permission

// CreateGenericPermission creates a permission
func CreateGenericPermission(permissionType string, permissionCategory string, permissionTenant string) *Permission {
	typ := &permissionType
	category := &permissionCategory
	tenant := &permissionTenant

	if *typ == "" {
		typ = nil
	}
	if *category == "" {
		category = nil
	}
	if *tenant == "" {
		tenant = nil
	}

	return &Permission{
		PermissionType:     typ,
		PermissionCategory: category,
		PermissionTenant:   tenant,
	}
}

func CreatePermission(permissionId string, permissionIdType string, permissionType string, permissionCategory string, permissionTenant string, permissionTenantId string) *Permission {
	typ := &permissionType
	category := &permissionCategory
	tenant := &permissionTenant
	tenantId := &permissionTenantId

	if *typ == "" {
		typ = nil
	}
	if *category == "" {
		category = nil
	}
	if *tenant == "" {
		tenant = nil
	}
	if *tenantId == "" {
		tenantId = nil
	}

	return &Permission{
		PermissionId:       &permissionId,
		PermissionIdType:   &permissionIdType,
		PermissionType:     typ,
		PermissionCategory: category,
		PermissionTenant:   tenant,
		PermissionTenantId: tenantId,
	}
}

func (permissions *Permissions) CompareTo(p *Permission) bool {
	for _, permission := range *permissions {
		if *p.PermissionType == *permission.PermissionType && *p.PermissionCategory == *permission.PermissionCategory && *p.PermissionTenant == *permission.PermissionTenant {
			return true
		}
	}
	return false
}

// PermissionDA provides access to the database for authentication purposes
type PermissionDA struct {
	access *db.Access
}

// NewPermissionDA creates a new data access from a underlying shared data access
func NewPermissionDA(access *db.Access) *PermissionDA {
	return &PermissionDA{
		access: access,
	}
}

// Commit commits the current implicit transaction
func (access *PermissionDA) Commit() error {
	return access.access.Commit()
}

//////////////////////////////////////////////////
// Mappers

func mapPermission(rows *sql.Rows) (interface{}, error) {
	var identifier Permission
	err := rows.Scan(
		&identifier.Id,
		&identifier.PermissionId,
		&identifier.PermissionIdType,
		&identifier.PermissionType,
		&identifier.PermissionCategory,
		&identifier.PermissionTenant,
		&identifier.PermissionTenantId,
		&identifier.DateAdded,
	)
	if err != nil {
		return nil, err
	}
	return identifier, nil
}

//////////////////////////////////////////////////
// Functions

// StorePermission stores an identifier
func (access *PermissionDA) StorePermission(identifier *Permission) error {
	_, err := access.access.Query(
		`SELECT 1 FROM permission.identifier_store($1, $2, $3, $4, $5, $6)`, nil,
		identifier.PermissionId, identifier.PermissionIdType, identifier.PermissionType, identifier.PermissionCategory, identifier.PermissionTenant, identifier.PermissionTenantId)
	if err != nil {
		return err
	}
	return nil
}

// FindPermission finds an identifier
func (access *PermissionDA) FindPermission(identifier *Permission, permissions *Permissions) (Permissions, error) {
	permissionContent, err := json.Marshal(*permissions)
	if err != nil {
		return nil, err
	}
	results, err := access.access.Query(
		`SELECT * FROM permission.identifier_find($1, $2, $3, $4, $5, $6, $7, $8, $9)`, mapPermission,
		identifier.Id, identifier.PermissionId, identifier.PermissionIdType, identifier.PermissionType, identifier.PermissionCategory, identifier.PermissionTenant, identifier.PermissionTenantId, identifier.DateAdded, permissionContent)
	if err != nil {
		return nil, err
	}
	tmp := make([]*Permission, 0)
	for r, _ := range results {
		if value, ok := results[r].(Permission); ok {
			tmp = append(tmp, &value)
		}
	}
	return tmp, nil
}

//DeletePermission deletes an identifier
func (access *PermissionDA) DeletePermission(identifier *Permission) (*Permission, error) {
	results, err := access.access.Query(
		`SELECT * FROM permission.identifier_remove($1)`, nil,
		identifier.Id)
	if err != nil {
		return nil, err
	}
	var tmp Permissions
	tmp = make([]*Permission, 0)
	for r, _ := range results {
		if value, ok := results[r].(Permission); ok {
			tmp = append(tmp, &value)
		}
	}
	return tmp.FindHead(), nil
}

//FindHead returns the first Permission
func (permissions Permissions) FindHead() *Permission {
	if len(permissions) == 0 {
		return nil
	}
	return permissions[0]
}
