package handler

import (
	"context"
	"encoding/csv"
	"io"
	"log"
	"net/http"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/balugcath/mongo-test/api"
	"github.com/balugcath/mongo-test/internal/csv2rec"
	"github.com/balugcath/mongo-test/pkg/types"
)

// GRPCHandler ...
type GRPCHandler struct {
	fetcher
	storer
}

type fetcher interface {
	FetchProductPrice(types.ProductPriceListParams) ([]types.ProductPrice, error)
}

type storer interface {
	StoreProductPrice([]types.ProductPrice) error
}

// NewGRPCHandler ...
func NewGRPCHandler(f fetcher, st storer) (*GRPCHandler, error) {
	s := &GRPCHandler{fetcher: f, storer: st}
	return s, nil
}

// List ...
func (s GRPCHandler) List(srv api.ProductPrice_ListServer) error {
	ctx := srv.Context()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("receive error %v", err)
			continue
		}

		pp, err := s.FetchProductPrice(types.ProductPriceListParams{
			SortOrder: int64(req.SortOrder),
			SortField: req.SortField,
			LenPage:   req.LenPage,
			Page:      req.Page,
		})
		if err != nil {
			switch err {
			default:
				return status.Errorf(codes.Internal, err.Error())
			case types.ErrLenPage, types.ErrSortField, types.ErrSortOrder:
				return status.Errorf(codes.InvalidArgument, err.Error())
			}
		}
		for _, v := range pp {
			pt, _ := ptypes.TimestampProto(v.LastUpdate)
			resp := api.ProductPriceListReply{
				ProductName: v.ProductName,
				Price:       v.Price,
				CountUpdate: v.CountUpdate,
				LastUpdate:  pt,
			}
			if err := srv.Send(&resp); err != nil {
				log.Printf("send error %v", err)
			}
		}
		if int64(len(pp)) < req.LenPage {
			return status.Errorf(codes.OK, io.EOF.Error())
		}
	}
}

// Fetch ...
func (s GRPCHandler) Fetch(_ context.Context, params *api.ProductPriceFetchParams) (*empty.Empty, error) {
	go func() {
		if err := s.fetch(params); err != nil {
			log.Println(err)
		}
	}()
	return &empty.Empty{}, nil
}

func (s GRPCHandler) fetch(params *api.ProductPriceFetchParams) error {
	resp, err := http.Get(params.Url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	a, err := csv.NewReader(resp.Body).ReadAll()
	if err != nil {
		return err
	}
	r, err := csv2rec.Array2Records(a)
	if err != nil {
		return err
	}
	if err := s.StoreProductPrice(r); err != nil {
		return err
	}
	return nil
}
