package helpers

import (
	"crypto/tls"
	"log"
)

func TlsTcpConn() (*tls.Conn, error) {
	log.SetFlags(log.Lshortfile)

	conf := &tls.Config{
		InsecureSkipVerify:          true,
		DynamicRecordSizingDisabled: true,
	}

	conn, err := tls.Dial("tcp", "127.0.0.1:4444", conf)
	if err != nil {
		log.Println("tls connection error", err)
		// log.Fatal("tls connection error", err)

		var tr *tls.Conn

		return tr, err
	}
	// defer conn.Close()

	// n, err := conn.Write([]byte("hello\n"))
	// if err != nil {
	// 	log.Println(n, err)
	// 	return
	// }

	// buf := make([]byte, 100)
	// n, err = conn.Read(buf)
	// if err != nil {
	// 	log.Println(n, err)
	// 	return
	// }

	// println(string(buf[:n]))

	return conn, nil
}
