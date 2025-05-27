package tables

import (
	"fmt"
	"migrator/configs"
	"os"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx"
)

func readDataFromDB[D any](DB *sqlx.DB, tablename, query string) ([]D, error) {
	values := []D{}

	if query == "" {
		query = fmt.Sprintf("SELECT * FROM %s", tablename)
	}

	err := DB.Select(&values, query)
	if err != nil {
		return nil, err
	}
	return values, nil
}

func readQueryFromFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	query := string(data)
	return query, nil
}

type TableRepositoryInterface[T any] interface {
	ReadData() error
	WriteData() error
}

type TableRepository[T any] struct {
	DB     *sqlx.DB
	Folder string
	Name   string
	Query  string
	Schema T
	Data   []T
}

func (t *TableRepository[T]) ReadData() error {

	data, err := readDataFromDB[T](t.DB, t.Name, t.Query)
	if err != nil {
		return fmt.Errorf("error reading data from table %s: %w", t.Name, err)
	}

	t.Data = data
	return nil
}

func (t *TableRepository[T]) ToString() string {

	linesList := []string{}
	for _, d := range t.Data {
		lineList := []string{}
		valueOf := reflect.ValueOf(d)

		for i := range valueOf.NumField() {
			lineList = append(lineList, fmt.Sprintf("%v", valueOf.Field(i)))
		}

		line := strings.Join(lineList, ";")
		linesList = append(linesList, line)
	}
	lines := strings.Join(linesList, "\n")

	headerList := []string{}
	fieldOf := reflect.TypeOf(t.Data[0])
	for i := range fieldOf.NumField() {
		headerList = append(headerList, fmt.Sprintf("%v", fieldOf.Field(i).Name))
	}
	header := strings.Join(headerList, ";")
	result := header + "\n" + lines
	return result
}

func (c *TableRepository[T]) WriteData() error {

	_, err := os.Stat(c.Folder)
	if os.IsNotExist(err) {
		os.MkdirAll(c.Folder, 0744)
	}

	path := fmt.Sprintf("%s/%s.csv", c.Folder, c.Name)
	return os.WriteFile(path, []byte(c.ToString()), 0744)
}

func NewTableRepository[T any](name, schema string, db *sqlx.DB, config *configs.Config) (*TableRepository[T], error) {

	queryPath := fmt.Sprintf("%s/%s/%s.sql", config.QueryBaseFolder, schema, name)
	query, _ := readQueryFromFile(queryPath)

	folder := fmt.Sprintf("%s/%s", config.DataBaseFolder, schema)

	repo := &TableRepository[T]{
		DB:     db,
		Name:   name,
		Query:  query,
		Schema: *new(T),
		Data:   []T{},
		Folder: folder,
	}

	return repo, nil
}
