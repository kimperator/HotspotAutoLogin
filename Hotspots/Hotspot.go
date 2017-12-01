package Hotspots

import (
	"HotspotAutoLogin/Base"
)

type Hotspot interface {
	Login() bool
	Logout() bool
	CanHandle(Base.WifiEnv) bool
	NeedsLogin() bool
}
