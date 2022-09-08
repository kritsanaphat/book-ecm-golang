# Ecommerce
# Routes
- SignUp()😀
- Login()🤝
- ProductViewAdmin()🍎
- SearchProduct()⛄
- SearchProductByQuery()☄️

# Model
- User struct
- Product struct 
- ProductUser struct
- Address struct
- Order struct
- Payment struct

# Main.go


# Controllers
***controllers.go***
- HashPassword(password string)
- VerifyPassword(userPassword string, givenPassword string) (bool, string) 
- Signup() 😀
- Login()🤝
- ProductViewerAdmin()🍎
- searchProduct()⛄
- SearchProductByQuery()☄️

***cart.go***
- AddToCart()
- RemoveItem() 
- GetItemFromCart()
- BuyFromCart()
- InstanBuy()\

***address.go***
- AddAddress()
- EditHomeAddress()
- EditWorkAddress()
- DeleteAddress()

# Database
***cart.go***
- AddProductToCart() 
- RemoveCartItem()
- BuyItemFromCart()
- InstantBuyer()

***databasesetup***
- DBSet()
- UserData(client *mongo.Client, collectionName string)
- ProductData(client *mongo.Client, collectionName string)