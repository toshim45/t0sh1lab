package test

import (
	"errors"
	//	"net/http"
	"testing"
)

// di tempat lain non test
type PCatClient interface {
	CountProductCatalogue() (int, error)
}

type PCatClientService struct {
	PCatClient
}

func (pcat *PCatClientService) CountProductCatalogue() (int, error) {
	//	http.Client.Get("http://localhost:8081/pcat/count").Body.Read()
	//  gw blom ngoprek http.Client
	return 12, nil
}

type ProductService struct {
	pCat PCatClient
}

func NewProductService(pcat PCatClient) *ProductService {
	return &ProductService{pCat: pcat}
}

func (product *ProductService) calculate() int {
	count, err := product.pCat.CountProductCatalogue() // harus di mock biar ga keluar
	if err != nil {
		return 0
	}
	return 10 + count
}

// di folder test
// mocking
type MockPcatClientService struct {
	PCatClient
	val int
}

func NewMockPcatClientService(number int) *MockPcatClientService {
	return &MockPcatClientService{val: number}
}

func (mPCat *MockPcatClientService) CountProductCatalogue() (int, error) {
	if mPCat.val == 0 {
		return 0, errors.New("network time out")
	}
	return mPCat.val, nil
}

// test case
func TestNormal(t *testing.T) {
	num := 15
	service := NewProductService(&MockPcatClientService{val: num})
	result := service.calculate()
	if result != 25 {
		t.Errorf("should return 25 since it's %d + 10 ", num)
	}
}

func TestTimeOut(t *testing.T) {
	num := 0
	service := NewProductService(&MockPcatClientService{val: num})
	result := service.calculate()
	if result != 0 {
		t.Error("should return 0 since it's network time out")
	}
}
