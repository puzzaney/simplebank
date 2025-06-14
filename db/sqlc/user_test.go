package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/puzzaney/simplebank/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		Email:          util.RandomEmail(),
		FullName:       util.RandomOwner(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.FullName, user.FullName)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := CreateRandomUser(t)

	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)

}

func TestUpdateUserFullName(t *testing.T) {
	user := CreateRandomUser(t)

	newName := util.RandomOwner()

	arg := UpdateUserParams{
		Username: user.Username,
		FullName: sql.NullString{
			String: newName,
			Valid:  true,
		},
	}

	updatedUser, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedUser)
	require.NotEqual(t, user.FullName, updatedUser.FullName)
	require.Equal(t, updatedUser.FullName, newName)
	require.Equal(t, user.Email, updatedUser.Email)
	require.Equal(t, user.HashedPassword, updatedUser.HashedPassword)

}

func TestUpdateUserEmail(t *testing.T) {
	user := CreateRandomUser(t)

	newEmail := util.RandomEmail()

	arg := UpdateUserParams{
		Username: user.Username,
		Email: sql.NullString{
			String: newEmail,
			Valid:  true,
		},
	}

	updatedUser, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedUser)
	require.NotEqual(t, user.Email, updatedUser.Email)
	require.Equal(t, updatedUser.Email, newEmail)
	require.Equal(t, user.FullName, updatedUser.FullName)
	require.Equal(t, user.HashedPassword, updatedUser.HashedPassword)
}

func TestUpdateUserPassword(t *testing.T) {
	user := CreateRandomUser(t)

	newPassword := util.RandomString(8)
	hashedPassword, err := util.HashPassword(newPassword)
	require.NoError(t, err)

	arg := UpdateUserParams{
		Username: user.Username,
		HashedPassword: sql.NullString{
			String: hashedPassword,
			Valid:  true,
		},
	}

	updatedUser, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedUser)
	require.NotEqual(t, user.HashedPassword, updatedUser.HashedPassword)
	require.Equal(t, updatedUser.HashedPassword, hashedPassword)
	require.Equal(t, user.FullName, updatedUser.FullName)
	require.Equal(t, user.Email, updatedUser.Email)

}

func TestUpdateUserAllFields(t *testing.T) {
	user := CreateRandomUser(t)

	newPassword := util.RandomString(8)
	newEmail := util.RandomEmail()
	newFullName := util.RandomOwner()
	hashedPassword, err := util.HashPassword(newPassword)
	require.NoError(t, err)

	arg := UpdateUserParams{
		Username: user.Username,
		HashedPassword: sql.NullString{
			String: hashedPassword,
			Valid:  true,
		},
		FullName: sql.NullString{
			String: newFullName,
			Valid:  true,
		},
		Email: sql.NullString{
			String: newEmail,
			Valid:  true,
		},
	}

	updatedUser, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedUser)
	require.NotEqual(t, user.HashedPassword, updatedUser.HashedPassword)
	require.NotEqual(t, user.Email, updatedUser.Email)
	require.NotEqual(t, user.FullName, updatedUser.FullName)
	require.Equal(t, updatedUser.HashedPassword, hashedPassword)
	require.Equal(t, updatedUser.FullName, newFullName)
	require.Equal(t, updatedUser.Email, newEmail)

}
