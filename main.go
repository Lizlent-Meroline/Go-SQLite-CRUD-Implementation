package main

func main() {
	db, err := InitDB("store.db")
	if err != nil {
		panic(err)
	}

	store := &Store{DB: db}

	id, _ := store.CreateProduct(&Product{
		Name: "Laptop", Price: 1200.50, Quantity: 5,
	})

	fmt.Println("Inserted ID:", id)

	p, _ := store.GetProduct(int(id))
	fmt.Println(p)

	store.UpdateProduct(&Product{
		ID: int(id), Name: "Gaming Laptop", Price: 1500, Quantity: 3,
	})

	products, _ := store.ListProducts(100, "Laptop")
	fmt.Println(products)

	store.DeleteProduct(int(id))
}
