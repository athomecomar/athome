package ent

import (
	"context"
	"fmt"

	"github.com/athomecomar/athome/pb/pbimages"
	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/athome/pb/pbsemantic"
	"github.com/athomecomar/athome/pb/pbusers"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/currency"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

type Product struct {
	Id     uint64 `json:"id,omitempty"`
	UserId uint64 `json:"user_id,omitempty"`

	Title      string `json:"title,omitempty"`
	CategoryId uint64 `json:"category_id,omitempty"`

	Price         currency.ARS `json:"price,omitempty"`
	Stock         uint64       `json:"stock,omitempty"`
	ReservedStock uint64       `json:"reserved_stock,omitempty"`
}

func (p *Product) ReserveStock(qt uint64) error {
	if p.Stock < qt {
		return fmt.Errorf("qt to reserve: %d while got: %v", qt, p.Stock)
	}
	p.Stock -= qt
	p.ReservedStock += qt
	return nil
}

func (p *Product) UndoStockReserve(qt uint64) error {
	if p.ReservedStock < qt {
		return fmt.Errorf("qt to undo: %d while got: %v", qt, p.ReservedStock)
	}
	p.ReservedStock -= qt
	p.Stock += qt
	return nil
}

func (p *Product) ConsumeStockReserve(qt uint64) error {
	if p.ReservedStock < qt {
		return fmt.Errorf("qt to consume: %d while got: %v", qt, p.ReservedStock)
	}
	p.ReservedStock -= qt
	return nil
}

func FindProduct(ctx context.Context, db *sqlx.DB, id uint64) (*Product, error) {
	row := db.QueryRowxContext(ctx, `SELECT * FROM products WHERE id=$1`, id)
	prod := &Product{}
	err := row.StructScan(prod)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return prod, nil
}

func FindProductsById(ctx context.Context, db *sqlx.DB, ids []uint64) ([]*Product, error) {
	rows, err := db.QueryxContext(ctx, `SELECT * FROM products WHERE id=any($1)`, pq.Array(ids))
	if err != nil {
		return nil, errors.Wrap(err, "QueryxContext")
	}
	defer rows.Close()
	var prods []*Product
	for rows.Next() {
		prod := &Product{}
		err = rows.StructScan(prod)
		if err != nil {
			return nil, errors.Wrap(err, "StructScan")
		}
		prods = append(prods, prod)
	}
	return prods, nil
}

func (p *Product) ToPb() *pbproducts.Product {
	return &pbproducts.Product{
		Title:      p.Title,
		CategoryId: p.CategoryId,
		Price:      p.Price.Float64(),
		Stock:      p.Stock,
	}
}

func (p *Product) User(ctx context.Context, users pbusers.ViewerClient) (*pbusers.User, error) {
	return users.RetrieveUser(ctx, &pbusers.RetrieveUserRequest{UserId: p.UserId})
}

func (p *Product) Images(ctx context.Context, img pbimages.ImagesClient) (map[string]*pbimages.Image, error) {
	resp, err := img.RetrieveImages(ctx, &pbimages.RetrieveImagesRequest{
		Entity: pbutil.ToPbEntity(p),
	})
	if err != nil {
		return nil, errors.Wrap(err, "RetrieveImages")
	}
	return resp.GetImages(), nil
}

func (p *Product) GetAttributes(ctx context.Context, sem pbsemantic.ProductsClient) (map[uint64]*pbproducts.Attribute, error) {
	schemas, err := sem.RetrieveAttributeSchemas(ctx, &pbsemantic.RetrieveAttributeSchemasRequest{CategoryId: p.CategoryId})
	if err != nil {
		return nil, errors.Wrap(err, "sem.RetrieveAttributesSchema")
	}
	datas, err := sem.RetrieveAttributeDatas(ctx, &pbsemantic.RetrieveAttributeDatasRequest{
		Entity: pbutil.ToPbEntity(p),
	})
	if err != nil {
		return nil, errors.Wrap(err, "sem.RetrieveAttributesData")
	}
	atts := make(map[uint64]*pbproducts.Attribute)
	for id, data := range datas.GetAttributes() {
		schema, ok := schemas.GetAttributes()[data.GetSchemaId()]
		if !ok {
			return nil, fmt.Errorf("couldnt find schema with id: %v for attribute data: %v", data.GetSchemaId(), id)
		}
		atts[id] = attributeFromDataAndSchema(data, schema)
	}

	return atts, nil
}

func attributeFromDataAndSchema(data *pbsemantic.AttributeData, schema *pbsemantic.AttributeSchema) *pbproducts.Attribute {
	return &pbproducts.Attribute{
		Name:      schema.GetName(),
		ValueType: schema.GetValueType(),
		Values:    data.GetValues(),
		SchemaId:  data.GetSchemaId(),
	}
}
