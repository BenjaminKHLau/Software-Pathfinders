import { useEffect } from "react"
import { PathGetAllThunk } from "../../store/paths"
import { useDispatch } from "react-redux"
import { CohortGetAllThunk } from "../../store/cohorts"
import { PostGetAllThunk } from "../../store/posts"



function NavBar(){

    const dispatch = useDispatch()

    useEffect(()=>{
        dispatch(PathGetAllThunk())
        .then(()=> dispatch(CohortGetAllThunk()))
        .then(()=> dispatch(PostGetAllThunk()))
    },[])

    return (
        <>
        HELLO
        
        </>
    )
}

export default NavBar