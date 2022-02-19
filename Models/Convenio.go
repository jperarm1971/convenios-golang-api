package Model

type Convenio struct {
	ConvenioID         int    `json:"convenioid,omitempty"`
	ConvenioName       string `json:"convenioname"`
	ConvenioFechaFirma string `json:"conveniofecha"`
	ConvenioGestor     string `json:"conveniogestor,omitempty"`
	ConvenioOrganismo  string `json:"convenioorganismo"`
}

type JsonResponseConvenios struct {
	Type    string `json:"type"`
	Data    string `json:"data"`
	Message string `json:"message"`
}
