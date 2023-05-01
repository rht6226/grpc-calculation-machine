package server

import (
	"io"
	"log"

	"github.com/rht6226/grpc-machine/machine"
	"github.com/rht6226/grpc-machine/pkg/stack"
	"github.com/rht6226/grpc-machine/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MachineServer struct {
	machine.UnimplementedMachineServer
}

func (s *MachineServer) Execute(stream machine.Machine_ExecuteServer) error {
	var stack stack.Stack

	for {
		instruction, err := stream.Recv()
		if err == io.EOF {
			log.Println("EOF")
			return nil
		}
		if err != nil {
			return err
		}

		operand := instruction.GetOperand()
		operator := instruction.GetOperator()
		opType := OperatorType(operator)

		log.Printf("Executing Instruction: Operand: %v, Operator: %v", operand, operator)

		switch opType {
		case PUSH:
			stack.Push(float32(operand))
		case POP:
			stack.Pop()
		case ADD, SUB, MUL, DIV:
			item2, popped1 := stack.Pop()
			item1, popped2 := stack.Pop()

			if !popped1 || !popped2 {
				return status.Error(codes.Aborted, "Invalide sets of instructions. Execution aborted")
			}

			var res float32
			if opType == ADD {
				res = item1 + item2
			} else if opType == SUB {
				res = item1 - item2
			} else if opType == MUL {
				res = item1 * item2
			} else if opType == DIV {
				res = item1 / item2
			}

			stack.Push(res)
			if err := stream.Send(&machine.Result{Output: float32(res)}); err != nil {
				return err
			}

		case FIB:
			n, popped := stack.Pop()

			if !popped {
				return status.Error(codes.Aborted, "Invalid sets of instructions. Execution aborted")
			}

			if opType == FIB {
				for f := range utils.FibonacciRange(int(n)) {
					if err := stream.Send(&machine.Result{Output: float32(f)}); err != nil {
						return err
					}
				}
			}

		default:
			return status.Errorf(codes.Unimplemented, "Operation '%s' not implemented yet", operator)
		}
	}
}
