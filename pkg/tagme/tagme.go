package tagme

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const defaultTagAPI = "https://tagme.d4science.org/tagme/tag"
const defaultSpotAPI = "https://tagme.d4science.org/tagme/spot"
const defaultRelAPI = "https://tagme.d4science.org/tagme/rel"

// Annotator
type Annotator struct {
	token string
}

// NewAnnotator creates a new TagmeAnnotator used to annotate text using TagMe
func NewAnnotator(token string) *Annotator {
	return &Annotator{
		token,
	}
}

// Annotate is used to annotate an input phrase
func (tAnn *Annotator) Annotate(text string) AnnotationResponse {
	requestParameters := map[string][]string{
		"gcube-token": {tAnn.token},
		"long_text":   {text},
		"text":        {text},
		"lang":        {"en"},
	}

	var decodedResponse AnnotationResponse
	apiRequest(defaultTagAPI, requestParameters, &decodedResponse)

	return decodedResponse
}

// Mentions is used to annotate an input phrase
func (tAnn *Annotator) Mentions(text string) SpotResponse {
	requestParameters := map[string][]string{
		"gcube-token": {tAnn.token},
		"text":        {text},
		"long_text":   {text},
		"lang":        {"en"},
	}

	var decodedResponse SpotResponse
	apiRequest(defaultSpotAPI, requestParameters, &decodedResponse)

	return decodedResponse
}

func apiRequest(API string, requestParameters map[string][]string, decodedResponse interface{}) {
	httpResponse, err := http.PostForm(API, requestParameters)
	if err != nil {
		fmt.Println("failed request")
	}

	responseBody, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		fmt.Println("Failed request")
	}

	defer httpResponse.Body.Close()

	err = json.Unmarshal(responseBody, &decodedResponse)
	if err != nil {
		fmt.Println("Failed request")
	}
}
