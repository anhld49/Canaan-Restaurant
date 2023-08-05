package constants

const (
	GetAllUsers                     = `select id, email, first_name, last_name, password, role, created_at, updated_at from public."user" u order by id`
	GetUser                         = `select id, email, first_name, last_name, password, role, created_at, updated_at from public."user" u where email = $1`
	GetUserByID                     = `select id, email, first_name, last_name, password, role, created_at, updated_at from public."user" u where id = $1`
	AddFriendToExistingFriendsArray = `update public.user
										set    friends = (select array_agg(distinct e) from unnest(friends || ARRAY[$2]) e)
										where  not friends @> ARRAY[$2] and email = $1
										and (not blocks @> ARRAY[$2] or blocks IS NULL);`
	AddFriendToNullFriendsArray = `UPDATE public.user SET friends = ARRAY[$2] where  email = $1 and friends IS NULL
									and (not blocks @> ARRAY[$2] or blocks IS NULL);`
	AddSubscribeToExistingSubscribeArray = `update public.user
											set    subscribe = (select array_agg(distinct e) from unnest(subscribe || ARRAY[$2]) e)
											where  not subscribe @> ARRAY[$2] and email = $1
											and (not blocks @> ARRAY[$2] or blocks IS NULL);`
	AddSubscribeToNullSubscribeArray = `UPDATE public.user SET subscribe = ARRAY[$2] where  email = $1 and subscribe IS NULL
										and (not blocks @> ARRAY[$2] or blocks IS NULL);`
	AddBlockToExistingSubscribeArray = `update public.user
										set    blocks = (select array_agg(distinct e) from unnest(blocks || ARRAY[$2]) e)
										where  not blocks @> ARRAY[$2] and email = $1`
	AddBlockToNullSubscribeArray = `UPDATE public.user SET blocks = ARRAY[$2] where  email = $1 and blocks IS NULL`
	VerifyBlock                  = `SELECT $2 = ANY (blocks::text[]) from public.user where email = $1`
)
