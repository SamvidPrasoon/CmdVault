package store

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	bolt "go.etcd.io/bbolt"
)

type Command struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Cmd         string    `json:"cmd"`
	Description string    `json:"description"`
	Tags        []string  `json:"tags"`
	CreatedAt   time.Time `json:"created_at"`
	RunCount    int       `json:"run_count"`
}
type DbClient struct {
	client *bolt.DB
}

var bucketName = []byte("commands")

func NewDbClient(path string) (*DbClient, error) {
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}
	//create bucket if not exists
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucketName)
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	return &DbClient{
		client: db,
	}, nil
}

// save command to db
func (db *DbClient) Save(cmd Command) error {
	return db.client.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		data, err := json.Marshal(cmd)
		if err != nil {
			return err
		}
		return b.Put([]byte(cmd.Name), data)
	})
}

//Get cmd

func (db *DbClient) GetCmd(name string) (*Command, error) {
	var c Command
	err := db.client.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		data := b.Get([]byte(name))
		if data == nil {
			return fmt.Errorf("command '%s' not found", name)
		}
		return json.Unmarshal(data, &c)
	})
	return &c, err
}

// list all cmds
func (db *DbClient) All() ([]Command, error) {
	var cmds []Command
	err := db.client.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		return b.ForEach(func(k, v []byte) error {
			var c Command
			if err := json.Unmarshal(v, &c); err != nil {
				return err
			}
			cmds = append(cmds, c)
			return nil
		})
	})
	return cmds, err
}

// Search  commands by name, description, or tag
func (db *DbClient) Search(query string) ([]Command, error) {
	all, err := db.All()
	if err != nil {
		return nil, err
	}
	query = strings.ToLower(query)
	var results []Command
	for _, c := range all {
		if strings.Contains(strings.ToLower(c.Name), query) ||
			strings.Contains(strings.ToLower(c.Description), query) ||
			containsTag(c.Tags, query) {
			results = append(results, c)
		}
	}
	return results, nil
}

// Delete removes a command by name
func (db *DbClient) Delete(name string) error {
	return db.client.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(bucketName).Delete([]byte(name))
	})
}

// IncrementRunCount tracks how many times a command was run
func (db *DbClient) IncrementRunCount(name string) {
	c, err := db.GetCmd(name)
	if err != nil {
		return
	}
	c.RunCount++
	db.Save(*c)
}

func containsTag(tags []string, query string) bool {
	for _, t := range tags {
		if strings.Contains(strings.ToLower(t), query) {
			return true
		}
	}
	return false
}

func (db *DbClient) Ingest(filePath string) (int, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return 0, fmt.Errorf("failed to read file: %w", err)
	}

	var cmds []Command
	if err := json.Unmarshal(data, &cmds); err != nil {
		return 0, fmt.Errorf("failed to parse JSON: %w", err)
	}

	ingested := 0
	for _, cmd := range cmds {
		if err := db.Save(cmd); err != nil {
			return ingested, fmt.Errorf("failed to save command '%s': %w", cmd.Name, err)
		}
		ingested++
	}

	return ingested, nil
}
