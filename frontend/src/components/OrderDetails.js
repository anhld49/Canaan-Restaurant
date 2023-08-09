import { useEffect, useState } from "react";
import {
  useParams,
  Link,
  useNavigate,
  useOutletContext,
} from "react-router-dom";

const OrderDetails = () => {
  const [orderDetails, setOrderDetails] = useState([]);
  let { id } = useParams();
  const { jwtToken } = useOutletContext();
  const navigate = useNavigate();

  useEffect(() => {
    if (jwtToken === "") {
      navigate("/login");
      return;
    }
    const headers = new Headers();
    headers.append("Content-Type", "application/json");
    headers.append("Authorization", "Bearer " + jwtToken);

    const requestOptions = {
      method: "GET",
      headers: headers,
    };

    fetch(`/admin/orders/${id}`, requestOptions)
      .then((response) => response.json())
      .then((data) => {
        setOrderDetails(data.order_dishes);
      })
      .catch((err) => {
        console.log(err);
      });
  }, [jwtToken, navigate]);

  return (
    <div>
      <h2>Order Details</h2>
      <hr />
      <table className="table table-striped table-hover">
        <thead>
          <tr>
            <th>Dish Id</th>
            <th>Quantity</th>
          </tr>
        </thead>
        <tbody>
          {orderDetails.map((m) => (
            <tr key={m.dish_id}>
              <td>{m.dish_id}</td>
              <td>{m.quantity}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default OrderDetails;
