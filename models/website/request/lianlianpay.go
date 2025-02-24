package request

//	export type llpayCardType = {
//		cardToken: string;
//		cardName: string;
//		token: string;
//		method: string;
//		orderId: string;
//	  };
type CreateLianLianPayType struct {
	CardToken string `json:"cardToken"`
	CardName  string `json:"cardName"`
	Token     string `json:"token"`
	Method    string `json:"method"`
	OrderId   string `json:"orderId"`
}
