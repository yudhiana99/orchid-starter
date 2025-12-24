package constants

import "github.com/mataharibiz/ward/logging"

// app origin buyer
const AppVladmir string = "vladmir"
const AppVembrace string = "vembrace"

// app origin seller
const AppVoodoo string = "voodoo"
const AppTrident string = "trident"

// app system
const AppSystem string = "system"
const AppTarrasque string = "tarrasque"
const AppWorker string = "worker"

// login as label
const BuyerLabel string = "buyer"
const SellerLabel string = "seller"
const SystemLabel string = "system"
const AdminLabel string = "admin"
const UnknownLabel string = "unknown"

type OriginID int8

const (
	Unknown  OriginID = 0
	BuyerID  OriginID = 1
	SellerID OriginID = 2
	SystemID OriginID = 3
	AdminID  OriginID = 4
)

func IsBuyer(appOrigin string) bool {
	clientList := map[string]bool{
		AppVladmir:  true,
		AppVembrace: true,
	}
	return clientList[appOrigin]
}

func IsSeller(appOrigin string) bool {
	clientList := map[string]bool{
		AppVoodoo:  true,
		AppTrident: true,
	}
	return clientList[appOrigin]
}

func IsNonUser(appOrigin string) bool {
	clientList := map[string]bool{
		AppSystem:    true,
		AppTarrasque: true,
		AppWorker:    true,
	}
	return clientList[appOrigin]
}

func LoginAs(appOrigin string) (loginAs OriginID) {
	switch appOrigin {
	case AppVladmir, AppVembrace:
		loginAs = BuyerID
	case AppTrident, AppVoodoo:
		loginAs = SellerID
	case AppTarrasque:
		loginAs = AdminID
	case AppSystem:
		loginAs = SystemID
	default:
		logging.NewLogger().Warn("unknown app origin", "app-origin", appOrigin)
		loginAs = Unknown
	}
	return loginAs
}

func LoginAsLabel(id OriginID) string {
	switch id {
	case BuyerID:
		return BuyerLabel
	case SellerID:
		return SellerLabel
	case AdminID:
		return AdminLabel
	case SystemID:
		return SystemLabel
	default:
		return UnknownLabel
	}
}
