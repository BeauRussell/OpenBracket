table "tournaments" {
  schema = schema.main
  column "id" {
    null = false
    type = integer
    auto_increment = true
  }
  column "name" {
    null = false
    type = text
  }
  column "slug" {
    null = false
    type = text
  }
  column "num_entrants" {
    null = false
    type = integer
  }
  column "started" {
    null = true
    type = boolean
  }
  primary_key {
    columns = [column.id]
  }
  index "slug" {
    columns = [column.slug]
    unique = true
  }
}

table "entrants" {
  schema = schema.main
  column "id" {
    null = false
    type = integer
    auto_increment = true
  }
  column "name" {
    null = false
    type = text
  }
  column "seed" {
    null = true
    type = integer
  }
  column "standing" {
    null = true
    type = integer
  }
  column "tournament_id" {
    null = false 
    type = integer
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "tournament_id" {
    columns = [column.tournament_id]
    ref_columns = [table.tournaments.column.id]
    on_update = NO_ACTION
    on_delete = CASCADE
  }
}

table "matches" {
  schema = schema.main
  column "id" {
    null = false
    type = integer
    auto_increment = true
  }
  column "match_id" {
    null = false
    type = integer
  }
  column "tournament_id" {
    null = false 
    type = integer
  }
  column "entrant1_id" {
    null = true
    type = integer
  }
  column "entrant2_id" {
    null = true
    type = integer
  }
  column "winner_id" {
    null = true
    type = integer
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "tournament_id" {
    columns = [column.tournament_id]
    ref_columns = [table.tournaments.column.id]
    on_update = NO_ACTION
    on_delete = CASCADE
  }
  foreign_key "entrant1_id" {
    columns = [column.entrant1_id]
    ref_columns = [table.entrants.column.id]
    on_update = NO_ACTION
    on_delete = NO_ACTION
  }
  foreign_key "entrant2_id" {
    columns = [column.entrant2_id]
    ref_columns = [table.entrants.column.id]
    on_update = NO_ACTION
    on_delete = NO_ACTION
  }
  foreign_key "winner_id" {
    columns = [column.winner_id]
    ref_columns = [table.entrants.column.id]
    on_update = NO_ACTION
    on_delete = NO_ACTION
  }
}

schema "main" {}
