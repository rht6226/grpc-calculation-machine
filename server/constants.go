package server

type OperatorType string

const (
	PUSH OperatorType = "PUSH"
	POP  OperatorType = "POP"
	ADD  OperatorType = "ADD"
	SUB  OperatorType = "SUB"
	MUL  OperatorType = "MUL"
	DIV  OperatorType = "DIV"
	FIB  OperatorType = "FIB"
)
