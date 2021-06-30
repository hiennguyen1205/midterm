package repo

import (
	"fmt"

	"github.com/TechMaster/golang/15GoMySQL/model"
)

func initCategory() {

	 categories := map[int]string{
		 1:"Sofa",
		 2:"Armchair",
		 3:"Table",
	 }
	fmt.Println(categories)

	for _, category := range categories {
		createCategory := model.Category{Name: category}
		result := Db.Create(&createCategory)
		if result.Error != nil {
			panic(result.Error)
		}
	}
}
