package Hotspots

func NewHotspot(provider string, username string, password string) Hotspot {
	if provider == "Telekom" {
		return NewHotspotTelekom(username, password)
	}
	return nil
}
