package main

import (
 db "myblog/database"
 r "myblog/router"
)

func main() {
	 defer db.SqlDB.Close()
	 router := r.InitRouter()
	 router.Run(":8000")
}

