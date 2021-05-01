package data

type Table int

const (
	DeleteAccounts Table = iota
	DeleteGenres
	DeleteHalls
	DeleteLocations
	DeletePerformances
	DeletePlaces
	DeletePosters
	DeletePrices
	DeleteRoles
	DeleteSchedules
	DeleteSectors
	DeleteTickets
	DeleteUsers
)

func (t Table) String() string {
	return []string{"accounts", "genres", "halls", "locations", "performances", "places",
		"posters", "prices", "roles", "schedules", "sectors", "tickets", "users"}[t]
}

const (
	readAllTicketsQuery = "SELECT tickets.id, p.name, g.name, p.duration, " +
		"s.date, h.name, h.capacity, l.address, l.phone_number, s2.name, p2.name, p3.price, " +
		"tickets.date_of_issue, tickets.paid, tickets.reservation, tickets.destroyed " +
		"FROM tickets JOIN schedules s on s.id = tickets.schedule_id " +
		"JOIN performances p on s.performance_id = p.id " +
		"JOIN genres g on p.genre_id = g.id " +
		"JOIN halls h on s.hall_id = h.id " +
		"JOIN locations l on h.location_id = l.id " +
		"JOIN places p2 on tickets.place_id = p2.id " +
		"JOIN sectors s2 on p2.sector_id = s2.id " +
		"JOIN prices p3 on p.id = p3.performance_id and s2.id = p3.sector_id"

	readAllPostersQuery = "SELECT posters.id, p.name, g.name, p.duration, s.date, h.name, " +
		"h.capacity, l.address, l.phone_number, posters.comment " +
		"FROM posters JOIN schedules s on s.id = posters.schedule_id " +
		"JOIN performances p on s.performance_id = p.id " +
		"JOIN genres g on p.genre_id = g.id " +
		"JOIN halls h on s.hall_id = h.id " +
		"JOIN locations l on h.location_id = l.id"
	readAllUsersQuery = "SELECT users.id, users.first_name, users.last_name, r.name, a.address, " +
		"a.phone_number, users.phone_number FROM users " +
		"JOIN roles r on users.role_id = r.id " +
		"JOIN locations a on a.id = users.account_id " +
		"WHERE users.account_id = $1"
	insertAccount = "INSERT INTO accounts(first_name, last_name, phone_number, email) " +
		"VALUES ($1, $2, $3, $4) RETURNING id"
	updateAccount = "UPDATE accounts SET first_name = $1, last_name = $2, phone_number = $3, email = $4 " +
		"WHERE id = $5"
	insertGenre       = "INSERT INTO genres (name) VALUES ($1) RETURNING id"
	updateGenre       = "UPDATE genres SET name = $1 WHERE id = $2"
	insertHall        = "INSERT INTO halls (account_id, name, capacity, location_id) VALUES ($1, $2, $3, $4) RETURNING id"
	updateHall        = "UPDATE halls SET account_id = $1, name = $2, capacity = $3, location_id = $4 WHERE id = $5"
	insertLocation    = "INSERT INTO locations (account_id, address, phone_number) VALUES ($1, $2, $3) RETURNING id"
	updateLocation    = "UPDATE locations SET account_id = $1, address = $2, phone_number = $3 WHERE id = $4"
	insertPerformance = "INSERT INTO performances (account_id, name, genre_id, duration) VALUES ($1, $2, $3, $4) RETURNING id"
	updatePerformance = "UPDATE performances SET account_id = $1, name = $2, genre_id = $3, duration = $4 WHERE id = $5"
	insertPlace       = "INSERT INTO places (sector_id, name) VALUES ($1, $2) RETURNING id"
	updatePlace       = "UPDATE places SET sector_id = $1, name = $2 WHERE id = $3"
	insertPoster      = "INSERT INTO posters (account_id, schedule_id, comment) VALUES ($1, $2, $3) RETURNING id"
	updatePoster      = "UPDATE posters SET account_id = $1, schedule_id = $2, comment = $3 WHERE id = $4"
	insertPrice       = "INSERT INTO prices (account_id, sector_id, performance_id, price) VALUES ($1, $2, $3, $4) RETURNING id"
	updatePrice       = "UPDATE prices SET account_id = $1, sector_id = $2, performance_id = $3, price = $4 WHERE id = $5"
	insertRole        = "INSERT INTO roles (name) VALUES ($1) RETURNING id"
	updateRole        = "UPDATE roles SET name = $1 WHERE id = $2"
	insertSchedule    = "INSERT INTO schedules (account_id, performance_id, date, hall_id) VALUES ($1, $2, $3, $4) RETURNING id"
	updateSchedule    = "UPDATE schedules SET account_id = $1, performance_id = $2, date = $3, hall_id = $4 WHERE id = $5"
	insertSector      = "INSERT INTO sectors (name) VALUES ($1) RETURNING id"
	updateSector      = "UPDATE sectors SET name = $1 WHERE id = $2"
	insertTicket      = "INSERT INTO tickets (account_id, schedule_id, place_id, date_of_issue, paid, reservation, destroyed) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	updateTicket = "UPDATE tickets SET account_id = $1, schedule_id = $2, place_id = $3, date_of_issue = $4, paid = $5, " +
		"reservation = $6, destroyed = $7 WHERE id = $8"
	insertUser = "INSERT INTO users (account_id, first_name, last_name, role_id, location_id, phone_number) " +
		"VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	updateUser = "UPDATE users SET account_id = $1, first_name = $2, last_name = $3, role_id = $4, location_id = $5, " +
		"phone_number = $6 WHERE id = $7"
	deleteBegin = "DELETE FROM "
	deleteEnd   = " WHERE id = $1"
)
