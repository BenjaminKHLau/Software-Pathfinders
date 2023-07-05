import { useSelector } from "react-redux"

function Splashpage(){

    const session = useSelector(state => state.session)
    const user = session.user ? session.user : null;

    return(
        <section>
        Welcome to Software Pathfinders!

        <div>{user?.profile.Email}</div>
        </section>
    )
}

export default Splashpage