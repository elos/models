package tag

import "github.com/elos/models"

type ByName []*models.Tag

func (b ByName) Len() int           { return len(b) }
func (b ByName) Less(i, j int) bool { return b[i].Name < b[j].Name }
func (b ByName) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

type ByCreatedAt []*models.Tag

func (b ByCreatedAt) Len() int           { return len(b) }
func (b ByCreatedAt) Less(i, j int) bool { return b[i].CreatedAt.Before(b[j].CreatedAt) }
func (b ByCreatedAt) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
