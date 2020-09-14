// +build integration_test
// see ../../build/.drone.yml

package storage

import (
	"github.com/balugcath/mongo-test/pkg/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	conn = "mongodb://root:root@localhost:27017"
	db   = "test"
	coll = "test"
)

func TestMongo_Client(t *testing.T) {
	type args struct {
		conn string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				conn: conn,
			},
			wantErr: false,
		},
		{
			name: "test2",
			args: args{
				conn: conn + "1",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli, err := NewMongoClient(tt.args.conn)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMongoClient() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				cli.Close()
			}
		})
	}
}

func TestMongo_Store(t *testing.T) {
	type args struct {
		data []types.ProductPrice
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				data: []types.ProductPrice{
					{ProductName: "1", Price: 1},
					{ProductName: "2", Price: 2},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli, err := NewMongoClient(conn)
			defer cli.Close()
			if err != nil {
				t.Errorf("NewMongoClient() error = %v", err)
			}
			coll := NewMongoCollection(cli, db, coll)
			if err := coll.StoreProductPrice(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("MongoCollection.StoreProductPrice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMongo_Fetch(t *testing.T) {
	type args struct {
		data       []types.ProductPrice
		pageParams types.ProductPriceListParams
		coll       string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    []types.ProductPrice
	}{
		{
			name: "test1",
			args: args{
				data: []types.ProductPrice{
					{ProductName: "1", Price: 1},
					{ProductName: "2", Price: 2},
				},
				pageParams: types.ProductPriceListParams{
					LenPage: 2, SortField: types.PriceMogoTag, SortOrder: types.ASC,
				},
				coll: coll + "1",
			},
			want: []types.ProductPrice{
				{ProductName: "1", Price: 1, CountUpdate: 1},
				{ProductName: "2", Price: 2, CountUpdate: 1},
			},
			wantErr: false,
		},
		{
			name: "test2",
			args: args{
				data: []types.ProductPrice{
					{ProductName: "1", Price: 1},
					{ProductName: "2", Price: 2},
				},
				coll: coll,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "test3",
			args: args{
				data: []types.ProductPrice{
					{ProductName: "1", Price: 1},
					{ProductName: "2", Price: 2},
				},
				pageParams: types.ProductPriceListParams{
					LenPage: 2, SortOrder: types.ASC,
				},
				coll: coll,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "test4",
			args: args{
				data: []types.ProductPrice{
					{ProductName: "1", Price: 1},
					{ProductName: "2", Price: 2},
				},
				pageParams: types.ProductPriceListParams{
					SortField: types.PriceMogoTag, SortOrder: types.ASC,
				},
				coll: coll,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "test5",
			args: args{
				data: []types.ProductPrice{
					{ProductName: "1", Price: 1},
					{ProductName: "2", Price: 2},
				},
				pageParams: types.ProductPriceListParams{
					LenPage: 2, SortField: types.PriceMogoTag,
				},
				coll: coll,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "test6",
			args: args{
				data: []types.ProductPrice{
					{ProductName: "1", Price: 1},
					{ProductName: "2", Price: 2},
				},
				pageParams: types.ProductPriceListParams{
					LenPage: 2, SortField: types.PriceMogoTag, SortOrder: types.DESC,
				},
				coll: coll + "2",
			},
			want: []types.ProductPrice{
				{ProductName: "2", Price: 2, CountUpdate: 1},
				{ProductName: "1", Price: 1, CountUpdate: 1},
			},
			wantErr: false,
		},
		{
			name: "test7",
			args: args{
				data: []types.ProductPrice{
					{ProductName: "1", Price: 1},
					{ProductName: "2", Price: 2},
				},
				pageParams: types.ProductPriceListParams{
					LenPage: 2, SortField: types.ProductNameMongoTag, SortOrder: types.DESC,
				},
				coll: coll + "3",
			},
			want: []types.ProductPrice{
				{ProductName: "2", Price: 2, CountUpdate: 1},
				{ProductName: "1", Price: 1, CountUpdate: 1},
			},
			wantErr: false,
		},
		{
			name: "test8",
			args: args{
				data: []types.ProductPrice{
					{ProductName: "1", Price: 1},
				},
				pageParams: types.ProductPriceListParams{
					LenPage: 2, SortField: types.ProductNameMongoTag, SortOrder: types.DESC,
				},
				coll: coll + "3",
			},
			want: []types.ProductPrice{
				{ProductName: "2", Price: 2, CountUpdate: 1},
				{ProductName: "1", Price: 1, CountUpdate: 2},
			},
			wantErr: false,
		},
		{
			name: "test9",
			args: args{
				pageParams: types.ProductPriceListParams{
					LenPage: 2, SortField: types.ProductNameMongoTag, SortOrder: 0,
				},
				coll: coll + "4",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli, err := NewMongoClient(conn)
			defer cli.Close()
			if err != nil {
				t.Errorf("NewMongoClient() error = %v", err)
			}
			coll := NewMongoCollection(cli, db, tt.args.coll)
			if err := coll.StoreProductPrice(tt.args.data); err != nil {
				t.Errorf("MongoCollection.StoreProductPrice() error = %v", err)
			}

			got, err := coll.FetchProductPrice(tt.args.pageParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("MongoCollection.FetchProductPrice() error = %v, wantErr %v", err, tt.wantErr)
			}

			removeTimestamp(got)
			if !assert.Equal(t, got, tt.want) {
				t.Errorf("MongoCollection.FetchProductPrice() wan %v, got %v", tt.want, got)
			}
		})
	}
}

func removeTimestamp(s []types.ProductPrice) {
	for i := range s {
		s[i].LastUpdate = time.Time{}
	}
}
