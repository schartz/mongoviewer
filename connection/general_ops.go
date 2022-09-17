package connection

import (
	"encoding/json"
	bolt "go.etcd.io/bbolt"
)

func GetConnectionList(db *bolt.DB) ([]map[string]string, error) {
	var connectionList []map[string]string
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Connections"))
		cursor := b.Cursor()
		for key, value := cursor.First(); key != nil; key, value = cursor.Next() {

			connectionList = append(connectionList, map[string]string{"name": string(key), "uri": string(value)})

		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return connectionList, nil

}

func AddConnection(connString string, db *bolt.DB) bool {
	var connection map[string]string
	err := json.Unmarshal([]byte(connString), &connection)
	if err != nil {
		return false
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("Connections"))
		if err != nil {
			return err
		}

		if err := b.Put([]byte(connection["name"]), []byte(connection["uri"])); err != nil {
			appLogger.Fatal(err)
		}

		return nil
	})
	if err != nil {
		return false
	}
	return true
}
