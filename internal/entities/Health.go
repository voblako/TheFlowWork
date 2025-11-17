package entities

type Health struct {
	Status       string `json:"Status"`
	CurrentTime  string `json:"CurrentTime"`
	Uptime       string `json:"Uptime"`
	Version      string `json:"BackendVersion"`
	DatabaseInfo string `json:"DBName"`
}
