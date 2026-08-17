package birthDay

type Birth struct {
	Name string
	Age  int
}

func AddOne(br *Birth) {
	br.Age = br.Age + 1
}
