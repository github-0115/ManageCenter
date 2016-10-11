package vars

var (
	ErrParameter                  = &Err{600, "Parameter err "}
	ErrParameterFormat            = &Err{601, "Parameter format err "}
	ErrParameterNil               = &Err{602, "Parameter is nil err "}
	ErrUsernameFormat             = &Err{603, "username format err "}
	ErrEmailFormat                = &Err{604, "email format err "}
	ErrUserInvoiceApplycount      = &Err{605, "user invoice Apply is exist"}
	ErrUserInvoiceApplyChange     = &Err{606, "Change invoice Apply status err"}
	ErrUserInvoiceApplyNotfount   = &Err{607, "user invoice Apply not find err"}
	ErrUserInvoiceApplyMoney      = &Err{608, "user invoice Apply money count err"}
	ErrNoticeEmailNotFound        = &Err{609, "notice Email not find  err"}
	ErrNoticeEmailNumNotFound     = &Err{610, "notice Email num not find  err"}
	ErrNoticeEmailExist           = &Err{611, "notice Email is exist  err"}
	ErrSubmitEmailNotFound        = &Err{612, "Submit Email not find  err"}
	ErrUserPakageApplyNotFound    = &Err{613, "user Pakage Apply Not Found"}
	ErrBindJson                   = &Err{614, "bind json err"}
	ErrUserNotFound               = &Err{700, "user not found"}
	ErrUserExist                  = &Err{701, "user is exist"}
	ErrUserTransferscount         = &Err{702, "user Transfers is exist"}
	ErrUserTransfersNotfount      = &Err{703, "user Transfers not find err"}
	ErrUserBillscount             = &Err{704, "user bill is exist"}
	ErrUserBillsNotfount          = &Err{705, "user bill not find err"}
	ErrUserBillsNotSettlement     = &Err{706, "user bill not Settlement err"}
	ErrKeyNotFound                = &Err{707, "user key not found"}
	ErrKeySave                    = &Err{708, "user key save err"}
	ErrBankCountFormat            = &Err{709, "BankCount Parameter format err "}
	ErrBankCountInvalid           = &Err{710, "BankCount Parameter invalid err "}
	ErrUserCursor                 = &Err{711, "Cursor err "}
	ErrUserGrant                  = &Err{713, "user grant update err"}
	ErrUserSave                   = &Err{714, "user info save err"}
	ErrPVcodeNotFound             = &Err{712, "Potential Vcode not found"}
	ErrPhoneNotFound              = &Err{715, "Potential phone not found"}
	ErrNeedToken                  = &Err{716, "need auth token"}
	ErrInvalidToken               = &Err{717, "invalid token"}
	ErrIncompleteToken            = &Err{718, "token incomplete"}
	ErrUserAdscount               = &Err{719, "user address is exist"}
	ErrUserAdsSave                = &Err{720, "save address err"}
	ErrUserInvoicecount           = &Err{721, "user invoice is exist"}
	ErrUserInvoiceSave            = &Err{722, "save invoice err"}
	ErrUserInvoiceNotfount        = &Err{723, "user invoice not find err"}
	ErrUserAdsNotfount            = &Err{724, "user address not find err"}
	ErrUserWalletNotfount         = &Err{725, "user wallet not find err"}
	ErrPornHourNotFount           = &Err{801, "result of pornHour not found"}
	ErrPornDayNotFount            = &Err{802, "result of pornday not found"}
	ErrPornMinuteNil              = &Err{803, "pornMinute is nil"}
	ErrPornMinuteNotFount         = &Err{804, "result of pornMinute not found"}
	ErrLiveReveiwNotFount         = &Err{805, "result of liveReview not found"}
	ErrSignPhone                  = &Err{813, "user phone exist"}
	ErrSignSave                   = &Err{814, "user info save err"}
	ErrSignEmail                  = &Err{815, "user email exist"}
	ErrSendSMS                    = &Err{816, "send verify sms failed, please try later."}
	ErrDetialNotFound             = &Err{817, "detail not found"}
	ErrDemoNotFound               = &Err{818, "demo not found"}
	ErrOther                      = &Err{400, "Other err"}
	ErrUserServicePackageNotFound = &Err{901, "user's service package not found"}
	ErrUserServicePackageCursor   = &Err{902, "user service package Cursor err"}
	ErrInterfaceConfigNotFound    = &Err{903, "interface config not found"}
	ErrInterfaceConfigCursor      = &Err{904, "interface config Cursor err"}
	ErrBillNotFound               = &Err{1001, "bill not found"}
	ErrMissRateNotFound           = &Err{1101, "miss_rate not found"}
)

type Err struct {
	Code int64
	Msg  string
}
