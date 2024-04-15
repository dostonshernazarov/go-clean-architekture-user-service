package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github/go-clean-template/internal/entity"
	"github/go-clean-template/internal/usecase"
)

var (
	errInternalServErr = errors.New("internal server error")
	errBadRequest = errors.New("bad request")
)

type test struct {
	name string
	mock func()
	res  interface{}
	err  error
}

func user(t *testing.T) (*usecase.UserUseCase, *MockUserRepo) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := NewMockUserRepo(mockCtl)

	user := usecase.New(repo)

	return user, repo
}

func TestUser(t *testing.T) {
	t.Parallel()

	user, repo := user(t)

	body := entity.User{
		Id:       "",
		FullName: "Doston Shernazarov",
		Username: "doston",
		Email:    "dostonshernazarov2001@gmail.com",
		Password: "1234abcd",
		Bio:      "Life is an unpredictable",
		Website:  "test.com",
	}

	createTest := test{
		name: "create with error",
		mock: func() {
			repo.EXPECT().Create(context.Background(), &body).Return(nil, errBadRequest)
		},
		res: (*entity.User)(nil),
		err: errBadRequest,
	}

	t.Run(createTest.name, func(t *testing.T) {
		t.Parallel()

		createTest.mock()

		res, err := user.CreateUser(context.Background(), &body)

		require.Equal(t, createTest.res, res)
		require.ErrorIs(t, err, createTest.err)
	})


}
