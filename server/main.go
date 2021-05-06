package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"

	"github.com/jackc/pgx/v4"
	grpc "google.golang.org/grpc"
	//"github.com/grpc/grpc-go"
)

// type Cat struct {
// 	Id    int
// 	Name  string
// 	Age   int
// 	Color string
// 	// []Images
// }

type CatsExplorerServerManager struct {
	dbConn *pgx.Conn
}

func NewCatsExplorerServerManager(dbConn *pgx.Conn) *CatsExplorerServerManager {

	return &CatsExplorerServerManager{
		dbConn: dbConn,
	}
}

func main() {

	urlPostgreS := "postgres://postgres:pgpass@localhost:5432"
	// connect string from the docker container
	//urlPostgreS := "postgres://postgres:pgpass@PostgresDB2:5432"
	conn, err := pgx.Connect(context.Background(), urlPostgreS)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	h := NewCatsExplorerServerManager(conn)

	h.PrepareTableIfNotExist()

	listenAddr := ":1323"

	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		fmt.Println("cant listen port", err)
		panic(err)
	}

	server := grpc.NewServer()
	RegisterCatsExplorerServer(server, h)

	fmt.Println("starting server at " + listenAddr)
	server.Serve(lis)

	// e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })

	// e.POST("/cats", h.UpdateCat)
	// e.PUT("/cats", h.CreateNewCat)
	// e.DELETE("/cats/:id", h.DeleteCat)
	// e.GET("/cats/:id", h.GetCat)
	// e.GET("/cats", h.GetAllCats)
	// e.GET("/cats/", h.GetAllCats)

	// e.Logger.Fatal(e.Start(":1323"))
}

func (h *CatsExplorerServerManager) PrepareTableIfNotExist() {
	qsCreateTable :=
		`CREATE TABLE IF NOT EXISTS cats (
				id SERIAL PRIMARY KEY,
				name VARCHAR ( 50 ) NOT NULL,
				age INT,
				color VARCHAR ( 50 )
			);`

	qsSelectInitial := "SELECT id, name, age, color FROM cats"
	qsInit :=
		`INSERT INTO cats (id, name, age, color) VALUES
	(DEFAULT, 'Dar', 6,	'black'),
	(DEFAULT, 'Fil',	1,	'black');`

	_, err := h.dbConn.Exec(context.Background(), qsCreateTable)
	if err != nil {
		panic(err)
	}

	cat := new(Cat)
	err = h.dbConn.QueryRow(context.Background(), qsSelectInitial).Scan(&cat.Id, &cat.Name, &cat.Age, &cat.Color)
	if err != nil {
		if err.Error() == "no rows in result set" {
			_, err = h.dbConn.Exec(context.Background(), qsInit)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

}

func (h *CatsExplorerServerManager) GetCat(ctx context.Context, idStruct *Id) (*Cat, error) {

	cat := new(Cat)

	id := idStruct.Id
	qs := "SELECT id, name, age, color FROM cats WHERE id=$1"

	//var result []byte
	err := h.dbConn.QueryRow(context.Background(), qs, id).Scan(&cat.Id, &cat.Name, &cat.Age, &cat.Color)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, errors.New("Cat was not found")
		} else {
			return nil, errors.New("QueryRow error: " + err.Error())
		}
	} else {
		// result, err = json.Marshal(cat)
		// if err != nil {
		// 	return nil, errors.New("Marshal error: " + err.Error())
		// }
		return cat, nil
	}

	//return c.JSONBlob(http.StatusOK, result)
}

func (h *CatsExplorerServerManager) GetAllCats(ctx context.Context, dummy *Nothing) (*Cats, error) {

	catsSlice := []*Cat{}

	qs := "select id, name, age, color from cats ORDER BY id"

	rows, err := h.dbConn.Query(context.Background(), qs)
	if err != nil {
		return nil, errors.New("Query error: " + err.Error())
	}

	for rows.Next() {

		cat := new(Cat)
		err := rows.Scan(&cat.Id, &cat.Name, &cat.Age, &cat.Color)
		if err != nil {
			return nil, errors.New("Scan error: " + err.Error())
		}
		catsSlice = append(catsSlice, cat)
	}

	cats := new(Cats)
	cats.Cats = catsSlice

	return cats, nil

	// result, err := json.Marshal(cats)
	// if err != nil {
	// 	return c.String(http.StatusInternalServerError, fmt.Sprintf("Marshal error: %v", err.Error()))
	// }

	//return c.JSONBlob(http.StatusOK, result)
}

func (h *CatsExplorerServerManager) UpdateCat(ctx context.Context, cat *Cat) (*Info, error) {

	// r := c.Request()
	// err := r.ParseForm()
	// if err != nil {
	// 	return c.String(http.StatusInternalServerError, fmt.Sprintf("ParseForm error: %v", err.Error()))
	// }
	//var updateCat Cat
	//err = json.NewDecoder(r.Body).Decode(&updateCat)
	//err = c.Bind(&updateCat)
	// if err != nil {
	// 	return c.String(http.StatusInternalServerError, fmt.Sprintf("Bind error: %v", err.Error()))
	// }

	resultInfo := new(Info)

	result, err := h.dbConn.Exec(context.Background(), "UPDATE cats SET name = $1, age = $2, color = $3 WHERE id = $4", cat.Name, cat.Age, cat.Color, cat.Id)
	if err != nil {
		resultInfo.Info = "Internal Server Error"
		return resultInfo, errors.New("Scan error: " + err.Error())
	}

	affected := result.RowsAffected()
	resultInfo.Info = fmt.Sprintf("Affected rows: %v", affected)
	return resultInfo, nil
}

func (h *CatsExplorerServerManager) CreateNewCat(ctx context.Context, cat *Cat) (*Info, error) {

	// r := c.Request()
	// err := r.ParseForm()
	// if err != nil {
	// 	return c.String(http.StatusInternalServerError, fmt.Sprintf("ParseForm error: %v", err.Error()))
	// }
	// var updateCat Cat
	// err = json.NewDecoder(r.Body).Decode(&updateCat)
	// if err != nil {
	// 	return c.String(http.StatusInternalServerError, fmt.Sprintf("Decode error: %v", err.Error()))
	// }

	resultInfo := new(Info)

	result, err := h.dbConn.Exec(context.Background(), "INSERT INTO cats VALUES (DEFAULT, $1, $2, $3)", cat.Name, cat.Age, cat.Color)
	if err != nil {
		resultInfo.Info = "Internal Server Error"
		return resultInfo, errors.New("Exec error: " + err.Error())
	}

	affected := result.RowsAffected()
	resultInfo.Info = fmt.Sprintf("Affected rows: %v", affected)
	return resultInfo, nil
}

func (h *CatsExplorerServerManager) DeleteCat(ctx context.Context, id *Id) (*Info, error) {

	//id := c.Param("id")

	resultInfo := new(Info)

	result, err := h.dbConn.Exec(context.Background(), "DELETE FROM cats WHERE id = $1", id)
	if err != nil {
		resultInfo.Info = "Internal Server Error"
		return resultInfo, errors.New("Exec error: " + err.Error())
	}

	affected := result.RowsAffected()
	resultInfo.Info = fmt.Sprintf("Affected rows: %v", affected)
	return resultInfo, nil
}

func (h *CatsExplorerServerManager) mustEmbedUnimplementedCatsExplorerServer() {}
