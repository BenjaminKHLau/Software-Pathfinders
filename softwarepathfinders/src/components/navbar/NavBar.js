import { useEffect } from "react"
import { PathGetAllThunk } from "../../store/paths"
import { useDispatch } from "react-redux"



function NavBar(){

    const dispatch = useDispatch()

    useEffect(()=>{
        dispatch(PathGetAllThunk())
    },[])

    return (
        <>
        HELLO
        
        </>
    )
}

export default NavBar