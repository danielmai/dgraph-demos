{
  spielberg as var(func: allofterms(name@en, "steven spielberg")) {
    uid
  }
  
  cameos(func: uid(spielberg)) @cascade @normalize {
    director : name@en
    director.film {
      film : name@en
      starring {
        performance.actor @filter(uid(spielberg)) {
          actor : name@en
        }
        performance.character {
          character : name@en
        }
      }
    }
  }
}
