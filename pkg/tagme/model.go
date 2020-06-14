package tagme

// AnnotationResponse is the response to an annotation request to TagMe
type AnnotationResponse struct {
	Time        int          `json:"time"`
	API         string       `json:"api"`
	Lang        string       `json:"lang"`
	Timestamp   string       `json:"timestamp"`
	Annotations []Annotation `json:"annotations"`
}

// Annotation ...
type Annotation struct {
	PageID          uint64  `json:"id"`
	Title           string  `json:"title"`
	Spot            string  `json:"spot"`
	Start           uint    `json:"start"`
	End             uint    `json:"end"`
	LinkProbability float64 `json:"link_probability"`
	Rho             float64 `json:"rho"`
}

func (ann *Annotation) GetPageID() uint64 {
	return ann.PageID
}

func (ann *Annotation) GetTitle() string {
	return ann.Title
}

func (ann *Annotation) GetSpot() string {
	return ann.Spot
}
func (ann *Annotation) GetStart() uint {
	return ann.Start
}

func (ann *Annotation) GetEnd() uint {
	return ann.End
}
func (ann *Annotation) GetLinkProbability() float64 {
	return ann.LinkProbability
}

func (ann *Annotation) GetRho() float64 {
	return ann.Rho
}

type Spot struct {
	Lp    float64 `json:"lp"`
	Spot  string  `json:"spot"`
	Start uint    `json:"start"`
	End   uint    `json:"end"`
}

type SpotResponse struct {
	Time      uint   `json:"time"`
	API       string `json:"api"`
	Lang      string `json:"lang"`
	Timestamp string `json:"timestamp"`
	Spots     []Spot `json:"spots"`
}
