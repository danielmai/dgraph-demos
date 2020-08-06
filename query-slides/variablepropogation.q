{
  num_roles(func: eq(name@en, "Warwick Davis")) @cascade @normalize {
    paths as math(1)             # records number of paths to each character

    actor : name@en
    actor.film {
      performance.film @filter(allofterms(name@en, "Harry Potter")) {
        film_name : name@en
        characters : math(paths) # how many paths (i.e. characters) reach this film
      }
    }
  }
}
