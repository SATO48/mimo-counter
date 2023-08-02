package lib

import (
	"fmt"
	"strconv"

	"github.com/dghubble/sling"
)

type Mimo struct {
	s sling.Sling
}

func NewMimo(host string, port int, document int) *Mimo {
	return &Mimo{
		s: *sling.New().Base(fmt.Sprintf("http://%s:%d/api/v1/documents/%d/", host, port, document)),
	}
}

func (m *Mimo) GetLayer(id string) (*Layer, error) {
	r := new(APIResponse[Layer])
	if _, err := m.s.New().Get(fmt.Sprintf("layers/%s", id)).ReceiveSuccess(&r); err != nil {
		return nil, err
	}

	return &r.Data, nil
}

func (m *Mimo) GetNumericText(id string) (int, error) {
	l, err := m.GetLayer(id)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(l.Attributes.InputValues.Text)
}

func (m *Mimo) SetNumericText(id string, nt int) error {
	nla := new(LayerAttributes)
	nla.InputValues.Text = strconv.Itoa(nt)

	nl := new(APIResponse[Layer])
	_, err := m.s.New().Put(fmt.Sprintf("layers/%s", id)).BodyJSON(nla).ReceiveSuccess(&nl)
	return err
}

func (m *Mimo) IncrementNumericText(id string) (int, error) {
	pnum, err := m.GetNumericText(id)
	if err != nil {
		return 0, err
	}

	pnum++

	if err := m.SetNumericText(id, pnum); err != nil {
		return 0, err
	}

	return pnum, nil
}

func (m *Mimo) DecrementNumericText(id string) (int, error) {
	pnum, err := m.GetNumericText(id)
	if err != nil {
		return 0, err
	}

	pnum--

	if err := m.SetNumericText(id, pnum); err != nil {
		return 0, err
	}

	return pnum, nil
}
