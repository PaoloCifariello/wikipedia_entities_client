package main

import (
	"fmt"
	"github.com/PaoloCifariello/wikipedia_entities_client/pkg/model"
	"github.com/PaoloCifariello/wikipedia_entities_client/pkg/tagme"
)

func main() {
	annotator := tagme.NewAnnotator("")
	text := "Maradona scored on the last match of the season"

	for _, annotation := range annotator.Annotate(text).Annotations {
		printEntity(&annotation)
	}
}

func printEntity(ann model.WikipediaAnnotation) {
	fmt.Printf("%d -> %s found in position (%d, %d)\n", ann.GetPageID(), ann.GetTitle(), ann.GetStart(), ann.GetEnd())

}
