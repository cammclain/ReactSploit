package models

// Config represents configuration settings that might be editable through the GUI.
type Config struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}
