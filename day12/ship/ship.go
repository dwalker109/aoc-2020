package ship

type Ship interface {
	Forward(int)
	Left(int)
	Right(int)
	Move(string, int)
	Manhattan() int
}
