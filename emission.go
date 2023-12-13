package main;

import("fmt")

func emission( rs ReservationStation ){
	fmt.Println("emission")
	
	if(rs.Vj == "" && registerMap[rs.Rj].Status == Ready){
		rs.Vj = registerMap[rs.Rj].Value

		reg := registerMap[rs.Rj]
		reg.Status = Busy
		registerMap[rs.Rj] = reg
	}
	if(rs.Vk == "" && registerMap[rs.Rk].Status == Ready){
		rs.Vk = registerMap[rs.Rk].Value

		reg := registerMap[rs.Rk]
		reg.Status = Busy
		registerMap[rs.Rk] = reg
	}
}