package service

type Mapper interface {
MapEvent(input map[string]interface{}, output map[string]interface{}, mapFunc func(input map[string]interface{}, output map[string]interface{}) error) error
}

type mapper struct {
}

func NewMapper() Mapper {

	return &mapper{}
}

func (m mapper) MapEvent(input map[string]interface{}, output map[string]interface{}, mapFunc func(input map[string]interface{}, output map[string]interface{}) error) error {

	err := mapFunc(input, output)

	if err != nil {
		return err
	}
	return nil
}