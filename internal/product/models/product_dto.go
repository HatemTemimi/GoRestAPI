package product

type ProductDTO struct {
	ID    uint   `json:"id,string,omitempty"`
	Code  string `json:"code,string" validate:"required, gt=0"`
	Price uint   `json:"price,string" validate:"required, gt=0"`
}
