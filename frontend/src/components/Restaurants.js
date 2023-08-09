import { useEffect, useState } from "react";
import { Link } from "react-router-dom";

const Restaurants = () => {
    const [restaurants, setRestaurants] = useState([]);

    useEffect( () => {
        const headers = new Headers();
        headers.append("Content-Type", "application/json");

        const requestOptions = {
            method: "GET",
            headers: headers,
        }

        fetch(`http://localhost:8080/restaurants`, requestOptions)
            .then((response) => response.json())
            .then((data) => {
                setRestaurants(data);
            })
            .catch(err => {
                console.log(err);
            })

    }, []);

    return(
        <div>
            <h2>Restaurants</h2>
            <hr />
            <table className="table table-striped table-hover">
                <thead>
                    <tr>
                        <th>Index</th>
                        <th>Name</th>
                        <th>Address</th>
                    </tr>
                </thead>
                <tbody>
                    {restaurants.map((m) => (
                        <tr key={m.id}>
                            <td>
                                <Link to={`/restaurants/${m.id}`}>
                                    {m.id}
                                </Link>
                            </td>
                            <td>{m.name}</td>
                            <td>{m.address}</td>
                        </tr>    
                    ))}
                </tbody>
            </table>
        </div>
    )
}

export default Restaurants;