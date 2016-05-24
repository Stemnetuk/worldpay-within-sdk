package hte
import "errors"

type Credential struct {

	MerchantClientKey string
	MerchantServiceKey string
}

func NewHTECredential(MerchantClientKey, MerchantServiceKey string) (Credential, error) {

	if(MerchantClientKey == "") {

		return Credential{}, errors.New("MerchantClientKey required, cannot be empty")

	} else if(MerchantServiceKey == "") {

		return Credential{}, errors.New("MerchantServiceKey required, cannot be empty")
	}

	result := Credential{
		MerchantClientKey:MerchantClientKey,
		MerchantServiceKey:MerchantServiceKey,
	}

	return result, nil
}