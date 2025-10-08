package main

//Тип для событий
type Event struct {
	Listener interface{} `json:"listener"`
	UseCapture bool `json:"useCapture"`
	Passive bool `json:"passive"`
	Once bool `json:"once"`
	Type string `json:"type"`
}

func (e *Event) changeUseCapture() bool {
	if !e.UseCapture {
	 	e.UseCapture = true
	}
	return e.UseCapture
}

//Коллекция событий
type EventListeners map[string][]*Event
