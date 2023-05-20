package main

import "github.com/google/uuid"

func inizializeService() {
	id := insertMovie(uuid.New(), "Inception", "2010", "A thief who steals corporate secrets through the use of dream-sharing technology is given the inverse task of planting an idea into the mind of a C.E.O.", 8.8)
	insertActor(uuid.New(), "Leonardo", "Di Caprio", id)
	insertActor(uuid.New(), "Joseph", "Gordon-Levitt", id)
	insertActor(uuid.New(), "Ellen", "Page", id)
	insertActor(uuid.New(), "Tom", "Hardy", id)
	insertDirector(uuid.New(), "Christopher", "Nolan", id)
}
