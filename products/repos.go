package products

import (
	"e-commerce-final/settings"
	"gorm.io/gorm"
	"log"
)

type ProductRepositoryInterface interface {
	CreateProduct(*Product) error
	GetProductByID(uint) (*Product, error)
	UpdateProduct(*Product) error
	DeleteProduct(uint) error
}
type CategoryRepositoryInterface interface {
	CreateCategory(category *Category) error
	GetCategoryByID(uint) (*Category, error)
	UpdateCategory(*Category) error
	DeleteCategory(uint) error
}

type ProductRepositoryV1 struct {
	DB *gorm.DB
}

type CategoryRepositoryV1 struct {
	DB *gorm.DB
}

func NewProductRepository() *ProductRepositoryV1 {
	db, err := settings.DbSetup()
	if err != nil {
		log.Fatal(err)
		return &ProductRepositoryV1{}
	}
	return &ProductRepositoryV1{DB: db}
}

func (p *ProductRepositoryV1) CreateProduct(product *Product) error {
	return p.DB.Create(product).Error
}

func (p *ProductRepositoryV1) UpdateProduct(product *Product) error {
	return p.DB.Save(product).Error
}

func (p *ProductRepositoryV1) DeleteProduct(id uint) error {
	return p.DB.Delete(&Product{}, id).Error
}

func (p *ProductRepositoryV1) GetProductByID(id uint) (*Product, error) {
	var product Product
	if err := p.DB.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func NewCategoryRepository() *CategoryRepositoryV1 {
	db, err := settings.DbSetup()
	if err != nil {
		log.Fatal(err)
		return &CategoryRepositoryV1{}
	}
	return &CategoryRepositoryV1{DB: db}
}

func (c *CategoryRepositoryV1) CreateCategory(category *Category) error {
	return c.DB.Create(category).Error
}

func (c *CategoryRepositoryV1) GetCategoryByID(id uint) (*Category, error) {
	var category Category
	if err := c.DB.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}
func (c *CategoryRepositoryV1) UpdateCategory(category *Category) error {
	return c.DB.Save(category).Error
}

func (c *CategoryRepositoryV1) DeleteCategory(id uint) error {
	return c.DB.Delete(&Category{}, id).Error
}
