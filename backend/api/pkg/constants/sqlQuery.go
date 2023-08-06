package constants

const (
	GetAllUsers = `select id, email, first_name, last_name, password, role, created_at, updated_at from public."user" u order by id`
	GetUser     = `select id, email, first_name, last_name, password, role, created_at, updated_at from public."user" u where email = $1`
	GetUserByID = `select id, email, first_name, last_name, password, role, created_at, updated_at from public."user" u where id = $1`

	GetAllRestaurants = `select id, owner_id, name, address, created_at, updated_at from public."restaurant" r order by id`
	GetRestaurantByID = `select id, owner_id, name, address, created_at, updated_at from public."restaurant" r where id = $1`
	CreateRestaurant  = `insert into public.restaurant (owner_id, name, address, created_at, updated_at) values ($1, $2, $3, $4, $5) returning id`
	UpdateRestaurant  = `update public.restaurant set owner_id = $1, name = $2, address = $3, updated_at = $4 where id = $5 returning id`
	DeleteRestaurant  = `delete from public.restaurant where id = $1 returning id`

	GetAllMenus = `select id, restaurant_id, name, created_at, updated_at from public."menu" r order by id`
	GetMenuByID = `select id, restaurant_id, name, created_at, updated_at from public."menu" r where id = $1`
	CreateMenu  = `insert into public.menu (restaurant_id, name, created_at, updated_at) values ($1, $2, $3, $4) returning id`
	UpdateMenu  = `update public.menu set restaurant_id = $1, name = $2, updated_at = $3 where id = $4 returning id`
	DeleteMenu  = `delete from public.menu where id = $1 returning id`
)
