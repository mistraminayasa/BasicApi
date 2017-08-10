package main
import (
 "fmt"
 "github.com/gin-gonic/gin"
 "github.com/jinzhu/gorm"
 _ "github.com/jinzhu/gorm/dialects/mysql"
)
var db *gorm.DB
var err error
type Paid struct {  
    ID uint `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
    Telepon string `gorm:"not null" form:"telepon" json:"telepon"`
    Name string `gorm:"not null" form:"name" json:"name"`
    Email string `gorm:"not null" form:"email" json:"email"`
    Code string `gorm:"not null" form:"code" json:"code"`
}
func main() {
  db, _ = gorm.Open("mysql", "root:sekarang9@tcp(127.0.0.1:3306)/data_karyawan?charset=utf8&parseTime=True&loc=Local")
 if err != nil {
    fmt.Println(err)
 }  
 defer db.Close()
 db.AutoMigrate(&Paid{})
 r := gin.Default()
 r.GET("/postpaid/", GetPaid)
 r.GET("/postpaid/:code", GetPostpaid)
 r.POST("/postpaid/take", CreatePostpaid)
 r.PUT("/postpaid/:id", UpdatePostpaid)
 r.DELETE("/postpaid/:id", DeletePaid)
 r.Run(":8080")
}
func DeletePaid(c *gin.Context) { 

 id := c.Params.ByName("id")
 var paid Paid
 d := db.Where("id = ?", id).Delete(&paid)
 fmt.Println(d)
 c.JSON(200, gin.H{"id #" + id: "deleted"})
}

func UpdatePostpaid(c *gin.Context) {  
 var paid Paid
 id := c.Params.ByName("id")
 if err := db.Where("id = ?", id).First(&paid).Error; err != nil {
    c.AbortWithStatus(404)
    fmt.Println(err)
 }
 c.BindJSON(&paid)
 db.Save(&paid)
 c.JSON(200, paid)
}
func CreatePostpaid(c *gin.Context) {
 var paid Paid
 c.BindJSON(&paid)
 db.Create(&paid)
 c.JSON(200, paid)
}
func GetPostpaid(c *gin.Context) {
 code := c.Params.ByName("code")
 var paid Paid
 if err := db.Where("code = ?", code).First(&paid).Error; err != nil {
    c.AbortWithStatus(404)
    fmt.Println(err)
 } else {
    c.JSON(200, paid)
 }
}
func GetPaid(c *gin.Context) {
 var paid []Paid
 if err := db.Find(&paid).Error; err != nil {
    c.AbortWithStatus(404)
    fmt.Println(err)
 } else {
    c.JSON(200, paid)
 }
}


