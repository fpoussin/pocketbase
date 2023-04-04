package logs

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/tools/migrate"
)

var LogsMigrations migrate.MigrationsList

func init() {
	LogsMigrations.Register(func(db dbx.Builder) (err error) {
		_, err = db.NewQuery(`
			CREATE TABLE {{_requests}} (
				[[id]]        TEXT PRIMARY KEY NOT NULL,
				[[url]]       TEXT DEFAULT '' NOT NULL,
				[[method]]    TEXT DEFAULT 'get' NOT NULL,
				[[status]]    INTEGER DEFAULT 200 NOT NULL,
				[[auth]]      TEXT DEFAULT 'guest' NOT NULL,
				[[ip]]        TEXT DEFAULT '127.0.0.1' NOT NULL,
				[[referer]]   TEXT DEFAULT '' NOT NULL,
				[[userAgent]] TEXT DEFAULT '' NOT NULL,
				[[meta]]      JSON DEFAULT '{}' NOT NULL,
				[[created]]   DATE DEFAULT CURRENT_DATE NOT NULL,
				[[updated]]   DATE DEFAULT CURRENT_DATE NOT NULL
			);

			CREATE INDEX _request_status_idx on {{_requests}} ([[status]]);
			CREATE INDEX _request_auth_idx on {{_requests}} ([[auth]]);
			CREATE INDEX _request_ip_idx on {{_requests}} ([[ip]]);
		`).Execute()

		return err
	}, func(db dbx.Builder) error {
		_, err := db.DropTable("_requests").Execute()
		return err
	})
}
