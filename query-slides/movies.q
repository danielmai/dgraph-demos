{
  director(func: allofterms(name@en, "steven spielberg")) {
    name@en
    director.film {
      name@en
      initial_release_date
      country {
        name@en
      }
      starring {
        performance.actor {
          name@en
        }
        performance.character {
          name@en
        }
      }
      genre {
        name@en
      }
    }
  }
}
