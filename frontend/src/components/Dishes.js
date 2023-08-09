import { useEffect, useState } from "react";
import {
  useParams,
  Link,
  useNavigate,
  useOutletContext,
} from "react-router-dom";

const Dishes = () => {
  const [dishes, setDishes] = useState([]);
  const cart = {};
  let { id } = useParams();
  const { jwtToken } = useOutletContext();
  const navigate = useNavigate();

  const updateDish = (updateId) => {
    console.log("updateId", updateId)
    navigate(`/dishes/update/${updateId}`);
  }

  const deleteDish = (deletedId) => {
    if (jwtToken === "") {
      navigate("/login");
      return;
    }
    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("Authorization", "Bearer " + jwtToken);

    const requestOptions = {
      method: "DELETE",
      headers: headers,
    };

    fetch(`/owner/dishes/${deletedId}`, requestOptions)
      .then((response) => response.json())
      .then((data) => {
        alert("order deleted successfully");
        navigate(`/dishes/${id}`);
      })
      .catch((err) => {
        alert("order deleted failed", err);
      });
  }

  const createOrder = () => {
    if (jwtToken === "") {
      navigate("/login");
      return;
    }
    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("Authorization", "Bearer " + jwtToken);

    let order_dish = [];
    for (const key in cart) {
      order_dish.push({
        dish_id: +key,
        quantity: cart[key],
      })
    }


    const requestOptions = {
      method: "PUT",
      headers: headers,
      body: JSON.stringify({
        driver_id: 2,
        order_dish: order_dish,
      }),
    };

    fetch(`/admin/orders`, requestOptions)
      .then((response) => response.json())
      .then((data) => {
        alert("order created successfully");
        navigate("/admin/orders");
      })
      .catch((err) => {
        console.log(err);
      });
  };

  const addToCart = (id) => {
    if (cart[`${id}`]) {
      cart[`${id}`] = +cart[`${id}`] + 1;
    } else {
      cart[`${id}`] = 1;
    }
  };

  useEffect(() => {
    const headers = new Headers();
    headers.append("Content-Type", "application/json");

    const requestOptions = {
      method: "GET",
      headers: headers,
    };

    fetch(
      `http://localhost:8080/dishes/getDishesByMenuId/${id}`,
      requestOptions
    )
      .then((response) => response.json())
      .then((data) => {
        setDishes(data);
        console.log(`http://localhost:8080/dishes/getDishesByMenuId/${id}`, data)
      })
      .catch((err) => {
        console.log(err);
      });
  }, []);

  return (
    <div>
      <h2>Dishes</h2>
      <h3>
        <button
          type="button"
          className="btn btn-primary"
          onClick={() => createOrder()}
        >
          Create Order
        </button>
      </h3>
      <hr />
      <table className="table table-striped table-hover">
        <thead>
          <tr>
            <th>Index</th>
            <th>Name</th>
            <th>Price</th>
            <th></th>
          </tr>
        </thead>
        <tbody>
          {dishes.map((m) => (
            <tr key={m.id}>
              <td>
                <Link to={`/dishes/${m.id}`}>{m.id}</Link>
              </td>
              <td>{m.name}</td>
              <td>${m.price}</td>
              <td>
                <button
                  type="button"
                  className="btn btn-primary"
                  onClick={() => addToCart(`${m.id}`)}
                >
                  Add To Cart
                </button>
                <button
                  type="button"
                  className="btn btn-primary"
                  onClick={() => updateDish(`${m.id}`)}
                >
                  Update this Dish
                </button>
                <button
                  type="button"
                  className="btn btn-primary"
                  onClick={() => deleteDish(`${m.id}`)}
                >
                  Delete this Dish
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default Dishes;
