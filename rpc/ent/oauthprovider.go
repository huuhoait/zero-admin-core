// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/suyuan32/simple-admin-core/rpc/ent/oauthprovider"
)

// OauthProvider is the model entity for the OauthProvider schema.
type OauthProvider struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// the provider's name | 提供商名称
	Name string `json:"name,omitempty"`
	// the client id | 客户端 id
	ClientID string `json:"client_id,omitempty"`
	// the client secret | 客户端密钥
	ClientSecret string `json:"client_secret,omitempty"`
	// the redirect url | 跳转地址
	RedirectURL string `json:"redirect_url,omitempty"`
	// the scopes | 权限范围
	Scopes string `json:"scopes,omitempty"`
	// the auth url of the provider | 认证地址
	AuthURL string `json:"auth_url,omitempty"`
	// the token url of the provider | 获取 token地址
	TokenURL string `json:"token_url,omitempty"`
	// the auth style, 0: auto detect 1: third party log in 2: log in with username and password
	AuthStyle uint64 `json:"auth_style,omitempty"`
	// the URL to request user information by token | 用户信息请求地址
	InfoURL string `json:"info_url,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OauthProvider) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case oauthprovider.FieldID, oauthprovider.FieldAuthStyle:
			values[i] = new(sql.NullInt64)
		case oauthprovider.FieldName, oauthprovider.FieldClientID, oauthprovider.FieldClientSecret, oauthprovider.FieldRedirectURL, oauthprovider.FieldScopes, oauthprovider.FieldAuthURL, oauthprovider.FieldTokenURL, oauthprovider.FieldInfoURL:
			values[i] = new(sql.NullString)
		case oauthprovider.FieldCreatedAt, oauthprovider.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OauthProvider", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OauthProvider fields.
func (op *OauthProvider) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case oauthprovider.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			op.ID = uint64(value.Int64)
		case oauthprovider.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				op.CreatedAt = value.Time
			}
		case oauthprovider.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				op.UpdatedAt = value.Time
			}
		case oauthprovider.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				op.Name = value.String
			}
		case oauthprovider.FieldClientID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_id", values[i])
			} else if value.Valid {
				op.ClientID = value.String
			}
		case oauthprovider.FieldClientSecret:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field client_secret", values[i])
			} else if value.Valid {
				op.ClientSecret = value.String
			}
		case oauthprovider.FieldRedirectURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field redirect_url", values[i])
			} else if value.Valid {
				op.RedirectURL = value.String
			}
		case oauthprovider.FieldScopes:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field scopes", values[i])
			} else if value.Valid {
				op.Scopes = value.String
			}
		case oauthprovider.FieldAuthURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field auth_url", values[i])
			} else if value.Valid {
				op.AuthURL = value.String
			}
		case oauthprovider.FieldTokenURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token_url", values[i])
			} else if value.Valid {
				op.TokenURL = value.String
			}
		case oauthprovider.FieldAuthStyle:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field auth_style", values[i])
			} else if value.Valid {
				op.AuthStyle = uint64(value.Int64)
			}
		case oauthprovider.FieldInfoURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field info_url", values[i])
			} else if value.Valid {
				op.InfoURL = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this OauthProvider.
// Note that you need to call OauthProvider.Unwrap() before calling this method if this OauthProvider
// was returned from a transaction, and the transaction was committed or rolled back.
func (op *OauthProvider) Update() *OauthProviderUpdateOne {
	return NewOauthProviderClient(op.config).UpdateOne(op)
}

// Unwrap unwraps the OauthProvider entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (op *OauthProvider) Unwrap() *OauthProvider {
	_tx, ok := op.config.driver.(*txDriver)
	if !ok {
		panic("ent: OauthProvider is not a transactional entity")
	}
	op.config.driver = _tx.drv
	return op
}

// String implements the fmt.Stringer.
func (op *OauthProvider) String() string {
	var builder strings.Builder
	builder.WriteString("OauthProvider(")
	builder.WriteString(fmt.Sprintf("id=%v, ", op.ID))
	builder.WriteString("created_at=")
	builder.WriteString(op.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(op.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(op.Name)
	builder.WriteString(", ")
	builder.WriteString("client_id=")
	builder.WriteString(op.ClientID)
	builder.WriteString(", ")
	builder.WriteString("client_secret=")
	builder.WriteString(op.ClientSecret)
	builder.WriteString(", ")
	builder.WriteString("redirect_url=")
	builder.WriteString(op.RedirectURL)
	builder.WriteString(", ")
	builder.WriteString("scopes=")
	builder.WriteString(op.Scopes)
	builder.WriteString(", ")
	builder.WriteString("auth_url=")
	builder.WriteString(op.AuthURL)
	builder.WriteString(", ")
	builder.WriteString("token_url=")
	builder.WriteString(op.TokenURL)
	builder.WriteString(", ")
	builder.WriteString("auth_style=")
	builder.WriteString(fmt.Sprintf("%v", op.AuthStyle))
	builder.WriteString(", ")
	builder.WriteString("info_url=")
	builder.WriteString(op.InfoURL)
	builder.WriteByte(')')
	return builder.String()
}

// OauthProviders is a parsable slice of OauthProvider.
type OauthProviders []*OauthProvider
