# Project


## Backend
	Code backend nằm trong thư mục backend
	Chạy code backend trên port 3000
	Chạy code bằng câu lệnh `go run app.go`
	
	API get all products: http://localhost:3000/api/product?limit={limit}&offset={offset}
		VD: http://localhost:3000/api/product?limit=17&offset=0
	API get product by id: http://localhost:3000/api/product/:id
		VD: http://localhost:3000/api/product/1
	API get product by category: http://localhost:3000/api/product/category/:name_category
		VD: http://localhost:3000/api/product/category/Armchair
		
## Frontend
	Chạy lệnh `npm install` để cài đặt các modules liên quan
	Code frontend nằm trong thư mục frontend
	Code frontend chạy trên cổng 8080
	