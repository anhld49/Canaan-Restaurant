import { useEffect, useState } from "react";
import { useParams, Link } from "react-router-dom";

const Menus = () => {
    const [menus, setMenus] = useState([]);
    let { id } = useParams();

    useEffect( () => {
        const headers = new Headers();
        headers.append("Content-Type", "application/json");

        const requestOptions = {
            method: "GET",
            headers: headers,
        }

        fetch(`http://localhost:8080/menus/getMenuByRestaurantId/${id}`, requestOptions)
            .then((response) => response.json())
            .then((data) => {
                setMenus(data);
            })
            .catch(err => {
                console.log(err);
            })

    }, []);

    return(
        <div>
            <h2>Menus</h2>
            <hr />
            <table className="table table-striped table-hover">
                <thead>
                    <tr>
                        <th>Index</th>
                        <th>Name</th>
                    </tr>
                </thead>
                <tbody>
                    {menus.map((m) => (
                        <tr key={m.id}>
                            <td>
                                <Link to={`/dishes/${m.id}`}>
                                    {m.id}
                                </Link>
                            </td>
                            <td>{m.name}</td>
                        </tr>    
                    ))}
                </tbody>
            </table>
        </div>
    )
}

export default Menus;