import { useSelector } from "react-redux";
import colorful from "../../pics/colorfullaptop.jpg";
import notes from "../../pics/noteslaptop.jpg";
import study from "../../pics/studylaptop.jpg";
import "./splashpage.css";

function Splashpage() {
  const session = useSelector((state) => state.session);
  const user = session.user ? session.user : null;

  return (
    <section>
      {/* <img src={notes} alt="hero" id="hero"/> */}
      <div id="hero">
        <div className="hero-text">Welcome to Software Pathfinders!</div>
      </div>

      <div>Hello, {user?.profile.FirstName}</div>
    </section>
  );
}

export default Splashpage;
