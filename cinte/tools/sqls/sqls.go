// Package sqls implements SQLite schema and queries.
package sqls

// Pragma is the per-connection database pragma.
const Pragma = `
	pragma busy_timeout = 5000;
	pragma foreign_keys = true;
	pragma synchronous  = normal;
`

// Schema is the default first-run database schema.
const Schema = `
	create table if not exists Notes (
		id   integer primary key asc,
		init integer not null default (unixepoch()),
		name text    not null,

		unique(name)
	);

	create table if not exists Pages (
		id   integer primary key asc,
		init integer not null default (unixepoch()),
		note integer not null,
		body text    not null,

		foreign key (note) references Notes(id) on delete cascade
	);

	create index if not exists PageNotes on Pages(note);
`
