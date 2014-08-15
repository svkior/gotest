package setups


type setups struct {
	name string
}

func Create() (setups){
	var st setups
	st.name = "Setups"
	return st
}
