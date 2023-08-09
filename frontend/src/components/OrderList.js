import { useEffect, useState } from "react";
import { Link, useNavigate, useOutletContext } from "react-router-dom";

const OrderList = () => {
    const [orders, setOrders] = useState([]);
    const { jwtToken } = useOutletContext();
    const navigate = useNavigate();

    useEffect( () => {
        if (jwtToken === "") {
            navigate("/login");
            return
        }
        const headers = new Headers();
        headers.append("Content-Type", "application/json");
        headers.append("Authorization", "Bearer " + jwtToken);

        const requestOptions = {
            method: "GET",
            headers: headers,
        }

        fetch(`/admin/orders`, requestOptions)
            .then((response) => response.json())
            .then((data) => {
                setOrders(data);
            })
            .catch(err => {
                console.log(err);
            })

    }, [jwtToken, navigate]);

    return(
        <div>
            <h2>Orders</h2>
            <hr />
            <table className="table table-striped table-hover">
                <thead>
                    <tr>
                        <th>Index</th>
                        <th>Driver Id</th>
                        <th>Amount</th>
                    </tr>
                </thead>
                <tbody>
                    {orders.map((m) => (
                        <tr key={m.id}>
                            <td>
                                <Link to={`/admin/orders/${m.id}`}>
                                    {m.id}
                                </Link>
                            </td>
                            <td>{m.driver_id}</td>
                            <td>{m.amount}</td> 
                        </tr>    
                    ))}
                </tbody>
            </table>
        </div>
    )
}

export default OrderList;