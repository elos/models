package main

import (
	"log"

	"github.com/elos/gen/metis"
	"github.com/elos/gen/metis/ego"
)

func main() {
	models, _ := metis.ParseGlob("models/*json")
	s := metis.BuildSchema(models...)

	if err := s.Valid(); err != nil {
		panic(err)
	} else {
		log.Print("ALL GOOD")
	}

	for _, m := range models {
		ego.MakeGo(m).WriteFile("../models/" + m.Kind + ".go")
	}

	ego.WriteKindsFile(s, "../models/kinds.go")
	ego.WriteDynamicFile(s, "../models/dynamic.go")
	ego.WriteDBsFile(s, "../models/dbs.go")

	/*
		for _, m := range models {
			textpath := filepath.Join("models/docs", m.Kind+".md")
			doc.MakeDoc(m, textpath).WriteFile("docs/" + m.Kind + ".md")
		}
	*/
}
