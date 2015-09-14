package main

import (
	"log"

	"github.com/elos/metis"
	"github.com/elos/metis/builtin/ego"
)

func main() {
	models, err := metis.ParseGlob("./definitions/models/*/*json")
	if err != nil {
		log.Fatalf("Error parsing the models: %s", err.Error())
	}

	s := metis.BuildSchema(models...)

	if err := s.Valid(); err != nil {
		log.Fatalf("Schema Invalid: %s", err)
	} else {
		log.Print("Schema Good")
	}

	for _, m := range models {
		ego.MakeGo(m).WriteFile("../" + m.Kind + ".go")
	}

	ego.WriteKindsFile(s, "../kinds.go")
	ego.WriteDynamicFile(s, "../dynamic.go")
	ego.WriteDBsFile(s, "../dbs.go")

	// Generate the documentation files
	/*
		for _, m := range models {
			textpath := filepath.Join("models/docs", m.Kind+".md")
			doc.MakeDoc(m, textpath).WriteFile("docs/" + m.Kind + ".md")
		}
	*/
}
