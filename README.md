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
- NewApplication(database.ProductData(database.Cliemt, "Product")🦀\
- database.UserData(database.Client, "User"🍄)


# Controllers
***controllers.go***
- HashPassword(password string)
- VerifyPassword(userPassword string, givenPassword string) (bool, string) 
- Signup() 😀
    - email check
    - phone check 
- Login()🤝
- ProductViewerAdmin()🍎
- searchProduct()⛄
- SearchProductByQuery()☄️

***cart.go***
- AddToCart()🌐
- RemoveItem() 🌐
- GetItemFromCart()
- BuyFromCart()🌐
- InstanBuy()🌐

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

***databasesetup.go***
- DBSet()
- UserData(client *mongo.Client, collectionName string)🍄
- ProductData(client *mongo.Client, collectionName string)🦀
- AddToCart()🌐
- RemoveItem()🌐
- BuyFromCart()🌐
- InstanBuy()🌐

# Middleware


# ทำไม gin.context ถึงสำคัญ
- เมื่อใดก็ตามที่คุณเรียกใช้เมื่อใดก็ตามที่คุณส่งคำขอไปยัง api ที่คำขอสามารถเข้าถึงได้ง่ายด้วยความช่วยเหลือของ gin.context

# i'm test