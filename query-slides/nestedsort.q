{
  movies as var(func: gt(count(~genre), 30000)) {
    name@en
    ~genre {
      date as initial_release_date
    }
    D as min(val(date))
  }

  movies(func: uid(movies), orderasc: val(D)) {
    name@en
  }
}
