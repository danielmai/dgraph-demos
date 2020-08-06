{
  movies(func: gt(count(~genre), 30000), first: 1) @recurse(depth: 5, loop: true) {
    name@en
    ~genre (first:10) @filter(gt(count(starring), 2))
    starring (first: 2)
    performance.actor
  }
}
