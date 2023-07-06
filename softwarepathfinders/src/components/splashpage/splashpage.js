import { useSelector } from "react-redux"

function Splashpage(){

    const session = useSelector(state => state.session)
    const user = session.user ? session.user : null;

    return(
        <section>
        Welcome to Software Pathfinders!

        <div>Hello, {user?.profile.FirstName}</div>
        </section>
    )
}

export default Splashpage