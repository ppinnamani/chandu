package daos

import (
	"database/sql"
	"errors"
	"github.com/ppinnamani/chandu/login/pkg/rest/server/daos/clients/sqls"
	"github.com/ppinnamani/chandu/login/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
)

type TestfieldsDao struct {
	sqlClient *sqls.SQLiteClient
}

func migrateTestfields(r *sqls.SQLiteClient) error {
	query := `
	CREATE TABLE IF NOT EXISTS testfields(
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
        
		Age TEXT NOT NULL,
		Password REAL NOT NULL,
		Usernawe TEXT NOT NULL,
        CONSTRAINT id_unique_key UNIQUE (Id)
	)
	`
	_, err1 := r.DB.Exec(query)
	return err1
}

func NewTestfieldsDao() (*TestfieldsDao, error) {
	sqlClient, err := sqls.InitSqliteDB()
	if err != nil {
		return nil, err
	}
	err = migrateTestfields(sqlClient)
	if err != nil {
		return nil, err
	}
	return &TestfieldsDao{
		sqlClient,
	}, nil
}

func (testfieldsDao *TestfieldsDao) CreateTestfields(m *models.Testfields) (*models.Testfields, error) {
	insertQuery := "INSERT INTO testfields(Age, Password, Usernawe)values(?, ?, ?)"
	res, err := testfieldsDao.sqlClient.DB.Exec(insertQuery, m.Age, m.Password, m.Usernawe)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	m.Id = id

	log.Debugf("testfields created")
	return m, nil
}

func (testfieldsDao *TestfieldsDao) UpdateTestfields(id int64, m *models.Testfields) (*models.Testfields, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	testfields, err := testfieldsDao.GetTestfields(id)
	if err != nil {
		return nil, err
	}
	if testfields == nil {
		return nil, sql.ErrNoRows
	}

	updateQuery := "UPDATE testfields SET Age = ?, Password = ?, Usernawe = ? WHERE Id = ?"
	res, err := testfieldsDao.sqlClient.DB.Exec(updateQuery, m.Age, m.Password, m.Usernawe, id)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, sqls.ErrUpdateFailed
	}

	log.Debugf("testfields updated")
	return m, nil
}

func (testfieldsDao *TestfieldsDao) DeleteTestfields(id int64) error {
	deleteQuery := "DELETE FROM testfields WHERE Id = ?"
	res, err := testfieldsDao.sqlClient.DB.Exec(deleteQuery, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sqls.ErrDeleteFailed
	}

	log.Debugf("testfields deleted")
	return nil
}

func (testfieldsDao *TestfieldsDao) ListTestfields() ([]*models.Testfields, error) {
	selectQuery := "SELECT * FROM testfields"
	rows, err := testfieldsDao.sqlClient.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	var testfields []*models.Testfields
	for rows.Next() {
		m := models.Testfields{}
		if err = rows.Scan(&m.Id, &m.Age, &m.Password, &m.Usernawe); err != nil {
			return nil, err
		}
		testfields = append(testfields, &m)
	}
	if testfields == nil {
		testfields = []*models.Testfields{}
	}

	log.Debugf("testfields listed")
	return testfields, nil
}

func (testfieldsDao *TestfieldsDao) GetTestfields(id int64) (*models.Testfields, error) {
	selectQuery := "SELECT * FROM testfields WHERE Id = ?"
	row := testfieldsDao.sqlClient.DB.QueryRow(selectQuery, id)
	m := models.Testfields{}
	if err := row.Scan(&m.Id, &m.Age, &m.Password, &m.Usernawe); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}

	log.Debugf("testfields retrieved")
	return &m, nil
}
