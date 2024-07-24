// Package sendresetpasswordemailuc contains the use case to reset the user password.
package sendresetpasswordemailuc

import (
	"github.com/cristiano-pacheco/pingo/internal/application/gateway/mailergw"
	"github.com/cristiano-pacheco/pingo/internal/application/gateway/mailertemplategw"
	"github.com/cristiano-pacheco/pingo/internal/application/repository/userrepo"
	"github.com/cristiano-pacheco/pingo/internal/domain/model/configdm"
	"github.com/cristiano-pacheco/pingo/internal/domain/model/userdm"
	"github.com/cristiano-pacheco/pingo/internal/domain/service/hashds"
)

type UseCase struct {
	userRepo         userrepo.UserRepository
	mailerGW         mailergw.MailerGateway
	mailerTemplateGW mailertemplategw.MailerTemplateGateway
	hashService      *hashds.Service
	config           *configdm.Config
}

func New(
	userRepo userrepo.UserRepository,
	mailerGW mailergw.MailerGateway,
	mailerTemplateGW mailertemplategw.MailerTemplateGateway,
	hs *hashds.Service,
	config *configdm.Config,
) *UseCase {
	return &UseCase{
		userRepo:         userRepo,
		mailerGW:         mailerGW,
		mailerTemplateGW: mailerTemplateGW,
		hashService:      hs,
		config:           config,
	}
}

func (uc *UseCase) Execute(input *Input) error {
	// Search the user by email
	email, err := userdm.NewEmail(input.Email)
	if err != nil {
		return err
	}

	user, err := uc.userRepo.FindByEmail(*email)
	if err != nil {
		return err
	}

	// Generate the reset password token
	resetPasswordToken, err := uc.hashService.GenerateRandomBytes()
	if err != nil {
		return err
	}

	// Persists the user reset password token in the database
	user.ResetPasswordToken = resetPasswordToken
	uc.userRepo.UpdateResetPasswordToken(*user)

	// Compile the email template
	tplVars := mapResetPasswordTemplVars(*user, uc.config.FrontEndBaseURL.String())
	content, err := uc.mailerTemplateGW.CompileTemplate("reset_password.gohtml", tplVars)
	if err != nil {
		return err
	}

	// Send the reset password email
	mailerData := mailergw.MailData{
		ToName:  user.Name.String(),
		ToEmail: user.Email.String(),
		Subject: "Reset Password",
		Content: content,
	}

	err = uc.mailerGW.Send(&mailerData)
	if err != nil {
		return err
	}

	return nil
}
