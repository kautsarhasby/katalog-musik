package memberships

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kautsarhasby/katalog-musik/internal/models/memberships"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_repository_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	assert.NoError(t, err)

	type args struct {
		model memberships.User
	}

	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		db *gorm.DB
		// Named input parameters for target function.
		model   memberships.User
		wantErr bool
		mockFn  func(args args)
		args    args
	}{
		// TODO: Add test cases.
		{
			name: "Success",

			model: memberships.User{
				Email:     "test@gmail.com",
				Username:  "testusername",
				Password:  "password",
				CreatedBy: "test@gmail.com",
				UpdateBy:  "test@gmail.com",
			},
			args: args{
				model: memberships.User{
					Email:     "test@gmail.com",
					Username:  "testusername",
					Password:  "password",
					CreatedBy: "test@gmail.com",
					UpdateBy:  "test@gmail.com",
				},
			},

			db:      gormDB,
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(`INSERT INTO "users" (.+) VALUES (.+)`).
					WithArgs(
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						args.model.Email,
						args.model.Username,
						args.model.Password,
						args.model.CreatedBy,
						args.model.UpdateBy,
					).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mock.ExpectCommit()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := NewRepository(tt.db)
			gotErr := r.CreateUser(tt.model)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("CreateUser() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("CreateUser() succeeded unexpectedly")
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func Test_repository_GetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	assert.NoError(t, err)

	now := time.Now()
	type args struct {
		email    string
		username string
		id       uint
	}
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		db *gorm.DB
		// Named input parameters for target function.
		model   memberships.User
		want    *memberships.User
		wantErr bool
		mockFn  func(args args)
		args    args
	}{
		{ // TODO: Add test cases.
			name: "Success",
			model: memberships.User{
				Email:     "test@gmail.com",
				Username:  "testusername",
				Password:  "password",
				CreatedBy: "test@gmail.com",
				UpdateBy:  "test@gmail.com",
			},
			db: gormDB,
			args: args{
				email:    "test@gmail.com",
				username: "testusername",
				id:       1,
			},
			want: &memberships.User{
				Model:     gorm.Model{ID: 1, CreatedAt: now, UpdatedAt: now},
				Email:     "test@gmail.com",
				Username:  "testusername",
				Password:  "test",
				CreatedBy: "test@gmail.com",
				UpdateBy:  "test@gmail.com",
			},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectQuery(`SELECT \* FROM "users" .+`).WithArgs(args.email, args.username, args.id, 1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "email", "username", "password",
						"created_by", "updated_by"}).AddRow(1, now, now, "test@gmail.com", "testusername", "test", "test@gmail.com", "test@gmail.com"))

			},
		},
		{ // TODO: Add test cases.
			name: "Failed",
			model: memberships.User{
				Email:     "test@gmail.com",
				Username:  "testusername",
				Password:  "password",
				CreatedBy: "test@gmail.com",
				UpdateBy:  "test@gmail.com",
			},
			db: gormDB,
			args: args{
				email:    "test@gmail.com",
				username: "testusername",
				id:       1,
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectQuery(`SELECT \* FROM "users" .+`).WithArgs(args.email, args.username, args.id, 1).
					WillReturnError(assert.AnError)

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := NewRepository(tt.db)
			_, gotErr := r.GetUser(tt.args.email, tt.args.username, tt.args.id)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetUser() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetUser() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
