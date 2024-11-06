package Constant

import (
	"itam/Model/Database"
	"itam/Model/Web"
)

func ToBarangResponse(barang Database.Barang) Web.BarangResponse {
	return Web.BarangResponse{
		Id:       barang.ID,
		Name:     barang.Name,
		Category: barang.Category,
		Price:    barang.Price,
	}
}

func ToBarangResponses(barangs []Database.Barang) []Web.BarangResponse {
	var barangResponses []Web.BarangResponse
	for _, barang := range barangs {
		barangResponses = append(barangResponses, ToBarangResponse(barang))
	}
	return barangResponses
}
