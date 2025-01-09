-- Create the Tournament table
CREATE TABLE tournaments ( 
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	slug TEXT UNIQUE NOT NULL,
	num_entrants INTEGER NOT NULL,
	started BOOLEAN
);

-- Create the Entrant table
CREATE TABLE entrants (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	seed INTEGER,
	name TEXT NOT NULL,
	standing INTEGER,
	tournament_id INTEGER NOT NULL,
	FOREIGN KEY (tournament_id) REFERENCES tournaments(id) ON DELETE CASCADE
);
