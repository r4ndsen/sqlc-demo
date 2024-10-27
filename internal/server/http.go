package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/google/uuid"
	"github.com/r4ndsen/sqlc-demo/internal/db"
	"log"
)

type Http struct {
	store db.Querier
	port  int
}

func New(store db.Querier, port int) *Http {
	return &Http{
		store: store,
		port:  port,
	}
}

func (s *Http) Start() {
	engine := html.New("./assets/views", ".html")
	engine.Reload(true)

	app := fiber.New(
		fiber.Config{
			Views:                 engine,
			DisableStartupMessage: true,
		},
	)

	app.Static("/", "./assets")

	app.Get("/", func(c *fiber.Ctx) error {
		entities, err := s.store.ListLinks(c.Context())

		if err != nil {
			return c.SendString(err.Error())
		}

		return c.Render("index", fiber.Map{
			"entities": entities,
		}, "layouts/main")
	})

	app.Delete("/delete/:id", func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))

		if err != nil {
			return err
		}

		_, err = s.store.DeleteLinkById(c.Context(), id)

		return err
	})

	log.Printf("server started. listening on port: %d", s.port)
	log.Fatalln(app.Listen(fmt.Sprintf(":%d", s.port)))
}
