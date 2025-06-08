package config

type Config struct {
	Mysql   Mysql         `json:"mysql"`
	Redis   Redis         `json:"redis"`
	System  System        `json:"system"`
	WXRobot WXRobot       `json:"wxRobot"`
	Job     CronJobConfig `json:"job"`
	Qny     Qny           `json:"qny"`
}

type Qny struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	Bucket    string `json:"bucket"`
	QnyServer string `json:"qnyServer"`
}

type Redis struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
	DB       int    `json:"db"`
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
	Name      string `json:"name"`
	Env       string `json:"env"`
	Port      string `json:"port"`
	Migration bool   `json:"migration"`
}

type CronJobConfig struct {
	JobStatus map[string]bool   `json:"jobStatus"`
	JobCron   map[string]string `json:"jobCron"`
}
