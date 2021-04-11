package main

    import (
        "database/sql"
		"fmt"
		"sync"
		"log"
        _ "github.com/mattn/go-sqlite3"
    )


	var lock = &sync.Mutex{}

	type Manager interface {
		getAll() userInfoList
		closeConnection() error
	}

	type single struct {
		db *sql.DB
	}


	type userInfo struct {
		Id     string `json:"id"`
		Username   string `json:"username"`
		Departname string `json: "departname"`
	}

	type userInfoList struct {
		UserinfoList []userInfo `json:"userinfoList"`
	}
	
	
	var Mgr Manager
	
	func getInstance() Manager {
		if Mgr == nil {
			lock.Lock()
			defer lock.Unlock()
			if Mgr == nil {
				fmt.Println("CREATING DB")
				db, err := sql.Open("sqlite3", "./Database.db")
				if err != nil {
					fmt.Println("Failed to init db:", err)
				}
				Mgr = &single{db: db}
			} else {
				fmt.Println("Single Instance already created-1")
			}
		} else {
			fmt.Println("Single Instance already created-2")
		}
		return Mgr
	}


	func (mgr *single) getAll() userInfoList {
		//selectAll
		rows, err := mgr.db.Query("SELECT * FROM userinfo")
		if err != nil {
			fmt.Println(err)
		}
		defer rows.Close()
		result := userInfoList{}

		for rows.Next() {
			userinfo := userInfo{}
			//var username, departname string
			//var id int
			if err := rows.Scan(&userinfo.Id, &userinfo.Username, &userinfo.Departname); err != nil {
				log.Fatal(err)
			}
			result.UserinfoList = append(result.UserinfoList, userinfo)
			fmt.Printf("id: %s\nusername: %s \ndepartname: %s \n \n---\n\n", userinfo.Id, userinfo.Username, userinfo.Departname)
		} 

		return result
	}

	func (mgr *single) closeConnection() (err error) {
		lock.Lock()
		defer lock.Unlock()
		if Mgr != nil {
			mgr.db.Close()
			Mgr = nil
			fmt.Println("Closed connection With DB")
		}
		return
	}




    func checkErr(err error) {
        if err != nil {
            panic(err)
        }
    }