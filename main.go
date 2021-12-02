package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

type Mahasiswa struct {
	Nama     string  `redis:"nama"`
	NIM      string  `redis:"nim"`
	IPK      float64 `redis:"ipk"`
	Semester int     `redis:"semester"`
}

func main() {
	// conn, err := redis.Dial("tcp", "localhost:6379")
	// if err != nil {
	// 	log.Panic(err)
	// }

	pool := redis.NewPool(
		func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		}, 0,
	)

	pool.MaxActive = 0
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("HSET", "mahasiswa:1", "nama", "Redha Juanda", "nim", "12345", "ipk", 3.34, "semester", 4)
	if err != nil {
		log.Panic(err)
	}

	// mengambil data dengan tipe string
	nama, err := redis.String(conn.Do("HGET", "mahasiswa:1", "nama"))
	if err != nil {
		log.Panic(err)
	}
	// mengambil data dengan tipe string
	nim, err := redis.String(conn.Do("HGET", "mahasiswa:1", "nim"))
	if err != nil {
		log.Panic(err)
	}
	// mengambil data dengan tipe float
	ipk, err := redis.Float64(conn.Do("HGET", "mahasiswa:1", "ipk"))
	if err != nil {
		log.Panic(err)
	}
	// mengambil data dengan tipe integer
	semester, err := redis.Int(conn.Do("HGET", "mahasiswa:1", "semester"))
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(nama)
	fmt.Println(nim)
	fmt.Println(ipk)
	fmt.Println(semester)

	// mengambil semua data berdasarkan id mahasiswa
	resp, err := redis.StringMap(conn.Do("HGETALL", "mahasiswa:1"))
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(resp)

	//assign ke struct
	// mengambil semua data berdasarkan id mahasiswa
	// HGETALL mengembalikan semua field yang ada pada objek tsb
	// redis.Values adalah reply helper yang mengconvert data
	// reply dengan tipe interface ke tipe []interface{}
	rep, err := redis.Values(conn.Do("HGETALL", "mahasiswa:1"))
	if err != nil {
		log.Panic(err)
	}

	mahasiswa := Mahasiswa{}
	err = redis.ScanStruct(rep, &mahasiswa)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%v", mahasiswa)
}
