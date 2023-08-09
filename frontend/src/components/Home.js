import { Link } from "react-router-dom";
import Ticket from "./../images/movie_tickets.jpg";

const Home = () => {
  return (
    <>
      <div className="text-center">
        <h2>Find a restaurant near you!</h2>
        <hr />
        <Link to="/movies">
          <img src={Ticket} alt="movie tickets"></img>
        </Link>
      </div>
    </>
  );
};

export default Home;
