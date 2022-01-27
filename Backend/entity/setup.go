package entity

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	//   "time"
	"golang.org/x/crypto/bcrypt"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-64.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(&Authorities{}, &Payment_status{}, &Prescription{},
		&MedicineLabel{}, &Suggestion{}, &Effect{},
		&MedicineRoom{}, &MedicineStorage{}, &MedicineType{}, &MedicineDisbursement{},
		&Packing{}, &ReceiveType{}, & Medicinereceive{},
	)

	db = database

	// Authority Data
	password1, err := bcrypt.GenerateFromPassword([]byte("86164091"), 14)
	db.Model(&Authorities{}).Create(&Authorities{
		AuthorityID: "A0001",
		FirstName:   "Chayodom",
		LastName:    "Heha",
		Email:       "chayodom@gmail.com",
		Password:    string(password1),
		TelNo:       "0911112233",
	})

	password2, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	db.Model(&Authorities{}).Create(&Authorities{
		AuthorityID: "A0002",
		FirstName:   "Demo",
		LastName:    "Account",
		Email:       "demo@gmail.com",
		Password:    string(password2),
		TelNo:       "0911112233",
	})

	//MedicineRoom
	// medicine1 := MedicineRoom{
	// 	Name:   "Paracetamol",
	// 	Amount: 1000,
	// 	Price: 2.0,
	// }
	// db.Model(&MedicineRoom{}).Create(&medicine1)

	// medicine2 := MedicineRoom{
	// 	Name:   "Chlorpheniramine",
	// 	Amount: 1000,
	// 	Price: 10.0,
	// }
	// db.Model(&MedicineRoom{}).Create(&medicine2)

	// medicine3 := MedicineRoom{
	// 	Name:   "Vitamin C",
	// 	Amount: 1000,
	// 	Price: 1.0,
	// }
	// db.Model(&MedicineRoom{}).Create(&medicine3)

	// medicine4 := MedicineRoom{
	// 	Name:   "Hydroxyzine",
	// 	Amount: 1000,
	// 	Price: 5.0,
	// }
	// db.Model(&MedicineRoom{}).Create(&medicine4)

	// medicine5 := MedicineRoom{
	// 	Name:   "Cetirizine",
	// 	Amount: 1000,
	// 	Price: 4.0,
	// }
	// db.Model(&MedicineRoom{}).Create(&medicine5)

	//Payment Status
	status1 := Payment_status{
		Status: "Not Paid",
	}
	db.Model(&Payment_status{}).Create(&status1)

	status2 := Payment_status{
		Status: "Paid",
	}
	db.Model(&Payment_status{}).Create(&status2)

	//Suggestion
	sug1 := Suggestion{
		SuggestionName: "ใช้ยาจนหมด",
	}
	db.Model(&Suggestion{}).Create(&sug1)

	sug2 := Suggestion{
		SuggestionName: "หยุดใช้เมื่อหาย",
	}
	db.Model(&Suggestion{}).Create(&sug2)

	//Effect
	effect1 := Effect{
		EffectName: "ทานแล้วอาจทำให้ง่วงซึม",
	}
	db.Model(&Effect{}).Create(&effect1)

	effect2 := Effect{
		EffectName: "ทำให้รู้สึกขมคอ",
	}
	db.Model(&Effect{}).Create(&effect2)

	//Medicinetype data
	Medicinetype1 := MedicineType{
		Name: "CAP",
	}
	db.Model(&MedicineType{}).Create(&Medicinetype1)
	Medicinetype2 := MedicineType{
		Name: "TAB",
	}
	db.Model(&MedicineType{}).Create(&Medicinetype2)
	Medicinetype3 := MedicineType{
		Name: "INJ",
	}
	db.Model(&MedicineType{}).Create(&Medicinetype3)

	//Medicinestorage data
	Medicinestorage1 := MedicineStorage{
		Name: "ASPIRIN",
		Count: 2000,
		MedicineType: Medicinetype2,
	}
	db.Model(&MedicineStorage{}).Create(&Medicinestorage1)

	Medicinestorage2 := MedicineStorage{
		Name: "GEMFIBROZIL",
		Count: 2800,
		MedicineType: Medicinetype1,
	}
	db.Model(&MedicineStorage{}).Create(&Medicinestorage2)

	//Medicineroom data
	Medicineroom1 := MedicineRoom{
		Name: "ASPIRIN",
		Amount: 100,
		Price: 35.00,
	}
	db.Model(&MedicineRoom{}).Create(&Medicineroom1)

	Medicineroom2 := MedicineRoom{
		Name: "GEMFIBROZIL",
		Amount: 100,
		Price: 54.20,
	}
	db.Model(&MedicineRoom{}).Create(&Medicineroom2)

	//Packing data
	Packing1 := Packing{
		Name: "กล่อง",
	}
	db.Model(&Packing{}).Create(&Packing1)

	Packing2 := Packing{
		Name: "ขวด",
	}
	db.Model(&Packing{}).Create(&Packing2)

	Packing3 := Packing{
		Name: "เม็ด",
	}
	db.Model(&Packing{}).Create(&Packing3)

	Receive1 := ReceiveType{
		Name: "โรงพยาบาล",
	}
	db.Model(&ReceiveType{}).Create(&Receive1)

	Receive2 := ReceiveType{
		Name: "บริษัท",
	}
	db.Model(&ReceiveType{}).Create(&Receive2)


}
