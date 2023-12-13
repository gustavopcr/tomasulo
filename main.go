package main
import("fmt")


//registers := make([]Register, 10)
var registerMap = make(map[string]Register, 10)

var registers [10]Register
func main(){
	//clock := 0
	instQueue := make([]Instruction, 0)
	reservationStation := make([]ReservationStation, 0, 6)

	inst := decode("SUB R1 R2 R3")
	instQueue = append(instQueue, inst)
	if(len(reservationStation)<cap(reservationStation)){
		fmt.Println("instQueue: ", instQueue )

		if(len(instQueue)>0){
			i := instQueue[0]
			//reservationStation = append(reservationStation, instQueue[0])
			rs := ReservationStation{
				Op:       i.Opcode,
				Rd:	i.DestinationReg,
				Rj: i.SourceReg1,
				Rk: i.SourceReg2,
				RemainingCycles: i.Cycles,
			}

			reservationStation = append(reservationStation, rs);
			instQueue = instQueue[1:]
		}

		fmt.Println("instQueue: ", instQueue )
		fmt.Println("reservationStation: " , reservationStation[0])
		emission( reservationStation[0] )
	}
}