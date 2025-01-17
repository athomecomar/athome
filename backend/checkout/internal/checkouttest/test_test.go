package checkouttest

var (
	gPurchases    GoldenPurchases
	gShippings    GoldenShippings
	gPbProducts   GoldenPbProducts
	gStateChanges GoldenStateChanges
	gPbUsers      GoldenPbUsers
	gPbAddresses  GoldenPbAddresses
	gPayments     GoldenPayments
	gCards        GoldenCards
	gPbEvents     GoldenPbEvents
	gPbServices   GoldenPbServices
)

func init() {
	Init(&gPurchases, &gShippings, &gPbProducts, &gStateChanges, &gPbUsers, &gPbAddresses, &gCards, &gPayments, &gPbEvents, &gPbServices)
}
