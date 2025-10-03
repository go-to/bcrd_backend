package output

import "github.com/go-to/bcrd_protobuf/pb"

type ShopsTotalOutput struct {
	ShopsTotalResponse pb.ShopsTotalResponse
}

type ShopsOutput struct {
	ShopsResponse pb.ShopsResponse
}

type ShopOutput struct {
	ShopResponse  pb.ShopResponse
	IsEventPeriod bool
}
