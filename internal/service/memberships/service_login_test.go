package memberships

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kautsarhasby/katalog-musik/internal/configs"
	"github.com/kautsarhasby/katalog-musik/internal/models/memberships"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_service_Login(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := NewMockrepository(ctrlMock)

	type args struct {
		request memberships.LoginRequest
	}

	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		cfg        *configs.Config
		repository repository
		// Named input parameters for target function.
		request memberships.LoginRequest
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		// TODO: Add test cases.

		{
			name: "Success",
			args: args{
				request: memberships.LoginRequest{
					Email:    "test@gmail.com",
					Password: "alamak123",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				fmt.Println(args.request.Email)
				mockRepo.EXPECT().GetUser(args.request.Email, "", uint(0)).Return(&memberships.User{
					Model: gorm.Model{
						ID: 1,
					},
					Email:    "test@gmail.com",
					Password: "$2a$10$JMHM5hcZzHxNUpf9.vunPe1EghJQcfv2fMCoGrEbqeaiIzag0yXXa",
					Username: "hasbydf",
				}, nil)

			},
		},
		{
			name: "Failed",
			args: args{
				request: memberships.LoginRequest{
					Email:    "test@gmail.com",
					Password: "wwrongg",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				fmt.Println(args.request.Email)
				mockRepo.EXPECT().GetUser(args.request.Email, "", uint(0)).Return(&memberships.User{
					Model: gorm.Model{
						ID: 1,
					},
					Email:    "test@gmail.com",
					Password: "$2a$10$JMHM5hcZzHxNUpf9.vunPe1EghJQcfv2fMCoGrEbqeaiIzag0yXXa",
					Username: "hasbydf",
				}, nil)

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := NewService(&configs.Config{Service: configs.Service{SecretKey: "abc"}}, mockRepo)
			got, gotErr := s.Login(tt.args.request)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Login() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Login() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if !tt.wantErr {
				assert.NotEmpty(t, got)
			} else {
				assert.Empty(t, got)
			}
		})
	}
}
