package src

type Friend struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Source string `json:"source"` //vk.com/id<...>
	Photo  string `json:"photo"`
	Sex    int    `json:"sex"` //1 - male, 2 - female
}
