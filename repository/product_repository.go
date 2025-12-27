package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, product_name, price FROM product"
	rows, err := pr.connection.Query(query)

	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}

func (pr *ProductRepository) DeleteProduct(id int) error {

	query, err := pr.connection.Exec("DELETE FROM product WHERE id = $1", id)
	if err != nil {
		fmt.Println(err)
	}
	println("id: ", id)
	rowsAffected, _ := query.RowsAffected()
	fmt.Println("Linhas afetadas: ", rowsAffected)

	return nil
}

func (pr *ProductRepository) ProductExists(id int) (bool, error) {
	var exists bool

	query, err := pr.connection.Prepare("SELECT COUNT(1) FROM product WHERE id = $1")

	if err != nil {
		fmt.Println(err)
		return false, err
	}

	err = query.QueryRow(id).Scan(&exists)

	if err != nil {
		fmt.Println(err)
		return false, err
	}

	query.Close()
	return exists, nil
}

func (pr *ProductRepository) UpdateProduct(product model.Product) error {
	var id int
	fmt.Println("infos: ", product.Name, product.Price, product.ID)

	query, err := pr.connection.Prepare("UPDATE product SET " +
		"product_name = $1, price = $2 WHERE id = $3 RETURNING id")

	if err != nil {
		fmt.Println(err)
		return err
	}

	err = query.QueryRow(product.Name, product.Price, product.ID).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	query.Close()
	fmt.Println("Atualizado: ", id)

	return nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	var id int
	query, err := pr.connection.Prepare("INSERT INTO product" +
		"(product_name, price)" +
		"VALUES ($1, $2) RETURNING id")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()

	return id, nil
}

func (pr *ProductRepository) GetProductById(id int) (*model.Product, error) {
	query, err := pr.connection.Prepare("SELECT * FROM product WHERE id = $1")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var product model.Product

	err = query.QueryRow(id).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		fmt.Println(err)
		return nil, err
	}

	query.Close()
	return &product, nil
}
