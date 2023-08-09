import { useEffect, useState } from "react";
import { useNavigate, useParams, useOutletContext } from "react-router-dom";

const UpdateDish = () => {
  const [restaurants, setRestaurants] = useState([]);
  const [menus, setMenus] = useState([]);
  const [currMenu, setCurrMenu] = useState(0);
  const [dishName, setDishName] = useState("");
  const [dishPrice, setDishPrice] = useState(0);

  let { id } = useParams();

  const [originalMenu, setOriginalMenu] = useState({});

  const { jwtToken } = useOutletContext();
  const navigate = useNavigate();

  const handleUpdateDish = () => {
    if (jwtToken === "") {
      navigate("/login");
      return;
    }

    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("Authorization", "Bearer " + jwtToken);

    const requestOptions = {
      method: "PUT",
      headers: headers,
      body: JSON.stringify({
        menu_id: currMenu,
        name: dishName,
        price: +dishPrice,
      }),
    };

    fetch(`http://localhost:8080/owner/dishes`, requestOptions)
      .then((response) => response.json())
      .then((data) => {
        alert("Dish was updated successfully");
        navigate(`/dishes/${currMenu}`);
      })
      .catch((err) => {
        alert("You don't have permission to do this: ", err);
        navigate(`/dishes/${currMenu}`);
      });
  };

  const handlerChangeDishName = (event) => {
    setDishName(event.target.value);
  };

  const handlerChangeDishPrice = (event) => {
    setDishPrice(event.target.value);
  };

  const selectMenu = (event) => {
    const menuId = event.target.value;

    const headers = new Headers();
    headers.append("Content-Type", "application/json");

    const requestOptions = {
      method: "GET",
      headers: headers,
    };

    fetch(
      `http://localhost:8080/menus/getMenuByRestaurantId/${menuId}`,
      requestOptions
    )
      .then((response) => response.json())
      .then((data) => {
        setMenus(data);
        setCurrMenu(data[0].id);
      })
      .catch((err) => {
        console.log(err);
      });
  };

  useEffect(() => {
    if (jwtToken === "") {
      navigate("/login");
      return;
    }

    const headers = new Headers();
    headers.append("Content-Type", "application/json");

    const requestOptions = {
      method: "GET",
      headers: headers,
    };

    fetch(`http://localhost:8080/restaurants`, requestOptions)
      .then((response) => response.json())
      .then((data) => {
        setRestaurants(data);
      })
      .catch((err) => {
        console.log(err);
      });

    fetch(`http://localhost:8080/menus/getMenuByRestaurantId/1`, requestOptions)
      .then((response) => response.json())
      .then((data) => {
        setMenus(data);
        setCurrMenu(data[0].id);
      })
      .catch((err) => {
        console.log(err);
      });

    fetch(`http://localhost:8080/dishes/${id}`, requestOptions)
      .then((response) => response.json())
      .then((data) => {
        setOriginalMenu({
          menu_id: data.menu_id,
          name: data.name,
          price: data.price,
        });
      })
      .catch((err) => {
        console.log(err);
      });
  }, []);

  return (
    <section className="order-form m-4">
      <div className="container pt-4">
        <div className="row">
          <div className="col-12 px-4">
            <h1>Update Dish</h1>
            <hr className="mt-1" />
          </div>

          <div className="col-12">
            <div className="row mt-3 mx-4">
              <div className="col-12">
                <label className="order-form-label">Select Restaurant</label>
              </div>
              <div className="col-12">
                <div className="form-outline">
                  <select
                    className="form-select"
                    aria-label="Default select example"
                    onChange={selectMenu}
                  >
                    {restaurants.map((restaurant) => (
                      <option key={restaurant.id} value={restaurant.id}>
                        {restaurant.name}
                      </option>
                    ))}
                  </select>
                </div>
              </div>
            </div>

            {menus.length > 0 && (
              <div className="row mt-3 mx-4">
                <div className="col-12">
                  <label className="order-form-label">Select Menu</label>
                </div>
                <div className="col-12">
                  <div className="form-outline">
                    <select
                      className="form-select"
                      aria-label="Default select example"
                    >
                      {menus.map((menu) => (
                        <option key={menu.id} value={menu.id}>
                          {menu.name}
                        </option>
                      ))}
                    </select>
                  </div>
                </div>
              </div>
            )}

            <div className="row mt-3 mx-4">
              <div className="col-12">
                <label className="order-form-label">Dish Name</label>
              </div>
              <div className="col-12">
                <div className="form-outline">
                  <input
                    type="text"
                    id="dish_name"
                    className="form-control order-form-input"
                    onChange={handlerChangeDishName}
                    defaultValue={originalMenu.name}
                  />
                </div>
              </div>
            </div>

            <div className="row mt-3 mx-4">
              <div className="col-12">
                <label className="order-form-label">Price ($)</label>
              </div>
              <div className="col-12">
                <div className="form-outline">
                  <input
                    type="text"
                    id="dish_price"
                    className="form-control order-form-input"
                    onChange={handlerChangeDishPrice}
                    defaultValue={originalMenu.price}
                  />
                </div>
              </div>
            </div>

            <div className="row mt-3">
              <div className="col-12">
                <button
                  type="button"
                  id="btnSubmit"
                  className="btn btn-primary d-block mx-auto btn-submit"
                  onClick={handleUpdateDish}
                >
                  Submit
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
};

export default UpdateDish;
