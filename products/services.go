package products

type ProductService interface {
	CreateProduct(*Product) error
	GetProductByID(uint) (*Product, error)
	UpdateProduct(*Product) error
	DeleteProduct(uint) error
}

type CategoryService interface {
	CreateCategory(category *Category) error
	GetCategoryByID(uint) (*Category, error)
	UpdateCategory(*Category) error
	DeleteCategory(uint) error
}
type productService struct {
	pr ProductRepository
}

type categoryService struct {
	cr CategoryRepository
}

func NewProductService(pr ProductRepository) ProductService {
	return &productService{pr: pr}
}

func NewCategoryService(cr CategoryRepository) CategoryService {
	return &categoryService{cr: cr}
}

func (p *productService) CreateProduct(product *Product) error {
	return p.pr.CreateProduct(product)
}

func (p *productService) GetProductByID(id uint) (*Product, error) {
	return p.pr.GetProductByID(id)
}

func (p *productService) UpdateProduct(product *Product) error {
	return p.pr.UpdateProduct(product)
}

func (p *productService) DeleteProduct(id uint) error {
	return p.pr.DeleteProduct(id)
}

func (c *categoryService) CreateCategory(category *Category) error {
	return c.cr.CreateCategory(category)
}

func (c *categoryService) GetCategoryByID(id uint) (*Category, error) {
	return c.cr.GetCategoryByID(id)
}

func (c *categoryService) UpdateCategory(category *Category) error {
	return c.cr.UpdateCategory(category)
}

func (c *categoryService) DeleteCategory(id uint) error {
	return c.cr.DeleteCategory(id)
}
