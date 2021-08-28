package repo

import (
	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo"
	"github.com/ozonva/ova-entertainment-api/internal/models"
	"github.com/stretchr/testify/assert"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
	"regexp"
)

var _ = Describe("Repo", func() {

	var (
		db   *sqlx.DB
		mock sqlxmock.Sqlmock
		err  error
	)

	BeforeEach(func() {
		db, mock, err = sqlxmock.Newx()
		if err != nil {
			GinkgoT().Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
	})

	AfterEach(func() {
		db.Close()
	})

	Context("Repo remove v1", func() {
		It("should not error", func() {

			ID := uint64(1)
			result := sqlxmock.NewResult(1, 1)
			mock.
				ExpectExec(regexp.QuoteMeta(`DELETE FROM entertainments WHERE id = $1`)).
				WithArgs(ID).
				WillReturnResult(result)

			r := NewRepo(db)
			err = r.RemoveEntertainment(ID)

			assert.Nil(GinkgoT(), err)
		})
	})

	Context("Repo create v1", func() {
		It("should not error", func() {

			model := models.New(1, "Title", "Description")

			result := sqlxmock.NewResult(1, 1)
			mock.
				ExpectExec(regexp.QuoteMeta(`INSERT INTO entertainments (user_id, title, description) VALUES (?, ?, ?)`)).
				WithArgs(model.UserID, model.Title, model.Description).
				WillReturnResult(result)

			r := NewRepo(db)
			err = r.AddEntertainments([]models.Entertainment{model})

			assert.Nil(GinkgoT(), err)
		})
	})

	Context("Repo describe v1", func() {
		It("should not error", func() {

			model := models.New(1, "Title", "Description")
			result := sqlxmock.NewResult(1, 1)
			mock.
				ExpectExec(regexp.QuoteMeta(`UPDATE entertainments SET title = ?, description = ? WHERE id = ?`)).
				WithArgs(model.Title, model.Description, model.ID).
				WillReturnResult(result)

			r := NewRepo(db)
			_, err = r.DescribeEntertainment(model)

			assert.Nil(GinkgoT(), err)
		})
	})

	Context("Repo list v1", func() {
		It("should not error", func() {

			rows := sqlxmock.NewRows([]string{"id", "user_id", "title", "description"}).AddRow(1, 1, "Title", "Description")
			mock.
				ExpectQuery(regexp.QuoteMeta(`SELECT * FROM entertainments ORDER BY id DESC LIMIT $1 OFFSET $2`)).
				WithArgs(100, 10).
				WillReturnRows(rows)

			r := NewRepo(db)
			_, err = r.ListEntertainments(100, 10) // @todo _

			assert.Nil(GinkgoT(), err)
		})
	})
})
