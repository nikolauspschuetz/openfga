package check

import (
	"context"
	"sync"

	authzGraph "github.com/openfga/language/pkg/go/graph"
	"github.com/openfga/openfga/internal/modelgraph"
	"github.com/openfga/openfga/pkg/storage"
)

type SqlStrategy struct {
	datastore storage.RelationshipTupleReader
	model     *modelgraph.AuthorizationModelGraph
}

func NewSql(model *modelgraph.AuthorizationModelGraph, datastore storage.RelationshipTupleReader) *SqlStrategy {
	return &SqlStrategy{
		model:     model,
		datastore: datastore,
	}
}

func (s *SqlStrategy) weight1(ctx context.Context, req *Request, edge *authzGraph.WeightedAuthorizationModelEdge) (*Response, error) {
	builder := s.datastore.Builder(req.Consistency)

	return &Response{}, nil
}

func (s *SqlStrategy) Resolve(ctx context.Context, req *Request, edge *authzGraph.WeightedAuthorizationModelEdge, i storage.TupleKeyIterator, visited *sync.Map) (*Response, error) {
	var weight int
	var ok bool

	if weight, ok = edge.GetWeight(req.GetUserType()); !ok {
		return &Response{}, nil
	}

	switch weight {
	case 1:
		return s.weight1(ctx, req, edge)
	case 2:
		fallthrough
	case authzGraph.Infinite:
		fallthrough
	default:
		return &Response{}, nil
	}
}

func (s *SqlStrategy) Userset(context.Context, *Request, *authzGraph.WeightedAuthorizationModelEdge, storage.TupleKeyIterator, *sync.Map) (*Response, error) {
	return &Response{}, nil
}

func (s *SqlStrategy) TTU(context.Context, *Request, *authzGraph.WeightedAuthorizationModelEdge, storage.TupleKeyIterator, *sync.Map) (*Response, error) {
	return &Response{}, nil
}
