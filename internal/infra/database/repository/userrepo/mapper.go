package userrepo

import "github.com/cristiano-pacheco/pingo/internal/domain/model/userdm"

func mapUserDBToUser(userdb *UserDB) (*userdm.User, error) {
	user, err := userdm.RestoreUser(
		userdb.ID,
		userdb.Name,
		userdb.Email,
		userdb.Status,
		userdb.PasswordHash,
		userdb.AccountConfirmationToken,
		userdb.ResetPasswordToken,
		userdb.CreatedAT,
		userdb.UpdatedAT,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
