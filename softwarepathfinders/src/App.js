import React, { useState, useEffect } from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
// import { Provider } from 'react-redux';
import { useDispatch } from 'react-redux';
import SignUpForm from './components/auth/SignUpForm';
import LoginForm from './components/auth/LoginForm';
import { authenticate } from './store/session';
import './App.css';
// import {store} from './store/store';
import NavBar from './components/navbar/NavBar';
import Splashpage from './components/splashpage/splashpage';

function App() {
  const [loaded, setLoaded] = useState(false);
  const dispatch = useDispatch();

  useEffect(() => {
      setLoaded(true);
  }, [dispatch]);


  return loaded && (
    <BrowserRouter>
        <NavBar />

        <Routes>
          <Route path='/signup' element={<SignUpForm />} exact={true} />
          <Route path='/login' element={<LoginForm />} exact={true} />
          <Route path='/' element={<Splashpage />} exact={true} />
        </Routes>

    </BrowserRouter>
  );
}

export default App;
