import { useEffect } from "react";
import { Link } from "react-router-dom";
import { PathGetAllThunk } from "../../store/paths";
import { useDispatch, useSelector } from "react-redux";
import { CohortGetAllThunk } from "../../store/cohorts";
import { PostGetAllThunk } from "../../store/posts";
import "./NavBar.css";
import Logout from "../auth/Logout";

function NavBar() {
  const dispatch = useDispatch();
  const session = useSelector((state) => state.session);
  const user = session.user ? session.user : null;
  console.log("USER", user);

  useEffect(() => {
    dispatch(PathGetAllThunk())
      .then(() => dispatch(CohortGetAllThunk()))
      .then(() => dispatch(PostGetAllThunk()));
  }, [dispatch]);

  return (
    <nav>
      <section className="nav-leftside">
        <button className="nav-button">
          <Link to="/">Home</Link>
        </button>
      </section>
      {!user && (
        <section className="nav-rightside">
          <button className="nav-button">
            <Link to="/signup">Sign Up</Link>
          </button>
          <button className="nav-button">
            <Link to="/login">Log In</Link>
          </button>
        </section>
      )}
      {user && (
        <section className="nav-rightside">
          <button className="nav-button">
            <Link to="/">PLACEHOLDER</Link>
          </button>
          <Logout />
        </section>
      )}
    </nav>
  );
}

export default NavBar;
