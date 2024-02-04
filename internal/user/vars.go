package user

import "time"

var (
	DefaultTimeSendEmails = time.Hour
	DefaultFromSendMail   = "marcellosanttos2014@gmail.com"

	DefaultServiceName         = "gmail"
	DefaultSubjectSendConfirm  = "Confirm signup"
	DefaulfBodySendConfirm     = "%s"
	DefaulfTemplateSendConfirm = ""

	DefaultSubjectResetPass  = "Reset Pass"
	DefaultBodyResetPass     = "%s"
	DefaultTemplateResetPass = ""
)
