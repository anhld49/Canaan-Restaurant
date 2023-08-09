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

	GetAllMenus           = `select id, restaurant_id, name, created_at, updated_at from public."menu" r order by id`
	GetMenuByID           = `select id, restaurant_id, name, created_at, updated_at from public."menu" r where id = $1`
	GetMenuByRestaurantId = `select id, restaurant_id, name, created_at, updated_at from public."menu" r where restaurant_id = $1`
	CreateMenu            = `insert into public.menu (restaurant_id, name, created_at, updated_at) values ($1, $2, $3, $4) returning id`
	UpdateMenu            = `update public.menu set restaurant_id = $1, name = $2, updated_at = $3 where id = $4 returning id`
	DeleteMenu            = `delete from public.menu where id = $1 returning id`

	GetAllDishs       = `select id, menu_id, name, price, created_at, updated_at from public."dish" r order by id`
	GetDishByID       = `select id, menu_id, name, price, created_at, updated_at from public."dish" r where id = $1`
	GetDishesByMenuId = `select id, menu_id, name, price, created_at, updated_at from public."dish" r where menu_id = $1`
	CreateDish        = `insert into public.dish (menu_id, name, price, created_at, updated_at) values ($1, $2, $3, $4, $5) returning id`
	UpdateDish        = `update public.dish set menu_id = $1, name = $2, price = $3, updated_at = $4 where id = $5 returning id`
	DeleteDish        = `delete from public.dish where id = $1 returning id`

	CreateOrderInit       = `insert into public.order (user_id, driver_id, amount, created_at, updated_at) values ($1, $2, 0.00, $3, $4) returning id`
	CreateOrderDish       = `insert into public.order_dish (order_id, dish_id, quantity, created_at, updated_at) values ($1, $2, $3, $4, $5)`
	GetDishPriceByDishId  = `select price from public.dish where id = $1`
	InsertAmountIntoOrder = `update public.order set amount = $1 where id = $2`
	GetOrderByID          = `select id, user_id, driver_id, amount, created_at, updated_at from public. order where id = $1`
	GetOrderDishByOrderId = `select dish_id, quantity from public.order_dish where order_id = $1`
	DeleteOrder           = `delete from public.order where id = $1 returning id`
	GetAllOrders          = `select id, user_id, driver_id, amount, created_at, updated_at from public. order where user_id = $1`
)
