package document

// Document refers the information of page viewing
type Document struct {
	Orientation  string `json:"orientation"`
	PageSize     string `json:"pageSize"`
	MarginBottom uint   `json:"marginBottom"`
	MarginTop    uint   `json:"marginTop"`
	MarginLeft   uint   `json:"marginLeft"`
	MarginRight  uint   `json:"marginRight"`
}
