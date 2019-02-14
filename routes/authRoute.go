package routes

func AuthRoute() *Routes{
	return Routes{
		Route{
			"Index",
			"GET",
			"/",
			Index
		},
		
	}
	
}