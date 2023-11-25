package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/sqlite3"
)

// Creates a Fiber store for keeping track of
// sessions data.
func NewSessionStore(sessionFilePath string) *session.Store {
	store := session.New(session.Config{
		Storage: sqlite3.New(sqlite3.Config{
			Database: sessionFilePath,
		}),
		Expiration:     time.Hour * 1,
		CookieHTTPOnly: true,
	})

	return store
}

func (core *Core) ManageSession(c *fiber.Ctx) error {
	sess, err := core.Store.Get(c)
	if err != nil {
		core.Logger.Info("store.get",
			"err", err,
		)
		return err
	}

	c.Locals(SESSION_KEY, sess)

	err = c.Next()
	if err != nil {
		return err
	}

	err = sess.Save()
	if err != nil {
		core.Logger.Error("session.save",
			"err", err,
		)
		panic(err)
	}

	return nil
}
