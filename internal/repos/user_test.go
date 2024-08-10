package repos

import (
	"errors"
	"marketplace/internal/mocks"
	"marketplace/internal/models"
	"testing"
)

func TestUserRepo_Create(t *testing.T) {

	successMockDB := new(mocks.MockDB)
	successMockStmt := new(mocks.MockStmt)
	successMockDB.On("Prepare", "INSERT INTO tcp_auth_db.user_tab(username, password_hash) VALUES ($1, $2)").Return(successMockStmt, nil)
	successMockStmt.On("Exec", "testuser", "testhash").Return(nil, nil)

	failPrepareMockDB := new(mocks.MockDB)
	failPrepareMockDB.On("Prepare", "INSERT INTO tcp_auth_db.user_tab(username, password_hash) VALUES ($1, $2)").Return(nil, errors.New("prepare error"))

	failCreateMockDB := new(mocks.MockDB)
	failCreateMockStmt := new(mocks.MockStmt)
	failCreateMockDB.On("Prepare", "INSERT INTO tcp_auth_db.user_tab(username, password_hash) VALUES ($1, $2)").Return(failCreateMockStmt, nil)
	failCreateMockStmt.On("Exec", "testuser", "testhash").Return(nil, errors.New("exec error"))

	type fields struct {
		db models.DB
	}
	type args struct {
		user models.UserAttributes
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				db: successMockDB,
			},
			args: args{
				user: models.UserAttributes{
					Username:     "testuser",
					PasswordHash: "testhash",
				},
			},
			wantErr: false,
		},
		{
			name: "FailPrepare",
			fields: fields{
				db: failPrepareMockDB,
			},
			args: args{
				user: models.UserAttributes{
					Username:     "testuser",
					PasswordHash: "testhash",
				},
			},
			wantErr: true,
		},
		{
			name: "FailCreate",
			fields: fields{
				db: failCreateMockDB,
			},
			args: args{
				user: models.UserAttributes{
					Username:     "testuser",
					PasswordHash: "testhash",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserRepo{
				db: tt.fields.db,
			}
			if err := u.Create(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
