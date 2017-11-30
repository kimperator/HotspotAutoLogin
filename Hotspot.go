package main

type Hotspot interface {
	Login() bool
	Logout() bool
	CanHandle(HotspotEnv) bool
	NeedsLogin() bool
}
