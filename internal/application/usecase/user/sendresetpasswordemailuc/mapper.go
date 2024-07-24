package sendresetpasswordemailuc

import (
	"encoding/base64"
	"fmt"

	"github.com/cristiano-pacheco/pingo/internal/domain/model/userdm"
)

func mapResetPasswordTemplVars(user userdm.User, baseURL string) *resetPasswordTemplateVars {
	resetPasswordToken := base64.StdEncoding.EncodeToString(user.ResetPasswordToken)
	resetPasswordLink := fmt.Sprintf(
		"%s/user/reset-password?id=%s&token=%s",
		baseURL,
		user.ID.String(),
		resetPasswordToken,
	)

	return &resetPasswordTemplateVars{
		Name:              user.Name.String(),
		ResetPasswordLink: resetPasswordLink,
	}
}
