package main

// Configuration struct
type Configuration struct {
	Debug          bool          `json:"debug"`
	Port	       int           `json:"port"`
	Db             DataSource    `json:"db"`
	Filesystem     bool          `json:"filesystem"`
}

// DataSource struct
type DataSource struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}