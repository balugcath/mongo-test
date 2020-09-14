package handler

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/balugcath/mongo-test/api"
	"github.com/balugcath/mongo-test/pkg/types"
	"google.golang.org/grpc/metadata"
)

var (
	countRecords = 10
)

type store struct {
	p []types.ProductPrice
}

func (s *store) StoreProductPrice(pp []types.ProductPrice) error {
	s.p = pp
	return nil
}

func (s *store) FetchProductPrice(p types.ProductPriceListParams) ([]types.ProductPrice, error) {
	var rep []types.ProductPrice
	for i := 0; i < countRecords; i++ {
		rep = append(rep, types.ProductPrice{ProductName: "name " + strconv.Itoa(i), Price: float64(i), CountUpdate: int64(i) + 1})
	}
	return rep, nil
}

func TestGRPCHandler_fetch(t *testing.T) {
	type fields struct {
		storer *store
	}
	type args struct {
		params *api.ProductPriceFetchParams
		fmtStr string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "test1",
			args:    args{fmtStr: "product %d,%f\n"},
			fields:  fields{storer: &store{p: []types.ProductPrice{}}},
			wantErr: false,
		},
		{
			name:    "test2",
			args:    args{fmtStr: "product %d, %f\n"},
			fields:  fields{storer: &store{p: []types.ProductPrice{}}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				for i := 0; i < countRecords; i++ {
					w.Write([]byte(fmt.Sprintf(tt.args.fmtStr, i, float64(i)+0.1)))
				}
			}))
			defer ts.Close()
			tt.args.params = &api.ProductPriceFetchParams{Url: ts.URL}

			s := GRPCHandler{
				storer: tt.fields.storer,
			}

			err := s.fetch(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GRPCHandler.Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			time.Sleep(time.Second)
			if !tt.wantErr && len(tt.fields.storer.p) != countRecords {
				t.Errorf("GRPCHandler.Fetch() = %v, want %v", len(tt.fields.storer.p), countRecords)
			}
		})
	}
}

type srv struct {
	p []*api.ProductPriceListReply
}

func (s *srv) SetHeader(_ metadata.MD) error { return nil }
func (s *srv) SendHeader(metadata.MD) error  { return nil }
func (s *srv) SetTrailer(metadata.MD)        {}
func (s *srv) Context() context.Context {
	ctx := context.Background()
	return ctx
}
func (s *srv) SendMsg(m interface{}) error { return nil }
func (s *srv) RecvMsg(m interface{}) error { return nil }

func (s *srv) Send(r *api.ProductPriceListReply) error {
	s.p = append(s.p, r)
	return nil
}

func (s *srv) Recv() (*api.ProductPriceListParams, error) {
	return &api.ProductPriceListParams{
		Page:      0,
		LenPage:   15,
		SortField: types.PriceMogoTag,
		SortOrder: types.ASC,
	}, nil
}

func TestGRPCHandler_List(t *testing.T) {
	type fields struct {
		fetcher fetcher
	}
	type args struct {
		srv *srv
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "test1",
			args:    args{srv: &srv{}},
			fields:  fields{fetcher: &store{}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := GRPCHandler{
				fetcher: tt.fields.fetcher,
			}
			if err := s.List(tt.args.srv); (err != nil) != tt.wantErr {
				t.Errorf("GRPCHandler.List() error = %v, wantErr %v", err, tt.wantErr)
			}
			if len(tt.args.srv.p) != countRecords {
				t.Errorf("GRPCHandler.Fetch() = %v, want %v", len(tt.args.srv.p), countRecords)
			}
		})
	}
}
