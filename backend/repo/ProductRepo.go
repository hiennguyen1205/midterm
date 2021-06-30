package repo

import (
	"fmt"

	"github.com/TechMaster/golang/15GoMySQL/model"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
)

func initProduct() {
	// var sony, xiaomi model.Manufacturer
	// Db.Where("name = ?", "Sony").First(&sony)
	// Db.Where("name = ?", "Xiaomi").First(&xiaomi)

	var sofa, armchair, table model.Category
	Db.Where("name LIKE ?", "sofa%").First(&sofa)
	Db.Where("name LIKE ?", "armchair%").First(&armchair)
	Db.Where("name LIKE ?", "table%").First(&table)

	armchair_1 := model.Product{
		Name:        "Vintage Luxury Chair",
		Description: "Đối với những tín đồ của phong cách bán cổ điển thì chiếc ghế armchair Polo này rất đáng để thử. Nhờ vào thiết kế với những đường cong nhẹ nhàng, vừa đủ và phần vải bọc trung tính, armchair có thể đi kèm với rất nhiều phong cách và không gian khác nhau, từ phòng khách đến phòng ngủ và cả phòng đọc sách.",
		Price:       8900000,
		Category:    &armchair,
		Image:       "public/img/armchair_1.jpg",
		Sale:        0,
	}
	Db.Create(&armchair_1)

	armchair_2 := model.Product{
		Name:        "Armchair Tudor Pink",
		Description: "Nhờ vào thiết kế tôn vinh đường cong tuyệt mỹ, các điểm nhấn hình kim cương ở phần tựa lưng và đệm ngồi cùng với chất liệu vải nhung cao cấp, Armchair Tudor mang lại cảm giác sang trọng và đẳng cấp. Đặc biệt khi được kết hợp với không gian mở, phá cách thì chiếc ghế bành này toát lên sự thư giãn và tận hưởng. Còn trong không gian phòng ngủ hoặc phòng đọc sách, chiếc ghế bành này còn tạo ra sự thoải mái vô cùng.",
		Price:       19900000,
		Category:    &armchair,
		Image:       "public/img/armchair_2.jpg",
		Sale:        20,
	}
	Db.Create(&armchair_2)

	armchair_3 := model.Product{
		Name:        "Armchair Maxine",
		Description: "Armchair Maxine với thiết kế tinh tế, trang nhã và nhỏ gọn phù hợp với căn hộ có diện tích vừa phải. Phần lưng tựa bằng da bò kết hợp gỗ Xà cừ màu nâu mang lại cảm giác ấm cúng cho không gian phòng khách. Maxine – Nét nâu trầm sang trọng Maxine, mang thiết kế vượt thời gian, gửi gắm và gói gọn lại những nét đẹp của thiên nhiên và con người nhưng vẫn đầy tính ứng dụng cao trong suốt hành trình phiêu lưu của nhà thiết kế người Pháp Dominique Moal. Bộ sưu tập nổi bật với màu sắc nâu trầm sang trọng, là sự kết hợp giữa gỗ, da bò và kim loại vàng bóng. Đặc biệt trên mỗi sản phẩm, những nét bo viền, chi tiết kết nối được sử dụng khéo léo tạo ra điểm nhất rất riêng cho bộ sưu tập. Maxine thể hiện nét trầm tư, thư giãn để tận hưởng không gian sống trong nhịp sống hiện đại. Sản phẩm thuộc BST Maxine được sản xuất tại Việt Nam.",
		Price:       13900000,
		Category:    &armchair,
		Image:       "public/img/armchair_3.jpg",
		Sale:        10,
	}
	Db.Create(&armchair_3)

	armchair_4 := model.Product{
		Name:        "Armchair Vải Santos",
		Description: "Kích thước nhỏ gọn phù hợp cho các căn hộ diện tích hẹp - nệm ngồi êm , chức năng xoay nhẹ nhàng nhờ chân khung kim loại , vải bọc màu trung tính dễ dàng phối với các mẫu sofa .",
		Price:       17400000,
		Category:    &armchair,
		Image:       "public/img/armchair_4.jpg",
		Sale:        40,
	}
	Db.Create(&armchair_4)

	armchair_5 := model.Product{
		Name:        "Armchair Hantz",
		Description: "Armchair Hantz có phần chân kim loại sơn đen và phần đệm ngồi bọc vải cao cấp, mang lại cảm giác nhẹ nhàng nhưng lại rất vừa vặn và thoải mài. Armchair phù hợp với những không gian nội thất hiện đại và chuộng sắc màu. Nhờ thiết kế nhỏ gọn, sản phẩm có thể được dùng trong cả phòng khách lẫn phòng ngủ.",
		Price:       4900000,
		Category:    &armchair,
		Image:       "public/img/armchair_5.jpg",
		Sale:        50,
	}
	Db.Create(&armchair_5)

	armchair_6 := model.Product{
		Name:        "Armchair Garbo",
		Description: "Armchair Garbo là sự kết hợp mới mẻ giữa phần chân kim loại sơn đen với phần đệm ngồi bọc vải cao cấp, mang lại cảm giác thanh mảnh nhưng rất chắc chắn. Armchair phù hợp với những không gian phòng khách hiện đại và phong cách. Với những gam màu sáng, sản phẩm sẽ là điểm nhấn cho nội thất nhà bạn.",
		Price:       10990000,
		Category:    &armchair,
		Image:       "public/img/armchair_6.jpg",
		Sale:        10,
	}
	Db.Create(&armchair_6)

	armchair_7 := model.Product{
		Name:        "Armchair Tudor Pink",
		Description: "Coco là chiếc ghế được thiết kế hình cong phễu ở phần lưng với khung thép kim loại dầy 20mm, phần đệm ngồi và phần lưng tựa được bọc hoàn toàn bởi chất liệu là nhung và da cho từng phiên bản. Calligaris đã ấn định màu sắc riêng biệt của mình vượt mọi thời đại với biểu tượng của thập niên 50. Cá tính nổi bật bởi phần khung kim loại sáng bóng và phần ghế ngồi hình vỏ xò hết sức lôi cuốn thì chiếc ghế này cũng có thể phối với nhiều phần vật liệu cũng như màu sắc khác nhau",
		Price:       29000000,
		Category:    &armchair,
		Image:       "public/img/armchair_7.jpg",
		Sale:        0,
	}
	Db.Create(&armchair_7)

	armchair_8 := model.Product{
		Name:        "Armchair String",
		Description: "Chiếc ghế Armchair được thiết kế đặc biệt với điểm nhấn là phần tay vịn được đan dây mang lại phong cách Retro cho phòng khách nhà bạn.",
		Price:       16500000,
		Category:    &armchair,
		Image:       "public/img/armchair_8.jpg",
		Sale:        50,
	}
	Db.Create(&armchair_8)

	sofa_1 := model.Product{
		Name:        "Sofa Miami",
		Description: "Sofa Miami 2 chỗ là một thiết kế tối giản cho không gian phòng khách hiện đại, chất liệu sofa vải cao cấp trên tông màu xám nhạt rất dễ dàng phối hợp cùng các sản phẩm khác. Kích thước nhỏ gọn 2 chỗ sẽ phù hợp cho các căn hộ, hoặc những góc nhỏ trong ngôi nhà của bạn.",
		Price:       12500000,
		Category:    &sofa,
		Image:       "public/img/sofa_1.jpg",
		Sale:        0,
	}
	Db.Create(&sofa_1)

	sofa_2 := model.Product{
		Name:        "Sofa Lovely",
		Description: "Sofa Lovely 3 chỗ với thiết kế đường nét thanh mảnh nhẹ nhàng. Chất liệu vải cao cấp, màu sắc tươi mới. Sofa Lovely là lựa chọn tối ưu cho phòng khách căn hộ có diện tích nhỏ.",
		Price:       16000000,
		Category:    &sofa,
		Image:       "public/img/sofa_2.jpg",
		Sale:        10,
	}
	Db.Create(&sofa_2)

	sofa_3 := model.Product{
		Name:        "Sofa John",
		Description: "Sofa John 3 chỗ với thiết kế mạnh mẽ - khỏe khoắn với hình khối vuông cạnh là điểm cộng của mẫu sofa John. Phần nệm ngồi rộng cho phép nằm thư giãn hoặc một giấc ngủ sâu sau một ngày bận rộn. Sofa có phần phối vải kẻ sọc cho sofa hình khối không những không bị cứng mà còn là điểm nhấn riêng biệt. Tại Nhà Xinh có đa dạng các mẫu sofa đẹp hiện đại, đa dạng kiểu dáng, phù hợp cho từng không gian nhà bạn.",
		Price:       25900000,
		Category:    &sofa,
		Image:       "public/img/sofa_3.jpg",
		Sale:        20,
	}
	Db.Create(&sofa_3)

	sofa_4 := model.Product{
		Name:        "Sofa Twoback",
		Description: "Sofa Twoback 2.5 chỗ với gam màu xanh tươi mát, nhã nhặn. Kết cấu khung làm từ gỗ thông của Newzerland, bọc nệm vải cao cấp tạo cảm giác thoải mái. Sofa Twoback là 1 lựa chọn tối ưu cho không gian phòng khách hiện đại.",
		Price:       31800000,
		Category:    &sofa,
		Image:       "public/img/sofa_4.jpg",
		Sale:        30,
	}
	Db.Create(&sofa_4)

	sofa_5 := model.Product{
		Name:        "Sofa Dubai",
		Description: "Sofa Dubai 2.5 chỗ với đường nét mỏng đảm bảo cái nhìn nhẹ nhàng và thanh thoát. Thiết kế sofa 2 chỗ nhưng vẫn mang lại cảm giác chỗ ngồi rộng hơn. Sofa Dubai có 2 màu nâu và kem để chọn lựa phù hợp cho không gian phòng khách hiện đại của gia đình bạn.",
		Price:       16500000,
		Category:    &sofa,
		Image:       "public/img/sofa_5.jpg",
		Sale:        10,
	}
	Db.Create(&sofa_5)

	table_1 := model.Product{
		Name:        "Bàn nước Pio",
		Description: "Bàn nước PIO thu hút ánh nhìn với mặt bàn bằng chất liệu hiện đại melamine marble. Thiết kế với kiểu dáng oval giúp, bàn nước Pio tạo điểm nhấn khác biệt cho không gian phòng khách của bạn. Việc bổ sung ngăn bên dưới là khu chứa đồ cũng như trưng bày các vật dụng trang trí. PIO – Vẻ đẹp yên bình giữa lối sống đô thị Pha trộn giữa phong cách scandinavian và sự mới lạ trong chọn lựa màu sắc, bộ sưu tập PIO toát lên vẻ đẹp nhẹ nhàng, thanh lịch và cũng rất gần gũi với thiên nhiên. PIO thể hiện lối sống của những người trẻ, biết chiêm nghiệm và thưởng ngoạn sự trở về bình yên giữa nhịp sống hiện đại. Thiết kế bởi những đường cong, điểm xuyến các chi tiết nhấn nhá bên cạnh sử dụng các vật liệu như gỗ beech, melamine marble.. giúp PIO trở nên cá tính và thu hút trong không gian hiện đại. Sản phẩm được thiết kế bởi đội ngũ Nhà Xinh và sản xuất tại Việt Nam.",
		Price:       6900000,
		Category:    &table,
		Image:       "public/img/table_1.jpg",
		Sale:        0,
	}
	Db.Create(&table_1)

	table_2 := model.Product{
		Name:        "Bàn nước Jazz",
		Description: "Bàn nước Jazz được ghép từ những thanh gỗ sồi già tự nhiên. Bề mặt đặc trưng với những đường nứt tét gỗ tự nhiên được xử lý khéo léo, kết hợp với chân sắt sơn tĩnh điện đầy mạnh mẽ sẽ mang lại nét cá tính độc đáo cho gia chủ.",
		Price:       9500000,
		Category:    &table,
		Image:       "public/img/table_2.jpg",
		Sale:        10,
	}
	Db.Create(&table_2)

	table_3 := model.Product{
		Name:        "Bàn nước Mây",
		Description: "Một chiếc bàn nước kết hợp nhịp nhàng bởi 2 khối hình khác nhau về độ cao. Bàn nước Mây giúp cho không gian phòng khách trở nên cá tính hơn. Sản phẩm sử dụng chất liệu đã marble cho phần mặt bàn, được bao quanh bởi kết cấu khung gỗ và nhấn nhá với phần chân kim loại đồng hiện đại.",
		Price:       10900000,
		Category:    &table,
		Image:       "public/img/table_3.jpg",
		Sale:        20,
	}
	Db.Create(&table_3)

	table_4 := model.Product{
		Name:        "Bàn nước Daylight",
		Description: "Bàn nước Daylight với thiết kế hiện đại, kết hợp giữa mặt đá và chân inox mạ sang trọng sẽ là điểm nhấn độc đáo cho phòng khách nhà bạn.",
		Price:       27500000,
		Category:    &table,
		Image:       "public/img/table_4.jpg",
		Sale:        30,
	}
	Db.Create(&table_4)

}

func getAllProducts() []model.Product{
	var products []model.Product
	Db.Preload("Category").Find(&products)
	return products
}

func FindProductById(Id int) *model.Product {
	var product *model.Product
	Db.Preload("Category").First(&product, Id)
	fmt.Println(product)
	return product
}

func FindProductByCategory(category string) (products []model.Product) {
	switch category {
	case "Armchair":
		{
			Db.Where("category_id=?", "1").Find(&products)
			return products
		}
	case "Table":
		{
			Db.Where("category_id=?", "2").Find(&products)
			return products
		}
	case "Sofa":
		{
			Db.Where("category_id=?", "3").Find(&products)
			return products
		}
	default:
		{
			Db.Find(&products)
			return products
		}
	}

}

/*
type Product struct {
	ID             uint
	Name           string
	Description    string
	Price          uint
	Madein         string
	Country        *Country `gorm:"foreignKey:Madein"`
	ManufacturerID uint
	Manufacturer   *Manufacturer
	CategoryID     uint
	Category       *Category
}
*/
