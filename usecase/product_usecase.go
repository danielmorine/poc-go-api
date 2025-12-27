package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(r repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: r,
	}
}

func (pu *ProductUseCase) GetProducts() (model.Response, error) {
	products, err := pu.repository.GetProducts()
	if err != nil {
		return model.Response{Status: false, Message: "Falha ao tentar obter produtos"}, err
	}
	return model.Response{Status: true, Data: products}, nil
}

func (pu *ProductUseCase) DeleteProduct(id int) (model.Response, error) {
	exists, err := pu.repository.ProductExists(id)
	if err != nil {
		return model.Response{Status: false, Message: "Falha ao buscar um produto", Data: nil}, err
	}

	if !exists {
		return model.Response{Status: false, Message: "Produto não encontrado", Data: nil}, err
	}

	err = pu.repository.DeleteProduct(id)
	if err != nil {
		return model.Response{Status: false, Message: err.Error(), Data: nil}, err
	}

	return model.Response{Status: true, Data: nil}, nil

}

func (pu *ProductUseCase) UpdateProduct(product model.Product) (model.Response, error) {
	exists, err := pu.repository.ProductExists(product.ID)
	if err != nil {
		return model.Response{Status: false, Message: "Falha ao tentar obter produto", Data: nil}, err
	}

	if !exists {
		return model.Response{Status: false, Message: "Produto não encontrado", Data: nil}, err
	}

	err = pu.repository.UpdateProduct(product)

	if err != nil {
		return model.Response{Status: false, Message: "Falha ao tentar atualizar produto", Data: nil}, err
	}

	return model.Response{Status: true, Data: product}, nil
}

func (pu *ProductUseCase) GetProductById(id int) (model.Response, error) {
	product, err := pu.repository.GetProductById(id)
	if err != nil {
		return model.Response{Status: false, Message: "Falha ao tentar obter produto"}, err
	}

	return model.Response{Status: true, Data: product}, nil
}

func (pu *ProductUseCase) CreateProduct(product model.Product) (model.Response, error) {
	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Response{Status: false, Data: "Falha ao tentar criar produto"}, err
	}

	product.ID = productId

	return model.Response{Status: true, Data: product}, nil
}
