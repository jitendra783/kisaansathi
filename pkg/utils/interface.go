package utils

/* Place all utils interfaces here
   so that there will be on mock up
   for all utils interfaces
*/
//go:generate mockgen -destination=mock/mock.go -package=mock equity-trading/pkg/utils  Utils
type Utils interface {
	RestCaller
	AesCipherGroup
}
