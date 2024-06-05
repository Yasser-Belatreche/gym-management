data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./src/lib/persistence/psql/gorm/models",
    "--dialect", "postgres",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  url = "postgres://postgres:postgres@localhost:5432/postgres?search_path=public&sslmode=disable"
  dev = "docker://postgres/latest/dev?search_path=public"
  migration {
    dir = "file://src/lib/persistence/psql/gorm/migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}