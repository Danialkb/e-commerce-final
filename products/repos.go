package products

import "gorm.io/gorm"

type ProductRepository interface {
	CreateProduct(*Product) error
	GetProductByID(uint) (*Product, error)
	UpdateProduct(*Product) error
	DeleteProduct(uint) error
}
type CategoryRepository interface {
	CreateCategory(category *Category) error
	GetCategoryByID(uint) (*Category, error)
	UpdateCategory(*Category) error
	DeleteCategory(uint) error
}

type productRepository struct {
	db *gorm.DB
}

type categoryRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) CreateProduct(product *Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) UpdateProduct(product *Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) DeleteProduct(id uint) error {
	return r.db.Delete(&Product{}, id).Error
}

func (r *productRepository) GetProductByID(id uint) (*Product, error) {
	var product Product
	if err := r.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) CreateCategory(category *Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) GetCategoryByID(id uint) (*Category, error) {
	var category Category
	if err := r.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}
func (r *categoryRepository) UpdateCategory(category *Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) DeleteCategory(id uint) error {
	return r.db.Delete(&Category{}, id).Error
}
