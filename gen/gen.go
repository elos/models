package main

import (
	"log"

	"github.com/elos/metis"
	"github.com/elos/metis/ego"
)

func main() {
	models, _ := metis.ParseGlob("./definitions/models/*json")
	s := metis.BuildSchema(models...)

	if err := s.Valid(); err != nil {
		panic(err)
	} else {
		log.Print("ALL GOOD")
	}

	for _, m := range models {
		ego.MakeGo(m).WriteFile("../" + m.Kind + ".go")
	}

	ego.WriteKindsFile(s, "../kinds.go")
	ego.WriteDynamicFile(s, "../dynamic.go")
	ego.WriteDBsFile(s, "../dbs.go")

	/*
		for _, m := range models {
			textpath := filepath.Join("models/docs", m.Kind+".md")
			doc.MakeDoc(m, textpath).WriteFile("docs/" + m.Kind + ".md")
		}
	*/
}
