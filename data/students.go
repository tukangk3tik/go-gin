package data

type Student struct {
	StudentId string `json:"student_id"`
	FullName  string `json:"full_name"`
	Address   string `json:"address"`
	Gender    string `json:"gender"`
}

var AllStudents = []Student{
	{
		StudentId: "123456",
		FullName:  "Budi Santoso",
		Address:   "Jl. By Pass No. 10",
		Gender:    "Laki-laki",
	},
	{
		StudentId: "87234",
		FullName:  "Nadia Smith",
		Address:   "Jl. Madio Santoso No. 10",
		Gender:    "Perempuan",
	},
	{
		StudentId: "44678",
		FullName:  "Iwan Dani",
		Address:   "Jl. Karya No. 6",
		Gender:    "Laki-laki",
	},
}
