package products

import "e-commerce-final/comments"

type ProductServiceInterface interface {
	CreateProduct(*Product) error
	GetProductByID(uint) (*Product, error)
	UpdateProduct(*Product) error
	DeleteProduct(uint) error
	GetProducts() ([]Product, error)
	GetCommentsByProductId(uint) ([]*comments.Comment, error)
	SearchByName(name string) ([]Product, error)
}

type CategoryServiceInterface interface {
	CreateCategory(category *Category) error
	GetCategoryByID(uint) (*Category, error)
	UpdateCategory(*Category) error
	DeleteCategory(uint) error
	GetCategories() ([]Category, error)
}

type ProductServiceV1 struct {
	productRepos ProductRepositoryInterface
}

type CategoryServiceV1 struct {
	categoryRepos CategoryRepositoryInterface
}

func NewProductService() ProductServiceV1 {
	return ProductServiceV1{productRepos: NewProductRepository()}
}

func NewCategoryService() CategoryServiceV1 {
	return CategoryServiceV1{categoryRepos: NewCategoryRepository()}
}

func (p ProductServiceV1) GetProducts() ([]Product, error) {
	return p.productRepos.GetProducts()
}

func (p ProductServiceV1) SearchByName(title string) ([]Product, error) {
	return p.productRepos.SearchByName(title)
}

func (p ProductServiceV1) GetCommentsByProductId(id uint) ([]*comments.Comment, error) {
	return p.productRepos.GetCommentsByProductId(id)
}

func (p ProductServiceV1) CreateProduct(product *Product) error {
	return p.productRepos.CreateProduct(product)
}

func (p ProductServiceV1) GetProductByID(id uint) (*Product, error) {
	return p.productRepos.GetProductByID(id)
}

func (p ProductServiceV1) UpdateProduct(product *Product) error {
	return p.productRepos.UpdateProduct(product)
}

func (p ProductServiceV1) DeleteProduct(id uint) error {
	return p.productRepos.DeleteProduct(id)
}

func (c CategoryServiceV1) CreateCategory(category *Category) error {
	return c.categoryRepos.CreateCategory(category)
}

func (c CategoryServiceV1) GetCategoryByID(id uint) (*Category, error) {
	return c.categoryRepos.GetCategoryByID(id)
}

func (c CategoryServiceV1) UpdateCategory(category *Category) error {
	return c.categoryRepos.UpdateCategory(category)
}

func (c CategoryServiceV1) DeleteCategory(id uint) error {
	return c.categoryRepos.DeleteCategory(id)
}

func (c CategoryServiceV1) GetCategories() ([]Category, error) {
	return c.categoryRepos.GetCategories()
}
