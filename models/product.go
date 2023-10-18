package models

type Product struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

//contoh modeling data
// type User struct {
// 	ID           uint
// 	Name         string
// 	Email        *string
// 	Age          uint8
// 	Birthday     *time.Time
// 	MemberNumber sql.NullString
// 	ActivatedAt  sql.NullTime
// 	CreatedAt    time.Time
// 	UpdatedAt    time.Time
//   }

//=======================================
// gorm.Model definition
// type Model struct {
// 	ID        uint           `gorm:"primaryKey"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
//   }

//=======================================
// type User struct {
// 	Name string `gorm:"<-:create"` // allow read and create
// 	Name string `gorm:"<-:update"` // allow read and update
// 	Name string `gorm:"<-"`        // allow read and write (create and update)
// 	Name string `gorm:"<-:false"`  // allow read, disable write permission
// 	Name string `gorm:"->"`        // readonly (disable write permission unless it configured)
// 	Name string `gorm:"->;<-:create"` // allow read and create
// 	Name string `gorm:"->:false;<-:create"` // createonly (disabled read from db)
// 	Name string `gorm:"-"`            // ignore this field when write and read with struct
// 	Name string `gorm:"-:all"`        // ignore this field when write, read and migrate with struct
// 	Name string `gorm:"-:migration"`  // ignore this field when migrate with struct
//   }

//=======================================
// type User struct {
// 	CreatedAt time.Time // Set to current time if it is zero on creating
// 	UpdatedAt int       // Set to current unix seconds on updating or if it is zero on creating
// 	Updated   int64 `gorm:"autoUpdateTime:nano"` // Use unix nano seconds as updating time
// 	Updated   int64 `gorm:"autoUpdateTime:milli"`// Use unix milli seconds as updating time
// 	Created   int64 `gorm:"autoCreateTime"`      // Use unix seconds as creating time
//   }

//=======================================
// type Author struct {
// 	Name  string
// 	Email string
//   }

//   type Blog struct {
// 	ID      int
// 	Author  Author `gorm:"embedded"`
// 	Upvotes int32
//   }
//   // equals
//   type Blog struct {
// 	ID    int64
// 	Name  string
// 	Email string
// 	Upvotes  int32
//   }
