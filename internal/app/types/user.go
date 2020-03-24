package types

import "time"

type (
	UserStatus string
	User       struct {
		ID          string     `json:"id,omitempty" bson:"_id,omitempty"`
		Email       string     `json:"email,omitempty" bson:"email,omitempty"`
		Password    string     `json:"-" bson:"password,omitempty"`
		FirstName   string     `json:"first_name,omitempty" bson:"first_name,omitempty"`
		LastName    string     `json:"last_name,omitempty" bson:"last_name,omitempty"`
		Status      UserStatus `json:"status,omitempty" bson:"status,omitempty"`
		Roles       []string   `json:"roles,omitempty" bson:"roles,omitempty"`
		Groups      []string   `json:"groups,omitempty" bson:"groups,omitempty"`
		Provider    string     `json:"provider,omitempty" bson:"provider,omitempty"`
		Name        string     `json:"name,omitempty" bson:"name,omitempty"`
		NickName    string     `json:"nick_name,omitempty" bson:"nick_name,omitempty"`
		Description string     `json:"desciption,omitempty" bson:"description,omnitempty"`
		UserID      string     `json:"user_id,omitempty" bson:"user_id,omitempty"`
		AvatarURL   string     `json:"avatar_url,omitempty" bson:"avatar_url,omitempty"`
		Location    string     `json:"location,omitempty" bson:"location,omitempty"`
		CreateAt    *time.Time `json:"create_at,omitempty" bson:"create_at,omitempty"`
		UpdateAt    *time.Time `json:"update_at,omitempty" bson:"create_at,omitempty"`
	}

	RegisterRequest struct {
		Email     string `json:"email,omitempty" validate:"required,email"`
		Password  string `json:"password" validate:"required,gt=6"`
		FirstName string `json:"first_name,omitempty"`
		LastName  string `json:"last_name,omitempty"`
	}

	UserInfo struct {
		ID          string     `json:"id,omitempty"`
		Email       string     `json:"email,omitempty"`
		FirstName   string     `json:"first_name,omitempty"`
		LastName    string     `json:"last_name,omitempty"`
		Name        string     `json:"name,omitempty"`
		NickName    string     `json:"nick_name,omitempty"`
		Description string     `json:"desciption,omitempty"`
		UserID      string     `json:"user_id,omitempty"`
		AvatarURL   string     `json:"avatar_url,omitempty"`
		Location    string     `json:"location,omitempty"`
		CreateAt    *time.Time `json:"create_at,omitempty"`
		UpdateAt    *time.Time `json:"update_at,omitempty"`
	}
)
