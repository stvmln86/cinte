// Package sqls implements SQLite schema and queries.
package sqls

// Pragma is the default always-on database pragma.
const Pragma = `
	pragma encoding = 'utf-8';
	pragma foreign_keys = on;
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

		foreign key (note) references Notes(id)
	);

	create index if not exists NoteNames on Notes(name);
	create index if not exists PageNotes on Pages(note);
`
