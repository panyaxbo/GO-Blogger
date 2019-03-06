package dbclient

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/panyaxbo/goblogger/accountservice/model"
)

type IBoltClient interface {
	OpenBoltDb()
	QueryAccount(account string) (model.Account, error)
	Seed()
}
type BoltClient struct {
	boltDB *bolt.DB
}

func (bc *BoltClient) QueryAccount(accountId string) (model.Account, error) {
	// panic("implement me")
	// Allocate an empty Account instance we'll let json.Unmarshal populate for us in a bit.
	account := model.Account{}

	err := bc.boltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("AccountBucket"))
		accountBytes := b.Get([]byte(accountId))
		if accountBytes == nil {
			return fmt.Errorf("No account found for "+ accountId)
		}
		// Unmarshal the return bytes into the account struct we create at
		// the top of the functions
		json.Unmarshal(accountBytes, &account)
		// Return nil to indicate nothing went wrong, e.g no error
		return nil
	})
	if err != nil {
		return model.Account{}, err
	}
	// Return the Account struct and nil as error.
	return account, nil
}

func (bc *BoltClient) OpenBoltDb() {
	var err error
	bc.boltDB, err = bolt.Open("accounts.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Creates an "AccountBucket" in our BoltDB. It will overwrite any existing bucket of the same name.
func (bc *BoltClient) initializeBucket() {
	bc.boltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("AccountBucket"))
		if err != nil {
			return fmt.Errorf("create bucket failed: %s", err)
		}
		return nil
	})
}

// Seed (n) make-believe account objects into the AcountBucket bucket.
func (bc *BoltClient) seedAccounts() {
	total := 100
	for i := 0; i < total; i++ {
		// Generate a key 10000 or larger
		key := strconv.Itoa(10000 + i)
		// Create an instance of our Account struct
		acc := model.Account{
			Id:   key,
			Name: "Person_" + strconv.Itoa(i),
		}
		// Serialize the struct to JSON
		jsonBytes, _ := json.Marshal(acc)
		// Write the data to the AccountBucket
		bc.boltDB.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("AccountBucket"))
			err := b.Put([]byte(key), jsonBytes)
			return err
		})
	}
	fmt.Printf("Seeded %v fake accounts...\n", total)
}
func (bc *BoltClient) Seed() {
	//initializeBucket()
	//seedAccounts()
	bc.initializeBucket()
	bc.seedAccounts()
}
