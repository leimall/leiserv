package utils

var (
	WebIdVerify               = WebRules{"ID": []string{NotEmpty()}}
	WebApiVerify              = WebRules{"Path": {NotEmpty()}, "Description": {NotEmpty()}, "ApiGroup": {NotEmpty()}, "Method": {NotEmpty()}}
	WebMenuVerify             = WebRules{"Path": {NotEmpty()}, "Name": {NotEmpty()}, "Component": {NotEmpty()}, "Sort": {Ge("0")}}
	WebMenuMetaVerify         = WebRules{"Title": {NotEmpty()}}
	WebLoginVerify            = WebRules{"CaptchaId": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}
	WebRegisterVerify         = WebRules{"Email": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}
	WebPageInfoVerify         = WebRules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	WebCustomerVerify         = WebRules{"CustomerName": {NotEmpty()}, "CustomerPhoneData": {NotEmpty()}}
	WebAutoCodeVerify         = WebRules{"Abbreviation": {NotEmpty()}, "StructName": {NotEmpty()}, "PackageName": {NotEmpty()}, "Fields": {NotEmpty()}}
	WebAutoPackageVerify      = WebRules{"PackageName": {NotEmpty()}}
	WebAuthorityVerify        = WebRules{"AuthorityId": {NotEmpty()}, "AuthorityName": {NotEmpty()}}
	WebAuthorityIdVerify      = WebRules{"AuthorityId": {NotEmpty()}}
	WebOldAuthorityVerify     = WebRules{"OldAuthorityId": {NotEmpty()}}
	WebChangePasswordVerify   = WebRules{"Password": {NotEmpty()}, "NewPassword": {NotEmpty()}}
	WebSetUserAuthorityVerify = WebRules{"AuthorityId": {NotEmpty()}}
)
