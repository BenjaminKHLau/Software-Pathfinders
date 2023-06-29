import React from 'react';
import { useDispatch } from 'react-redux';
import { useNavigate } from 'react-router-dom';
import { logout } from '../../store/session';


const Logout = () => {
  const dispatch = useDispatch()
  const navigate = useNavigate()
  const onLogout = () => {
    dispatch(logout())
    navigate("/")
};

  return <button onClick={onLogout} className="nav-button">Logout</button>;
};

export default Logout;