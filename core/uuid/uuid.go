package uuid

import uuid "github.com/nu7hatch/gouuid"

type Generator struct {
}

func New() *Generator {
	return &Generator{}
}

func (g *Generator) GenerateUUID() string {
	generatedUuid, _ := uuid.NewV4()
	return generatedUuid.String()
}
