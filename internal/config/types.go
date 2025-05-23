package config

type Config struct {
	Mysql   Mysql   `json:"mysql"`
	System  System  `json:"system"`
	WXRobot WXRobot `json:"wxRobot"`
}

type WXRobot struct {
	ErrorRobot string `json:"errorRobot"`
}

type Mysql struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type System struct {
	Env       string `json:"env"`
	Port      string `json:"port"`
	Migration bool   `json:"migration"`
}
