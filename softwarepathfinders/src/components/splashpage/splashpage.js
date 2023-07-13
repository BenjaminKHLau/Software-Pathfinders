import { useSelector } from "react-redux"
import colorful from "../../pics/colorfullaptop.jpg"
import notes from "../../pics/noteslaptop.jpg"
import study from "../../pics/studylaptop.jpg"
import "./splashpage.css"

function Splashpage(){

    const session = useSelector(state => state.session)
    const user = session.user ? session.user : null;

    return(
        <section>
            <img src={notes} alt="hero" id="hero"/>
        Welcome to Software Pathfinders!

        <div>Hello, {user?.profile.FirstName}</div>
        </section>
    )
}

export default Splashpage