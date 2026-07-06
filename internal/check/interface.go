package check

import (
	"context"
	"sync"

	authzGraph "github.com/openfga/language/pkg/go/graph"

	"github.com/openfga/openfga/pkg/storage"
)

type CheckResolver interface {
	ResolveCheck(context.Context, *Request) (*Response, error)
	ResolveUnion(context.Context, *Request, *authzGraph.WeightedAuthorizationModelNode, *sync.Map) (*Response, error)
	SpecificType(context.Context, *Request, *authzGraph.WeightedAuthorizationModelEdge) (*Response, error)
	SpecificTypeAndRelation(context.Context, *Request, *authzGraph.WeightedAuthorizationModelEdge, *sync.Map) (*Response, error)
	SpecificTypeWildcard(context.Context, *Request, *authzGraph.WeightedAuthorizationModelEdge) (*Response, error)
	TTU(context.Context, *Request, *authzGraph.WeightedAuthorizationModelEdge, *sync.Map) (*Response, error)
	ResolveRewrite(context.Context, *Request, *authzGraph.WeightedAuthorizationModelNode, *sync.Map) (*Response, error)
}

type Strategy interface {
	Resolve(context.Context, *Request, *authzGraph.WeightedAuthorizationModelEdge, storage.TupleKeyIterator, *sync.Map) (*Response, error)
	Userset(context.Context, *Request, *authzGraph.WeightedAuthorizationModelEdge, storage.TupleKeyIterator, *sync.Map) (*Response, error)
	TTU(context.Context, *Request, *authzGraph.WeightedAuthorizationModelEdge, storage.TupleKeyIterator, *sync.Map) (*Response, error)
}
