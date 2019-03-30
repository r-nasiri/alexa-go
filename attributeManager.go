package alexa

type BasicAttributeManger struct {
	Attributes map[string]interface{}
}

func (bam BasicAttributeManger) GetSessionAttributes() map[string]interface{} {
	return bam.Attributes
}

func (bam BasicAttributeManger) SetSessionAttributes(attributes map[string]interface{}) {

	for k, v := range attributes {
		bam.Attributes[k] = v
	}

}
