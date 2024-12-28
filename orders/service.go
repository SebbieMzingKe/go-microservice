package main

import (
	"common"
	pb "common/api"
	"context"
	"log"
)

type service struct {
	store OrdersStore
}

func NewService(store OrdersStore) *service  {
	return &service{store}
}

func (s *service) CreateOrder(context.Context) error {
	return nil
}

func (s *service) ValidateOrder(ctx context.Context, p *pb.CreateOrderRequest) error {
	if len(p.Items) == 0 {
		return common.ErrNoItems
	}
	mergedItems := mergeItemsQuantities(p.Items)
	log.Println(mergedItems)
	// validate with stock service
	return nil
}


func mergeItemsQuantities(items []*pb.ItemsWithQuantity) []*pb.ItemsWithQuantity {
	merged := make([]*pb.ItemsWithQuantity, 0)

	for _, item := range items {
		found := false
		for _, finalItem := range merged {
			if finalItem.CustomerID == item.CustomerID {
				finalItem.Quantity += item.Quantity
				found = true
				break
			}
		}
		if !found {
			merged = append(merged, item)
		}
	}
	return merged
}