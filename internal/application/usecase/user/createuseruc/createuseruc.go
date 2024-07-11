// Package createuseruc contains the use case to create the user.
package createuseruc

import (
	"github.com/cristiano-pacheco/pingo/internal/application/gateway/mailergw"
	"github.com/cristiano-pacheco/pingo/internal/application/gateway/mailertemplategw"
	"github.com/cristiano-pacheco/pingo/internal/application/repository/userrepo"
	"github.com/cristiano-pacheco/pingo/internal/domain/model/configdm"
)

type UseCase struct {
	userRepo         userrepo.UserRepository
	mailerGW         mailergw.MailerGateway
	mailerTemplateGW mailertemplategw.MailerTemplateGateway
	config           *configdm.Config
	mapper           Mapper
}

func New(
	userRepo userrepo.UserRepository,
	mailerGW mailergw.MailerGateway,
	mailerTemplateGW mailertemplategw.MailerTemplateGateway,
	config *configdm.Config,
	mapper Mapper,
) *UseCase {
	return &UseCase{
		userRepo:         userRepo,
		mailerGW:         mailerGW,
		mailerTemplateGW: mailerTemplateGW,
		config:           config,
		mapper:           mapper,
	}
}

func (uc *UseCase) Execute(input *Input) (*Output, error) {
	user, err := uc.mapper.mapInputToNewUser(input)
	if err != nil {
		return nil, err
	}

	err = uc.userRepo.Create(*user)
	if err != nil {
		return nil, err
	}

	tplVars := uc.mapper.mapAccountConfTemplVars(*user, uc.config.BaseURL.String())

	// template content
	content, err := uc.mailerTemplateGW.CompileTemplate("account_confirmation.gohtml", tplVars)
	if err != nil {
		return nil, err
	}

	mailerData := mailergw.MailData{
		ToName:  user.Name.String(),
		ToEmail: user.Email.String(),
		Subject: "Account Confirmation",
		Content: content,
	}

	err = uc.mailerGW.Send(&mailerData)
	if err != nil {
		return nil, err
	}

	output := uc.mapper.mapUserToOutput(*user)

	return output, nil
}
