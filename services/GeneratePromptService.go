package services

import (
	"fmt"
	"github.com/sashabaranov/go-openai"
	"log"
)

type PromptListProjectStruct struct {
	TypeProject  string `gorm:"not null" json:"type_project" valid:"required"`
	Approach     string `gorm:"not null" json:"approach" valid:"required"`
	Requirements string `gorm:"not null" json:"requirements" valid:"required"`
	Course       string `gorm:"not null" json:"course" valid:"required"`
	Technology   string `gorm:"not null" json:"technology" valid:"required"`
	Context      string `gorm:"not null" json:"context" valid:"required"`
}

func (p *PromptListProjectStruct) generateListProject() string {
	return fmt.Sprintf(`Realiza una lista de 10 items sobre temas de investigación para una %s que se adapten a estos temas para su 
		desarrollo: 
			tecnología: %s, 
			enfoque: %s, 
			exigencias: %s, 
			curso: %s
		para lograr esto: %s`,
		p.TypeProject,
		p.Technology,
		p.Approach,
		p.Requirements,
		p.Course,
		p.Context,
	)
}

func (p *PromptListProjectStruct) GetListOfPossibleProject(beforeMessage []openai.ChatCompletionMessage) (openai.ChatCompletionResponse, error) {
	prompt := p.generateListProject()
	ms, err := SendMessage(append(beforeMessage, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: prompt,
	}), 500)
	if err != nil {
		log.Fatal(err)
		return openai.ChatCompletionResponse{}, err
	}

	return ms, nil

}
