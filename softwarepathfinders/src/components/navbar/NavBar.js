import { useEffect } from "react";
import { Link } from "react-router-dom";
import { PathGetAllThunk } from "../../store/paths";
import { useDispatch, useSelector } from "react-redux";
import { CohortGetAllThunk } from "../../store/cohorts";
import { PostGetAllThunk } from "../../store/posts";
import "./NavBar.css";
import Logout from "../auth/Logout";
// import { allUsersThunk } from "../../store/allUsers";

function NavBar() {
  const dispatch = useDispatch();
  const session = useSelector((state) => state.session);
  const user = session.user ? session.user : null;
  //   console.log("USER", user);

  useEffect(() => {
    dispatch(PathGetAllThunk())
      .then(() => dispatch(CohortGetAllThunk()))
      .then(() => dispatch(PostGetAllThunk()));
    //   .then(() => dispatch(allUsersThunk()))
  }, [dispatch]);

  return (
    <nav>
      <section className="nav-leftside">
        <Link to="/">
          <button className="nav-button">Home</button>
        </Link>
      </section>
      {!user && (
        <section className="nav-rightside">
          <Link to="/signup">
            <button className="nav-button">Sign Up</button>
          </Link>
          <Link to="/login">
            <button className="nav-button">Log In</button>
          </Link>
        </section>
      )}
      {user && (
        <section className="nav-rightside">
          {user.profile.Admin && (
            <Link to="/admin">
              <button className="nav-button">Admin</button>
            </Link>
          )}
          <Link to="/material">
            <button className="nav-button">Learn</button>
          </Link>
          <Link to="/cohort">
            <button className="nav-button">Cohort</button>
          </Link>
          <Logout />
        </section>
      )}
    </nav>
  );
}

export default NavBar;
